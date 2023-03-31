package main

import (
	"clippy/brain"
	"flag"
	"fmt"
)

const clippy = `/‾‾\
|  |
O  O
|| |/
|| ||
|\_/|
\___/`

func main() {
	// check args
	goal := flag.String("goal", "Run a minecraft server", "What would you like clippy to help you with today?")
	flag.Parse()

	brain, err := brain.NewBrain("gpt-3.5-turbo")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// print clippy
	fmt.Println(clippy)

	fmt.Printf("\nHello there.\n")
	intepretedGoal, err := brain.InterpretGoal(*goal)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("%s\n", intepretedGoal)

}
