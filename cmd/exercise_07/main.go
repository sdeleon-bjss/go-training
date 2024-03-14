package main

import (
	"fmt"
	"github.com/sdeleon-bjss/pkg"
)

// Exercise 07:
// Create a program that rolls two dice (1 to 6) fifty times. Display the number rolls and the outcomes in sequential order.
// Resulting rolls are to be processed in the following manner: [Random Numbers][Switches]
// - 7 and 11 are to be called NATURAL
// - 2 is called SNAKE-EYES-CRAPS
// - 3 and 12 is called LOSS-CRAPS
// - Any other combination is called NEUTRAL.
func main() {
	for i := 0; i < 50; i++ {
		rolls := pkg.ComboRolls()
		rollsCombined := rolls[0] + rolls[1]
		rollResult := pkg.ProcessRoll(rollsCombined)

		output := fmt.Sprintf("Roll attempt #%d: You rolled a (%d) = %s", i+1, rollsCombined, rollResult)

		fmt.Println(output)
	}
}
