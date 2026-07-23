package main

import (
	"fmt"
	"math/big"
)

func main() {

	n := big.NewInt(0)
	n.SetString("273966616513101251352941655302036077733021013991", 10)

	i := big.NewInt(0)

	i.SetString("496968652506233112158689", 10)

	temp := big.NewInt(0)

	two := big.NewInt(2)
	for {

		temp.Mod(n, i)
		if temp.Sign() == 0 {
			fmt.Println(i)
			break
		}

		i.Add(i, two)
	}
}
