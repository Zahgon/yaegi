package main

import (
	"fmt"
	"runtime"
	"sync"
)

func humanizeBytes(bytes uint64) string { _ = "STUB: not implemented"; return "" }

func main() {
	i := 0
	wg := sync.WaitGroup{}

	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("#%d: alloc = %s, routines = %d, gc = %d\n", i, humanizeBytes(m.Alloc), runtime.NumGoroutine(), m.NumGC)

		wg.Add(1)
		go func() {
			wg.Done()
		}()
		wg.Wait()
		i = i + 1
	}
}
