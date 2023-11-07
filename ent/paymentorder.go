// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/paymentorder"
	"github.com/paycrest/protocol/ent/paymentorderrecipient"
	"github.com/paycrest/protocol/ent/receiveaddress"
	"github.com/paycrest/protocol/ent/senderprofile"
	"github.com/paycrest/protocol/ent/token"
	"github.com/shopspring/decimal"
)

// PaymentOrder is the model entity for the PaymentOrder schema.
type PaymentOrder struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// AmountPaid holds the value of the "amount_paid" field.
	AmountPaid decimal.Decimal `json:"amount_paid,omitempty"`
	// TxHash holds the value of the "tx_hash" field.
	TxHash string `json:"tx_hash,omitempty"`
	// ReceiveAddressText holds the value of the "receive_address_text" field.
	ReceiveAddressText string `json:"receive_address_text,omitempty"`
	// Label holds the value of the "label" field.
	Label string `json:"label,omitempty"`
	// Status holds the value of the "status" field.
	Status paymentorder.Status `json:"status,omitempty"`
	// LastUsed holds the value of the "last_used" field.
	LastUsed time.Time `json:"last_used,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentOrderQuery when eager-loading is set.
	Edges                         PaymentOrderEdges `json:"edges"`
	api_key_payment_orders        *uuid.UUID
	sender_profile_payment_orders *uuid.UUID
	token_payment_orders          *int
	selectValues                  sql.SelectValues
}

// PaymentOrderEdges holds the relations/edges for other nodes in the graph.
type PaymentOrderEdges struct {
	// SenderProfile holds the value of the sender_profile edge.
	SenderProfile *SenderProfile `json:"sender_profile,omitempty"`
	// Token holds the value of the token edge.
	Token *Token `json:"token,omitempty"`
	// ReceiveAddress holds the value of the receive_address edge.
	ReceiveAddress *ReceiveAddress `json:"receive_address,omitempty"`
	// Recipient holds the value of the recipient edge.
	Recipient *PaymentOrderRecipient `json:"recipient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// SenderProfileOrErr returns the SenderProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentOrderEdges) SenderProfileOrErr() (*SenderProfile, error) {
	if e.loadedTypes[0] {
		if e.SenderProfile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: senderprofile.Label}
		}
		return e.SenderProfile, nil
	}
	return nil, &NotLoadedError{edge: "sender_profile"}
}

// TokenOrErr returns the Token value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentOrderEdges) TokenOrErr() (*Token, error) {
	if e.loadedTypes[1] {
		if e.Token == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: token.Label}
		}
		return e.Token, nil
	}
	return nil, &NotLoadedError{edge: "token"}
}

// ReceiveAddressOrErr returns the ReceiveAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentOrderEdges) ReceiveAddressOrErr() (*ReceiveAddress, error) {
	if e.loadedTypes[2] {
		if e.ReceiveAddress == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: receiveaddress.Label}
		}
		return e.ReceiveAddress, nil
	}
	return nil, &NotLoadedError{edge: "receive_address"}
}

