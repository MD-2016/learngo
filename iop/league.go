package iop

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(rd io.Reader) (League, error) {
	var leag []Player
	err := json.NewDecoder(rd).Decode(&leag)

	if err != nil {
		err = fmt.Errorf("problem parsing league %v", err)
	}

	return leag, err
}
