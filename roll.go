package main

import (
	"fmt"
	"os"
)

func main() {
	specs := os.Args[1:]
	detailed := false
	requestedRolls := make([]RollSpec, 5)

	for _, eachSpec := range specs {
		if eachSpec == "-d" {
			detailed = true
			continue
		}

		spec, specErr := Parse(eachSpec)
		if specErr != nil {
			fmt.Printf("%s: roll specification in incorrect format: %s\n", eachSpec, specErr)
		} else {
			requestedRolls = append(requestedRolls, *spec)
		}
	}

	for _, eachRequest := range requestedRolls {
		results := DoRolls(eachRequest)
		for eachRoll := 0; eachRoll < results.Count; eachRoll++ {
			roll := results.Rolls[eachRoll]
			fmt.Printf("%d ", roll.Total)
			if detailed {
				fmt.Printf(":")
				for _, eachDie := range roll.Dies {
					fmt.Printf(" [%d]", eachDie)
				}
			}
			fmt.Printf("\n")
		}
	}
}
