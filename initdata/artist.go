package initdata

import (
	"context"
	_ "embed"

	"github.com/Arthur1/economical-graphql-server-demo/ent"
	"github.com/jszwec/csvutil"
)

//go:embed "csv/artists.csv"
var artistsBytes []byte

type Artist struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
}

func initArtists(ctx context.Context, tx *ent.Tx) error {
	artists := []*Artist{}
	if err := csvutil.Unmarshal(artistsBytes, &artists); err != nil {
		return err
	}

	builders := make([](*ent.ArtistCreate), 0, len(artists))
	for _, a := range artists {
		builder := tx.Artist.Create().
			SetID(a.ID).
			SetName(a.Name)
		builders = append(builders, builder)
	}
	return tx.Artist.CreateBulk(builders...).Exec(ctx)
}
