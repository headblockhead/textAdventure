package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	//see cls.go for definition
	Cls()
	//defines all of the commands for each location (room) in one chunk
	mainPath0.commands["west"] = func(s *state) {
		Cls()
		fmt.Println("You go west")
		s.room = mRoom
		s.gateEntered = true
		s.RoomNo = 4
		s.HiddenCommands["Mausoleum/go through tunnel"] = struct{}{}
	}
	mainPath0.commands["north"] = func(s *state) {
		Cls()
		fmt.Println("You Move north")
		s.gateEntered = true
		s.room = mainPath2
		s.RoomNo = 7
	}

	mRoom.commands["east"] = func(s *state) {
		Cls()
		fmt.Println("You go east")
		s.room = mainPath0
		s.RoomNo = 3
	}
	mRoom.commands["pick up hammer"] = func(s *state) {
		Cls()
		fmt.Println("You Pick up the hammer")
		s.Hammergot = true
		s.HiddenCommands["Mausoleum/pick up hammer"] = struct{}{}
	}

	mainPath2.commands["east"] = func(s *state) {
		Cls()
		fmt.Println("You go east")
		s.HiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
		s.room = gTRoom
		s.RoomNo = 5
	}

	mainPath2.commands["south"] = func(s *state) {
		Cls()
		fmt.Println("You go south")
		s.room = mainPath0
		s.RoomNo = 3
	}
	mainPath2.commands["north"] = func(s *state) {
		Cls()
		fmt.Println("You go north")
		s.room = mainPath3
		s.RoomNo = 9
	}

	gTRoom.commands["west"] = func(s *state) {
		Cls()
		fmt.Println("You go west")
		s.room = mainPath2
		s.RoomNo = 7
	}
	gTRoom.commands["pick up safe code"] = func(s *state) {
		Cls()
		fmt.Println("You pick up the safe code")
		s.Safecodegot = true
		s.HiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
	}
	mainPath3.commands["south"] = func(s *state) {
		Cls()
		fmt.Println("You go south")
		s.room = mainPath2
		s.RoomNo = 7
	}
	mainPath3.commands["west"] = func(s *state) {
		Cls()
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
		fmt.Println("You go west")
		s.room = maroom
		s.RoomNo = 6
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
	}
	mainPath3.commands["east"] = func(s *state) {
		Cls()
		fmt.Println("You go east")
		s.room = gRoom
		s.RoomNo = 8
	}
	gRoom.commands["turn on power"] = func(s *state) {
		Cls()
		fmt.Println("You turn on the power")
		s.Electricity = true
		s.HiddenCommands["Generator/turn on power"] = struct{}{}
	}
	gRoom.commands["west"] = func(s *state) {
		Cls()
		fmt.Println("You go west")
		s.room = mainPath3
		s.RoomNo = 9
	}
	maroom.commands["south"] = func(s *state) {
		Cls()
		fmt.Println("You go south")
		s.room = mainPath3
		s.RoomNo = 9
	}
	mRoom.commands["north"] = func(s *state) {
		Cls()
		fmt.Println("You go through the tunnel")
		s.room = mABRoom
		s.RoomNo = 14
	}
	mABRoom.commands["pull lever"] = func(s *state) {
		Cls()
		fmt.Println("You pull the lever and rocks fall into a pit in the center of the room")
		s.room = mABRoom
		s.RoomNo = 14
		s.RocksFallen = true
		s.HiddenCommands["Basement/pull lever"] = struct{}{}
	}
	mABRoom.commands["south"] = func(s *state) {
		Cls()
		fmt.Println("You go back to the mausoleum")
		s.room = mRoom
		s.RoomNo = 4
		s.HiddenCommands["Mausoleum/go through tunnel"] = struct{}{}
	}
	mainPath3.commands["north"] = func(s *state) {
		Cls()
		fmt.Println("You go north")
		s.room = mainpath5
		s.RoomNo = 10
	}
	maroom.commands["grab key"] = func(s *state) {
		Cls()
		fmt.Println("You grab the key")
		s.KeyGot = true
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
		delete(s.HiddenCommands, "mainpath5/exit")
	}
	mainpath5.commands["exit"] = func(s *state) {
		Cls()
		fmt.Println("You leave.")
		s.room = gamefinish
		s.RoomNo = 12

	}
	mainpath5.commands["south"] = func(s *state) {
		Cls()
		fmt.Println("You go south")
		s.room = mainPath3
		s.RoomNo = 9
	}
	mainpath5.commands["west"] = func(s *state) {
		Cls()
		fmt.Println("You go west")
		s.room = bRoom
		s.RoomNo = 11
	}
	bRoom.commands["switch off"] = func(s *state) {
		Cls()
		fmt.Println("You switch off the electromagnet")
		s.BreakerRoomUsed = true
		s.HiddenCommands["Breaker Room/switch off"] = struct{}{}
	}
	bRoom.commands["east"] = func(s *state) {
		Cls()
		fmt.Println("You go east")
		s.room = mainpath5
		s.RoomNo = 10
	}
	s := &state{
		room: titleRoom,
		HiddenCommands: map[string]struct{}{
			"End of path/exit": {},
		},
		Movestaken: 0,
	}
	//creates a reader to read input from the terminal
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			//increase time if the game is not paused
			if s.room != titleRoom && s.room != gamefinish && s.room != stats && s.room != pause {
				s.Time++
			} else if s.room == titleRoom {
				s.Time = 0
			}
		}
	}()

	for {
		//renderRoom takes a room and prints the contents to the screen in a user friendly way
		renderRoom(s.room, s)
		fmt.Print("> ")
		//reads terminal input until a new line (enter being pressed)
		text, _ := reader.ReadString('\n')
		action, ok := s.room.commands[strings.TrimSpace(text)]
		//special commands:
		if strings.EqualFold(strings.TrimSpace(text), "quit") && s.room == titleRoom {
			if s.room != stats {
				Cls()
				fmt.Println("\n You Quit the game.\n ")
				os.Exit(0)
			}
		} else if strings.EqualFold(strings.TrimSpace(text), "quit") && (s.room == pause || s.room == leftStartPath) {
			if s.room != stats {
				Cls()
				fmt.Println("\n You Quit the game.\n ")
				os.Exit(0)
			}
		} else if strings.EqualFold(strings.TrimSpace(text), "save") && s.room != titleRoom && s.room != leftStartPath && s.room == pause {
			Cls()
			save(s)
			fmt.Println("\n You save the game.\n ")
			Cls()
		} else if ok && !commandIsHidden(strings.TrimSpace(text), s) {
			s.Movestaken++
			action(s)
		} else if strings.EqualFold(strings.TrimSpace(text), "escape") && s.room == mainpath5 {
			//renders escape message
			fmt.Println("\n You escape the area")
			fmt.Println(gamefinish.Title)
			fmt.Println()
			fmt.Println(gamefinish.Desc)
			//exits the program after 10 seconds
			time.Sleep(10 * time.Second)
			//exits with an error of 0 (no error)
			os.Exit(0)
		} else if strings.EqualFold(strings.TrimSpace(text), "stats") && s.room != titleRoom && s.room != stats && s.room == pause {
			Cls()
			s.room = stats
			s.onpause = false
		} else if strings.EqualFold(strings.TrimSpace(text), "go back") && (s.room == stats || s.room == pause) {
			if s.onpause == true {
				s.onpause = false
				s.room = getRoomFromR(s.RoomNo)
			} else {
				s.room = pause
				s.onpause = true
			}

		} else if strings.EqualFold(strings.TrimSpace(text), "quit to title") && s.room == leftStartPath || (s.room != titleRoom && s.room == pause && strings.EqualFold(strings.TrimSpace(text), "quit to title")) {
			s.room = titleRoom
			s.RoomNo = 0
		} else if strings.EqualFold(strings.TrimSpace(text), "pause") && s.room != titleRoom && s.room != pause && s.room != leftStartPath && s.room != stats {
			s.room = pause
			s.onpause = true
		} else {
			//if the command is not a special command or a standard room command, tell the user
			Cls()
			fmt.Println()
			fmt.Println("The command You have entered is not a command that is possible!")
		}

	}
}

