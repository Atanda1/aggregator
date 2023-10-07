package services

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/paycrest/paycrest-protocol/ent"
	"github.com/paycrest/paycrest-protocol/ent/fiatcurrency"
	"github.com/paycrest/paycrest-protocol/ent/lockpaymentorder"
	"github.com/paycrest/paycrest-protocol/ent/provideravailability"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
	"github.com/paycrest/paycrest-protocol/ent/providerrating"
	"github.com/paycrest/paycrest-protocol/ent/provisionbucket"
	"github.com/paycrest/paycrest-protocol/storage"
	"github.com/paycrest/paycrest-protocol/types"
	"github.com/paycrest/paycrest-protocol/utils"
	"github.com/paycrest/paycrest-protocol/utils/logger"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type PriorityQueueService struct{}

// NewPriorityQueueService creates a new instance of PriorityQueueService
func NewPriorityQueueService() *PriorityQueueService {
	return &PriorityQueueService{}
}

// ProcessBucketQueues creates a priority queue for each bucket and saves it to redis
func (s *PriorityQueueService) ProcessBucketQueues(ctx context.Context) error {

	buckets, err := s.GetProvidersByBucket(ctx)
	if err != nil {
		return fmt.Errorf("failed to process bucket queues: %w", err)
	}

	var wg sync.WaitGroup

	for _, bucket := range buckets {
		wg.Add(1)
		go s.CreatePriorityQueueForBucket(ctx, bucket)
	}

	wg.Wait()

	return nil
}

