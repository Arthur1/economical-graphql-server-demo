package initdata

import (
	"context"
	_ "embed"

	"github.com/Arthur1/economical-graphql-server-demo/ent"
	"github.com/jszwec/csvutil"
)

//go:embed "csv/songs.csv"
var songsBytes []byte

type Song struct {
	ID           int    `csv:"id"`
	Title        string `csv:"title"`
	ReleasedYear int    `csv:"released_year"`
	ArtistID     int    `csv:"artist_id"`
}

func initSongs(ctx context.Context, tx *ent.Tx) error {
	songs := []*Song{}
	if err := csvutil.Unmarshal(songsBytes, &songs); err != nil {
		return err
	}

	builders := make([](*ent.SongCreate), 0, len(songs))
	for _, s := range songs {
		builder := tx.Song.Create().
			SetID(s.ID).
			SetTitle(s.Title).
			SetReleasedYear(s.ReleasedYear).
			SetArtistID(s.ArtistID)
		builders = append(builders, builder)
	}
	return tx.Song.CreateBulk(builders...).Exec(ctx)
}
