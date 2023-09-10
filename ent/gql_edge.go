// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Artist) Songs(ctx context.Context) (result []*Song, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedSongs(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.SongsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QuerySongs().All(ctx)
	}
	return result, err
}

func (s *Song) Artist(ctx context.Context) (*Artist, error) {
	result, err := s.Edges.ArtistOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryArtist().Only(ctx)
	}
	return result, err
}