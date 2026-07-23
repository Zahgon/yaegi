package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		print("test")
	}()

	wg.Wait()
}

func print(state string) { _ = "STUB: not implemented"; return }
