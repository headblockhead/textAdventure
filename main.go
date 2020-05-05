package main

// TO DO : GLOBAL COMMANDS! SAVE QUIT.
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

	mainPath0.commands["left"] = func(s *state) {
		fmt.Println("You Turn Left")
		s.Room = mRoom
		s.HiddenCommands["Mausoleum/go through tunnel"] = struct{}{}
	}
	mainPath0.commands["forward"] = func(s *state) {
		fmt.Println("You Move forward")
		s.Room = mainPath2
	}

	mRoom.commands["backward"] = func(s *state) {
		fmt.Println("You Move Back to where you came from")
		s.Room = mainPath0
	}
	mRoom.commands["pick up hammer"] = func(s *state) {
		fmt.Println("You Pick up the hammer")
		s.Hammergot = true
		s.HiddenCommands["Mausoleum/pick up hammer"] = struct{}{}
	}

	mainPath2.commands["right"] = func(s *state) {
		fmt.Println("You Turn Right")
		s.HiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
		s.Room = gTRoom
	}

	mainPath2.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.Room = mainPath0
	}
	mainPath2.commands["forward"] = func(s *state) {
		fmt.Println("You continue walking")
		s.Room = mainPath3
	}

	gTRoom.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.Room = mainPath2
	}
	gTRoom.commands["pick up safe code"] = func(s *state) {
		fmt.Println("You pick up the safe code")
		s.Safecodegot = true
		s.HiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
	}
	mainPath3.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.Room = mainPath2
	}
	mainPath3.commands["left"] = func(s *state) {
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
		fmt.Println("You go left")
		s.Room = maroom
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
	}
	mainPath3.commands["right"] = func(s *state) {
		fmt.Println("You go right")
		s.Room = gRoom
	}
	gRoom.commands["turn on power"] = func(s *state) {
		fmt.Println("You turn on the power")
		s.Electricity = true
		s.HiddenCommands["Generator/turn on power"] = struct{}{}
	}
	gRoom.commands["backward"] = func(s *state) {
		fmt.Println("You turn back")
		s.Room = mainPath3
	}
	maroom.commands["backward"] = func(s *state) {
		fmt.Println("You turn back")
		s.Room = mainPath3
	}
	mRoom.commands["go through tunnel"] = func(s *state) {
		fmt.Println("You go through the tunnel")
		s.Room = mABRoom
	}
	mABRoom.commands["pull lever"] = func(s *state) {
		fmt.Println("You pull the lever and rocks fall into a pit in the center of the room")
		s.Room = mABRoom
		s.RocksFallen = true
		s.HiddenCommands["Basement/pull lever"] = struct{}{}
	}
	mABRoom.commands["backward"] = func(s *state) {
		fmt.Println("You go back to the mausoleum")
		s.Room = mRoom
		s.HiddenCommands["Mausoleum/go through tunnel"] = struct{}{}
	}
	mainPath3.commands["forward"] = func(s *state) {
		fmt.Println("You forward")
		s.Room = mainpath5
		s.HiddenCommands["End of path/exit"] = struct{}{}
	}
	maroom.commands["grab key"] = func(s *state) {
		fmt.Println("You grab the key")
		s.KeyGot = true
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
	}
	mainpath5.commands["exit"] = func(s *state) {
		fmt.Println("You leave.")
		s.Room = gamefinish

	}
	mainpath5.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.Room = mainPath3
	}
	mainpath5.commands["left"] = func(s *state) {
		fmt.Println("You go left")
		s.Room = bRoom
	}
	bRoom.commands["switch off electromagnet"] = func(s *state) {
		fmt.Println("You switch off the electromagnet")
		s.BreakerRoomUsed = true
		s.HiddenCommands["Breaker room/switch off electromagnet"] = struct{}{}
	}
	bRoom.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.Room = mainpath5
	}
	s := &state{
		Room:           titleRoom,
		HiddenCommands: map[string]struct{}{},
		Movestaken:     0,
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		renderRoom(s.Room, s)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		action, ok := s.Room.commands[strings.TrimSpace(text)]
		if ok && !commandIsHidden(strings.TrimSpace(text), s) {
			s.Movestaken++
			action(s)
		} else {
			fmt.Println()
			fmt.Println("that is not a valid command!")
		}
	}
}

type state struct {
	//game_data
	Room            *room
	Hammergot       bool
	Safecodegot     bool
	Electricity     bool
	BreakerRoomUsed bool
	RocksFallen     bool
	KeyGot          bool
	HiddenCommands  map[string]struct{}
	//info_data
	Movestaken int
}

type action func(s *state)

type room struct {
	Title     string
	Desc      string
	commands  map[string]action
	stateDesc func(s *state) string
}

func renderRoom(r *room, s *state) {
	fmt.Println(r.Title)
	fmt.Println()
	if r.stateDesc != nil {
		fmt.Println(r.stateDesc(s))
	} else {
		fmt.Println(r.Desc)
	}
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println(strings.Join(getCommands(r.commands, s), "|"))
	fmt.Println()
}

func commandIsHidden(cmd string, s *state) bool {
	_, ok := s.HiddenCommands[s.Room.Title+"/"+cmd]
	return ok
}

func getCommands(m map[string]action, s *state) (commands []string) {
	for k := range m {
		if _, ok := s.HiddenCommands[s.Room.Title+"/"+k]; !ok {
			commands = append(commands, k)
		}
	}
	sort.Strings(commands)
	return
}

