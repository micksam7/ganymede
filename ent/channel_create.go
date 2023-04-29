// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/vod"
)

// ChannelCreate is the builder for creating a Channel entity.
type ChannelCreate struct {
	config
	mutation *ChannelMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetExtID sets the "ext_id" field.
func (cc *ChannelCreate) SetExtID(s string) *ChannelCreate {
	cc.mutation.SetExtID(s)
	return cc
}

// SetNillableExtID sets the "ext_id" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableExtID(s *string) *ChannelCreate {
	if s != nil {
		cc.SetExtID(*s)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ChannelCreate) SetName(s string) *ChannelCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDisplayName sets the "display_name" field.
func (cc *ChannelCreate) SetDisplayName(s string) *ChannelCreate {
	cc.mutation.SetDisplayName(s)
	return cc
}

// SetImagePath sets the "image_path" field.
func (cc *ChannelCreate) SetImagePath(s string) *ChannelCreate {
	cc.mutation.SetImagePath(s)
	return cc
}

// SetRetention sets the "retention" field.
func (cc *ChannelCreate) SetRetention(b bool) *ChannelCreate {
	cc.mutation.SetRetention(b)
	return cc
}

// SetNillableRetention sets the "retention" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableRetention(b *bool) *ChannelCreate {
	if b != nil {
		cc.SetRetention(*b)
	}
	return cc
}

// SetRetentionDays sets the "retention_days" field.
func (cc *ChannelCreate) SetRetentionDays(i int64) *ChannelCreate {
	cc.mutation.SetRetentionDays(i)
	return cc
}

// SetNillableRetentionDays sets the "retention_days" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableRetentionDays(i *int64) *ChannelCreate {
	if i != nil {
		cc.SetRetentionDays(*i)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ChannelCreate) SetUpdatedAt(t time.Time) *ChannelCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableUpdatedAt(t *time.Time) *ChannelCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ChannelCreate) SetCreatedAt(t time.Time) *ChannelCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableCreatedAt(t *time.Time) *ChannelCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ChannelCreate) SetID(u uuid.UUID) *ChannelCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableID(u *uuid.UUID) *ChannelCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (cc *ChannelCreate) AddVodIDs(ids ...uuid.UUID) *ChannelCreate {
	cc.mutation.AddVodIDs(ids...)
	return cc
}

// AddVods adds the "vods" edges to the Vod entity.
func (cc *ChannelCreate) AddVods(v ...*Vod) *ChannelCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cc.AddVodIDs(ids...)
}

// AddLiveIDs adds the "live" edge to the Live entity by IDs.
func (cc *ChannelCreate) AddLiveIDs(ids ...uuid.UUID) *ChannelCreate {
	cc.mutation.AddLiveIDs(ids...)
	return cc
}

// AddLive adds the "live" edges to the Live entity.
func (cc *ChannelCreate) AddLive(l ...*Live) *ChannelCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cc.AddLiveIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cc *ChannelCreate) Mutation() *ChannelMutation {
	return cc.mutation
}

