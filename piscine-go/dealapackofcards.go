package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	players := 4
	cardsPerPlayer := 3

	for i := 0; i < players; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < cardsPerPlayer; j++ {
			card := deck[i*cardsPerPlayer+j]
			if j == cardsPerPlayer-1 {
				fmt.Printf("%d", card)
			} else {
				fmt.Printf("%d, ", card)
			}
		}
		fmt.Printf("\n")
	}
}
