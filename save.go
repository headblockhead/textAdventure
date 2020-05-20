package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func save(s *state) error {
	start:
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if strings.EqualFold(strings.TrimSpace(text), "quit") {
		fmt.Println("That is a reseved name, try a different one.")
		goto start
	}
	strings.TrimSpace(text)
	f, err := os.Create(strings.TrimSpace(text)+ ".ETA")
	if err != nil {
	}
	defer f.Close()
	e := json.NewEncoder(f)
	err = e.Encode(s)
	return err
}
