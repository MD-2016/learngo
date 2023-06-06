package iop

import (
	"os"
	"testing"
)

func createTempFile(t testing.TB, initData string) (*os.File, func()) {
	t.Helper()

	tmp, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp %v", err)
	}

	tmp.Write([]byte(initData))

	removeFile := func() {
		tmp.Close()
		os.Remove(tmp.Name())
	}

	return tmp, removeFile
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league sorted", func(t *testing.T) {
		db, cleanDatabase := createTempFile(t, `[
			{"Name": "Jim", "Wins": 10},
			{"Name": "Don", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(db)

		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Don", 33},
			{"Jim", 10},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDatabase := createTempFile(t, `[
			{"Name": "Jim", "Wins": 10},
			{"Name": "Don", "Wins": 33},
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(db)

		assertNoError(t, err)

		got := store.GetPlayerScore("Don")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jim", "Wins": 10},
			{"Name": "Don", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		store.RecordWin("Don")

		got := store.GetPlayerScore("Don")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jim", "Wins": 10},
			{"Name": "Don", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		store.RecordWin("Tim")

		got := store.GetPlayerScore("Tim")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
	}
}
