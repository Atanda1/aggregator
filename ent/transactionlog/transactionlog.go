// Code generated by ent, DO NOT EDIT.

package transactionlog

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the transactionlog type in the database.
	Label = "transaction_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGatewayID holds the string denoting the gateway_id field in the database.
	FieldGatewayID = "gateway_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldNetwork holds the string denoting the network field in the database.
	FieldNetwork = "network"
	// FieldTxHash holds the string denoting the tx_hash field in the database.
	FieldTxHash = "tx_hash"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the transactionlog in the database.
	Table = "transaction_logs"
)

// Columns holds all SQL columns for transactionlog fields.
var Columns = []string{
	FieldID,
	FieldGatewayID,
	FieldStatus,
	FieldNetwork,
	FieldTxHash,
	FieldMetadata,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transaction_logs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"lock_payment_order_transactions",
	"payment_order_transactions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusOrderInitiated is the default value of the Status enum.
const DefaultStatus = StatusOrderInitiated

// Status values.
const (
	StatusOrderInitiated  Status = "order_initiated"
	StatusCryptoDeposited Status = "crypto_deposited"
	StatusOrderCreated    Status = "order_created"
	StatusOrderSettled    Status = "order_settled"
	StatusOrderRefunded   Status = "order_refunded"
	StatusOrderReverted   Status = "order_reverted"
	StatusGasPrefunded    Status = "gas_prefunded"
	StatusGatewayApproved Status = "gateway_approved"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOrderInitiated, StatusCryptoDeposited, StatusOrderCreated, StatusOrderSettled, StatusOrderRefunded, StatusOrderReverted, StatusGasPrefunded, StatusGatewayApproved:
		return nil
	default:
		return fmt.Errorf("transactionlog: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the TransactionLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGatewayID orders the results by the gateway_id field.
func ByGatewayID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGatewayID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByNetwork orders the results by the network field.
func ByNetwork(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNetwork, opts...).ToFunc()
}

// ByTxHash orders the results by the tx_hash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
