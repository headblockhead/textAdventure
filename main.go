package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear() // Clear current screen
	tm.MoveCursor(1, 1)
	tm.Flush()

	mainPath0.Commands["left"] = func(s *state) {
		fmt.Println("You Turn Left")
		s.room = mRoom
	}
	mainPath0.Commands["forward"] = func(s *state) {
		fmt.Println("You Move forward")
		s.room = mainPath2
	}

	mRoom.Commands["backwards"] = func(s *state) {
		fmt.Println("You Move Back to where you came from")
		s.room = mainPath0
	}
	mRoom.Commands["pick up hammer"] = func(s *state) {
		fmt.Println("You Pick up the hammer")
		s.hammergot = true
		s.hiddenCommands["Mausoleum/pick up hammer"] = struct{}{}
	}

	mainPath2.Commands["right"] = func(s *state) {
		fmt.Println("You Turn Right")
		s.hiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
		s.room = gTRoom
	}

	mainPath2.Commands["back"] = func(s *state) {
		fmt.Println("You go back")
		s.room = mainPath0
	}

	gTRoom.Commands["back"] = func(s *state) {
		fmt.Println("You go back")
		s.room = mainPath2
	}
	gTRoom.Commands["pick up safe code"] = func(s *state) {
		fmt.Println("You pick up the safe code")
		s.Safecodegot = true
		s.hiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
	}

	s := &state{
		room:           startRoom,
		hiddenCommands: map[string]struct{}{},
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		renderRoom(s.room, s)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		action, ok := s.room.Commands[strings.TrimSpace(text)]
		if ok && !commandIsHidden(strings.TrimSpace(text), s) {
			action(s)
		} else {
			fmt.Println()
			fmt.Println("that is not a valid command!")
		}
	}
}

type state struct {
	room            *room
	hammergot       bool
	Safecodegot     bool
	electricityon   bool
	breakerroomused bool
	rocksFallen     bool
	hiddenCommands  map[string]struct{}
}

type action func(s *state)

type room struct {
	Title     string
	Desc      string
	Commands  map[string]action
	StateDesc func(s *state) string
}

func renderRoom(r *room, s *state) {
	fmt.Println(r.Title)
	fmt.Println()
	if r.StateDesc != nil {
		fmt.Println(r.StateDesc(s))
	} else {
		fmt.Println(r.Desc)
	}
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println(strings.Join(getCommands(r.Commands, s), "|"))
}

func commandIsHidden(cmd string, s *state) bool {
	_, ok := s.hiddenCommands[s.room.Title+"/"+cmd]
	return ok
}

func getCommands(m map[string]action, s *state) (commands []string) {
	for k := range m {
		if _, ok := s.hiddenCommands[s.room.Title+"/"+k]; !ok {
			commands = append(commands, k)
		}
	}
	sort.Strings(commands)
	return
}

var startRoom = &room{
	Title: "Road",
	Desc:  "You are standing on a road. You don't know why you are here, and doubt you ever will know. 3 men in a vehicle are chasing you and you come accross a split in the path. You see a long winding path to your left, and an entrance to what seems like a park on your right. The entrence to the park has a metal gate which you can lock, But once you are inside, there seems to be no way out. \n Which direction do you choose?",
	Commands: map[string]action{
		"left": func(s *state) {
			fmt.Println("\n You turn left.\n ")
			s.room = leftStartPath
		},
		"right": func(s *state) {
			fmt.Println("\n You turn right\n ")
			s.room = mainPath0
		},
	},
}

var leftStartPath = &room{
	Title: "Left Path",
	Desc:  " You run as fast as you can along the left path. You notice the group in the car continue approching. You are fast, but not fast enough, the car stops and the men get out. They seem to want to kill you. You cannot escape as they drag you into their van. This is the end for you.\n You died.",
	Commands: map[string]action{
		"quit": func(s *state) {
			fmt.Println("\n You Quit the game.\n ")
			os.Exit(1)
		},
	},
}

var mainPath0 = &room{
	Title:    "Right Path",
	Desc:     " You run as fast as you can along the right path. You dive into the enclosed space and lock the gate. You're safe for the moment, but you know that they will wait for you to come out, no matter what. It is getting dark now and you pull out your lantern. You look around and see a path to your left, There is also a path forewards.\n Which Direction do you go in?",
	Commands: map[string]action{},
}
var mRoom = &room{
	Title:    "Mausoleum",
	Commands: map[string]action{},
	StateDesc: func(s *state) string {
		general := "You turn left and walk along the path towards a building that looks like a mausoleum. You look around yourself and notice graves, all over the floor some with bones still half sticking out. You begin to wonder what you got yourself into as you approach the Mausoleum.\n"
		if s.hammergot == true {
			return general + "There used to be a hammer on the floor, but you picked it up."
		}
		return general + "There is a hammer on the floor."
	},
}
var mainPath2 = &room{
	Title:    "Path",
	Desc:     "You walk to the next intersection on the path and notice a guard tower to your right. You can continue walking foreward, or take the path to your right.",
	Commands: map[string]action{},
}
var gTRoom = &room{
	Title:    "Guard Tower",
	Commands: map[string]action{},
	StateDesc: func(s *state) string {
		general := "You Approach the guard tower"
		if s.electricityon == true && s.breakerroomused == true {
			delete(s.hiddenCommands, "Guard Tower/pick up safe code")
			return general + ". The door is unlocked, You enter and find a meatal block on the floor with a safe code sticking out of it"
		}
		if s.electricityon == true {

			return general + ". The door is unlocked, You enter and find a meatal block on the floor with a safe code sticking out of it. You cannot lift it as an electromagnet has it stuck to the floor"
		}

		return general + ". The Door is locked, it requires electricity to function"
	},
}
