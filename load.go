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

//PrintFiles read files in the selected directory that end with .ETA
func PrintFiles(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if path.Ext(file.Name()) == ".save" {
			fmt.Println(file.Name()[0 : len(file.Name())-4])
		}
	}
}

func load(s *state) (ok bool, isquit bool, err error) {
Cls()
	fmt.Println("Choose savefile:")
	PrintFiles(".")
	fmt.Println("Type \"quit\" to go back to the menu")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if strings.EqualFold(strings.TrimSpace(text), "quit") {
		isquit = true
		Cls()
		return
	}
	var filename = strings.TrimSpace(text) + ".save"

	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	e := json.NewDecoder(f)
	err = e.Decode(s)
	return err == nil, isquit, err
}
