// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/paycrest/paycrest-protocol/ent/provideravailability"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
)

// ProviderAvailability is the model entity for the ProviderAvailability schema.
type ProviderAvailability struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Cadence holds the value of the "cadence" field.
	Cadence provideravailability.Cadence `json:"cadence,omitempty"`
	// StartTime holds the value of the "start_time" field.
	StartTime time.Time `json:"start_time,omitempty"`
	// EndTime holds the value of the "end_time" field.
	EndTime time.Time `json:"end_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProviderAvailabilityQuery when eager-loading is set.
	Edges                         ProviderAvailabilityEdges `json:"edges"`
	provider_profile_availability *string
	selectValues                  sql.SelectValues
}

// ProviderAvailabilityEdges holds the relations/edges for other nodes in the graph.
type ProviderAvailabilityEdges struct {
	// Provider holds the value of the provider edge.
	Provider *ProviderProfile `json:"provider,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProviderOrErr returns the Provider value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProviderAvailabilityEdges) ProviderOrErr() (*ProviderProfile, error) {
	if e.loadedTypes[0] {
		if e.Provider == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: providerprofile.Label}
		}
		return e.Provider, nil
	}
	return nil, &NotLoadedError{edge: "provider"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProviderAvailability) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case provideravailability.FieldID:
			values[i] = new(sql.NullInt64)
		case provideravailability.FieldCadence:
			values[i] = new(sql.NullString)
		case provideravailability.FieldStartTime, provideravailability.FieldEndTime:
			values[i] = new(sql.NullTime)
		case provideravailability.ForeignKeys[0]: // provider_profile_availability
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProviderAvailability fields.
func (pa *ProviderAvailability) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provideravailability.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case provideravailability.FieldCadence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cadence", values[i])
			} else if value.Valid {
				pa.Cadence = provideravailability.Cadence(value.String)
			}
		case provideravailability.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				pa.StartTime = value.Time
			}
		case provideravailability.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				pa.EndTime = value.Time
			}
		case provideravailability.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_profile_availability", values[i])
			} else if value.Valid {
				pa.provider_profile_availability = new(string)
				*pa.provider_profile_availability = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProviderAvailability.
// This includes values selected through modifiers, order, etc.
func (pa *ProviderAvailability) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// QueryProvider queries the "provider" edge of the ProviderAvailability entity.
func (pa *ProviderAvailability) QueryProvider() *ProviderProfileQuery {
	return NewProviderAvailabilityClient(pa.config).QueryProvider(pa)
}

// Update returns a builder for updating this ProviderAvailability.
// Note that you need to call ProviderAvailability.Unwrap() before calling this method if this ProviderAvailability
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *ProviderAvailability) Update() *ProviderAvailabilityUpdateOne {
	return NewProviderAvailabilityClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the ProviderAvailability entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *ProviderAvailability) Unwrap() *ProviderAvailability {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProviderAvailability is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *ProviderAvailability) String() string {
	var builder strings.Builder
	builder.WriteString("ProviderAvailability(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("cadence=")
	builder.WriteString(fmt.Sprintf("%v", pa.Cadence))
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(pa.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(pa.EndTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProviderAvailabilities is a parsable slice of ProviderAvailability.
type ProviderAvailabilities []*ProviderAvailability
