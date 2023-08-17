// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ofdl/ofdl/ent/media"
	"github.com/ofdl/ofdl/ent/post"
	"github.com/ofdl/ofdl/ent/predicate"
)

// MediaUpdate is the builder for updating Media entities.
type MediaUpdate struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where appends a list predicates to the MediaUpdate builder.
func (mu *MediaUpdate) Where(ps ...predicate.Media) *MediaUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetPostID sets the "post_id" field.
func (mu *MediaUpdate) SetPostID(i int) *MediaUpdate {
	mu.mutation.SetPostID(i)
	return mu
}

// SetType sets the "type" field.
func (mu *MediaUpdate) SetType(s string) *MediaUpdate {
	mu.mutation.SetType(s)
	return mu
}

// SetFull sets the "full" field.
func (mu *MediaUpdate) SetFull(s string) *MediaUpdate {
	mu.mutation.SetFull(s)
	return mu
}

// SetNillableFull sets the "full" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableFull(s *string) *MediaUpdate {
	if s != nil {
		mu.SetFull(*s)
	}
	return mu
}

// ClearFull clears the value of the "full" field.
func (mu *MediaUpdate) ClearFull() *MediaUpdate {
	mu.mutation.ClearFull()
	return mu
}

// SetPostedAt sets the "posted_at" field.
func (mu *MediaUpdate) SetPostedAt(s string) *MediaUpdate {
	mu.mutation.SetPostedAt(s)
	return mu
}

// SetDownloadedAt sets the "downloaded_at" field.
func (mu *MediaUpdate) SetDownloadedAt(t time.Time) *MediaUpdate {
	mu.mutation.SetDownloadedAt(t)
	return mu
}

// SetNillableDownloadedAt sets the "downloaded_at" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableDownloadedAt(t *time.Time) *MediaUpdate {
	if t != nil {
		mu.SetDownloadedAt(*t)
	}
	return mu
}

// ClearDownloadedAt clears the value of the "downloaded_at" field.
func (mu *MediaUpdate) ClearDownloadedAt() *MediaUpdate {
	mu.mutation.ClearDownloadedAt()
	return mu
}

// SetStashID sets the "stash_id" field.
func (mu *MediaUpdate) SetStashID(s string) *MediaUpdate {
	mu.mutation.SetStashID(s)
	return mu
}

// SetNillableStashID sets the "stash_id" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableStashID(s *string) *MediaUpdate {
	if s != nil {
		mu.SetStashID(*s)
	}
	return mu
}

// ClearStashID clears the value of the "stash_id" field.
func (mu *MediaUpdate) ClearStashID() *MediaUpdate {
	mu.mutation.ClearStashID()
	return mu
}

// SetOrganizedAt sets the "organized_at" field.
func (mu *MediaUpdate) SetOrganizedAt(t time.Time) *MediaUpdate {
	mu.mutation.SetOrganizedAt(t)
	return mu
}

// SetNillableOrganizedAt sets the "organized_at" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableOrganizedAt(t *time.Time) *MediaUpdate {
	if t != nil {
		mu.SetOrganizedAt(*t)
	}
	return mu
}

// ClearOrganizedAt clears the value of the "organized_at" field.
func (mu *MediaUpdate) ClearOrganizedAt() *MediaUpdate {
	mu.mutation.ClearOrganizedAt()
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MediaUpdate) SetCreatedAt(t time.Time) *MediaUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableCreatedAt(t *time.Time) *MediaUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MediaUpdate) SetUpdatedAt(t time.Time) *MediaUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetPost sets the "post" edge to the Post entity.
func (mu *MediaUpdate) SetPost(p *Post) *MediaUpdate {
	return mu.SetPostID(p.ID)
}

