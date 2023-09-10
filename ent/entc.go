//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ext, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./graph/ent.graphqls"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	cnf := &gen.Config{
		Features: []gen.Feature{
			gen.FeatureExecQuery,
		},
	}
	if err := entc.Generate("./ent/schema", cnf, entc.Extensions(ext)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
