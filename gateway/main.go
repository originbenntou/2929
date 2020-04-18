package main

import (
	"github.com/originbenntou/2929BE/gateway/interfaces/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/originbenntou/2929BE/gateway/graphql/account"
	accountGen "github.com/originbenntou/2929BE/gateway/graphql/account/generated"
	"github.com/originbenntou/2929BE/gateway/graphql/trend"
	trendGen "github.com/originbenntou/2929BE/gateway/graphql/trend/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter()
	r.Use(middleware.Tracing)

	accountSrv := handler.NewDefaultServer(accountGen.NewExecutableSchema(accountGen.Config{Resolvers: account.NewAccountResolver()}))
	trendSrv := handler.NewDefaultServer(trendGen.NewExecutableSchema(trendGen.Config{Resolvers: trend.NewTrendResolver()}))

	r.Path("/").HandlerFunc(playground.Handler("GraphQL playground", "/account"))
	r.Path("/account").Handler(accountSrv)
	r.Path("/trend").Handler(trendSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
