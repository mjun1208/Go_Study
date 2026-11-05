// 2026-10-23
// Go Lang
// Mutex

package main

import (
	"fmt"
	"sync"
)

type Player struct {
	HP int
	mu sync.Mutex
}

func (p *Player) Damage(amount int) error {
	if amount < 0 {
		return fmt.Errorf("Damage Error")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.HP -= amount
	return nil
}

func (p *Player) Heal(amount int) error {
	if amount < 0 {
		return fmt.Errorf("Heal Error")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.HP += amount
	return nil
}

func (p *Player) GetHP() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.HP
}

func main() {
	player := &Player{HP: 1000}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			player.Damage(1)
		}

		// deferも使える
		// mu.Unlock()
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			player.Heal(1)
		}

		// deferも使える
		// mu.Unlock()
	}()

	wg.Wait()

	fmt.Println("Player HP:", player.GetHP())
}
