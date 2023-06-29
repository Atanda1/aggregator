// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
	"github.com/paycrest/paycrest-protocol/ent/providerprofile"
)

// ProviderProfileDelete is the builder for deleting a ProviderProfile entity.
type ProviderProfileDelete struct {
	config
	hooks    []Hook
	mutation *ProviderProfileMutation
}

// Where appends a list predicates to the ProviderProfileDelete builder.
func (ppd *ProviderProfileDelete) Where(ps ...predicate.ProviderProfile) *ProviderProfileDelete {
	ppd.mutation.Where(ps...)
	return ppd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ppd *ProviderProfileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ppd.sqlExec, ppd.mutation, ppd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ppd *ProviderProfileDelete) ExecX(ctx context.Context) int {
	n, err := ppd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ppd *ProviderProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(providerprofile.Table, sqlgraph.NewFieldSpec(providerprofile.FieldID, field.TypeString))
	if ps := ppd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ppd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ppd.mutation.done = true
	return affected, err
}

// ProviderProfileDeleteOne is the builder for deleting a single ProviderProfile entity.
type ProviderProfileDeleteOne struct {
	ppd *ProviderProfileDelete
}

// Where appends a list predicates to the ProviderProfileDelete builder.
func (ppdo *ProviderProfileDeleteOne) Where(ps ...predicate.ProviderProfile) *ProviderProfileDeleteOne {
	ppdo.ppd.mutation.Where(ps...)
	return ppdo
}

// Exec executes the deletion query.
func (ppdo *ProviderProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := ppdo.ppd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{providerprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ppdo *ProviderProfileDeleteOne) ExecX(ctx context.Context) {
	if err := ppdo.Exec(ctx); err != nil {
		panic(err)
	}
}
