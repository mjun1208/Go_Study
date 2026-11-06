package player

import (
	"fmt"
	"slices"
)

const MaxItems = 5

type Player struct {
	ID    int
	Name  string
	items []string
}

func New(id int, name string) *Player {
	return &Player{
		ID:   id,
		Name: name,
	}
}

func (p *Player) AddItem(item string) error {
	if len(p.items) >= MaxItems {
		return fmt.Errorf("inventory full")
	}

	p.items = append(p.items, item)
	return nil
}

func (p *Player) GetItems() []string {
	return slices.Clone(p.items)
}
