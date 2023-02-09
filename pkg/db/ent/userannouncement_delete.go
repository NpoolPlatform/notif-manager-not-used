// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/userannouncement"
)

// UserAnnouncementDelete is the builder for deleting a UserAnnouncement entity.
type UserAnnouncementDelete struct {
	config
	hooks    []Hook
	mutation *UserAnnouncementMutation
}

// Where appends a list predicates to the UserAnnouncementDelete builder.
func (uad *UserAnnouncementDelete) Where(ps ...predicate.UserAnnouncement) *UserAnnouncementDelete {
	uad.mutation.Where(ps...)
	return uad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uad *UserAnnouncementDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uad.hooks) == 0 {
		affected, err = uad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uad.mutation = mutation
			affected, err = uad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uad.hooks) - 1; i >= 0; i-- {
			if uad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uad *UserAnnouncementDelete) ExecX(ctx context.Context) int {
	n, err := uad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uad *UserAnnouncementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userannouncement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userannouncement.FieldID,
			},
		},
	}
	if ps := uad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserAnnouncementDeleteOne is the builder for deleting a single UserAnnouncement entity.
type UserAnnouncementDeleteOne struct {
	uad *UserAnnouncementDelete
}

// Exec executes the deletion query.
func (uado *UserAnnouncementDeleteOne) Exec(ctx context.Context) error {
	n, err := uado.uad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userannouncement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uado *UserAnnouncementDeleteOne) ExecX(ctx context.Context) {
	uado.uad.ExecX(ctx)
}