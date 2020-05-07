package main

import (
	"encoding/json"
	"os"
)

func save(s *state) error {
	f, err := os.Create("savedata.ETA")
	if err != nil {
	}
	defer f.Close()
	e := json.NewEncoder(f)
	err = e.Encode(s)
	return err
}
