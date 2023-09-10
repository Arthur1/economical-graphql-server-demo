package schema

import "entgo.io/ent"

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return nil
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return nil
}
