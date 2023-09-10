// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Arthur1/economical-graphql-server-demo/ent/artist"
	"github.com/Arthur1/economical-graphql-server-demo/ent/song"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *ArtistQuery) CollectFields(ctx context.Context, satisfies ...string) (*ArtistQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *ArtistQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(artist.Columns))
		selectedFields = []string{artist.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "songs":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SongClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedSongs(alias, func(wq *SongQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[artist.FieldName]; !ok {
				selectedFields = append(selectedFields, artist.FieldName)
				fieldSeen[artist.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type artistPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ArtistPaginateOption
}

func newArtistPaginateArgs(rv map[string]any) *artistPaginateArgs {
	args := &artistPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*ArtistWhereInput); ok {
		args.opts = append(args.opts, WithArtistFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SongQuery) CollectFields(ctx context.Context, satisfies ...string) (*SongQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return s, nil
	}
	if err := s.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SongQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(song.Columns))
		selectedFields = []string{song.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "artist":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ArtistClient{config: s.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			s.withArtist = query
			if _, ok := fieldSeen[song.FieldArtistID]; !ok {
				selectedFields = append(selectedFields, song.FieldArtistID)
				fieldSeen[song.FieldArtistID] = struct{}{}
			}
		case "title":
			if _, ok := fieldSeen[song.FieldTitle]; !ok {
				selectedFields = append(selectedFields, song.FieldTitle)
				fieldSeen[song.FieldTitle] = struct{}{}
			}
		case "releasedYear":
			if _, ok := fieldSeen[song.FieldReleasedYear]; !ok {
				selectedFields = append(selectedFields, song.FieldReleasedYear)
				fieldSeen[song.FieldReleasedYear] = struct{}{}
			}
		case "artistID":
			if _, ok := fieldSeen[song.FieldArtistID]; !ok {
				selectedFields = append(selectedFields, song.FieldArtistID)
				fieldSeen[song.FieldArtistID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		s.Select(selectedFields...)
	}
	return nil
}

type songPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SongPaginateOption
}

func newSongPaginateArgs(rv map[string]any) *songPaginateArgs {
	args := &songPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*SongWhereInput); ok {
		args.opts = append(args.opts, WithSongFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