// GetProvidersByBucket returns a list of providers grouped by bucket
func (s *PriorityQueueService) GetProvidersByBucket(ctx context.Context) ([]*ent.ProvisionBucket, error) {
	buckets, err := storage.Client.ProvisionBucket.
		Query().
		Select(provisionbucket.EdgeProviderProfiles).
		WithProviderProfiles(func(ppq *ent.ProviderProfileQuery) {
			ppq.WithProviderRating(func(prq *ent.ProviderRatingQuery) {
				prq.Select(providerrating.FieldTrustScore)
			})
			ppq.Select(providerprofile.FieldID)

			// Filter only providers that are always available
			// or are available until one hour from now
			// TODO: the duration should be a config setting
			oneHourFromNow := time.Now().Add(time.Hour)
			ppq.Where(
				providerprofile.HasAvailabilityWith(
					provideravailability.And(
						provideravailability.CadenceEQ(provideravailability.CadenceAlways),
						provideravailability.EndTimeGTE(oneHourFromNow),
					),
				),
			)
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

// getProviderRate returns the rate for a provider
func (s *PriorityQueueService) getProviderRate(provider *ent.ProviderProfile) (decimal.Decimal, error) {
	// TODO: implement fetching of provider rate. this also includes calculating floating rate
	return decimal.Decimal{}, nil
}

// CreatePriorityQueueForBucket creates a priority queue for a bucket and saves it to redis
func (s *PriorityQueueService) CreatePriorityQueueForBucket(ctx context.Context, bucket *ent.ProvisionBucket) {
	// Create a slice to store the provider profiles sorted by trust score
	providers := bucket.Edges.ProviderProfiles
	sort.SliceStable(providers, func(i, j int) bool {
		trustScoreI, _ := providers[i].Edges.ProviderRating.TrustScore.Float64()
		trustScoreJ, _ := providers[j].Edges.ProviderRating.TrustScore.Float64()
		return trustScoreI > trustScoreJ // Sort in descending order
	})

	// Enqueue provider ID and rate as a single string into the circular queue
	redisKey := fmt.Sprintf("bucket_%s_%d_%d", bucket.Currency, bucket.MinAmount, bucket.MaxAmount)

	_, err := storage.RedisClient.Del(ctx, redisKey).Result() // delete existing queue
	if err != nil {
		logger.Errorf("failed to delete existing circular queue: %v", err)
	}

	for _, provider := range providers {
		providerID := provider.ID
		rate, _ := s.getProviderRate(provider)

		// Serialize the provider ID and rate into a single string
		data := fmt.Sprintf("%s:%s", providerID, rate)

		// Enqueue the serialized data into the circular queue
		err := storage.RedisClient.RPush(ctx, redisKey, data).Err()
		if err != nil {
			logger.Errorf("failed to enqueue provider data to circular queue: %v", err)
		}
	}
}

// CreatePriorityQueueForDefaultBucket creates a priority queue for the default bucket and saves it to redis
func (s *PriorityQueueService) CreatePriorityQueueForDefaultBucket(ctx context.Context, currency *ent.FiatCurrency) {
	// Fetch providers with is_partner == true
	providers, err := storage.Client.ProviderProfile.
		Query().
		Where(
			providerprofile.IsPartnerEQ(true),
			providerprofile.HasCurrencyWith(fiatcurrency.IDEQ(currency.ID)),
		).
		// WithProviderRating(). //  should we sort partner providers by trust score?
		All(ctx)
	if err != nil {
		logger.Errorf("failed to fetch partner providers: %v", err)
		return
	}

	// Create default buckets
	for _, provider := range providers {
		providerID := provider.ID

		// Enqueue provider ID and market rate as a single string into the circular queue
		redisKey := fmt.Sprintf("bucket_%s_default", currency.Code)
		data := fmt.Sprintf("%s:%s", providerID, currency.MarketRate)
		err := storage.RedisClient.RPush(ctx, redisKey, data).Err()
		if err != nil {
			logger.Errorf("failed to enqueue provider data to circular queue: %v", err)
		}
	}
}

// AssignLockPaymentOrders assigns lock payment orders to providers
func (s *PriorityQueueService) AssignLockPaymentOrder(ctx context.Context, order types.LockPaymentOrderFields) error {
	go s.ReassignUnfulfilledLockOrders(ctx)

	// Get the first provider from the circular queue
	redisKey := fmt.Sprintf("bucket_%s_%d_%d", order.ProvisionBucket.Currency, order.ProvisionBucket.MinAmount, order.ProvisionBucket.MaxAmount)

	// Start a Redis transaction
	pipe := storage.RedisClient.TxPipeline()

	for index := 0; ; index++ {
		providerData, err := pipe.LIndex(ctx, redisKey, int64(index)).Result()
		if err != nil {
			if err == redis.Nil {
				// Assign to top provider in default bucket and rotate the default bucket queue
				s.AssignLockOrderToDefaultBucket(ctx, order)
				logger.Errorf("failed to access index %d from circular queue: %v", index, err)
			}
			break
		}

		// Extract the rate from the data (assuming it's in the format "providerID:rate")
		parts := strings.Split(providerData, ":")
		if len(parts) != 2 {
			logger.Errorf("invalid data format at index %d: %s", index, providerData)
			continue // Skip this entry due to invalid format
		}
		providerID := parts[0]

		// Skip entry is provider is excluded
		excludeList, err := pipe.LRange(ctx, fmt.Sprintf("order_exclude_list_%d", order.ID), 0, -1).Result()
		if err != nil {
			logger.Errorf("failed to get exclude list for order %d: %v", order.ID, err)
			return err
		}
		if utils.ContainsString(excludeList, providerID) {
			continue
		}

		rate, err := decimal.NewFromString(parts[1])
		if err != nil {
			logger.Errorf("failed to parse rate at index %d: %v", index, err)
			continue // Skip this entry due to parsing error
		}

		if rate.Equal(order.Rate) {
			// Found a match for the rate
			if index == 0 {
				// Match found at index 0, perform LPOP to dequeue
				data, err := pipe.LPop(ctx, redisKey).Result()
				if err != nil {
					logger.Errorf("failed to dequeue from circular queue: %v", err)
					return err
				}

				// Enqueue data to the end of the queue
				err = pipe.RPush(ctx, redisKey, data).Err()
				if err != nil {
					logger.Errorf("failed to enqueue to circular queue: %v", err)
					return err
				}
			}

			// Assign the order to the provider and save it to Redis
			orderKey := fmt.Sprintf("order_request_%d", order.ID)
			orderRequestData := map[string]interface{}{
				"amount":      order.Amount.Mul(order.Rate),
				"token":       order.Token.Symbol,
				"institution": order.Institution,
				"provider_id": providerID,
			}

			err = pipe.HSet(ctx, orderKey, orderRequestData).Err()
			if err != nil {
				logger.Errorf("failed to map order to a provider in Redis: %v", err)
				return err
			}

			// Set a TTL for the order request
			err = pipe.ExpireAt(ctx, orderKey, time.Now().Add(OrderConf.OrderRequestValidity)).Err()
			if err != nil {
				logger.Errorf("failed to set TTL for order request: %v", err)
				return err
			}

			// TODO: Send POST request to the provider's webhook (for automatic provider case)

			// TODO: Send out a push notification to the provider (for manual provider case)
		}

	}

	// Execute all Redis commands within the transaction atomically
	_, err := pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("failed to execute Redis transaction: %v", err)
		return err
	}

	return nil
}

// AssignLockOrderToDefaultBucket assigns lock payment orders to providers in the default bucket
func (s *PriorityQueueService) AssignLockOrderToDefaultBucket(ctx context.Context, order types.LockPaymentOrderFields) error {
	pipe := storage.RedisClient.TxPipeline()
	data, err := pipe.LPop(ctx, fmt.Sprintf("bucket_%s_default", order.ProvisionBucket.Currency)).Result()
	if err != nil {
		logger.Errorf("failed to dequeue from circular queue: %v", err)
		return err
	}

	// Enqueue data to the end of the queue
	err = pipe.RPush(ctx, fmt.Sprintf("bucket_%s_default", order.ProvisionBucket.Currency), data).Err()
	if err != nil {
		logger.Errorf("failed to enqueue to circular queue: %v", err)
		return err
	}

	// Extract the rate from the data (assuming it's in the format "providerID:rate")
	parts := strings.Split(data, ":")
	if len(parts) != 2 {
		logger.Errorf("invalid data format: %s", data)
		return err
	}
	providerID := parts[0]

	// Assign the order to the provider and save it to Redis
	orderKey := fmt.Sprintf("order_request_%d", order.ID)
	orderRequestData := map[string]interface{}{
		"amount":      order.Amount.Mul(order.Rate),
		"token":       order.Token.Symbol,
		"institution": order.Institution,
		"provider_id": providerID,
	}

	err = pipe.HSet(ctx, orderKey, orderRequestData).Err()
	if err != nil {
		logger.Errorf("failed to map order to a provider in Redis: %v", err)
		return err
	}

	// Set a TTL for the order request
	err = pipe.ExpireAt(ctx, orderKey, time.Now().Add(OrderConf.OrderRequestValidity)).Err()
	if err != nil {
		logger.Errorf("failed to set TTL for order request: %v", err)
		return err
	}

	// Execute all Redis commands within the transaction atomically
	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Errorf("failed to execute Redis transaction: %v", err)
		return err
	}

	return nil
}

// ReassignStaleOrderRequest reassigns expired order requests to providers
func (s *PriorityQueueService) ReassignStaleOrderRequest(ctx context.Context, orderRequestChan <-chan *redis.Message) {
	for msg := range orderRequestChan {
		key := strings.Split(msg.Payload, "_")
		orderID := key[len(key)-1]

		orderUUID, err := uuid.Parse(orderID)
		if err != nil {
			logger.Errorf("ReassignStaleOrderRequest: %v", err)
			return
		}

		// Get the order from the database
		order, err := storage.Client.LockPaymentOrder.
			Query().
			Where(
				lockpaymentorder.IDEQ(orderUUID),
			).
			WithProvisionBucket().
			Only(context.Background())
		if err != nil {
			logger.Errorf("ReassignStaleOrderRequest: %v", err)
			return
		}

		orderFields := types.LockPaymentOrderFields{
			ID:                order.ID,
			OrderID:           order.OrderID,
			Amount:            order.Amount,
			Rate:              order.Rate,
			BlockNumber:       order.BlockNumber,
			Institution:       order.Institution,
			AccountIdentifier: order.AccountIdentifier,
			AccountName:       order.AccountName,
			ProvisionBucket:   order.Edges.ProvisionBucket,
		}

		// Assign the order to a provider
		err = s.AssignLockPaymentOrder(ctx, orderFields)
		if err != nil {
			logger.Errorf("ReassignStaleOrderRequest: %v", err)
			return
		}
	}
}

// ReassignUnfulfilledLockOrders reassigns lockOrder unfulfilled within a time frame.
func (s *PriorityQueueService) ReassignUnfulfilledLockOrders(ctx context.Context) error {
	// query unfulfilled lock orders.
	lockOrders, err := storage.Client.LockPaymentOrder.
		Query().
		Where(
			lockpaymentorder.Not(lockpaymentorder.HasFulfillment()),
			lockpaymentorder.Or(
				lockpaymentorder.And(
					lockpaymentorder.StatusEQ(lockpaymentorder.StatusProcessing),
					lockpaymentorder.UpdatedAtLTE(time.Now().Add(-OrderConf.OrderFulfillmentValidity*time.Minute)),
				),
				lockpaymentorder.StatusEQ(lockpaymentorder.StatusCancelled),
			),
		).
		WithToken().
		WithProvider().
		WithProvisionBucket().
		All(ctx)
	if err != nil {
		return err
	}

	// unassign unfulfilled lock orders.
	_, err = storage.Client.LockPaymentOrder.
		Update().
		Where(
			lockpaymentorder.StatusEQ(lockpaymentorder.StatusProcessing),
			lockpaymentorder.UpdatedAtLTE(time.Now().Add(-OrderConf.OrderFulfillmentValidity*time.Minute)),
			lockpaymentorder.Not(lockpaymentorder.HasFulfillment()),
		).
		SetStatus(lockpaymentorder.StatusPending).
		Save(ctx)
	if err != nil {
		return err
	}

	for _, order := range lockOrders {
		lockPaymentOrder := types.LockPaymentOrderFields{
			Token:             order.Edges.Token,
			OrderID:           order.OrderID,
			Amount:            order.Amount,
			Rate:              order.Rate,
			BlockNumber:       order.BlockNumber,
			Institution:       order.Institution,
			AccountIdentifier: order.AccountIdentifier,
			AccountName:       order.AccountName,
			ProviderID:        order.Edges.Provider.ID,
			ProvisionBucket:   order.Edges.ProvisionBucket,
		}

		err := s.AssignLockPaymentOrder(ctx, lockPaymentOrder)
		if err != nil {
			logger.Errorf("task reassign unfulfilled lock order with id: %s => %v\n", order.OrderID, err)
		}
	}

	return nil
}
