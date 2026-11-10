package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					fmt.Println("Worker", id, "終了")
					return

				case <-ticker.C:
					fmt.Println("Worker", id, "実行中")
				}
			}
		}(i)
	}

	time.Sleep(5 * time.Second)

	cancel()
	wg.Wait()

	fmt.Println("Server Shutdown Complete")
}
