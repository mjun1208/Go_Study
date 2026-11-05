// 2026-10-23
// Go Lang
// CompareAndSwap

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Player struct {
	HP atomic.Int64
}

func (p *Player) Damage(amount int64) error {
	if amount < 0 {
		return fmt.Errorf("Invalid Damage : %d", amount)
	}

	for {
		old := p.HP.Load()

		newHP := old - amount

		if newHP < 0 {
			newHP = 0
		}

		hpChanged := p.HP.CompareAndSwap(old, newHP)

		if hpChanged {
			break
		}
	}

	return nil
}

func main() {
	player := &Player{}
	player.HP.Store(100)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		_ = player.Damage(80)
	}()

	go func() {
		defer wg.Done()
		_ = player.Damage(100)
	}()

	wg.Wait()
	fmt.Println("Final HP : ", player.HP.Load())
}
