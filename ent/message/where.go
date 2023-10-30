// Code generated by ent, DO NOT EDIT.

package message

import (
	"github.com/rianfowler/fist-bump-chat/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldID, id))
}

// Message applies equality check predicate on the "Message" field. It's identical to MessageEQ.
func Message(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldMessage, v))
}

// MessageEQ applies the EQ predicate on the "Message" field.
func MessageEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "Message" field.
func MessageNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "Message" field.
func MessageIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "Message" field.
func MessageNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "Message" field.
func MessageGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "Message" field.
func MessageGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "Message" field.
func MessageLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "Message" field.
func MessageLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "Message" field.
func MessageContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "Message" field.
func MessageHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "Message" field.
func MessageHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "Message" field.
func MessageEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "Message" field.
func MessageContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldMessage, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(sql.NotPredicates(p))
}
