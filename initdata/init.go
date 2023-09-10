package initdata

import (
	"context"
	"fmt"

	"github.com/Arthur1/economical-graphql-server-demo/ent"
)

type Initializer struct {
	client *ent.Client
}

func NewInitializer(client *ent.Client) *Initializer {
	return &Initializer{client}
}

func (in *Initializer) InitWithContext(ctx context.Context) error {
	tx, err := in.client.Tx(ctx)
	if err != nil {
		return err
	}
	if _, err := in.client.ExecContext(ctx, "PRAGMA foreign_keys = OFF;"); err != nil {
		return err
	}
	if err = initArtists(ctx, tx); err != nil {
		return rollback(tx, fmt.Errorf("failed iniArtists: %w", err))
	}
	if err = initSongs(ctx, tx); err != nil {
		return rollback(tx, fmt.Errorf("failed initSongs: %w", err))
	}
	if _, err := in.client.ExecContext(ctx, "PRAGMA foreign_keys = ON;"); err != nil {
		return err
	}
	return tx.Commit()
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
