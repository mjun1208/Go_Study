// 2026-10-27
// Go Lang
// Slice
// Map

package main

import "fmt"

type Player struct {
	ID    int
	Name  string
	Items []string
}

type Room struct {
	Players map[int]*Player
}

func NewRoom() *Room {
	return &Room{
		Players: make(map[int]*Player),
	}
}

func (room *Room) Join(player *Player) error {
	if player == nil {
		return fmt.Errorf("player is nil")
	}

	if _, exists := room.Players[player.ID]; exists {
		return fmt.Errorf("duplicate player ID: %d", player.ID)
	}

	room.Players[player.ID] = player
	return nil
}

func (room *Room) Leave(id int) {
	delete(room.Players, id)
}

func (room *Room) GetPlayer(id int) (*Player, bool) {
	player, ok := room.Players[id]
	return player, ok
}

func (room *Room) AddItem(id int, item string) {
	player, ok := room.GetPlayer(id)

	if !ok {
		return
	}

	player.Items = append(player.Items, item)
}

func main() {
	room := NewRoom()

	room.Join(&Player{ID: 1, Name: "Reimu"})
	room.Join(&Player{ID: 2, Name: "Marisa"})

	room.AddItem(1, "Potion")
	room.AddItem(1, "Sword")

	p, ok := room.GetPlayer(1)

	if ok {
		fmt.Println("Name:", p.Name)
		fmt.Println("Items:", p.Items)
	}

	room.Leave(2)

	fmt.Println("Player Count:", len(room.Players))
}
