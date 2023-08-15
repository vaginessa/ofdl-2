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
	"github.com/ofdl/ofdl/ent/message"
	"github.com/ofdl/ofdl/ent/post"
	"github.com/ofdl/ofdl/ent/predicate"
	"github.com/ofdl/ofdl/ent/subscription"
)

// SubscriptionUpdate is the builder for updating Subscription entities.
type SubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriptionMutation
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (su *SubscriptionUpdate) Where(ps ...predicate.Subscription) *SubscriptionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetAvatar sets the "avatar" field.
func (su *SubscriptionUpdate) SetAvatar(s string) *SubscriptionUpdate {
	su.mutation.SetAvatar(s)
	return su
}

// SetHeader sets the "header" field.
func (su *SubscriptionUpdate) SetHeader(s string) *SubscriptionUpdate {
	su.mutation.SetHeader(s)
	return su
}

// SetName sets the "name" field.
func (su *SubscriptionUpdate) SetName(s string) *SubscriptionUpdate {
	su.mutation.SetName(s)
	return su
}

// SetUsername sets the "username" field.
func (su *SubscriptionUpdate) SetUsername(s string) *SubscriptionUpdate {
	su.mutation.SetUsername(s)
	return su
}

// SetHeadMarker sets the "head_marker" field.
func (su *SubscriptionUpdate) SetHeadMarker(s string) *SubscriptionUpdate {
	su.mutation.SetHeadMarker(s)
	return su
}

// SetNillableHeadMarker sets the "head_marker" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableHeadMarker(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetHeadMarker(*s)
	}
	return su
}

// ClearHeadMarker clears the value of the "head_marker" field.
func (su *SubscriptionUpdate) ClearHeadMarker() *SubscriptionUpdate {
	su.mutation.ClearHeadMarker()
	return su
}

// SetStashID sets the "stash_id" field.
func (su *SubscriptionUpdate) SetStashID(s string) *SubscriptionUpdate {
	su.mutation.SetStashID(s)
	return su
}

// SetNillableStashID sets the "stash_id" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableStashID(s *string) *SubscriptionUpdate {
	if s != nil {
		su.SetStashID(*s)
	}
	return su
}

// ClearStashID clears the value of the "stash_id" field.
func (su *SubscriptionUpdate) ClearStashID() *SubscriptionUpdate {
	su.mutation.ClearStashID()
	return su
}

// SetOrganizedAt sets the "organized_at" field.
func (su *SubscriptionUpdate) SetOrganizedAt(t time.Time) *SubscriptionUpdate {
	su.mutation.SetOrganizedAt(t)
	return su
}

// SetNillableOrganizedAt sets the "organized_at" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableOrganizedAt(t *time.Time) *SubscriptionUpdate {
	if t != nil {
		su.SetOrganizedAt(*t)
	}
	return su
}

// ClearOrganizedAt clears the value of the "organized_at" field.
func (su *SubscriptionUpdate) ClearOrganizedAt() *SubscriptionUpdate {
	su.mutation.ClearOrganizedAt()
	return su
}

// SetEnabled sets the "enabled" field.
func (su *SubscriptionUpdate) SetEnabled(b bool) *SubscriptionUpdate {
	su.mutation.SetEnabled(b)
	return su
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableEnabled(b *bool) *SubscriptionUpdate {
	if b != nil {
		su.SetEnabled(*b)
	}
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SubscriptionUpdate) SetCreatedAt(t time.Time) *SubscriptionUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SubscriptionUpdate) SetNillableCreatedAt(t *time.Time) *SubscriptionUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SubscriptionUpdate) SetUpdatedAt(t time.Time) *SubscriptionUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (su *SubscriptionUpdate) AddPostIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.AddPostIDs(ids...)
	return su
}

// AddPosts adds the "posts" edges to the Post entity.
func (su *SubscriptionUpdate) AddPosts(p ...*Post) *SubscriptionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddPostIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (su *SubscriptionUpdate) AddMessageIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.AddMessageIDs(ids...)
	return su
}

