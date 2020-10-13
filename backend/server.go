//go:generate sqlboiler psql --no-back-referencing --wipe --add-soft-deletes
package main

import (
	// "context"
	"log"
	"net/http"
	"os"

	"github.com/sjhitchner/soapcalc/backend/generated/graphql"
	"github.com/sjhitchner/soapcalc/backend/graphql/resolver"

	// ggql "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := graphql.Config{
		Resolvers: &resolver.Resolver{},
	}

	/*
		config.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next ggql.Resolver) (interface{}, error) {
				if !getCurrentUser(ctx).HasRole(role) {
					return nil, nil // fmt.Errorf("Access denied")
				}

			// or let it pass through
			return next(ctx)
		}
	*/

	srv := handler.NewDefaultServer(
		graphql.NewExecutableSchema(config),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