// Save creates the Channel in the database.
func (cc *ChannelCreate) Save(ctx context.Context) (*Channel, error) {
	cc.defaults()
	return withHooks[*Channel, ChannelMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChannelCreate) SaveX(ctx context.Context) *Channel {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChannelCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChannelCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChannelCreate) defaults() {
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := channel.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := channel.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := channel.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChannelCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Channel.name"`)}
	}
	if _, ok := cc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "Channel.display_name"`)}
	}
	if _, ok := cc.mutation.ImagePath(); !ok {
		return &ValidationError{Name: "image_path", err: errors.New(`ent: missing required field "Channel.image_path"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Channel.updated_at"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Channel.created_at"`)}
	}
	return nil
}

func (cc *ChannelCreate) sqlSave(ctx context.Context) (*Channel, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChannelCreate) createSpec() (*Channel, *sqlgraph.CreateSpec) {
	var (
		_node = &Channel{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(channel.Table, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.ExtID(); ok {
		_spec.SetField(channel.FieldExtID, field.TypeString, value)
		_node.ExtID = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(channel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.DisplayName(); ok {
		_spec.SetField(channel.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := cc.mutation.ImagePath(); ok {
		_spec.SetField(channel.FieldImagePath, field.TypeString, value)
		_node.ImagePath = value
	}
	if value, ok := cc.mutation.Retention(); ok {
		_spec.SetField(channel.FieldRetention, field.TypeBool, value)
		_node.Retention = value
	}
	if value, ok := cc.mutation.RetentionDays(); ok {
		_spec.SetField(channel.FieldRetentionDays, field.TypeInt64, value)
		_node.RetentionDays = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(channel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(channel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := cc.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.LiveIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(live.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Channel.Create().
//		SetExtID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChannelUpsert) {
//			SetExtID(v+v).
//		}).
//		Exec(ctx)
func (cc *ChannelCreate) OnConflict(opts ...sql.ConflictOption) *ChannelUpsertOne {
	cc.conflict = opts
	return &ChannelUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Channel.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ChannelCreate) OnConflictColumns(columns ...string) *ChannelUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ChannelUpsertOne{
		create: cc,
	}
}

type (
	// ChannelUpsertOne is the builder for "upsert"-ing
	//  one Channel node.
	ChannelUpsertOne struct {
		create *ChannelCreate
	}

	// ChannelUpsert is the "OnConflict" setter.
	ChannelUpsert struct {
		*sql.UpdateSet
	}
)

// SetExtID sets the "ext_id" field.
func (u *ChannelUpsert) SetExtID(v string) *ChannelUpsert {
	u.Set(channel.FieldExtID, v)
	return u
}

// UpdateExtID sets the "ext_id" field to the value that was provided on create.
func (u *ChannelUpsert) UpdateExtID() *ChannelUpsert {
	u.SetExcluded(channel.FieldExtID)
	return u
}

// ClearExtID clears the value of the "ext_id" field.
func (u *ChannelUpsert) ClearExtID() *ChannelUpsert {
	u.SetNull(channel.FieldExtID)
	return u
}

// SetName sets the "name" field.
func (u *ChannelUpsert) SetName(v string) *ChannelUpsert {
	u.Set(channel.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ChannelUpsert) UpdateName() *ChannelUpsert {
	u.SetExcluded(channel.FieldName)
	return u
}

// SetDisplayName sets the "display_name" field.
func (u *ChannelUpsert) SetDisplayName(v string) *ChannelUpsert {
	u.Set(channel.FieldDisplayName, v)
	return u
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *ChannelUpsert) UpdateDisplayName() *ChannelUpsert {
	u.SetExcluded(channel.FieldDisplayName)
	return u
}

// SetImagePath sets the "image_path" field.
func (u *ChannelUpsert) SetImagePath(v string) *ChannelUpsert {
	u.Set(channel.FieldImagePath, v)
	return u
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *ChannelUpsert) UpdateImagePath() *ChannelUpsert {
	u.SetExcluded(channel.FieldImagePath)
	return u
}

// SetRetention sets the "retention" field.
func (u *ChannelUpsert) SetRetention(v bool) *ChannelUpsert {
	u.Set(channel.FieldRetention, v)
	return u
}

// UpdateRetention sets the "retention" field to the value that was provided on create.
func (u *ChannelUpsert) UpdateRetention() *ChannelUpsert {
	u.SetExcluded(channel.FieldRetention)
	return u
}

// ClearRetention clears the value of the "retention" field.
func (u *ChannelUpsert) ClearRetention() *ChannelUpsert {
	u.SetNull(channel.FieldRetention)
	return u
}

// SetRetentionDays sets the "retention_days" field.
func (u *ChannelUpsert) SetRetentionDays(v int64) *ChannelUpsert {
	u.Set(channel.FieldRetentionDays, v)
	return u
}

// UpdateRetentionDays sets the "retention_days" field to the value that was provided on create.
func (u *ChannelUpsert) UpdateRetentionDays() *ChannelUpsert {
	u.SetExcluded(channel.FieldRetentionDays)
	return u
}

// AddRetentionDays adds v to the "retention_days" field.
func (u *ChannelUpsert) AddRetentionDays(v int64) *ChannelUpsert {
	u.Add(channel.FieldRetentionDays, v)
	return u
}

// ClearRetentionDays clears the value of the "retention_days" field.
func (u *ChannelUpsert) ClearRetentionDays() *ChannelUpsert {
	u.SetNull(channel.FieldRetentionDays)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelUpsert) SetUpdatedAt(v time.Time) *ChannelUpsert {
	u.Set(channel.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelUpsert) UpdateUpdatedAt() *ChannelUpsert {
	u.SetExcluded(channel.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Channel.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(channel.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChannelUpsertOne) UpdateNewValues() *ChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(channel.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(channel.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Channel.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ChannelUpsertOne) Ignore() *ChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChannelUpsertOne) DoNothing() *ChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChannelCreate.OnConflict
// documentation for more info.
func (u *ChannelUpsertOne) Update(set func(*ChannelUpsert)) *ChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChannelUpsert{UpdateSet: update})
	}))
	return u
}

// SetExtID sets the "ext_id" field.
func (u *ChannelUpsertOne) SetExtID(v string) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.SetExtID(v)
	})
}

// UpdateExtID sets the "ext_id" field to the value that was provided on create.
func (u *ChannelUpsertOne) UpdateExtID() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateExtID()
	})
}

// ClearExtID clears the value of the "ext_id" field.
func (u *ChannelUpsertOne) ClearExtID() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.ClearExtID()
	})
}

// SetName sets the "name" field.
func (u *ChannelUpsertOne) SetName(v string) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ChannelUpsertOne) UpdateName() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateName()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *ChannelUpsertOne) SetDisplayName(v string) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *ChannelUpsertOne) UpdateDisplayName() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateDisplayName()
	})
}

// SetImagePath sets the "image_path" field.
func (u *ChannelUpsertOne) SetImagePath(v string) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.SetImagePath(v)
	})
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *ChannelUpsertOne) UpdateImagePath() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateImagePath()
	})
}

// SetRetention sets the "retention" field.
func (u *ChannelUpsertOne) SetRetention(v bool) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.SetRetention(v)
	})
}

// UpdateRetention sets the "retention" field to the value that was provided on create.
func (u *ChannelUpsertOne) UpdateRetention() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateRetention()
	})
}

// ClearRetention clears the value of the "retention" field.
func (u *ChannelUpsertOne) ClearRetention() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.ClearRetention()
	})
}

// SetRetentionDays sets the "retention_days" field.
func (u *ChannelUpsertOne) SetRetentionDays(v int64) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.SetRetentionDays(v)
	})
}

// AddRetentionDays adds v to the "retention_days" field.
func (u *ChannelUpsertOne) AddRetentionDays(v int64) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.AddRetentionDays(v)
	})
}

// UpdateRetentionDays sets the "retention_days" field to the value that was provided on create.
func (u *ChannelUpsertOne) UpdateRetentionDays() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateRetentionDays()
	})
}

// ClearRetentionDays clears the value of the "retention_days" field.
func (u *ChannelUpsertOne) ClearRetentionDays() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.ClearRetentionDays()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelUpsertOne) SetUpdatedAt(v time.Time) *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelUpsertOne) UpdateUpdatedAt() *ChannelUpsertOne {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ChannelUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChannelCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChannelUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ChannelUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ChannelUpsertOne.ID is not supported by MySQL driver. Use ChannelUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ChannelUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ChannelCreateBulk is the builder for creating many Channel entities in bulk.
type ChannelCreateBulk struct {
	config
	builders []*ChannelCreate
	conflict []sql.ConflictOption
}

// Save creates the Channel entities in the database.
func (ccb *ChannelCreateBulk) Save(ctx context.Context) ([]*Channel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Channel, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChannelCreateBulk) SaveX(ctx context.Context) []*Channel {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChannelCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Channel.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChannelUpsert) {
//			SetExtID(v+v).
//		}).
//		Exec(ctx)
func (ccb *ChannelCreateBulk) OnConflict(opts ...sql.ConflictOption) *ChannelUpsertBulk {
	ccb.conflict = opts
	return &ChannelUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Channel.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ChannelCreateBulk) OnConflictColumns(columns ...string) *ChannelUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ChannelUpsertBulk{
		create: ccb,
	}
}

// ChannelUpsertBulk is the builder for "upsert"-ing
// a bulk of Channel nodes.
type ChannelUpsertBulk struct {
	create *ChannelCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Channel.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(channel.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChannelUpsertBulk) UpdateNewValues() *ChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(channel.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(channel.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Channel.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ChannelUpsertBulk) Ignore() *ChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChannelUpsertBulk) DoNothing() *ChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChannelCreateBulk.OnConflict
// documentation for more info.
func (u *ChannelUpsertBulk) Update(set func(*ChannelUpsert)) *ChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChannelUpsert{UpdateSet: update})
	}))
	return u
}

// SetExtID sets the "ext_id" field.
func (u *ChannelUpsertBulk) SetExtID(v string) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.SetExtID(v)
	})
}

// UpdateExtID sets the "ext_id" field to the value that was provided on create.
func (u *ChannelUpsertBulk) UpdateExtID() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateExtID()
	})
}

// ClearExtID clears the value of the "ext_id" field.
func (u *ChannelUpsertBulk) ClearExtID() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.ClearExtID()
	})
}

// SetName sets the "name" field.
func (u *ChannelUpsertBulk) SetName(v string) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ChannelUpsertBulk) UpdateName() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateName()
	})
}

// SetDisplayName sets the "display_name" field.
func (u *ChannelUpsertBulk) SetDisplayName(v string) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.SetDisplayName(v)
	})
}

// UpdateDisplayName sets the "display_name" field to the value that was provided on create.
func (u *ChannelUpsertBulk) UpdateDisplayName() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateDisplayName()
	})
}

// SetImagePath sets the "image_path" field.
func (u *ChannelUpsertBulk) SetImagePath(v string) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.SetImagePath(v)
	})
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *ChannelUpsertBulk) UpdateImagePath() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateImagePath()
	})
}

// SetRetention sets the "retention" field.
func (u *ChannelUpsertBulk) SetRetention(v bool) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.SetRetention(v)
	})
}

// UpdateRetention sets the "retention" field to the value that was provided on create.
func (u *ChannelUpsertBulk) UpdateRetention() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateRetention()
	})
}

// ClearRetention clears the value of the "retention" field.
func (u *ChannelUpsertBulk) ClearRetention() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.ClearRetention()
	})
}

// SetRetentionDays sets the "retention_days" field.
func (u *ChannelUpsertBulk) SetRetentionDays(v int64) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.SetRetentionDays(v)
	})
}

// AddRetentionDays adds v to the "retention_days" field.
func (u *ChannelUpsertBulk) AddRetentionDays(v int64) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.AddRetentionDays(v)
	})
}

// UpdateRetentionDays sets the "retention_days" field to the value that was provided on create.
func (u *ChannelUpsertBulk) UpdateRetentionDays() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateRetentionDays()
	})
}

// ClearRetentionDays clears the value of the "retention_days" field.
func (u *ChannelUpsertBulk) ClearRetentionDays() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.ClearRetentionDays()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChannelUpsertBulk) SetUpdatedAt(v time.Time) *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChannelUpsertBulk) UpdateUpdatedAt() *ChannelUpsertBulk {
	return u.Update(func(s *ChannelUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ChannelUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ChannelCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChannelCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChannelUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
