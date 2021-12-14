package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nguyenvd27/graphql-test/graph/directive"
	"github.com/nguyenvd27/graphql-test/graph/generated"
	"github.com/nguyenvd27/graphql-test/graph/resolver"
	"github.com/nguyenvd27/graphql-test/initialize/db"
	"github.com/nguyenvd27/graphql-test/interactor"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gormDB := db.Initialize()
	reg := interactor.NewRegistry(gormDB)
	initInteractor := interactor.NewInteractor(reg)

	config := generated.Config{
		Resolvers: &resolver.Resolver{
			Interactor: initInteractor,
		},
	}
	config.Directives.NeedAuthentication = directive.GenerateNeedAuthenticationDirective()
	config.Directives.NeedRole = directive.GenerateNeedRoleDirective()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
