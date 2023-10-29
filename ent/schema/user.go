package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type User struct {
	ent.Schema
}

// Fields of the Message.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("Email").Unique(),
		// field.Int("GithubID").Unique(),
		field.String("Name").Optional(),
		// field.String("OAuthToken").Sensitive().Optional(),
	}
}

// Edges of the Message.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type), // One-to-Many relationship from User to Message
	}
}