// Mutation returns the MediaMutation object of the builder.
func (mu *MediaUpdate) Mutation() *MediaMutation {
	return mu.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (mu *MediaUpdate) ClearPost() *MediaUpdate {
	mu.mutation.ClearPost()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MediaUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MediaUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MediaUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MediaUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MediaUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := media.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MediaUpdate) check() error {
	if _, ok := mu.mutation.PostID(); mu.mutation.PostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Media.post"`)
	}
	return nil
}

func (mu *MediaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.SetField(media.FieldType, field.TypeString, value)
	}
	if value, ok := mu.mutation.Full(); ok {
		_spec.SetField(media.FieldFull, field.TypeString, value)
	}
	if mu.mutation.FullCleared() {
		_spec.ClearField(media.FieldFull, field.TypeString)
	}
	if value, ok := mu.mutation.PostedAt(); ok {
		_spec.SetField(media.FieldPostedAt, field.TypeString, value)
	}
	if value, ok := mu.mutation.DownloadedAt(); ok {
		_spec.SetField(media.FieldDownloadedAt, field.TypeTime, value)
	}
	if mu.mutation.DownloadedAtCleared() {
		_spec.ClearField(media.FieldDownloadedAt, field.TypeTime)
	}
	if value, ok := mu.mutation.StashID(); ok {
		_spec.SetField(media.FieldStashID, field.TypeString, value)
	}
	if mu.mutation.StashIDCleared() {
		_spec.ClearField(media.FieldStashID, field.TypeString)
	}
	if value, ok := mu.mutation.OrganizedAt(); ok {
		_spec.SetField(media.FieldOrganizedAt, field.TypeTime, value)
	}
	if mu.mutation.OrganizedAtCleared() {
		_spec.ClearField(media.FieldOrganizedAt, field.TypeTime)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(media.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(media.FieldUpdatedAt, field.TypeTime, value)
	}
	if mu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MediaUpdateOne is the builder for updating a single Media entity.
type MediaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MediaMutation
}

// SetPostID sets the "post_id" field.
func (muo *MediaUpdateOne) SetPostID(i int) *MediaUpdateOne {
	muo.mutation.SetPostID(i)
	return muo
}

// SetType sets the "type" field.
func (muo *MediaUpdateOne) SetType(s string) *MediaUpdateOne {
	muo.mutation.SetType(s)
	return muo
}

// SetFull sets the "full" field.
func (muo *MediaUpdateOne) SetFull(s string) *MediaUpdateOne {
	muo.mutation.SetFull(s)
	return muo
}

// SetNillableFull sets the "full" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableFull(s *string) *MediaUpdateOne {
	if s != nil {
		muo.SetFull(*s)
	}
	return muo
}

// ClearFull clears the value of the "full" field.
func (muo *MediaUpdateOne) ClearFull() *MediaUpdateOne {
	muo.mutation.ClearFull()
	return muo
}

// SetPostedAt sets the "posted_at" field.
func (muo *MediaUpdateOne) SetPostedAt(s string) *MediaUpdateOne {
	muo.mutation.SetPostedAt(s)
	return muo
}

// SetDownloadedAt sets the "downloaded_at" field.
func (muo *MediaUpdateOne) SetDownloadedAt(t time.Time) *MediaUpdateOne {
	muo.mutation.SetDownloadedAt(t)
	return muo
}

// SetNillableDownloadedAt sets the "downloaded_at" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableDownloadedAt(t *time.Time) *MediaUpdateOne {
	if t != nil {
		muo.SetDownloadedAt(*t)
	}
	return muo
}

// ClearDownloadedAt clears the value of the "downloaded_at" field.
func (muo *MediaUpdateOne) ClearDownloadedAt() *MediaUpdateOne {
	muo.mutation.ClearDownloadedAt()
	return muo
}

// SetStashID sets the "stash_id" field.
func (muo *MediaUpdateOne) SetStashID(s string) *MediaUpdateOne {
	muo.mutation.SetStashID(s)
	return muo
}