//if the bool is true the return "yes" if false, "no" if true nor false, force quit and print debug info
func trueOrFalse(b bool) (s string) {
	if b {
		return "Yes"
	} else if !b {
		return "No"
	} else {
		panic(1)
	}
}

func ftfioi(t int) (ft string) {
	secs := time.Duration(time.Duration(t) * time.Second)
	return secs.String()
}

type state struct {
	//important info to save between playthroughs
	room            *room
	Hammergot       bool
	Safecodegot     bool
	Electricity     bool
	BreakerRoomUsed bool
	RocksFallen     bool
	KeyGot          bool
	RoomNo          int
	HiddenCommands  map[string]struct{}
	Movestaken      int
	Time            int
	gateEntered     bool
	onpause         bool
}

type action func(s *state)

//defines what a room can contain
type room struct {
	Title     string
	Desc      string
	commands  map[string]action
	stateDesc func(s *state) string
}

//renderRoom takes all aspects of a room and prints them to the screen
// it also prints all available commands from the list of commands output by getCommands()
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
	fmt.Println(strings.Join(getCommands(r.commands, s), " | "))
	fmt.Println()
}

//if a command is in the list of hidden commands, don't show it to the user
func commandIsHidden(cmd string, s *state) bool {
	_, ok := s.HiddenCommands[s.room.Title+"/"+cmd]
	return ok
}

