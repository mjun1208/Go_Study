package room

import (
	"fmt"

	"example.com/lesson10/player"
)

type Room struct {
	players map[int]*player.Player
}

func New() *Room {
	return &Room{
		players: make(map[int]*player.Player),
	}
}

func (r *Room) Join(p *player.Player) error {
	if p == nil {
		return fmt.Errorf("player is nil")
	}

	if _, ok := r.players[p.ID]; ok {
		return fmt.Errorf("player %d already joined", p.ID)
	}

	r.players[p.ID] = p

	return nil
}

func (r *Room) Leave(id int) {
	delete(r.players, id)
}

func (r *Room) GetPlayer(id int) (*player.Player, bool) {
	player, ok := r.players[id]
	return player, ok
}

func (r *Room) Count() int {
	return len(r.players)
}
