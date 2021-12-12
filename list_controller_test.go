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
	payload := []byte(`{"name":"Test-list","tasks":[{"label":"First entry"}]}`)

	req, err := http.NewRequest("POST", "/lists", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func TestGetList(t *testing.T) {
	req, _ := http.NewRequest("GET", "/lists/Test-list", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestUpdateList(t *testing.T) {
	payload := []byte(`{"name":"Test-list","tasks":[{"label":"First entry"}, {"label": "Adding a second entry"}]}`)

	req, err := http.NewRequest("PUT", "/lists/Test-list", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteList(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/lists/Test-list", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusNoContent, response.Code)
}
