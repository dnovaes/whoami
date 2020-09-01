package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/dnovaes/portfolio/database"
	"github.com/dnovaes/portfolio/gqlgen/graph"
	"github.com/dnovaes/portfolio/gqlgen/graph/generated"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const host string = "localhost"

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", 8090)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func init() {
	SetupCloseHandler()
	db.StartConnection()
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		db.StopConnection()
		db.Db.Cancel()
		os.Exit(0)
	}()
}
