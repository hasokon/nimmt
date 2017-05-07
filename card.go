package nimmt

import (
	"fmt"
)

// Card express a Nimmt Card
type Card struct {
	Number int
	Cows   int
}

func (c Card) String() string {
	cows := []byte{}
	for i := 0; i < c.Cows; i++ {
		cows = append(cows, byte('*'))
	}

	return fmt.Sprintf("%d (%s)", c.Number, cows)
}
