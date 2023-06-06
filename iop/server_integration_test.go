package iop

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := NewFileSystemPlayerStore(database)

	assertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreRequest(player))
		assertStatus(t, res.Code, http.StatusOK)

		assertResponseBody(t, res.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()

		server.ServeHTTP(res, newLeagueRequest())
		assertStatus(t, res.Code, http.StatusOK)

		got := getLeagueFromResponse(t, res.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
