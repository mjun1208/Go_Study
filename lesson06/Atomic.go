// 2026-10-23
// Go Lang
// Atomic

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Player struct {
	HP atomic.Int64
}

func (p *Player) Damage(amount int64) {
	p.HP.Add(-amount)
}

func (p *Player) Heal(amount int64) {
	p.HP.Add(amount)
}

func (p *Player) GetHP() int64 {
	return p.HP.Load()
}

func Atomic() {
	player := &Player{}
	player.HP.Store(1000)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			player.Damage(1)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			player.Heal(1)
		}
	}()

	wg.Wait()

	fmt.Println("HP : ", player.GetHP())
}
