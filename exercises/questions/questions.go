package questions

import (
	"fmt"
	"math/rand"
)

// QuestionOne performs addition without using arithmetic operators
func QuestionOne(x, y int64) int64 {
	/*
		 No arithmetic, so binary operations required
		 addition:
			 0 + 0 = 00
			 0 + 1 = 01
			 1 + 0 = 01
			 1 + 1 = 10
		can split addition into adding the result without carrying one
		with the result of carrying the one, without the addition
		this is adding an XOR of the two numbers with and AND shifted by one bit
		101 + 011 = 110 + 010 = 100 + 100 = 000 + 1000

		Alternatively you can iterate through the bits in increasing order of significance and perform AND, carrying the one to the next bit if required.
	*/

	if y == 0 {
		return x
	}
	addition := x ^ y
	carry := (x & y) << 1
	return QuestionOne(addition, carry)
}

const (
	Spade = iota
	Club
	Diamond
	Heart
)

type Card struct {
	suit  int
	value int
}

func (c Card) Print() string {
	switch c.suit {
	case Spade:
		return fmt.Sprintf("S%d ", c.value)
	case Club:
		return fmt.Sprintf("C%d ", c.value)
	case Diamond:
		return fmt.Sprintf("D%d ", c.value)
	case Heart:
		return fmt.Sprintf("H%d ", c.value)
	default:
		return "nope"
	}
}

// QuestionTwo performs a shuffle of a deck of cards, using a perfect random number generator
func QuestionTwo(deck []*Card, position int) []*Card {
	/*
		There are 52! permutations of a deck, which are assembled by choosing randomly without replacement from the initial deck
	*/
	if position == 0 {
		return deck
	}
	deck = QuestionTwo(deck, position-1)
	randomPosition := random(position)
	temp := deck[randomPosition]
	deck[randomPosition] = deck[position]
	deck[position] = temp
	return deck
}

func initDeck() []*Card {
	deck := []*Card{}
	for _, suit := range []int{Spade, Club, Diamond, Heart} {
		for i := 1; i < 14; i++ {
			deck = append(deck, &Card{suit: suit, value: i})
		}
	}
	return deck
}

// random returns a random int between the 0 and x
func random(x int) int64 {
	return rand.Int63n(int64(x))
}

// QuestionThree generates a random subset of the provided set
func QuestionThree(set []int, count int) []int {
	subset := make([]int, count)
	copy(subset, set[:count])

	for i := count; i < len(set); i++ {
		random := random(len(set))
		if random < int64(count) {
			subset[random] = set[i]
		}
	}
	return subset
}
