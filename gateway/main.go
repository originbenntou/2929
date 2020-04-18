package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	accountSrv := handler.NewDefaultServer(accountGen.NewExecutableSchema(accountGen.Config{Resolvers: account.NewAccountResolver()}))
	trendSrv := handler.NewDefaultServer(trendGen.NewExecutableSchema(trendGen.Config{Resolvers: trend.NewTrendResolver()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/account"))
	http.Handle("/account", accountSrv)
	http.Handle("/trend", trendSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