// RecipientOrErr returns the Recipient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentOrderEdges) RecipientOrErr() (*PaymentOrderRecipient, error) {
	if e.loadedTypes[3] {
		if e.Recipient == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: paymentorderrecipient.Label}
		}
		return e.Recipient, nil
	}
	return nil, &NotLoadedError{edge: "recipient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PaymentOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case paymentorder.FieldAmount, paymentorder.FieldAmountPaid:
			values[i] = new(decimal.Decimal)
		case paymentorder.FieldTxHash, paymentorder.FieldReceiveAddressText, paymentorder.FieldLabel, paymentorder.FieldStatus:
			values[i] = new(sql.NullString)
		case paymentorder.FieldCreatedAt, paymentorder.FieldUpdatedAt, paymentorder.FieldLastUsed:
			values[i] = new(sql.NullTime)
		case paymentorder.FieldID:
			values[i] = new(uuid.UUID)
		case paymentorder.ForeignKeys[0]: // api_key_payment_orders
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case paymentorder.ForeignKeys[1]: // sender_profile_payment_orders
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case paymentorder.ForeignKeys[2]: // token_payment_orders
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PaymentOrder fields.
func (po *PaymentOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case paymentorder.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				po.ID = *value
			}
		case paymentorder.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case paymentorder.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case paymentorder.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				po.Amount = *value
			}
		case paymentorder.FieldAmountPaid:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount_paid", values[i])
			} else if value != nil {
				po.AmountPaid = *value
			}
		case paymentorder.FieldTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value.Valid {
				po.TxHash = value.String
			}
		case paymentorder.FieldReceiveAddressText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receive_address_text", values[i])
			} else if value.Valid {
				po.ReceiveAddressText = value.String
			}
		case paymentorder.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				po.Label = value.String
			}
		case paymentorder.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				po.Status = paymentorder.Status(value.String)
			}
		case paymentorder.FieldLastUsed:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_used", values[i])
			} else if value.Valid {
				po.LastUsed = value.Time
			}
		case paymentorder.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field api_key_payment_orders", values[i])
			} else if value.Valid {
				po.api_key_payment_orders = new(uuid.UUID)
				*po.api_key_payment_orders = *value.S.(*uuid.UUID)
			}
		case paymentorder.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field sender_profile_payment_orders", values[i])
			} else if value.Valid {
				po.sender_profile_payment_orders = new(uuid.UUID)
				*po.sender_profile_payment_orders = *value.S.(*uuid.UUID)
			}
		case paymentorder.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field token_payment_orders", value)
			} else if value.Valid {
				po.token_payment_orders = new(int)
				*po.token_payment_orders = int(value.Int64)
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PaymentOrder.
// This includes values selected through modifiers, order, etc.
func (po *PaymentOrder) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QuerySenderProfile queries the "sender_profile" edge of the PaymentOrder entity.
func (po *PaymentOrder) QuerySenderProfile() *SenderProfileQuery {
	return NewPaymentOrderClient(po.config).QuerySenderProfile(po)
}

// QueryToken queries the "token" edge of the PaymentOrder entity.
func (po *PaymentOrder) QueryToken() *TokenQuery {
	return NewPaymentOrderClient(po.config).QueryToken(po)
}

// QueryReceiveAddress queries the "receive_address" edge of the PaymentOrder entity.
func (po *PaymentOrder) QueryReceiveAddress() *ReceiveAddressQuery {
	return NewPaymentOrderClient(po.config).QueryReceiveAddress(po)
}

// QueryRecipient queries the "recipient" edge of the PaymentOrder entity.
func (po *PaymentOrder) QueryRecipient() *PaymentOrderRecipientQuery {
	return NewPaymentOrderClient(po.config).QueryRecipient(po)
}

// Update returns a builder for updating this PaymentOrder.
// Note that you need to call PaymentOrder.Unwrap() before calling this method if this PaymentOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *PaymentOrder) Update() *PaymentOrderUpdateOne {
	return NewPaymentOrderClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the PaymentOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *PaymentOrder) Unwrap() *PaymentOrder {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: PaymentOrder is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *PaymentOrder) String() string {
	var builder strings.Builder
	builder.WriteString("PaymentOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", po.Amount))
	builder.WriteString(", ")
	builder.WriteString("amount_paid=")
	builder.WriteString(fmt.Sprintf("%v", po.AmountPaid))
	builder.WriteString(", ")
	builder.WriteString("tx_hash=")
	builder.WriteString(po.TxHash)
	builder.WriteString(", ")
	builder.WriteString("receive_address_text=")
	builder.WriteString(po.ReceiveAddressText)
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(po.Label)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", po.Status))
	builder.WriteString(", ")
	builder.WriteString("last_used=")
	builder.WriteString(po.LastUsed.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PaymentOrders is a parsable slice of PaymentOrder.
type PaymentOrders []*PaymentOrder
