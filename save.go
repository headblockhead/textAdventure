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
	fmt.Println("Enter the name of the savefile you want to save the game into (existing saves will be overwritten if selected)")
	fmt.Println("Existing saves:")
	PrintFiles(".")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if strings.EqualFold(strings.TrimSpace(text), "quit") {
		fmt.Println("That is a reserved name, try a different one.")
		goto start
	}
	strings.TrimSpace(text)
	f, err := os.Create(strings.TrimSpace(text) + ".ETA")
	if err != nil {
	}
	defer f.Close()
	e := json.NewEncoder(f)
	err = e.Encode(s)
	return err
}
