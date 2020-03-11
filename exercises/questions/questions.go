package questions

import (
	"fmt"
	"math"
	"math/bits"
	"math/rand"
	"sort"
	"strconv"
	"strings"
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
	for i := len(deck) - 1; i > 0; i-- {
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
	for i := 0; i < len(strconv.Itoa(N)); i++ {
		count = count + CountTwosAtDigit(N, i)
	}
	return count
}

func CountTwosAtDigit(N, d int) int {
	power := int(math.Pow10(d))
	next := power * 10
	right := N % power

	roundDown := N - N%next
	roundUp := roundDown + next

	digit := (N / power) % 10
	if digit < 2 {
		return roundDown / 10
	} else if digit == 2 {
		return roundDown/10 + right + 1
	} else {
		return roundUp / 10
	}
}

/*
QuestionSeven

Given a list of keys and frequencies, and a separate list of synonym key lists
construct a true frequency list
*/
func QuestionSeven(freq map[string]int, synonyms [][]string) map[string]int {
	/*
		key to this is the data structure we use for converting from synonyms
	*/
	g := InitGraph(freq)
	AddSynonyms(g, synonyms)
	return g.CountFrequencies()
}

func InitGraph(freq map[string]int) *FrequencyGraph {
	g := &FrequencyGraph{nodes: map[string]*FrequencyNode{}}
	for k, v := range freq {
		g.AddNode(&FrequencyNode{
			name: k,
			freq: v,
		})
	}
	return g
}

func AddSynonyms(graph *FrequencyGraph, synonyms [][]string) {
	for _, pair := range synonyms {
		graph.AddEdge(pair[0], pair[1])
	}
}

type FrequencyGraph struct {
	nodes map[string]*FrequencyNode
}

func (g *FrequencyGraph) AddNode(node *FrequencyNode) {
	g.nodes[node.name] = node
}

func (g *FrequencyGraph) AddEdge(a, b string) {
	aNode, ok := g.nodes[a]
	if !ok {
		aNode = &FrequencyNode{
			name: a,
			freq: 0,
		}
		g.AddNode(aNode)
	}
	bNode, ok := g.nodes[b]
	if !ok {
		bNode = &FrequencyNode{
			name: b,
			freq: 0,
		}
		g.AddNode(bNode)
	}
	aNode.AddChild(bNode)
}

func (g *FrequencyGraph) Print() string {
	nodes := []string{}
	for k, node := range g.nodes {
		nodes = append(nodes, fmt.Sprintf("%s: %s", k, node.Print()))
	}
	return strings.Join(nodes, "\n")
}

func (g *FrequencyGraph) CountFrequencies() map[string]int {
	fmt.Println(g.Print())
	counts := map[string]int{}
	for _, node := range g.nodes {
		fmt.Println(node.visited, node.name, node.freq)
		if !node.visited {
			counts[node.name] = node.CountFrequencies()
		}
	}
	return counts
}

type FrequencyNode struct {
	name     string
	freq     int
	children []*FrequencyNode
	visited  bool
}

func (n *FrequencyNode) AddChild(node *FrequencyNode) {
	n.children = append(n.children, node)
	node.children = append(node.children, n)
}

func (n *FrequencyNode) CountFrequencies() int {
	n.visited = true
	count := n.freq
	for _, child := range n.children {
		if !child.visited {
			count = count + child.CountFrequencies()
		}
	}
	return count
}

func (n *FrequencyNode) Print() string {
	names := []string{}
	for _, child := range n.children {
		names = append(names, child.name)
	}
	return fmt.Sprintf("%s %d %t children: %s", n.name, n.freq, n.visited, strings.Join(names, ","))
}

/*
QuestionEight

Given a set of people with height and weight, build the highest tower of people
lighter and shorter than the people below
*/
func QuestionEight(staff []*CircusPerson) []*CircusPerson {
	/*
		Could start by sorting the set by height and weight, then incrementing across both lists.
		Might need to increment weight first and hieght first to avoid edge cases
	*/
	sort.Slice(staff, func(i, j int) bool {
		return staff[i].weight < staff[j].weight
	})
	return LongestSubTower(staff, []*CircusPerson{}, 0)
}

func LongestSubTower(array []*CircusPerson, sequence []*CircusPerson, index int) []*CircusPerson {
	if index >= len(array) {
		return sequence
	}
	bestWith := []*CircusPerson{}
	if canAppend(sequence, array[index]) {
		sequenceWith := append(sequence, array[index])
		bestWith = LongestSubTower(array, sequenceWith, index+1)
	}
	bestWithout := LongestSubTower(array, sequence, index+1)
	if len(bestWith) > len(bestWithout) {
		return bestWith
	}
	return bestWithout

}

func canAppend(solution []*CircusPerson, person *CircusPerson) bool {
	if len(solution) == 0 {
		return true
	}
	return solution[len(solution)-1].isSmaller(person)
}

type CircusPerson struct {
	weight int
	height int
}

func (c *CircusPerson) isSmaller(p *CircusPerson) bool {
	return c.height < p.height && c.weight < p.weight
}

/*
QuestionNine

Design an algorithm to find the kth number such that the only prime factors are
3, 5, 7. 3, 5, 7 need not be factors but no other prime
*/
func QuestionNine(k int) []int {
	results, primes := []int{}, []int{}
	factor := 1
	for i := 0; i < k; i++ {
		factor = GetNextNumber(factor, &primes)
		results = append(results, factor)
		factor++
	}
	return results
}

func GetNextNumber(factor int, primes *[]int) int {
	for true {
		prime := CheckFactor(factor, primes)
		if prime {
			fmt.Println(factor, "prime")
			return factor
		}
		factor++
	}
	return 0
}

func CheckFactor(i int, primes *[]int) bool {
	fmt.Println(i, primes)

	for _, p := range *primes {
		if i%p == 0 {
			return false
		}
	}
	if i == 1 || i%3 == 0 || i%5 == 0 || i%7 == 0 {
		return true
	}
	*primes = append(*primes, i)
	return false
}