var startRoom = &room{
	Title: "Road",
	Desc:  "You are standing on a road. You don't know why you are here, and doubt you ever will know. 3 men in a vehicle are chasing you and you come accross a split in the path. You see a long winding path to your left, and an entrance to what seems like a park on your right. The entrance to the park has a metal gate which you can lock, But once you are inside, there seems to be no way out. \n Which direction do you choose?",
	commands: map[string]action{
		"left": func(s *state) {
			fmt.Println("\n You turn left.\n ")
			s.Room = leftStartPath
		},
		"right": func(s *state) {
			fmt.Println("\n You turn right\n ")
			s.Room = mainPath0
		},
	},
}

var leftStartPath = &room{
	Title: "Left Path",
	Desc:  " You run as fast as you can along the left path. You notice the group in the car continue approching. You are fast, but not fast enough, the car stops and the men get out. They seem to want to kill you. You cannot escape as they drag you into their van. This is the end for you.\n You died.",
	commands: map[string]action{
		"quit": func(s *state) {
			fmt.Println("\n You Quit the game.\n ")
			os.Exit(1)
		},
	},
}

var mainPath0 = &room{
	Title:    "Right Path",
	Desc:     " You run as fast as you can along the right path. You dive into the enclosed space and lock the gate. You're safe for the moment, but you know that they will wait for you to come out, no matter what. It is getting dark now and you pull out your lantern. You look around and see a path to your left, There is also a path forewards.\n Which Direction do you go in?",
	commands: map[string]action{},
}
var mRoom = &room{
	Title:    "Mausoleum",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You turn left and walk along the path towards a building that looks like a mausoleum. You look around yourself and notice graves, all over the floor some with bones still half sticking out. You begin to wonder what you got yourself into as you approach the Mausoleum.\n"
		if s.Hammergot == true {
			delete(s.HiddenCommands, "Mausoleum/go through tunnel")
			return general + "There used to be a hammer on the floor, but you picked it up. There is also now a tunnel that you can go through."
		}

		return general + "There is a hammer on the floor."
	},
}
var mainPath2 = &room{
	Title:    "Path",
	Desc:     "You walk to the next intersection on the path and notice a prison guard tower to your right. It streaches far into the sky but has no staircase. You can continue walking forward, or take the path to your right.",
	commands: map[string]action{},
}
var gTRoom = &room{
	Title:    "Guard Tower",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk towards the guard tower, the night falling around you, and see that"
		if s.Electricity == true && s.BreakerRoomUsed == true {
			delete(s.HiddenCommands, "Guard Tower/pick up safe code")
			return general + " the door is unlocked, You enter and find a metal block on the floor with a safe code sticking out of it"
		}
		if s.Electricity == true {

			return general + " the door is unlocked, You enter and find a metal block on the floor with a safe code sticking out of it. You cannot lift it as an electromagnet has it stuck to the floor"
		}

		return general + " the door is locked, it requires electricity to function"
	},
}
var mainPath3 = &room{
	Title:    "Path",
	Desc:     "You walk to the next intersection on the path and look around yourself. You look to your left and see an abandoned mansion, covered in winding vines. To your left you see a small house, Presumably where the generator is kept.",
	commands: map[string]action{},
}
var maroom = &room{
	Title:    "Mansion",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk along the short path, observing your surroundings. You notice the distant screams coming from the car that was chasing you. Shivers make their way down your spine as you peek inside the old building. Inside you notice"
		if s.RocksFallen == true && s.Safecodegot == true && s.Hammergot == true {
			delete(s.HiddenCommands, "Mansion/grab key")
			return general + " a key guarded by a safe. You smash the glass, unlock the safe with the code and see a key to the exit."
		}
		if s.RocksFallen == true {
			if s.Hammergot {
				return general + " a safe guarded by a sheet of glass. You smash the glass using a hammer to reveal a safe with a code."
			}
			return general + " a safe guarded by a sheet of glass. A hammer could be useful."
		}

		return general + " a pile of rocks blocking your path."
	},
}
var gRoom = &room{
	Title:    "Generator",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You enter the room with the power generator, the switch on the wall controlling the power output is set to "
		if !s.Electricity {
			delete(s.HiddenCommands, "Generator/turn on power")
			return general + "off."
		}

		return general + "on."
	},
}
var mABRoom = &room{
	Title:    "Basement",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You crawl through the tunnel and reach the basement of the Mansion. There is a lever controlling a trapdoor in the floor above. The lever is "
		if !s.RocksFallen {
			delete(s.HiddenCommands, "Basement/pull lever")
			return general + "not pulled."
		}

		return general + "pulled."
	},
}

var mainpath5 = &room{
	Title:    "End of path",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk along the path and reach a gate. This appears to be your way out. You can turn right."
		if !s.KeyGot {
			delete(s.HiddenCommands, "mainpath5/exit")
		}
		return general
	},
}

var bRoom = &room{
	Title:    "Breaker Room",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You Enter the breaker room and notice a panel of switches. You see one labeled \"Electromagnet\" ."
		if !s.BreakerRoomUsed {
			delete(s.HiddenCommands, "Breaker Room/switch off electromagnet")
		}
		return general
	},
}
var titleRoom = &room{
	Title: "Eddie's Text Adventure",
	Desc:  "MENU",
	commands: map[string]action{
		"quit": func(s *state) {
			fmt.Println("\n You Quit the game.\n ")
			os.Exit(1)
		},
		"start": func(s *state) {
			s.Room = startRoom
		},
		"load": func(s *state) {
			ok, err := load(s)
			if err != nil {
				fmt.Println(err)
			}
			if !ok {
				fmt.Println("No file to load.")
			}
			s.Room = mainPath0
		},
	},
}
var gamefinish = &room{
	Title: "Escape",
	Desc:  "You escape from the place you were in and run from the people in the car. You escape and get home. The end",
	commands: map[string]action{
		"Title": func(s *state) {
			fmt.Println("\n You Go to the title screen.\n ")
			s.Room = titleRoom
		},
	},
}
