package main

import (
	"fmt"
	"os"
	"roller"
)

func main() {
	specs := os.Args[1:]
	for eachSpec := 0; eachSpec < len(specs); eachSpec++ {
		spec, specErr := roller.Parse(specs[eachSpec])
		if specErr != nil {
			fmt.Printf("Roll specification in incorrect format: %s", specErr)
		} else {
			results := roller.DoRolls(*spec)
			for eachRoll := 0; eachRoll < results.Count; eachRoll++ {
				roll := results.Rolls[eachRoll]
				fmt.Printf("%d ", roll.Total)
			}
			fmt.Printf("\n")
		}
	}
}
