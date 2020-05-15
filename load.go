package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func load(s *state) (ok bool, err error, isquit bool) {
	reader := bufio.NewReader(os.Stdin)
	dirname := "."
	fmt.Println("Choose savefile:")
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if path.Ext(file.Name()) == "ETA" {
			fmt.Println(file.Name()[0 : len(file.Name())-4])
		}
	}
	fmt.Println("Type \"quit\" to exit")
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	if strings.TrimSpace(text) == "exit"{
		isquit = true
		return
	}
	var filename = strings.TrimSpace(text) + ".eta"

	f, err = os.Open(filename)
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
	return err == nil, err, isquit
}
