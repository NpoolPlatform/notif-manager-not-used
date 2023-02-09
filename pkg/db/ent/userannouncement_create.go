// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/userannouncement"
	"github.com/google/uuid"
)

// UserAnnouncementCreate is the builder for creating a UserAnnouncement entity.
type UserAnnouncementCreate struct {
	config
	mutation *UserAnnouncementMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (uac *UserAnnouncementCreate) SetCreatedAt(u uint32) *UserAnnouncementCreate {
	uac.mutation.SetCreatedAt(u)
	return uac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uac *UserAnnouncementCreate) SetNillableCreatedAt(u *uint32) *UserAnnouncementCreate {
	if u != nil {
		uac.SetCreatedAt(*u)
	}
	return uac
}

// SetUpdatedAt sets the "updated_at" field.
func (uac *UserAnnouncementCreate) SetUpdatedAt(u uint32) *UserAnnouncementCreate {
	uac.mutation.SetUpdatedAt(u)
	return uac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uac *UserAnnouncementCreate) SetNillableUpdatedAt(u *uint32) *UserAnnouncementCreate {
	if u != nil {
		uac.SetUpdatedAt(*u)
	}
	return uac
}

// SetDeletedAt sets the "deleted_at" field.
func (uac *UserAnnouncementCreate) SetDeletedAt(u uint32) *UserAnnouncementCreate {
	uac.mutation.SetDeletedAt(u)
	return uac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uac *UserAnnouncementCreate) SetNillableDeletedAt(u *uint32) *UserAnnouncementCreate {
	if u != nil {
		uac.SetDeletedAt(*u)
	}
	return uac
}

// SetAppID sets the "app_id" field.
func (uac *UserAnnouncementCreate) SetAppID(u uuid.UUID) *UserAnnouncementCreate {
	uac.mutation.SetAppID(u)
	return uac
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (uac *UserAnnouncementCreate) SetNillableAppID(u *uuid.UUID) *UserAnnouncementCreate {
	if u != nil {
		uac.SetAppID(*u)
	}
	return uac
}

// SetUserID sets the "user_id" field.
func (uac *UserAnnouncementCreate) SetUserID(u uuid.UUID) *UserAnnouncementCreate {
	uac.mutation.SetUserID(u)
	return uac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uac *UserAnnouncementCreate) SetNillableUserID(u *uuid.UUID) *UserAnnouncementCreate {
	if u != nil {
		uac.SetUserID(*u)
	}
	return uac
}

// SetAnnouncementID sets the "announcement_id" field.
func (uac *UserAnnouncementCreate) SetAnnouncementID(u uuid.UUID) *UserAnnouncementCreate {
	uac.mutation.SetAnnouncementID(u)
	return uac
}

// SetNillableAnnouncementID sets the "announcement_id" field if the given value is not nil.
func (uac *UserAnnouncementCreate) SetNillableAnnouncementID(u *uuid.UUID) *UserAnnouncementCreate {
	if u != nil {
		uac.SetAnnouncementID(*u)
	}
	return uac
}

// SetID sets the "id" field.
func (uac *UserAnnouncementCreate) SetID(u uuid.UUID) *UserAnnouncementCreate {
	uac.mutation.SetID(u)
	return uac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uac *UserAnnouncementCreate) SetNillableID(u *uuid.UUID) *UserAnnouncementCreate {
	if u != nil {
		uac.SetID(*u)
	}
	return uac
}

// Mutation returns the UserAnnouncementMutation object of the builder.
func (uac *UserAnnouncementCreate) Mutation() *UserAnnouncementMutation {
	return uac.mutation
}

// Save creates the UserAnnouncement in the database.
func (uac *UserAnnouncementCreate) Save(ctx context.Context) (*UserAnnouncement, error) {
	var (
		err  error
		node *UserAnnouncement
	)
	if err := uac.defaults(); err != nil {
		return nil, err
	}
	if len(uac.hooks) == 0 {
		if err = uac.check(); err != nil {
			return nil, err
		}
		node, err = uac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAnnouncementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uac.check(); err != nil {
				return nil, err
			}
			uac.mutation = mutation
			if node, err = uac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uac.hooks) - 1; i >= 0; i-- {
			if uac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserAnnouncement)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserAnnouncementMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uac *UserAnnouncementCreate) SaveX(ctx context.Context) *UserAnnouncement {
	v, err := uac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uac *UserAnnouncementCreate) Exec(ctx context.Context) error {
	_, err := uac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uac *UserAnnouncementCreate) ExecX(ctx context.Context) {
	if err := uac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uac *UserAnnouncementCreate) defaults() error {
	if _, ok := uac.mutation.CreatedAt(); !ok {
		if userannouncement.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized userannouncement.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := userannouncement.DefaultCreatedAt()
		uac.mutation.SetCreatedAt(v)
	}
	if _, ok := uac.mutation.UpdatedAt(); !ok {
		if userannouncement.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userannouncement.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userannouncement.DefaultUpdatedAt()
		uac.mutation.SetUpdatedAt(v)
	}
	if _, ok := uac.mutation.DeletedAt(); !ok {
		if userannouncement.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized userannouncement.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := userannouncement.DefaultDeletedAt()
		uac.mutation.SetDeletedAt(v)
	}
	if _, ok := uac.mutation.AppID(); !ok {
		if userannouncement.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized userannouncement.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := userannouncement.DefaultAppID()
		uac.mutation.SetAppID(v)
	}
	if _, ok := uac.mutation.UserID(); !ok {
		if userannouncement.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized userannouncement.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := userannouncement.DefaultUserID()
		uac.mutation.SetUserID(v)
	}
	if _, ok := uac.mutation.AnnouncementID(); !ok {
		if userannouncement.DefaultAnnouncementID == nil {
			return fmt.Errorf("ent: uninitialized userannouncement.DefaultAnnouncementID (forgotten import ent/runtime?)")
		}
		v := userannouncement.DefaultAnnouncementID()
		uac.mutation.SetAnnouncementID(v)
	}
	if _, ok := uac.mutation.ID(); !ok {
		if userannouncement.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized userannouncement.DefaultID (forgotten import ent/runtime?)")
		}
		v := userannouncement.DefaultID()
		uac.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uac *UserAnnouncementCreate) check() error {
	if _, ok := uac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserAnnouncement.created_at"`)}
	}
	if _, ok := uac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserAnnouncement.updated_at"`)}
	}
	if _, ok := uac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "UserAnnouncement.deleted_at"`)}
	}
	return nil
}

func (uac *UserAnnouncementCreate) sqlSave(ctx context.Context) (*UserAnnouncement, error) {
	_node, _spec := uac.createSpec()
	if err := sqlgraph.CreateNode(ctx, uac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (uac *UserAnnouncementCreate) createSpec() (*UserAnnouncement, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAnnouncement{config: uac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userannouncement.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userannouncement.FieldID,
			},
		}
	)
	_spec.OnConflict = uac.conflict
	if id, ok := uac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userannouncement.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userannouncement.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userannouncement.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := uac.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userannouncement.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := uac.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userannouncement.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uac.mutation.AnnouncementID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userannouncement.FieldAnnouncementID,
		})
		_node.AnnouncementID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserAnnouncement.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserAnnouncementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (uac *UserAnnouncementCreate) OnConflict(opts ...sql.ConflictOption) *UserAnnouncementUpsertOne {
	uac.conflict = opts
	return &UserAnnouncementUpsertOne{
		create: uac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserAnnouncement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uac *UserAnnouncementCreate) OnConflictColumns(columns ...string) *UserAnnouncementUpsertOne {
	uac.conflict = append(uac.conflict, sql.ConflictColumns(columns...))
	return &UserAnnouncementUpsertOne{
		create: uac,
	}
}

type (
	// UserAnnouncementUpsertOne is the builder for "upsert"-ing
	//  one UserAnnouncement node.
	UserAnnouncementUpsertOne struct {
		create *UserAnnouncementCreate
	}

	// UserAnnouncementUpsert is the "OnConflict" setter.
	UserAnnouncementUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UserAnnouncementUpsert) SetCreatedAt(v uint32) *UserAnnouncementUpsert {
	u.Set(userannouncement.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsert) UpdateCreatedAt() *UserAnnouncementUpsert {
	u.SetExcluded(userannouncement.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserAnnouncementUpsert) AddCreatedAt(v uint32) *UserAnnouncementUpsert {
	u.Add(userannouncement.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserAnnouncementUpsert) SetUpdatedAt(v uint32) *UserAnnouncementUpsert {
	u.Set(userannouncement.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsert) UpdateUpdatedAt() *UserAnnouncementUpsert {
	u.SetExcluded(userannouncement.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserAnnouncementUpsert) AddUpdatedAt(v uint32) *UserAnnouncementUpsert {
	u.Add(userannouncement.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserAnnouncementUpsert) SetDeletedAt(v uint32) *UserAnnouncementUpsert {
	u.Set(userannouncement.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsert) UpdateDeletedAt() *UserAnnouncementUpsert {
	u.SetExcluded(userannouncement.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserAnnouncementUpsert) AddDeletedAt(v uint32) *UserAnnouncementUpsert {
	u.Add(userannouncement.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserAnnouncementUpsert) SetAppID(v uuid.UUID) *UserAnnouncementUpsert {
	u.Set(userannouncement.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsert) UpdateAppID() *UserAnnouncementUpsert {
	u.SetExcluded(userannouncement.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserAnnouncementUpsert) ClearAppID() *UserAnnouncementUpsert {
	u.SetNull(userannouncement.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserAnnouncementUpsert) SetUserID(v uuid.UUID) *UserAnnouncementUpsert {
	u.Set(userannouncement.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsert) UpdateUserID() *UserAnnouncementUpsert {
	u.SetExcluded(userannouncement.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserAnnouncementUpsert) ClearUserID() *UserAnnouncementUpsert {
	u.SetNull(userannouncement.FieldUserID)
	return u
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *UserAnnouncementUpsert) SetAnnouncementID(v uuid.UUID) *UserAnnouncementUpsert {
	u.Set(userannouncement.FieldAnnouncementID, v)
	return u
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsert) UpdateAnnouncementID() *UserAnnouncementUpsert {
	u.SetExcluded(userannouncement.FieldAnnouncementID)
	return u
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (u *UserAnnouncementUpsert) ClearAnnouncementID() *UserAnnouncementUpsert {
	u.SetNull(userannouncement.FieldAnnouncementID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserAnnouncement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userannouncement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserAnnouncementUpsertOne) UpdateNewValues() *UserAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userannouncement.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserAnnouncement.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserAnnouncementUpsertOne) Ignore() *UserAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserAnnouncementUpsertOne) DoNothing() *UserAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserAnnouncementCreate.OnConflict
// documentation for more info.
func (u *UserAnnouncementUpsertOne) Update(set func(*UserAnnouncementUpsert)) *UserAnnouncementUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserAnnouncementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserAnnouncementUpsertOne) SetCreatedAt(v uint32) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserAnnouncementUpsertOne) AddCreatedAt(v uint32) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsertOne) UpdateCreatedAt() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserAnnouncementUpsertOne) SetUpdatedAt(v uint32) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserAnnouncementUpsertOne) AddUpdatedAt(v uint32) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsertOne) UpdateUpdatedAt() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserAnnouncementUpsertOne) SetDeletedAt(v uint32) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserAnnouncementUpsertOne) AddDeletedAt(v uint32) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsertOne) UpdateDeletedAt() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserAnnouncementUpsertOne) SetAppID(v uuid.UUID) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsertOne) UpdateAppID() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserAnnouncementUpsertOne) ClearAppID() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserAnnouncementUpsertOne) SetUserID(v uuid.UUID) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsertOne) UpdateUserID() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserAnnouncementUpsertOne) ClearUserID() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.ClearUserID()
	})
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *UserAnnouncementUpsertOne) SetAnnouncementID(v uuid.UUID) *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetAnnouncementID(v)
	})
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsertOne) UpdateAnnouncementID() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateAnnouncementID()
	})
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (u *UserAnnouncementUpsertOne) ClearAnnouncementID() *UserAnnouncementUpsertOne {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.ClearAnnouncementID()
	})
}

// Exec executes the query.
func (u *UserAnnouncementUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserAnnouncementCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserAnnouncementUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserAnnouncementUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserAnnouncementUpsertOne.ID is not supported by MySQL driver. Use UserAnnouncementUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserAnnouncementUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserAnnouncementCreateBulk is the builder for creating many UserAnnouncement entities in bulk.
type UserAnnouncementCreateBulk struct {
	config
	builders []*UserAnnouncementCreate
	conflict []sql.ConflictOption
}

// Save creates the UserAnnouncement entities in the database.
func (uacb *UserAnnouncementCreateBulk) Save(ctx context.Context) ([]*UserAnnouncement, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uacb.builders))
	nodes := make([]*UserAnnouncement, len(uacb.builders))
	mutators := make([]Mutator, len(uacb.builders))
	for i := range uacb.builders {
		func(i int, root context.Context) {
			builder := uacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAnnouncementMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uacb *UserAnnouncementCreateBulk) SaveX(ctx context.Context) []*UserAnnouncement {
	v, err := uacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uacb *UserAnnouncementCreateBulk) Exec(ctx context.Context) error {
	_, err := uacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uacb *UserAnnouncementCreateBulk) ExecX(ctx context.Context) {
	if err := uacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserAnnouncement.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserAnnouncementUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (uacb *UserAnnouncementCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserAnnouncementUpsertBulk {
	uacb.conflict = opts
	return &UserAnnouncementUpsertBulk{
		create: uacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserAnnouncement.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uacb *UserAnnouncementCreateBulk) OnConflictColumns(columns ...string) *UserAnnouncementUpsertBulk {
	uacb.conflict = append(uacb.conflict, sql.ConflictColumns(columns...))
	return &UserAnnouncementUpsertBulk{
		create: uacb,
	}
}

// UserAnnouncementUpsertBulk is the builder for "upsert"-ing
// a bulk of UserAnnouncement nodes.
type UserAnnouncementUpsertBulk struct {
	create *UserAnnouncementCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserAnnouncement.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userannouncement.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserAnnouncementUpsertBulk) UpdateNewValues() *UserAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userannouncement.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserAnnouncement.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserAnnouncementUpsertBulk) Ignore() *UserAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserAnnouncementUpsertBulk) DoNothing() *UserAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserAnnouncementCreateBulk.OnConflict
// documentation for more info.
func (u *UserAnnouncementUpsertBulk) Update(set func(*UserAnnouncementUpsert)) *UserAnnouncementUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserAnnouncementUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserAnnouncementUpsertBulk) SetCreatedAt(v uint32) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *UserAnnouncementUpsertBulk) AddCreatedAt(v uint32) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsertBulk) UpdateCreatedAt() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserAnnouncementUpsertBulk) SetUpdatedAt(v uint32) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserAnnouncementUpsertBulk) AddUpdatedAt(v uint32) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsertBulk) UpdateUpdatedAt() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserAnnouncementUpsertBulk) SetDeletedAt(v uint32) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserAnnouncementUpsertBulk) AddDeletedAt(v uint32) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserAnnouncementUpsertBulk) UpdateDeletedAt() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserAnnouncementUpsertBulk) SetAppID(v uuid.UUID) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsertBulk) UpdateAppID() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *UserAnnouncementUpsertBulk) ClearAppID() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.ClearAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserAnnouncementUpsertBulk) SetUserID(v uuid.UUID) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsertBulk) UpdateUserID() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserAnnouncementUpsertBulk) ClearUserID() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.ClearUserID()
	})
}

// SetAnnouncementID sets the "announcement_id" field.
func (u *UserAnnouncementUpsertBulk) SetAnnouncementID(v uuid.UUID) *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.SetAnnouncementID(v)
	})
}

// UpdateAnnouncementID sets the "announcement_id" field to the value that was provided on create.
func (u *UserAnnouncementUpsertBulk) UpdateAnnouncementID() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.UpdateAnnouncementID()
	})
}

// ClearAnnouncementID clears the value of the "announcement_id" field.
func (u *UserAnnouncementUpsertBulk) ClearAnnouncementID() *UserAnnouncementUpsertBulk {
	return u.Update(func(s *UserAnnouncementUpsert) {
		s.ClearAnnouncementID()
	})
}

// Exec executes the query.
func (u *UserAnnouncementUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserAnnouncementCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserAnnouncementCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserAnnouncementUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}