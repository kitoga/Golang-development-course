package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkRoomIsClean(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {

	//give some value to the struct
	rooms := [...]Room{
		Room{name: "bedroom"},
		Room{name: "kitchen"},
		Room{name: "ops"},
		Room{name: "reception"},
	}

	checkRoomIsClean(rooms)

	fmt.Println("cleaning taking place...")
	rooms[2].cleaned = true
	rooms[0].cleaned = true

	checkRoomIsClean(rooms)
}
