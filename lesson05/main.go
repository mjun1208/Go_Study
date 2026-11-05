// 2026-10-23
// Go Lang
// time

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const N = 200000

	hp := 0
	var mu sync.Mutex

	increment := func() {
		mu.Lock()
		hp++
		mu.Unlock()
	}

	// 1. 일반 for문
	start := time.Now()

	for i := 0; i < N; i++ {
		increment()
	}

	elapsed := time.Since(start)
	fmt.Println("For:", elapsed, "HP:", hp)

	// 2. 고루틴 2개
	hp = 0
	var wg sync.WaitGroup
	wg.Add(2)

	start = time.Now()

	for j := 0; j < 2; j++ {
		go func() {
			defer wg.Done()

			for i := 0; i < N/2; i++ {
				increment()
			}
		}()
	}

	wg.Wait()

	elapsed = time.Since(start)
	fmt.Println("Go 2:", elapsed, "HP:", hp)

	// 3. 고루틴 2000개
	hp = 0
	wg.Add(2000)

	start = time.Now()

	for j := 0; j < 2000; j++ {
		go func() {
			defer wg.Done()

			for i := 0; i < N/2000; i++ {
				increment()
			}
		}()
	}

	wg.Wait()

	elapsed = time.Since(start)
	fmt.Println("Go 2000:", elapsed, "HP:", hp)
}
