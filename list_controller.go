package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/**
 * @api {get} /lists Get Lists
 * @apiName GetLists
 * @apiGroup Lists
 */
func GetLists(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	res, _ := app.RedisClient.Keys(ctx, "*").Result()

	keys, err := json.Marshal(res)

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while getting the keys"))
	} else {

		w.Write([]byte(keys))
	}

}

/**
 * @api {get} /lists/:listName Get List
 * @apiName GetList
 * @apiGroup Lists
 */
func GetList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	list, err := app.RedisClient.LRange(ctx, vars["listName"], 0, -1).Result()

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while getting the list"))
		return
	}

	listJson, _ := json.Marshal(list)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(listJson))
}

/**
 * @api {post} /lists/Create a List
 * @apiName CreateList
 * @apiGroup Lists
 */
func CreateList(w http.ResponseWriter, r *http.Request) {

	var list List

	err := json.NewDecoder(r.Body).Decode(&list)

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while decoding the body"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if len(list.Tasks) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The list must contain at least one task"))
		return
	}

	items, err := json.Marshal(list.Tasks)

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while marshalling the tasks"))
		return
	}

	_, err = app.RedisClient.LPush(ctx, list.Name, items).Result()

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while saving the list"))
		return
	}

	listJson, _ := json.Marshal(list)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(listJson))
}

/**
 * @api {put} /lists/:listName Update a List
 * @apiName UpdateList
 * @apiGroup Lists
 */
func UpdateList(w http.ResponseWriter, r *http.Request) {

	var list List

	err := json.NewDecoder(r.Body).Decode(&list)

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while decoding the body"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// Remove the old list
	// Obviously a better way to do this would be acting upon items directly
	_, err = app.RedisClient.Del(ctx, list.Name).Result()
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while deleting the list"))
		return
	}

	if len(list.Tasks) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The list must contain at least one task"))
		return
	}

	items, _ := json.Marshal(list.Tasks)

	_, err = app.RedisClient.RPush(ctx, list.Name, items).Result()

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while updating the list"))
		return
	}

	listJson, _ := json.Marshal(list)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(listJson))
}

/**
 * @api {delete} /lists/:listName Delete a List
 * @apiName DeleteList
 * @apiGroup Lists
 */
func DeleteList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := app.RedisClient.Del(ctx, vars["listName"]).Result()

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error while deleting the list"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
