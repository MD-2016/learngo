package iop

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(f *os.File) (*FileSystemPlayerStore, error) {
	err := initPlayerDBFile(f)

	if err != nil {
		return nil, fmt.Errorf("problem reading db file, %v", err)
	}

	leag, err := NewLeague(f)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from db %s, %v", f.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{f}),
		league:   leag,
	}, nil
}

func initPlayerDBFile(f *os.File) error {
	f.Seek(0, 0)

	inf, err := f.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info %s, %v", f.Name(), err)
	}

	if inf.Size() == 0 {
		f.Write([]byte("[]"))
		f.Seek(0, 0)
	}

	return nil
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
