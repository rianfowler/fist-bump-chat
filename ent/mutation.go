// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fistbump/ent/message"
	"fistbump/ent/predicate"
	"fistbump/ent/user"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMessage = "Message"
	TypeUser    = "User"
)

// MessageMutation represents an operation that mutates the Message nodes in the graph.
type MessageMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_Message      *string
	clearedFields map[string]struct{}
	user          *int
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Message, error)
	predicates    []predicate.Message
}

var _ ent.Mutation = (*MessageMutation)(nil)

// messageOption allows management of the mutation configuration using functional options.
type messageOption func(*MessageMutation)

// newMessageMutation creates new mutation for the Message entity.
func newMessageMutation(c config, op Op, opts ...messageOption) *MessageMutation {
	m := &MessageMutation{
		config:        c,
		op:            op,
		typ:           TypeMessage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMessageID sets the ID field of the mutation.
func withMessageID(id int) messageOption {
	return func(m *MessageMutation) {
		var (
			err   error
			once  sync.Once
			value *Message
		)
		m.oldValue = func(ctx context.Context) (*Message, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Message.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMessage sets the old Message of the mutation.
func withMessage(node *Message) messageOption {
	return func(m *MessageMutation) {
		m.oldValue = func(context.Context) (*Message, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MessageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MessageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MessageMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MessageMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Message.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetMessage sets the "Message" field.
func (m *MessageMutation) SetMessage(s string) {
	m._Message = &s
}

// Message returns the value of the "Message" field in the mutation.
func (m *MessageMutation) Message() (r string, exists bool) {
	v := m._Message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "Message" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "Message" field.
func (m *MessageMutation) ResetMessage() {
	m._Message = nil
}

// SetUserID sets the "user" edge to the User entity by id.
func (m *MessageMutation) SetUserID(id int) {
	m.user = &id
}

// ClearUser clears the "user" edge to the User entity.
func (m *MessageMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *MessageMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *MessageMutation) UserID() (id int, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *MessageMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the MessageMutation builder.
func (m *MessageMutation) Where(ps ...predicate.Message) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MessageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MessageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Message, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MessageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MessageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Message).
func (m *MessageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MessageMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m._Message != nil {
		fields = append(fields, message.FieldMessage)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MessageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case message.FieldMessage:
		return m.Message()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MessageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case message.FieldMessage:
		return m.OldMessage(ctx)
	}
	return nil, fmt.Errorf("unknown Message field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case message.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MessageMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MessageMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Message numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MessageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MessageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MessageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Message nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MessageMutation) ResetField(name string) error {
	switch name {
	case message.FieldMessage:
		m.ResetMessage()
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MessageMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, message.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MessageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case message.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MessageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MessageMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MessageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, message.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MessageMutation) EdgeCleared(name string) bool {
	switch name {
	case message.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MessageMutation) ClearEdge(name string) error {
	switch name {
	case message.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Message unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MessageMutation) ResetEdge(name string) error {
	switch name {
	case message.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Message edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op              Op
	typ             string
	id              *int
	_Email          *string
	_Name           *string
	clearedFields   map[string]struct{}
	messages        map[int]struct{}
	removedmessages map[int]struct{}
	clearedmessages bool
	done            bool
	oldValue        func(context.Context) (*User, error)
	predicates      []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "Email" field.
func (m *UserMutation) SetEmail(s string) {
	m._Email = &s
}

// Email returns the value of the "Email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m._Email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "Email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "Email" field.
func (m *UserMutation) ResetEmail() {
	m._Email = nil
}

// SetName sets the "Name" field.
func (m *UserMutation) SetName(s string) {
	m._Name = &s
}

// Name returns the value of the "Name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m._Name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "Name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "Name" field.
func (m *UserMutation) ClearName() {
	m._Name = nil
	m.clearedFields[user.FieldName] = struct{}{}
}

// NameCleared returns if the "Name" field was cleared in this mutation.
func (m *UserMutation) NameCleared() bool {
	_, ok := m.clearedFields[user.FieldName]
	return ok
}

// ResetName resets all changes to the "Name" field.
func (m *UserMutation) ResetName() {
	m._Name = nil
	delete(m.clearedFields, user.FieldName)
}

// AddMessageIDs adds the "messages" edge to the Message entity by ids.
func (m *UserMutation) AddMessageIDs(ids ...int) {
	if m.messages == nil {
		m.messages = make(map[int]struct{})
	}
	for i := range ids {
		m.messages[ids[i]] = struct{}{}
	}
}

// ClearMessages clears the "messages" edge to the Message entity.
func (m *UserMutation) ClearMessages() {
	m.clearedmessages = true
}

// MessagesCleared reports if the "messages" edge to the Message entity was cleared.
func (m *UserMutation) MessagesCleared() bool {
	return m.clearedmessages
}

// RemoveMessageIDs removes the "messages" edge to the Message entity by IDs.
func (m *UserMutation) RemoveMessageIDs(ids ...int) {
	if m.removedmessages == nil {
		m.removedmessages = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.messages, ids[i])
		m.removedmessages[ids[i]] = struct{}{}
	}
}

// RemovedMessages returns the removed IDs of the "messages" edge to the Message entity.
func (m *UserMutation) RemovedMessagesIDs() (ids []int) {
	for id := range m.removedmessages {
		ids = append(ids, id)
	}
	return
}

// MessagesIDs returns the "messages" edge IDs in the mutation.
func (m *UserMutation) MessagesIDs() (ids []int) {
	for id := range m.messages {
		ids = append(ids, id)
	}
	return
}

// ResetMessages resets all changes to the "messages" edge.
func (m *UserMutation) ResetMessages() {
	m.messages = nil
	m.clearedmessages = false
	m.removedmessages = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m._Email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m._Name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldName) {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldName:
		m.ClearName()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.messages != nil {
		edges = append(edges, user.EdgeMessages)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.messages))
		for id := range m.messages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedmessages != nil {
		edges = append(edges, user.EdgeMessages)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeMessages:
		ids := make([]ent.Value, 0, len(m.removedmessages))
		for id := range m.removedmessages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedmessages {
		edges = append(edges, user.EdgeMessages)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeMessages:
		return m.clearedmessages
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeMessages:
		m.ResetMessages()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
