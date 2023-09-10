package main

import (
	"log"
	"net/http"
	"os"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Arthur1/economical-graphql-server-demo/ent"
	"github.com/Arthur1/economical-graphql-server-demo/graph"
	"github.com/Arthur1/economical-graphql-server-demo/graph/generated"
	"github.com/Arthur1/economical-graphql-server-demo/sqlite"
)

const (
	defaultPort = "8080"
)

func init() {
	sqlite.Init()
}

func main() {
	client, err := ent.Open(dialect.SQLite, "file:db.sqlite?cache=shared&mode=ro")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
