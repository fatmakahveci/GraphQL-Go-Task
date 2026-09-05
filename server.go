package main

import (
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"

	"github.com/rs/cors"

	"github.com/pvormste/task-gql-full-stack/graph"
	"github.com/pvormste/task-gql-full-stack/graph/generated"
)

func newHandler() http.Handler {
	gqlHandler := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	gqlHandler.AddTransport(transport.POST{})
	gqlHandler.Use(extension.Introspection{})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	mux := http.NewServeMux()
	mux.Handle("/query", c.Handler(gqlHandler))
	return mux
}

func main() {
	log.Println("Send operations to: http://localhost:8085/query")
	server := &http.Server{
		Addr:              ":8085",
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
