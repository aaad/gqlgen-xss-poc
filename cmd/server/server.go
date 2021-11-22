package main

import (
	"fmt"
	"log"
	"net/http"

	iotuil "io/ioutil"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aaad/gqlgen-xss-poc/internal/graph"
	ggen "github.com/aaad/gqlgen-xss-poc/internal/graph/generated"
)

func main() {
	port := "7777"

	srv := handler.NewDefaultServer(
		ggen.NewExecutableSchema(
			ggen.Config{
				Resolvers: &graph.Resolver{},
			},
		))

	http.Handle("/playground", playground.Handler("Poc", "/query"))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		body, err := iotuil.ReadFile("assets/poc.html")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		fmt.Fprintf(w, string(body))
	}))
	http.Handle("/query", srv)

	log.Printf("connect to http://0.0.0.0:%s/ for POC or http://0.0.0.0:%s/playground for GraphQL playground", port, port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
