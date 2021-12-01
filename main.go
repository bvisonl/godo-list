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

var rdb *redis.Client

func main() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r := mux.NewRouter()
	r.HandleFunc("/lists", GetLists).Methods("GET")
	r.HandleFunc("/lists/{listName}", GetList).Methods("GET")
	r.HandleFunc("/lists", CreateList).Methods("POST")
	r.HandleFunc("/lists/{listName}", UpdateList).Methods("PUT", "PATCH")
	r.HandleFunc("/lists/{listName}", DeleteList).Methods("DELETE")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}

	c := make(chan os.Signal, 1)

	go func(c chan os.Signal) {
		if err := srv.ListenAndServe(); err != nil {
			c <- os.Interrupt
		}
	}(c)

	signal.Notify(c, os.Interrupt)

	<-c

	srv.Shutdown(ctx)
	log.Println("Shutting down")

	os.Exit(0)
}
