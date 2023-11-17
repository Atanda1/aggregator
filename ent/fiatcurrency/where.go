// Code generated by ent, DO NOT EDIT.

package fiatcurrency

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldUpdatedAt, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldCode, v))
}

// ShortName applies equality check predicate on the "short_name" field. It's identical to ShortNameEQ.
func ShortName(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldShortName, v))
}

// Decimals applies equality check predicate on the "decimals" field. It's identical to DecimalsEQ.
func Decimals(v int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldDecimals, v))
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldSymbol, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldName, v))
}

// MarketRate applies equality check predicate on the "market_rate" field. It's identical to MarketRateEQ.
func MarketRate(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldMarketRate, v))
}

// IsEnabled applies equality check predicate on the "is_enabled" field. It's identical to IsEnabledEQ.
func IsEnabled(v bool) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldIsEnabled, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldUpdatedAt, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContainsFold(FieldCode, v))
}

// ShortNameEQ applies the EQ predicate on the "short_name" field.
func ShortNameEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldShortName, v))
}

// ShortNameNEQ applies the NEQ predicate on the "short_name" field.
func ShortNameNEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldShortName, v))
}

// ShortNameIn applies the In predicate on the "short_name" field.
func ShortNameIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldShortName, vs...))
}

// ShortNameNotIn applies the NotIn predicate on the "short_name" field.
func ShortNameNotIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldShortName, vs...))
}

// ShortNameGT applies the GT predicate on the "short_name" field.
func ShortNameGT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldShortName, v))
}

// ShortNameGTE applies the GTE predicate on the "short_name" field.
func ShortNameGTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldShortName, v))
}

// ShortNameLT applies the LT predicate on the "short_name" field.
func ShortNameLT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldShortName, v))
}

// ShortNameLTE applies the LTE predicate on the "short_name" field.
func ShortNameLTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldShortName, v))
}

// ShortNameContains applies the Contains predicate on the "short_name" field.
func ShortNameContains(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContains(FieldShortName, v))
}

// ShortNameHasPrefix applies the HasPrefix predicate on the "short_name" field.
func ShortNameHasPrefix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasPrefix(FieldShortName, v))
}

// ShortNameHasSuffix applies the HasSuffix predicate on the "short_name" field.
func ShortNameHasSuffix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasSuffix(FieldShortName, v))
}

// ShortNameEqualFold applies the EqualFold predicate on the "short_name" field.
func ShortNameEqualFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEqualFold(FieldShortName, v))
}

// ShortNameContainsFold applies the ContainsFold predicate on the "short_name" field.
func ShortNameContainsFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContainsFold(FieldShortName, v))
}

// DecimalsEQ applies the EQ predicate on the "decimals" field.
func DecimalsEQ(v int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldDecimals, v))
}

// DecimalsNEQ applies the NEQ predicate on the "decimals" field.
func DecimalsNEQ(v int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldDecimals, v))
}

// DecimalsIn applies the In predicate on the "decimals" field.
func DecimalsIn(vs ...int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldDecimals, vs...))
}

// DecimalsNotIn applies the NotIn predicate on the "decimals" field.
func DecimalsNotIn(vs ...int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldDecimals, vs...))
}

// DecimalsGT applies the GT predicate on the "decimals" field.
func DecimalsGT(v int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldDecimals, v))
}

// DecimalsGTE applies the GTE predicate on the "decimals" field.
func DecimalsGTE(v int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldDecimals, v))
}

// DecimalsLT applies the LT predicate on the "decimals" field.
func DecimalsLT(v int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldDecimals, v))
}

// DecimalsLTE applies the LTE predicate on the "decimals" field.
func DecimalsLTE(v int) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldDecimals, v))
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldSymbol, v))
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldSymbol, v))
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldSymbol, vs...))
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldSymbol, vs...))
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldSymbol, v))
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldSymbol, v))
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldSymbol, v))
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldSymbol, v))
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContains(FieldSymbol, v))
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasPrefix(FieldSymbol, v))
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasSuffix(FieldSymbol, v))
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEqualFold(FieldSymbol, v))
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContainsFold(FieldSymbol, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldContainsFold(FieldName, v))
}

// MarketRateEQ applies the EQ predicate on the "market_rate" field.
func MarketRateEQ(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldMarketRate, v))
}

// MarketRateNEQ applies the NEQ predicate on the "market_rate" field.
func MarketRateNEQ(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldMarketRate, v))
}

// MarketRateIn applies the In predicate on the "market_rate" field.
func MarketRateIn(vs ...decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldIn(FieldMarketRate, vs...))
}

// MarketRateNotIn applies the NotIn predicate on the "market_rate" field.
func MarketRateNotIn(vs ...decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNotIn(FieldMarketRate, vs...))
}

// MarketRateGT applies the GT predicate on the "market_rate" field.
func MarketRateGT(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGT(FieldMarketRate, v))
}

// MarketRateGTE applies the GTE predicate on the "market_rate" field.
func MarketRateGTE(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldGTE(FieldMarketRate, v))
}

// MarketRateLT applies the LT predicate on the "market_rate" field.
func MarketRateLT(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLT(FieldMarketRate, v))
}

// MarketRateLTE applies the LTE predicate on the "market_rate" field.
func MarketRateLTE(v decimal.Decimal) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldLTE(FieldMarketRate, v))
}

// IsEnabledEQ applies the EQ predicate on the "is_enabled" field.
func IsEnabledEQ(v bool) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldEQ(FieldIsEnabled, v))
}

// IsEnabledNEQ applies the NEQ predicate on the "is_enabled" field.
func IsEnabledNEQ(v bool) predicate.FiatCurrency {
	return predicate.FiatCurrency(sql.FieldNEQ(FieldIsEnabled, v))
}

// HasProviders applies the HasEdge predicate on the "providers" edge.
func HasProviders() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvidersTable, ProvidersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProvidersWith applies the HasEdge predicate on the "providers" edge with a given conditions (other predicates).
func HasProvidersWith(preds ...predicate.ProviderProfile) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		step := newProvidersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProvisionBuckets applies the HasEdge predicate on the "provision_buckets" edge.
func HasProvisionBuckets() predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvisionBucketsTable, ProvisionBucketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProvisionBucketsWith applies the HasEdge predicate on the "provision_buckets" edge with a given conditions (other predicates).
func HasProvisionBucketsWith(preds ...predicate.ProvisionBucket) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		step := newProvisionBucketsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FiatCurrency) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FiatCurrency) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FiatCurrency) predicate.FiatCurrency {
	return predicate.FiatCurrency(func(s *sql.Selector) {
		p(s.Not())
	})
}
