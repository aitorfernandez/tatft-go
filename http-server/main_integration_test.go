package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := newInMemoryStore()
	server := PlayerServer{store}
	player := "Bob"

	server.ServeHTTP(httptest.NewRecorder(), newWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newWinRequest(player))

	res := httptest.NewRecorder()
	server.ServeHTTP(res, newScoreRequest(player))

	assertStatus(t, res.Code, http.StatusOK)
	assertResBody(t, res.Body.String(), "3")
}