// SetNillableStashID sets the "stash_id" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableStashID(s *string) *MediaUpdateOne {
	if s != nil {
		muo.SetStashID(*s)
	}
	return muo
}

// ClearStashID clears the value of the "stash_id" field.
func (muo *MediaUpdateOne) ClearStashID() *MediaUpdateOne {
	muo.mutation.ClearStashID()
	return muo
}

// SetOrganizedAt sets the "organized_at" field.
func (muo *MediaUpdateOne) SetOrganizedAt(t time.Time) *MediaUpdateOne {
	muo.mutation.SetOrganizedAt(t)
	return muo
}

// SetNillableOrganizedAt sets the "organized_at" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableOrganizedAt(t *time.Time) *MediaUpdateOne {
	if t != nil {
		muo.SetOrganizedAt(*t)
	}
	return muo
}

// ClearOrganizedAt clears the value of the "organized_at" field.
func (muo *MediaUpdateOne) ClearOrganizedAt() *MediaUpdateOne {
	muo.mutation.ClearOrganizedAt()
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MediaUpdateOne) SetCreatedAt(t time.Time) *MediaUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableCreatedAt(t *time.Time) *MediaUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MediaUpdateOne) SetUpdatedAt(t time.Time) *MediaUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetPost sets the "post" edge to the Post entity.
func (muo *MediaUpdateOne) SetPost(p *Post) *MediaUpdateOne {
	return muo.SetPostID(p.ID)
}

// Mutation returns the MediaMutation object of the builder.
func (muo *MediaUpdateOne) Mutation() *MediaMutation {
	return muo.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (muo *MediaUpdateOne) ClearPost() *MediaUpdateOne {
	muo.mutation.ClearPost()
	return muo
}

// Where appends a list predicates to the MediaUpdate builder.
func (muo *MediaUpdateOne) Where(ps ...predicate.Media) *MediaUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MediaUpdateOne) Select(field string, fields ...string) *MediaUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Media entity.
func (muo *MediaUpdateOne) Save(ctx context.Context) (*Media, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MediaUpdateOne) SaveX(ctx context.Context) *Media {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MediaUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MediaUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MediaUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := media.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MediaUpdateOne) check() error {
	if _, ok := muo.mutation.PostID(); muo.mutation.PostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Media.post"`)
	}
	return nil
}

func (muo *MediaUpdateOne) sqlSave(ctx context.Context) (_node *Media, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Media.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, media.FieldID)
		for _, f := range fields {
			if !media.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != media.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.SetField(media.FieldType, field.TypeString, value)
	}
	if value, ok := muo.mutation.Full(); ok {
		_spec.SetField(media.FieldFull, field.TypeString, value)
	}
	if muo.mutation.FullCleared() {
		_spec.ClearField(media.FieldFull, field.TypeString)
	}
	if value, ok := muo.mutation.PostedAt(); ok {
		_spec.SetField(media.FieldPostedAt, field.TypeString, value)
	}
	if value, ok := muo.mutation.DownloadedAt(); ok {
		_spec.SetField(media.FieldDownloadedAt, field.TypeTime, value)
	}
	if muo.mutation.DownloadedAtCleared() {
		_spec.ClearField(media.FieldDownloadedAt, field.TypeTime)
	}
	if value, ok := muo.mutation.StashID(); ok {
		_spec.SetField(media.FieldStashID, field.TypeString, value)
	}
	if muo.mutation.StashIDCleared() {
		_spec.ClearField(media.FieldStashID, field.TypeString)
	}
	if value, ok := muo.mutation.OrganizedAt(); ok {
		_spec.SetField(media.FieldOrganizedAt, field.TypeTime, value)
	}
	if muo.mutation.OrganizedAtCleared() {
		_spec.ClearField(media.FieldOrganizedAt, field.TypeTime)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(media.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(media.FieldUpdatedAt, field.TypeTime, value)
	}
	if muo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Media{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}