// AddMessages adds the "messages" edges to the Message entity.
func (su *SubscriptionUpdate) AddMessages(m ...*Message) *SubscriptionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return su.AddMessageIDs(ids...)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (su *SubscriptionUpdate) Mutation() *SubscriptionMutation {
	return su.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (su *SubscriptionUpdate) ClearPosts() *SubscriptionUpdate {
	su.mutation.ClearPosts()
	return su
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (su *SubscriptionUpdate) RemovePostIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.RemovePostIDs(ids...)
	return su
}

// RemovePosts removes "posts" edges to Post entities.
func (su *SubscriptionUpdate) RemovePosts(p ...*Post) *SubscriptionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemovePostIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (su *SubscriptionUpdate) ClearMessages() *SubscriptionUpdate {
	su.mutation.ClearMessages()
	return su
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (su *SubscriptionUpdate) RemoveMessageIDs(ids ...int) *SubscriptionUpdate {
	su.mutation.RemoveMessageIDs(ids...)
	return su
}

// RemoveMessages removes "messages" edges to Message entities.
func (su *SubscriptionUpdate) RemoveMessages(m ...*Message) *SubscriptionUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return su.RemoveMessageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscriptionUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubscriptionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SubscriptionUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := subscription.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *SubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Avatar(); ok {
		_spec.SetField(subscription.FieldAvatar, field.TypeString, value)
	}
	if value, ok := su.mutation.Header(); ok {
		_spec.SetField(subscription.FieldHeader, field.TypeString, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(subscription.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Username(); ok {
		_spec.SetField(subscription.FieldUsername, field.TypeString, value)
	}
	if value, ok := su.mutation.HeadMarker(); ok {
		_spec.SetField(subscription.FieldHeadMarker, field.TypeString, value)
	}
	if su.mutation.HeadMarkerCleared() {
		_spec.ClearField(subscription.FieldHeadMarker, field.TypeString)
	}
	if value, ok := su.mutation.StashID(); ok {
		_spec.SetField(subscription.FieldStashID, field.TypeString, value)
	}
	if su.mutation.StashIDCleared() {
		_spec.ClearField(subscription.FieldStashID, field.TypeString)
	}
	if value, ok := su.mutation.OrganizedAt(); ok {
		_spec.SetField(subscription.FieldOrganizedAt, field.TypeTime, value)
	}
	if su.mutation.OrganizedAtCleared() {
		_spec.ClearField(subscription.FieldOrganizedAt, field.TypeTime)
	}
	if value, ok := su.mutation.Enabled(); ok {
		_spec.SetField(subscription.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(subscription.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(subscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.PostsTable,
			Columns: []string{subscription.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedPostsIDs(); len(nodes) > 0 && !su.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.PostsTable,
			Columns: []string{subscription.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.PostsTable,
			Columns: []string{subscription.PostsColumn},
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
	if su.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.MessagesTable,
			Columns: []string{subscription.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !su.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.MessagesTable,
			Columns: []string{subscription.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.MessagesTable,
			Columns: []string{subscription.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SubscriptionUpdateOne is the builder for updating a single Subscription entity.
type SubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriptionMutation
}

// SetAvatar sets the "avatar" field.
func (suo *SubscriptionUpdateOne) SetAvatar(s string) *SubscriptionUpdateOne {
	suo.mutation.SetAvatar(s)
	return suo
}

// SetHeader sets the "header" field.
func (suo *SubscriptionUpdateOne) SetHeader(s string) *SubscriptionUpdateOne {
	suo.mutation.SetHeader(s)
	return suo
}

// SetName sets the "name" field.
func (suo *SubscriptionUpdateOne) SetName(s string) *SubscriptionUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetUsername sets the "username" field.
func (suo *SubscriptionUpdateOne) SetUsername(s string) *SubscriptionUpdateOne {
	suo.mutation.SetUsername(s)
	return suo
}

// SetHeadMarker sets the "head_marker" field.
func (suo *SubscriptionUpdateOne) SetHeadMarker(s string) *SubscriptionUpdateOne {
	suo.mutation.SetHeadMarker(s)
	return suo
}

// SetNillableHeadMarker sets the "head_marker" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableHeadMarker(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetHeadMarker(*s)
	}
	return suo
}

// ClearHeadMarker clears the value of the "head_marker" field.
func (suo *SubscriptionUpdateOne) ClearHeadMarker() *SubscriptionUpdateOne {
	suo.mutation.ClearHeadMarker()
	return suo
}

// SetStashID sets the "stash_id" field.
func (suo *SubscriptionUpdateOne) SetStashID(s string) *SubscriptionUpdateOne {
	suo.mutation.SetStashID(s)
	return suo
}

// SetNillableStashID sets the "stash_id" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableStashID(s *string) *SubscriptionUpdateOne {
	if s != nil {
		suo.SetStashID(*s)
	}
	return suo
}

// ClearStashID clears the value of the "stash_id" field.
func (suo *SubscriptionUpdateOne) ClearStashID() *SubscriptionUpdateOne {
	suo.mutation.ClearStashID()
	return suo
}

// SetOrganizedAt sets the "organized_at" field.
func (suo *SubscriptionUpdateOne) SetOrganizedAt(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetOrganizedAt(t)
	return suo
}

// SetNillableOrganizedAt sets the "organized_at" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableOrganizedAt(t *time.Time) *SubscriptionUpdateOne {
	if t != nil {
		suo.SetOrganizedAt(*t)
	}
	return suo
}

// ClearOrganizedAt clears the value of the "organized_at" field.
func (suo *SubscriptionUpdateOne) ClearOrganizedAt() *SubscriptionUpdateOne {
	suo.mutation.ClearOrganizedAt()
	return suo
}

// SetEnabled sets the "enabled" field.
func (suo *SubscriptionUpdateOne) SetEnabled(b bool) *SubscriptionUpdateOne {
	suo.mutation.SetEnabled(b)
	return suo
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableEnabled(b *bool) *SubscriptionUpdateOne {
	if b != nil {
		suo.SetEnabled(*b)
	}
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SubscriptionUpdateOne) SetCreatedAt(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SubscriptionUpdateOne) SetNillableCreatedAt(t *time.Time) *SubscriptionUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SubscriptionUpdateOne) SetUpdatedAt(t time.Time) *SubscriptionUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (suo *SubscriptionUpdateOne) AddPostIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.AddPostIDs(ids...)
	return suo
}

// AddPosts adds the "posts" edges to the Post entity.
func (suo *SubscriptionUpdateOne) AddPosts(p ...*Post) *SubscriptionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddPostIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (suo *SubscriptionUpdateOne) AddMessageIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.AddMessageIDs(ids...)
	return suo
}

// AddMessages adds the "messages" edges to the Message entity.
func (suo *SubscriptionUpdateOne) AddMessages(m ...*Message) *SubscriptionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return suo.AddMessageIDs(ids...)
}

// Mutation returns the SubscriptionMutation object of the builder.
func (suo *SubscriptionUpdateOne) Mutation() *SubscriptionMutation {
	return suo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (suo *SubscriptionUpdateOne) ClearPosts() *SubscriptionUpdateOne {
	suo.mutation.ClearPosts()
	return suo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (suo *SubscriptionUpdateOne) RemovePostIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.RemovePostIDs(ids...)
	return suo
}

// RemovePosts removes "posts" edges to Post entities.
func (suo *SubscriptionUpdateOne) RemovePosts(p ...*Post) *SubscriptionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemovePostIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (suo *SubscriptionUpdateOne) ClearMessages() *SubscriptionUpdateOne {
	suo.mutation.ClearMessages()
	return suo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (suo *SubscriptionUpdateOne) RemoveMessageIDs(ids ...int) *SubscriptionUpdateOne {
	suo.mutation.RemoveMessageIDs(ids...)
	return suo
}

// RemoveMessages removes "messages" edges to Message entities.
func (suo *SubscriptionUpdateOne) RemoveMessages(m ...*Message) *SubscriptionUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return suo.RemoveMessageIDs(ids...)
}

// Where appends a list predicates to the SubscriptionUpdate builder.
func (suo *SubscriptionUpdateOne) Where(ps ...predicate.Subscription) *SubscriptionUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscriptionUpdateOne) Select(field string, fields ...string) *SubscriptionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscription entity.
func (suo *SubscriptionUpdateOne) Save(ctx context.Context) (*Subscription, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) SaveX(ctx context.Context) *Subscription {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SubscriptionUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := subscription.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *SubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *Subscription, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscription.Table, subscription.Columns, sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscription.FieldID)
		for _, f := range fields {
			if !subscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Avatar(); ok {
		_spec.SetField(subscription.FieldAvatar, field.TypeString, value)
	}
	if value, ok := suo.mutation.Header(); ok {
		_spec.SetField(subscription.FieldHeader, field.TypeString, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(subscription.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Username(); ok {
		_spec.SetField(subscription.FieldUsername, field.TypeString, value)
	}
	if value, ok := suo.mutation.HeadMarker(); ok {
		_spec.SetField(subscription.FieldHeadMarker, field.TypeString, value)
	}
	if suo.mutation.HeadMarkerCleared() {
		_spec.ClearField(subscription.FieldHeadMarker, field.TypeString)
	}
	if value, ok := suo.mutation.StashID(); ok {
		_spec.SetField(subscription.FieldStashID, field.TypeString, value)
	}
	if suo.mutation.StashIDCleared() {
		_spec.ClearField(subscription.FieldStashID, field.TypeString)
	}
	if value, ok := suo.mutation.OrganizedAt(); ok {
		_spec.SetField(subscription.FieldOrganizedAt, field.TypeTime, value)
	}
	if suo.mutation.OrganizedAtCleared() {
		_spec.ClearField(subscription.FieldOrganizedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.Enabled(); ok {
		_spec.SetField(subscription.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(subscription.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(subscription.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.PostsTable,
			Columns: []string{subscription.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !suo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.PostsTable,
			Columns: []string{subscription.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.PostsTable,
			Columns: []string{subscription.PostsColumn},
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
	if suo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.MessagesTable,
			Columns: []string{subscription.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !suo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.MessagesTable,
			Columns: []string{subscription.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscription.MessagesTable,
			Columns: []string{subscription.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subscription{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
