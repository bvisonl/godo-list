package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestGetLists(t *testing.T) {
	req, _ := http.NewRequest("GET", "/lists", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestCreateList(t *testing.T) {
	payload := []byte(`{"name":"Listsss","tasks":[{"label":"First entry"}]}`)

	req, err := http.NewRequest("POST", "/lists", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
}
