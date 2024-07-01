// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/paycrest/protocol/ent/senderordertoken"
	"github.com/paycrest/protocol/ent/senderprofile"
	"github.com/paycrest/protocol/ent/token"
	"github.com/shopspring/decimal"
)

// SenderOrderToken is the model entity for the SenderOrderToken schema.
type SenderOrderToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FeePerTokenUnit holds the value of the "fee_per_token_unit" field.
	FeePerTokenUnit decimal.Decimal `json:"fee_per_token_unit,omitempty"`
	// FeeAddress holds the value of the "fee_address" field.
	FeeAddress string `json:"fee_address,omitempty"`
	// RefundAddress holds the value of the "refund_address" field.
	RefundAddress string `json:"refund_address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SenderOrderTokenQuery when eager-loading is set.
	Edges                       SenderOrderTokenEdges `json:"edges"`
	sender_profile_order_tokens *uuid.UUID
	token_sender_settings       *int
	selectValues                sql.SelectValues
}

// SenderOrderTokenEdges holds the relations/edges for other nodes in the graph.
type SenderOrderTokenEdges struct {
	// Sender holds the value of the sender edge.
	Sender *SenderProfile `json:"sender,omitempty"`
	// Token holds the value of the token edge.
	Token *Token `json:"token,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SenderOrderTokenEdges) SenderOrErr() (*SenderProfile, error) {
	if e.loadedTypes[0] {
		if e.Sender == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: senderprofile.Label}
		}
		return e.Sender, nil
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// TokenOrErr returns the Token value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SenderOrderTokenEdges) TokenOrErr() (*Token, error) {
	if e.loadedTypes[1] {
		if e.Token == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: token.Label}
		}
		return e.Token, nil
	}
	return nil, &NotLoadedError{edge: "token"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SenderOrderToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case senderordertoken.FieldFeePerTokenUnit:
			values[i] = new(decimal.Decimal)
		case senderordertoken.FieldID:
			values[i] = new(sql.NullInt64)
		case senderordertoken.FieldFeeAddress, senderordertoken.FieldRefundAddress:
			values[i] = new(sql.NullString)
		case senderordertoken.FieldCreatedAt, senderordertoken.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case senderordertoken.ForeignKeys[0]: // sender_profile_order_tokens
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case senderordertoken.ForeignKeys[1]: // token_sender_settings
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SenderOrderToken fields.
func (sot *SenderOrderToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case senderordertoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sot.ID = int(value.Int64)
		case senderordertoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sot.CreatedAt = value.Time
			}
		case senderordertoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sot.UpdatedAt = value.Time
			}
		case senderordertoken.FieldFeePerTokenUnit:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field fee_per_token_unit", values[i])
			} else if value != nil {
				sot.FeePerTokenUnit = *value
			}
		case senderordertoken.FieldFeeAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fee_address", values[i])
			} else if value.Valid {
				sot.FeeAddress = value.String
			}
		case senderordertoken.FieldRefundAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refund_address", values[i])
			} else if value.Valid {
				sot.RefundAddress = value.String
			}
		case senderordertoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field sender_profile_order_tokens", values[i])
			} else if value.Valid {
				sot.sender_profile_order_tokens = new(uuid.UUID)
				*sot.sender_profile_order_tokens = *value.S.(*uuid.UUID)
			}
		case senderordertoken.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field token_sender_settings", value)
			} else if value.Valid {
				sot.token_sender_settings = new(int)
				*sot.token_sender_settings = int(value.Int64)
			}
		default:
			sot.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SenderOrderToken.
// This includes values selected through modifiers, order, etc.
func (sot *SenderOrderToken) Value(name string) (ent.Value, error) {
	return sot.selectValues.Get(name)
}

// QuerySender queries the "sender" edge of the SenderOrderToken entity.
func (sot *SenderOrderToken) QuerySender() *SenderProfileQuery {
	return NewSenderOrderTokenClient(sot.config).QuerySender(sot)
}

// QueryToken queries the "token" edge of the SenderOrderToken entity.
func (sot *SenderOrderToken) QueryToken() *TokenQuery {
	return NewSenderOrderTokenClient(sot.config).QueryToken(sot)
}

// Update returns a builder for updating this SenderOrderToken.
// Note that you need to call SenderOrderToken.Unwrap() before calling this method if this SenderOrderToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (sot *SenderOrderToken) Update() *SenderOrderTokenUpdateOne {
	return NewSenderOrderTokenClient(sot.config).UpdateOne(sot)
}

// Unwrap unwraps the SenderOrderToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sot *SenderOrderToken) Unwrap() *SenderOrderToken {
	_tx, ok := sot.config.driver.(*txDriver)
	if !ok {
		panic("ent: SenderOrderToken is not a transactional entity")
	}
	sot.config.driver = _tx.drv
	return sot
}

// String implements the fmt.Stringer.
func (sot *SenderOrderToken) String() string {
	var builder strings.Builder
	builder.WriteString("SenderOrderToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sot.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sot.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sot.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("fee_per_token_unit=")
	builder.WriteString(fmt.Sprintf("%v", sot.FeePerTokenUnit))
	builder.WriteString(", ")
	builder.WriteString("fee_address=")
	builder.WriteString(sot.FeeAddress)
	builder.WriteString(", ")
	builder.WriteString("refund_address=")
	builder.WriteString(sot.RefundAddress)
	builder.WriteByte(')')
	return builder.String()
}

// SenderOrderTokens is a parsable slice of SenderOrderToken.
type SenderOrderTokens []*SenderOrderToken
