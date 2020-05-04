package main

import (
	"encoding/json"
	"os"
)

func load(s *state) (ok bool, err error) {
	f, err := os.Open("savedata.savTada")
	if err == os.ErrNotExist {
		err = nil
		return
	}
	if err != nil {
		return
	}
	defer f.Close()
	e := json.NewDecoder(f)
	err = e.Decode(s)
	return err == nil, err
}
