package main

import "fmt"

func main() {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 1; i < 10; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Printf("%v\n", dp)
}
