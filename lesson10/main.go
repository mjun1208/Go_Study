package main

import (
	"fmt"

	"example.com/lesson10/player"
	"example.com/lesson10/room"
)

func main() {
	gameRoom := room.New()

	p1 := player.New(1, "Reimu")
	p2 := player.New(2, "Marisa")

	if err := gameRoom.Join(p1); err != nil {
		fmt.Println(err)
	}

	if err := gameRoom.Join(p2); err != nil {
		fmt.Println(err)
	}

	_ = p1.AddItem("Potion")
	_ = p1.AddItem("Sword")

	fmt.Println("Players:", gameRoom.Count())

	gameRoom.Leave(2)

	fmt.Println("Players:", gameRoom.Count())
}
