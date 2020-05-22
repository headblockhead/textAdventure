package main

import (
	"fmt"
	"os"
	"os/exec"
)
//Cls Clears the screen for all operating systems
func Cls() {
	fmt.Println("\033[2J")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
