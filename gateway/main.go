package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/originbenntou/2929BE/gateway/interfaces/middleware"
	"github.com/originbenntou/2929BE/shared/logger"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
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
	// playground起動中は切る
	//r.Use(middleware.Logging)

	accountSrv := handler.NewDefaultServer(accountGen.NewExecutableSchema(accountGen.Config{Resolvers: account.NewAccountResolver()}))
	trendSrv := handler.NewDefaultServer(trendGen.NewExecutableSchema(trendGen.Config{Resolvers: trend.NewTrendResolver()}))

	r.Path("/account").Handler(accountSrv)
	r.Path("/trend").Handler(trendSrv)

	if os.Getenv("ENV") == "LOCAL" {
		r.Path("/").HandlerFunc(playground.Handler("GraphQL playground", "/account"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Common.Fatal(err.Error())
	}
}
