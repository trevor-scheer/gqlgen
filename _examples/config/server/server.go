package main

import (
	"log"
	"net/http"

	todo "github.com/trevor-scheer/gqlgen/_examples/config"
	"github.com/trevor-scheer/gqlgen/graphql/handler"
	"github.com/trevor-scheer/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
