// 2026-10-22
// Go Lang
// select

package main

import (
	"fmt"
	"sync"
)

func main() {
	damageCh := make(chan int)
	healCh := make(chan int)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		damageCh <- 30
		damageCh <- 10
	}()

	go func() {
		defer wg.Done()

		healCh <- 30
		healCh <- 10
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	hp := 100

	for {
		select {
		case damage := <-damageCh:
			hp -= damage
			fmt.Println("Damage", damage)

		case heal := <-healCh:
			hp += heal
			fmt.Println("Heal:", heal)

		case <-done:
			fmt.Println("Final HP:", hp)
			return
		}
	}
}
