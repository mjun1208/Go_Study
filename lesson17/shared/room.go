package game

import "sync"

type Room struct {
	mu      sync.Mutex
	players map[int64]string
}

func NewRoom() *Room {
	return &Room{
		players: make(map[int64]string),
	}
}

func (r *Room) Join(id int64, name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.players[id] = name
}

func (r *Room) Leave(id int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.players, id)
}

func (r *Room) Count() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return len(r.players)
}
