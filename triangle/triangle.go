package triangle

import (
	"fmt"
)

//Display Triangle upside down
func Display() {

	upperBound := 5
	outerCounter := 0

	for outerCounter <= upperBound {
		for innerCounter := 0; innerCounter < upperBound; innerCounter++ {
			if innerCounter >= outerCounter {
				fmt.Print("* ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		outerCounter = outerCounter + 1
	}
}
