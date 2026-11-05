// 2026-10-22
// Go Lang
// gorutine
// channel
// WaitGroup

package main

import (
	"fmt"
	"sync"
)

type Player struct {
	HP int
}

type DamageEvent struct {
	AttackerID int
	Damage     int
}

func main() {
	player := &Player{HP: 100}

	damageCh := make(chan DamageEvent)
	var wg sync.WaitGroup

	wg.Add(2)

	// Monster A
	go func() {
		defer wg.Done()

		damageCh <- DamageEvent{1, 20}
		damageCh <- DamageEvent{1, 10}
	}()

	// Monster B
	go func() {
		defer wg.Done()

		damageCh <- DamageEvent{2, 15}
		damageCh <- DamageEvent{2, 5}
	}()

	// Close Channel
	go func() {
		wg.Wait()
		close(damageCh)
	}()

	for event := range damageCh {
		player.HP -= event.Damage

		fmt.Printf("Monster %d attacked! HP: %d\n", event.AttackerID, player.HP)
	}

	fmt.Println("Final HP: ", player.HP)
}
