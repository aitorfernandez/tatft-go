package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"foo": 20,
			"bar": 10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("returns Foo's score", func(t *testing.T) {
		req := newScoreRequest("foo")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusOK)
		assertResBody(t, res.Body.String(), "20")
	})

	t.Run("returns Bar's score", func(t *testing.T) {
		req := newScoreRequest("bar")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusOK)
		assertResBody(t, res.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newScoreRequest("baz")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &PlayerServer{&store}

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "qux"

		req := newWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func newScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get the correct status, got %d, want %d", got, want)
	}
}

func assertResBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("body is wrong, got %q, want %q", got, want)
	}
}
