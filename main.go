package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var app App

type App struct {
	RedisClient *redis.Client
	Router      *mux.Router
	Server      http.Server
}

func (app *App) Initialize(c chan os.Signal) {

	// Redis setup
	app.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Router
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/lists", GetLists).Methods("GET")
	app.Router.HandleFunc("/lists/{listName}", GetList).Methods("GET")
	app.Router.HandleFunc("/lists", CreateList).Methods("POST")
	app.Router.HandleFunc("/lists/{listName}", UpdateList).Methods("PUT", "PATCH")
	app.Router.HandleFunc("/lists/{listName}", DeleteList).Methods("DELETE")

	// Server
	app.Server = http.Server{
		Handler: app.Router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

}

func main() {

	log.Println("⏱  Starting the server...")

	// Exit channel
	c := make(chan os.Signal, 1)

	app = App{}
	app.Initialize(c)

	go func(c chan os.Signal) {
		log.Println("👂 Listening to connections.")
		if err := app.Server.ListenAndServe(); err != nil {
			c <- os.Interrupt
		}
	}(c)

	<-c

	signal.Notify(c, os.Interrupt)

	log.Println("🚨  Stopping the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.Server.Shutdown(ctx)

	os.Exit(0)
}
