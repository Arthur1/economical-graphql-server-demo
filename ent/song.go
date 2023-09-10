// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Arthur1/economical-graphql-server-demo/ent/artist"
	"github.com/Arthur1/economical-graphql-server-demo/ent/song"
)

// Song is the model entity for the Song schema.
type Song struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// ReleasedYear holds the value of the "released_year" field.
	ReleasedYear int `json:"released_year,omitempty"`
	// ArtistID holds the value of the "artist_id" field.
	ArtistID int `json:"artist_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SongQuery when eager-loading is set.
	Edges        SongEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SongEdges holds the relations/edges for other nodes in the graph.
type SongEdges struct {
	// Artist holds the value of the artist edge.
	Artist *Artist `json:"artist,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// ArtistOrErr returns the Artist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SongEdges) ArtistOrErr() (*Artist, error) {
	if e.loadedTypes[0] {
		if e.Artist == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: artist.Label}
		}
		return e.Artist, nil
	}
	return nil, &NotLoadedError{edge: "artist"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Song) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case song.FieldID, song.FieldReleasedYear, song.FieldArtistID:
			values[i] = new(sql.NullInt64)
		case song.FieldTitle:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Song fields.
func (s *Song) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case song.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case song.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case song.FieldReleasedYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field released_year", values[i])
			} else if value.Valid {
				s.ReleasedYear = int(value.Int64)
			}
		case song.FieldArtistID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field artist_id", values[i])
			} else if value.Valid {
				s.ArtistID = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Song.
// This includes values selected through modifiers, order, etc.
func (s *Song) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryArtist queries the "artist" edge of the Song entity.
func (s *Song) QueryArtist() *ArtistQuery {
	return NewSongClient(s.config).QueryArtist(s)
}

// Update returns a builder for updating this Song.
// Note that you need to call Song.Unwrap() before calling this method if this Song
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Song) Update() *SongUpdateOne {
	return NewSongClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Song entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Song) Unwrap() *Song {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Song is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Song) String() string {
	var builder strings.Builder
	builder.WriteString("Song(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteString(", ")
	builder.WriteString("released_year=")
	builder.WriteString(fmt.Sprintf("%v", s.ReleasedYear))
	builder.WriteString(", ")
	builder.WriteString("artist_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ArtistID))
	builder.WriteByte(')')
	return builder.String()
}

// Songs is a parsable slice of Song.
type Songs []*Song