package schema

import "entgo.io/ent"

// Song holds the schema definition for the Song entity.
type Song struct {
	ent.Schema
}

// Fields of the Song.
func (Song) Fields() []ent.Field {
	return nil
}

// Edges of the Song.
func (Song) Edges() []ent.Edge {
	return nil
}
