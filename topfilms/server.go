package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aayy07/topfilms/graph"
	"github.com/aayy07/topfilms/graph/generated"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	environment := ""
	if len(os.Args) < 2 {
		log.Println("missing application environment method \nUsage: go run server.go {dev | uat | prod} ")
		os.Exit(1)
	} else {
		environment = os.Args[1]
	}

	if environment == "dev" {
		err := godotenv.Load(".dev-env")
		if err != nil {
			log.Println("failed to load .dev-env file")
		}
	} else{
		log.Println("Environment '" + environment + "' not available.\nUsage: go run server.go dev ")
		os.Exit(1)
	}





	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}