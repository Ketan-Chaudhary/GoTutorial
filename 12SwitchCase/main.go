package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Go")

	rand.NewSource(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("The Dice Value is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1 Nice")
	case 2:
		fmt.Println("You got 2 move 2 place")
	case 3:
		fmt.Println("You got 3 move 3 place")
		fallthrough // Now the next case will also run
	case 4:
		fmt.Println("You got 4 move 4 place")
	case 5:
		fmt.Println("You got 5 move 5 place")
	case 6:
		fmt.Println("Lucky Roll the Dice Again")
	default:
		fmt.Println("Something bizzard Happens")
	}
}
