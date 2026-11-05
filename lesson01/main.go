// 2026-10-22
// Go Lang
// struct
// interface
// value

package main

import (
	"fmt"
)

type Player struct {
	ID    int
	Name  string
	HP    int
	Items []string
}

const MAXHP int = 100

func (p *Player) Damage(amount int) error {
	if amount < 0 {
		return fmt.Errorf("Invalid Damage: %d", amount)
	}

	p.HP -= amount

	if p.HP < 0 {
		p.HP = 0
	}

	return nil
}

func (p *Player) Heal(amount int) error {
	if amount < 0 {
		return fmt.Errorf("Invalid Heal: %d", amount)
	}

	p.HP += amount

	if p.HP > MAXHP {
		p.HP = MAXHP
	}

	return nil
}

type Damageable interface {
	Damage(int) error
}

func Attack(target Damageable, amount int) error {
	return target.Damage(amount)
}

func FindPlayer(players map[int]*Player, id int) (*Player, error) {
	p, ok := players[id]

	if !ok {
		return nil, fmt.Errorf("Player ID : %d Not Found", id)
	}

	return p, nil
}

func main() {
	p := &Player{
		ID:   1,
		Name: "Player1",
		HP:   MAXHP,
	}

	// 1番にｐを入れる
	players := map[int]*Player{p.ID: p}

	if err := Attack(p, 25); err != nil {
		fmt.Println("Damage Error: ", err)
	}

	p.Heal(15)

	p.Items = append(p.Items, "Potion", "Sword")

	for i, item := range p.Items {
		fmt.Printf("Item %d: %s\n", i, item)
	}

	found, err := FindPlayer(players, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s HP: %d\n", found.Name, found.HP)
}
