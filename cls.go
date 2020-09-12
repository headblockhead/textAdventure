package main

import (
	"fmt"
	"os"
	"os/exec"
)
//Cls Clears the screen for all operating systems by printing an escape code for mac and linux, then running the cls terminal command for windows.
// this is run after the escape code so that it clears up any weird charectors left behind by windows not understanding the escape code.
func Cls() {
	fmt.Println("\033[2J")
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
