package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"github.com/Arthur1/economical-graphql-server-demo/ent"
	"github.com/Arthur1/economical-graphql-server-demo/initdata"
	"github.com/Arthur1/economical-graphql-server-demo/sqlite"
)

func init() {
	sqlite.Init()
}

func main() {
	client, err := ent.Open(dialect.SQLite, "file:db.sqlite")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schemas: %v", err)
	}

	in := initdata.NewInitializer(client)
	if err := in.InitWithContext(ctx); err != nil {
		log.Fatalf("failed creating records: %v", err)
	}

	fmt.Println("success")
}
