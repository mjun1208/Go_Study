// 2026-11-05
// Slice (copy)

package main

import (
	"fmt"
)

type Player struct {
	Items []string
}

const MaxItems = 5

func (p *Player) AddItem(item string) error {
	if len(p.Items) >= MaxItems {
		return fmt.Errorf("inventory full")
	}

	p.Items = append(p.Items, item)

	return nil
}

func (p *Player) GetItems() []string {
	result := make([]string, len(p.Items))
	copy(result, p.Items)
	return result

	// Cloneもできる
	// return slices.Clone(p.Items)
}

func main() {
	player := &Player{}

	for _, item := range []string{
		"Sword", "Potion", "Shield",
		"Armor", "Ring", "Extra",
	} {
		if err := player.AddItem(item); err != nil {
			fmt.Println("Error:", err)
		}
	}

	// Blocked
	items := player.GetItems()
	items[0] = "Hacked"

	fmt.Println("Original:", player.GetItems())
}
