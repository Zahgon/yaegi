package main

import (
	"fmt"
	"time"
)

func get10Hours() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func main() {
	fmt.Println(get10Hours().String())
}
