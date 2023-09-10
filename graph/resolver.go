package graph

import (
	"github.com/Arthur1/economical-graphql-server-demo/ent"
	"github.com/Arthur1/economical-graphql-server-demo/ent/artist"
	"github.com/Arthur1/economical-graphql-server-demo/ent/song"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *ent.Client
}

func getTypeFromID(id int) string {
	if id > 1000 {
		return artist.Table
	} else {
		return song.Table
	}
}
