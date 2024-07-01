// Code generated by ent, DO NOT EDIT.

package senderordertoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/paycrest/protocol/ent/predicate"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// FeePerTokenUnit applies equality check predicate on the "fee_per_token_unit" field. It's identical to FeePerTokenUnitEQ.
func FeePerTokenUnit(v decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldFeePerTokenUnit, v))
}

// FeeAddress applies equality check predicate on the "fee_address" field. It's identical to FeeAddressEQ.
func FeeAddress(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldFeeAddress, v))
}

// RefundAddress applies equality check predicate on the "refund_address" field. It's identical to RefundAddressEQ.
func RefundAddress(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldRefundAddress, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// FeePerTokenUnitEQ applies the EQ predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitEQ(v decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldFeePerTokenUnit, v))
}

// FeePerTokenUnitNEQ applies the NEQ predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitNEQ(v decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNEQ(FieldFeePerTokenUnit, v))
}

// FeePerTokenUnitIn applies the In predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitIn(vs ...decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldIn(FieldFeePerTokenUnit, vs...))
}

// FeePerTokenUnitNotIn applies the NotIn predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitNotIn(vs ...decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNotIn(FieldFeePerTokenUnit, vs...))
}

// FeePerTokenUnitGT applies the GT predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitGT(v decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGT(FieldFeePerTokenUnit, v))
}

// FeePerTokenUnitGTE applies the GTE predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitGTE(v decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGTE(FieldFeePerTokenUnit, v))
}

// FeePerTokenUnitLT applies the LT predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitLT(v decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLT(FieldFeePerTokenUnit, v))
}

// FeePerTokenUnitLTE applies the LTE predicate on the "fee_per_token_unit" field.
func FeePerTokenUnitLTE(v decimal.Decimal) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLTE(FieldFeePerTokenUnit, v))
}

// FeeAddressEQ applies the EQ predicate on the "fee_address" field.
func FeeAddressEQ(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldFeeAddress, v))
}

// FeeAddressNEQ applies the NEQ predicate on the "fee_address" field.
func FeeAddressNEQ(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNEQ(FieldFeeAddress, v))
}

// FeeAddressIn applies the In predicate on the "fee_address" field.
func FeeAddressIn(vs ...string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldIn(FieldFeeAddress, vs...))
}

// FeeAddressNotIn applies the NotIn predicate on the "fee_address" field.
func FeeAddressNotIn(vs ...string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNotIn(FieldFeeAddress, vs...))
}

// FeeAddressGT applies the GT predicate on the "fee_address" field.
func FeeAddressGT(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGT(FieldFeeAddress, v))
}

// FeeAddressGTE applies the GTE predicate on the "fee_address" field.
func FeeAddressGTE(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGTE(FieldFeeAddress, v))
}

// FeeAddressLT applies the LT predicate on the "fee_address" field.
func FeeAddressLT(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLT(FieldFeeAddress, v))
}

// FeeAddressLTE applies the LTE predicate on the "fee_address" field.
func FeeAddressLTE(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLTE(FieldFeeAddress, v))
}

// FeeAddressContains applies the Contains predicate on the "fee_address" field.
func FeeAddressContains(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldContains(FieldFeeAddress, v))
}

// FeeAddressHasPrefix applies the HasPrefix predicate on the "fee_address" field.
func FeeAddressHasPrefix(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldHasPrefix(FieldFeeAddress, v))
}

// FeeAddressHasSuffix applies the HasSuffix predicate on the "fee_address" field.
func FeeAddressHasSuffix(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldHasSuffix(FieldFeeAddress, v))
}

// FeeAddressEqualFold applies the EqualFold predicate on the "fee_address" field.
func FeeAddressEqualFold(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEqualFold(FieldFeeAddress, v))
}

// FeeAddressContainsFold applies the ContainsFold predicate on the "fee_address" field.
func FeeAddressContainsFold(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldContainsFold(FieldFeeAddress, v))
}

// RefundAddressEQ applies the EQ predicate on the "refund_address" field.
func RefundAddressEQ(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEQ(FieldRefundAddress, v))
}

// RefundAddressNEQ applies the NEQ predicate on the "refund_address" field.
func RefundAddressNEQ(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNEQ(FieldRefundAddress, v))
}

// RefundAddressIn applies the In predicate on the "refund_address" field.
func RefundAddressIn(vs ...string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldIn(FieldRefundAddress, vs...))
}

// RefundAddressNotIn applies the NotIn predicate on the "refund_address" field.
func RefundAddressNotIn(vs ...string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldNotIn(FieldRefundAddress, vs...))
}

// RefundAddressGT applies the GT predicate on the "refund_address" field.
func RefundAddressGT(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGT(FieldRefundAddress, v))
}

// RefundAddressGTE applies the GTE predicate on the "refund_address" field.
func RefundAddressGTE(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldGTE(FieldRefundAddress, v))
}

// RefundAddressLT applies the LT predicate on the "refund_address" field.
func RefundAddressLT(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLT(FieldRefundAddress, v))
}

// RefundAddressLTE applies the LTE predicate on the "refund_address" field.
func RefundAddressLTE(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldLTE(FieldRefundAddress, v))
}

// RefundAddressContains applies the Contains predicate on the "refund_address" field.
func RefundAddressContains(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldContains(FieldRefundAddress, v))
}

// RefundAddressHasPrefix applies the HasPrefix predicate on the "refund_address" field.
func RefundAddressHasPrefix(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldHasPrefix(FieldRefundAddress, v))
}

// RefundAddressHasSuffix applies the HasSuffix predicate on the "refund_address" field.
func RefundAddressHasSuffix(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldHasSuffix(FieldRefundAddress, v))
}

// RefundAddressEqualFold applies the EqualFold predicate on the "refund_address" field.
func RefundAddressEqualFold(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldEqualFold(FieldRefundAddress, v))
}

// RefundAddressContainsFold applies the ContainsFold predicate on the "refund_address" field.
func RefundAddressContainsFold(v string) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(sql.FieldContainsFold(FieldRefundAddress, v))
}

// HasSender applies the HasEdge predicate on the "sender" edge.
func HasSender() predicate.SenderOrderToken {
	return predicate.SenderOrderToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SenderTable, SenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSenderWith applies the HasEdge predicate on the "sender" edge with a given conditions (other predicates).
func HasSenderWith(preds ...predicate.SenderProfile) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(func(s *sql.Selector) {
		step := newSenderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasToken applies the HasEdge predicate on the "token" edge.
func HasToken() predicate.SenderOrderToken {
	return predicate.SenderOrderToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TokenTable, TokenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTokenWith applies the HasEdge predicate on the "token" edge with a given conditions (other predicates).
func HasTokenWith(preds ...predicate.Token) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(func(s *sql.Selector) {
		step := newTokenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SenderOrderToken) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SenderOrderToken) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(func(s *sql.Selector) {
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
func Not(p predicate.SenderOrderToken) predicate.SenderOrderToken {
	return predicate.SenderOrderToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
