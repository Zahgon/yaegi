package main

import (
	"fmt"
	"time"
)

const df = time.Minute * 30

func main() {
	fmt.Printf("df: %v %T\n", df, df)
}