//gets commands from the room & add special commands if necessary
func getCommands(m map[string]action, s *state) (commands []string) {
	if s.room == titleRoom {
		commands = append(commands, "quit")
	}
	if s.room == pause {
		commands = append(commands, "quit")
	}
	if s.room == leftStartPath {
		commands = append(commands, "quit")
	}
	if s.room != titleRoom && s.room == pause {
		commands = append(commands, "save")
	}
	if s.room == mainpath5 && s.KeyGot == true {
		commands = append(commands, "escape")
	}
	if s.room == stats || s.room == pause {
		commands = append(commands, "go back")
	}
	if s.room != titleRoom && s.room == pause {
		commands = append(commands, "stats")
	}
	if s.room != titleRoom && s.room == pause {
		commands = append(commands, "quit to title")
	}
	if s.room != titleRoom && s.room != pause && s.room != leftStartPath {
		commands = append(commands, "pause")
	}
	for k := range m {
		if _, ok := s.HiddenCommands[s.room.Title+"/"+k]; !ok {
			commands = append(commands, k)

		}
	}
	sort.Strings(commands)
	return
}

//takes the room number stored in a save file or the state and sets the room to the appropriate value
func getRoomFromR(r int) *room {
	if r == 1 {
		var n = startRoom
		return n
	} else if r == 2 {
		var n = leftStartPath
		return n
	} else if r == 3 {
		var n = mainPath0
		return n
	} else if r == 4 {
		var n = mRoom
		return n
	} else if r == 5 {
		var n = gTRoom
		return n
	} else if r == 6 {
		var n = maroom
		return n
	} else if r == 7 {
		var n = mainPath2
		return n
	} else if r == 8 {
		var n = gRoom
		return n
	} else if r == 9 {
		var n = mainPath3
		return n
	} else if r == 10 {
		var n = mainpath5
		return n
	} else if r == 11 {
		var n = bRoom
		return n
	} else if r == 12 {
		var n = gamefinish
		return n
	} else if r == 14 {
		var n = mABRoom
		return n
	}
	return startRoom
}

//the room where you start
var startRoom = &room{
	Title: "Road",
	Desc:  "You are running on a road. You don't know why you are here, and doubt you ever will know. 3 men in a vehicle are chasing you and you come accross a split in the path. You see a long winding path to the west, and an entrance to what seems like an abandoned graveyard to the east. The entrance to the graveyard has a metal gate which you can lock, But once you are inside, there seems to be no way out.\n",
	commands: map[string]action{
		"west": func(s *state) {
			Cls()
			fmt.Println("\n You go west.\n ")
			s.room = leftStartPath
			s.RoomNo = 2
		},
		"east": func(s *state) {
			Cls()
			fmt.Println("\n You go east\n ")
			s.gateEntered = false
			s.room = mainPath0
			s.RoomNo = 3
		},
	},
}

//path to the left of the start
var leftStartPath = &room{
	Title:    "West Path",
	Desc:     " You run as fast as you can along the path going west. You notice the group in the car continue approching. You are fast, but not fast enough, the car stops and the men get out. They seem to want to kill you. You cannot escape as they drag you into their vehicle. \n You died.\n",
	commands: map[string]action{},
}

// the middle path at the first pos
var mainPath0 = &room{
	Title: "East Path",
	stateDesc: func(s *state) string {
		general := " You run as fast as you can along the path going east. You dive into the enclosed space and lock the gate. You're safe for the moment, but you know that they will wait for you to come out from that gate, no matter what. It is getting dark now and you pull out your lantern. \n You look around and see a path to the west, There is also a path north.\n"
		general1 := "You look around and see a path going west, There is also a path north."
		if s.gateEntered == true {
			s.gateEntered = true
			return general1
		}
		return general
	},
	commands: map[string]action{},
}

//the mausoleum
var mRoom = &room{
	Title:    "Mausoleum",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You go west and walk along the path towards a building that looks like an old mausoleum. You look around yourself and notice graves in the woods some with bones still half sticking out. You begin to wonder what you got yourself into as you enter.\n"
		if s.Hammergot == true {
			delete(s.HiddenCommands, "Mausoleum/go through tunnel")
			return general + "There used to be a hammer on the floor, but you picked it up. A trapdoor opened when you removed the hammer."
		}

		return general + "There is a hammer on the floor."
	},
}

// the middle path at the second pos
var mainPath2 = &room{
	Title:    "Path",
	Desc:     "You walk to the next intersection on the path and see a guard tower to the east. It stretches far into the sky but has no window. You can continue walking north, or take the path towards it.",
	commands: map[string]action{},
}

