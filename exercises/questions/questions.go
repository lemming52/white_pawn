package questions

import (
	"fmt"
	"math/bits"
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
	for i := len(deck) - 1; i > 1; i-- {
		k := random(i)
		temp := deck[k]
		deck[k] = deck[i]
		deck[i] = temp
	}
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

// Question Four,
func QuestionFour(list []*IntMod, column int64) int64 {
	/*
		An array contains integers from 0 to N, except one is missing
		the integers cannot be accessed directly, but only by the jth bit of array[i]
		find the missing integer in O(n) time

		the missing number will be revealed based on
	*/
	if column >= bits.UintSize {
		return 0
	}

	zeroes := []*IntMod{}
	ones := []*IntMod{}

	for _, i := range list {
		if i.Get(column) {
			// column bit is 1
			ones = append(ones, i)
		} else {
			zeroes = append(zeroes, i)
		}
	}

	if len(zeroes) > len(ones) {
		return (QuestionFour(ones, column+1) << 1) | 1
	} else {
		return (QuestionFour(zeroes, column+1) << 1) | 0
	}
}

type IntMod struct {
	value int64
}

func (i *IntMod) Get(j int64) bool {
	return bits.OnesCount(uint((1<<j)&i.value)) > 0
}

/*
QuestionFive

Given an array of A and B, build the longest sublist containing an equal
number of letters and numbers
*/
func QuestionFive(list []string) []string {
	/*
		Sub string must be even in length
		brute force by checking all subarrays, with some optimisations to allow for early exit

		complexity: N3
	*/
	var subArray []string
	differences := map[int]int{0: -1}
	aCount := 0
	bCount := 0
	for i, element := range list {
		if element == "A" {
			aCount++
		} else {
			bCount++
		}
		difference := aCount - bCount
		marker, ok := differences[difference]
		if !ok {
			differences[difference] = i
		} else {
			subLength := i - marker
			if len(subArray) < subLength {
				subArray = list[marker+1 : i+1]
			}
		}
	}
	return subArray
}

/*
QuestionSix

Write a method to count the total number of 2s between 0 and N inclusive
i.e. 22 -> 2, 12, 20, 21, 22 -> 6
*/
func QuestionSix(N int) int {
	count := 0
	for i := 0; i < N+1; i++ {
		count = count + CountTwos(i)
	}
	return count
}

func CountTwos(x int) int {
	/*
		Count the number of twos in a number
		thinking in terms of base 10
	*/
	power := 1
	count := 0
	for power < x {
		digit := x / power
		fmt.Println(count, x, digit, power)

		if digit == 2 {
			count++
		}
		power = power * 10
	}
	return count
}
