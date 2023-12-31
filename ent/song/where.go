// Code generated by ent, DO NOT EDIT.

package song

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Arthur1/economical-graphql-server-demo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Song {
	return predicate.Song(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Song {
	return predicate.Song(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Song {
	return predicate.Song(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Song {
	return predicate.Song(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Song {
	return predicate.Song(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Song {
	return predicate.Song(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Song {
	return predicate.Song(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldTitle, v))
}

// ReleasedYear applies equality check predicate on the "released_year" field. It's identical to ReleasedYearEQ.
func ReleasedYear(v int) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldReleasedYear, v))
}

// ArtistID applies equality check predicate on the "artist_id" field. It's identical to ArtistIDEQ.
func ArtistID(v int) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldArtistID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Song {
	return predicate.Song(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Song {
	return predicate.Song(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Song {
	return predicate.Song(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Song {
	return predicate.Song(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Song {
	return predicate.Song(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Song {
	return predicate.Song(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Song {
	return predicate.Song(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Song {
	return predicate.Song(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Song {
	return predicate.Song(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Song {
	return predicate.Song(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Song {
	return predicate.Song(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Song {
	return predicate.Song(sql.FieldContainsFold(FieldTitle, v))
}

// ReleasedYearEQ applies the EQ predicate on the "released_year" field.
func ReleasedYearEQ(v int) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldReleasedYear, v))
}

// ReleasedYearNEQ applies the NEQ predicate on the "released_year" field.
func ReleasedYearNEQ(v int) predicate.Song {
	return predicate.Song(sql.FieldNEQ(FieldReleasedYear, v))
}

// ReleasedYearIn applies the In predicate on the "released_year" field.
func ReleasedYearIn(vs ...int) predicate.Song {
	return predicate.Song(sql.FieldIn(FieldReleasedYear, vs...))
}

// ReleasedYearNotIn applies the NotIn predicate on the "released_year" field.
func ReleasedYearNotIn(vs ...int) predicate.Song {
	return predicate.Song(sql.FieldNotIn(FieldReleasedYear, vs...))
}

// ReleasedYearGT applies the GT predicate on the "released_year" field.
func ReleasedYearGT(v int) predicate.Song {
	return predicate.Song(sql.FieldGT(FieldReleasedYear, v))
}

// ReleasedYearGTE applies the GTE predicate on the "released_year" field.
func ReleasedYearGTE(v int) predicate.Song {
	return predicate.Song(sql.FieldGTE(FieldReleasedYear, v))
}

// ReleasedYearLT applies the LT predicate on the "released_year" field.
func ReleasedYearLT(v int) predicate.Song {
	return predicate.Song(sql.FieldLT(FieldReleasedYear, v))
}

// ReleasedYearLTE applies the LTE predicate on the "released_year" field.
func ReleasedYearLTE(v int) predicate.Song {
	return predicate.Song(sql.FieldLTE(FieldReleasedYear, v))
}

// ArtistIDEQ applies the EQ predicate on the "artist_id" field.
func ArtistIDEQ(v int) predicate.Song {
	return predicate.Song(sql.FieldEQ(FieldArtistID, v))
}

// ArtistIDNEQ applies the NEQ predicate on the "artist_id" field.
func ArtistIDNEQ(v int) predicate.Song {
	return predicate.Song(sql.FieldNEQ(FieldArtistID, v))
}

// ArtistIDIn applies the In predicate on the "artist_id" field.
func ArtistIDIn(vs ...int) predicate.Song {
	return predicate.Song(sql.FieldIn(FieldArtistID, vs...))
}

// ArtistIDNotIn applies the NotIn predicate on the "artist_id" field.
func ArtistIDNotIn(vs ...int) predicate.Song {
	return predicate.Song(sql.FieldNotIn(FieldArtistID, vs...))
}

// HasArtist applies the HasEdge predicate on the "artist" edge.
func HasArtist() predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArtistTable, ArtistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtistWith applies the HasEdge predicate on the "artist" edge with a given conditions (other predicates).
func HasArtistWith(preds ...predicate.Artist) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		step := newArtistStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Song) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Song) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Song) predicate.Song {
	return predicate.Song(func(s *sql.Selector) {
		p(s.Not())
	})
}
