package main

import (
	"fmt"
)

var optionsG map[string]string

var roundG int = 30

func strToInt(s string, defaultValue int) int { _ = "STUB: not implemented"; return 0 }

func main() {
	optionsG := map[string]string{"round": "12", "b": "one"}
	roundG = strToInt(optionsG["round"], 50)
	fmt.Println(roundG)
	fmt.Println(optionsG)
}