//the guard tower room
var gTRoom = &room{
	Title:    "Guard Tower",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk towards the guard tower, the night falling around you, and see that"
		if s.Electricity == true && s.BreakerRoomUsed == true && s.Safecodegot == false {
			delete(s.HiddenCommands, "Guard Tower/pick up safe code")
			return general + " the door is unlocked, You enter and find a metal block on the floor with a safe code sticking out of it."
		}
		if s.Electricity == true && s.Safecodegot == false {

			return general + " the door is unlocked, You enter and find a metal block on the floor with a safe code that is currently unreadable sticking out of it. You cannot lift it as an electromagnet has it stuck to the floor."
		} else if s.Safecodegot == true {
			return general + " the door is unlocked, You enter and find a metal block on the floor. The safe code that was once here is now gone"
		}

		return general + " the door is locked, it has a lock that requires electricity to function."
	},
}

// the middle path at the third pos
var mainPath3 = &room{
	Title:    "Path",
	Desc:     "You walk to the next intersection on the path and look around. You look west and see an abandoned mansion, covered in winding vines. To the east you see a small house.",
	commands: map[string]action{},
}

//the main puzzle room with the emergency exit key
var maroom = &room{
	Title:    "Mansion",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk along the short path, observing your surroundings. You notice the distant screams coming from the car that was chasing you. Shivers make their way down your spine as you peek inside the old building. Inside, you notice"
		if s.RocksFallen == true && s.Safecodegot == true && s.Hammergot == true && s.KeyGot == false {
			delete(s.HiddenCommands, "Mansion/grab key")
			return general + " a key guarded by a safe. You smash the glass, unlock the safe with the code you found and see a key with \"Escape\" written on it."
		}
		if s.RocksFallen == true {
			if s.Hammergot {
				return general + " a safe guarded by a sheet of glass. You smash the glass using a hammer to reveal a safe that appears to need a code."
			}
			return general + " a safe needing a code guarded by a sheet of glass."
		}

		return general + " a pile of rocks that have seemingly fallen from the celling, blocking your path."
	},
}

//generator room
var gRoom = &room{
	Title:    "Generator",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You enter the room and comfirm your suspicions, this was a generator room. The generator was labled \"Backup power \" and connected (via a switch) to the rest of the graveyard. The switch is set to "
		if !s.Electricity {
			delete(s.HiddenCommands, "Generator/turn on power")
			return general + "off."
		}

		return general + "on."
	},
}

//mansion basement room
var mABRoom = &room{
	Title:    "Basement",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You crawl through the trapdoor and spy a tunnel snaking off past your view. You enter the tunnel and reach the basement of a Mansion. There is a lever seemingly controlling a trapdoor in the floor above. The lever is "
		if !s.RocksFallen {
			delete(s.HiddenCommands, "Basement/pull lever")
			return general + "up (not flipped)."
		}

		return general + "down (flipped)."
	},
}

// the middle path at the last pos
var mainpath5 = &room{
	Title:    "End of path",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk along the path and reach a towering gate labled \"Exit\". This appears to be your way out of the horrific place you trapped yourself in. You can turn west."
		if s.KeyGot {
			return general + " You have the key for the gate."
		}
		return general
	},
}

//the room contaning the power breakers
var bRoom = &room{
	Title:    "Breaker Room",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You trudge along the path, turning a corner you spy a room. You enter the room and see an array of confusing buttons, dials and switches. It appears to be a breaker room. You a switch labeled \"Electromagnet\"."
		if !s.BreakerRoomUsed {
			delete(s.HiddenCommands, "Breaker Room/switch off electromagnet")
		}
		return general
	},
}

// the title screen
var titleRoom = &room{
	Title: "Eddie's Text Adventure",
	Desc:  "MENU",
	commands: map[string]action{
		"start": func(s *state) {
			Cls()
			s.room = startRoom
			s.RoomNo = 1
		},
		"load": func(s *state) {
			Cls()
			for {
				ok, isQuit, err := load(s)
				if err != nil {
					fmt.Println(err)
					continue
				}
				if isQuit {
					break
				}
				if ok {
					s.room = getRoomFromR(s.RoomNo)
					break
				}
				fmt.Println("That is not a valid savefile, please try another one")
			}
		},
	},
}

//the last, inaccessible room, printed outside of the main script after finishing the game
var gamefinish = &room{
	Title:    "Escape",
	Desc:     "You dash out of the gate and launch yourself from the Graveyard, looking all over for any signs of the shadowy figures. They had not expected you to exit though this way. You sprint to your car and drive away before any of them get wiser. You Win!",
	commands: map[string]action{},
}

//shows statistics about the game
var stats = &room{
	Title:    "Stats",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		var Info1 = "Moves Taken: "
		var Info2 = "Time Spent Playing: "
		var output = Info1 + strconv.Itoa(s.Movestaken)
		output = output + "\n" + Info2
		output = output + ftfioi(s.Time)
		return output
	},
}

//pause screen
var pause = &room{
	Title:    "Pause Menu",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		var output = "The Game is paused."
		return output
	},
}
