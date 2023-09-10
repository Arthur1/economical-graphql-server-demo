package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/Arthur1/economical-graphql-server-demo/ent"
	"github.com/Arthur1/economical-graphql-server-demo/graph/generated"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id, ent.WithFixedNodeType(getTypeFromID(id)))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

func (r *queryResolver) Artists(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.ArtistWhereInput) (*ent.ArtistConnection, error) {
	return r.Client.Artist.Query().Paginate(ctx, after, first, before, last, ent.WithArtistFilter(where.Filter))
}

func (r *queryResolver) Songs(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.SongWhereInput) (*ent.SongConnection, error) {
	return r.Client.Song.Query().Paginate(ctx, after, first, before, last, ent.WithSongFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
