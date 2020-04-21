package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/originbenntou/2929BE/gateway/interfaces/middleware"
	"github.com/originbenntou/2929BE/shared/logger"
	"log"
	"net/http"
	"os"
	"time"

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
	// playground起動中は鬱陶しいので切っておく
	//r.Use(middleware.Logging)
	r.Use(middleware.Tracing)
	r.Use(middleware.NewCORS)

	auth := middleware.NewAuthentication()

	aSvc := handler.NewDefaultServer(accountGen.NewExecutableSchema(accountGen.Config{Resolvers: account.NewAccountResolver()}))
	tSvc := handler.NewDefaultServer(trendGen.NewExecutableSchema(trendGen.Config{Resolvers: trend.NewTrendResolver()}))

	r.Path("/account").Handler(aSvc)
	r.Path("/trend").Handler(auth(tSvc))

	if os.Getenv("ENV") == "LOCAL" {
		r.Path("/").HandlerFunc(playground.Handler("GraphQL playground", "/account"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Common.Fatal(err.Error())
	}
}
