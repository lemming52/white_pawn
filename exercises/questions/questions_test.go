package questions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestionOne(t *testing.T) {
	tests := []struct {
		name     string
		x        int64
		y        int64
		expected int64
	}{
		{
			name:     "simple",
			x:        1,
			y:        1,
			expected: 2,
		}, {
			name:     "larger",
			x:        759,
			y:        674,
			expected: 1433,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionOne(tt.x, tt.y)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestQuestionTwo(t *testing.T) {
	deck := initDeck()
	fmt.Println("\ninitial")
	for _, card := range deck {
		fmt.Print(card.Print())
	}
	result := QuestionTwo(deck, 51)

	fmt.Println("\nshuffled")
	for _, card := range result {
		fmt.Print(card.Print())
	}
	assert.Equal(t, len(deck), len(result))
	if !testDecks(deck, result) {
		t.Error("decks don't match")
	}
}

func testDecks(a, b []*Card) bool {
	for _, card := range a {
		if !checkCard(card, b) {
			return false
		}
	}
	return true
}

func checkCard(card *Card, deck []*Card) bool {
	for _, c := range deck {
		if (card.suit == c.suit) && (card.value == c.value) {
			return true
		}
	}
	return false
}

func TestQuestionThree(t *testing.T) {
	test := []int{1, 4, 6, 7, 8}
	fmt.Println(QuestionThree(test, 3))
	fmt.Println(QuestionThree(test, 3))
	fmt.Println(QuestionThree(test, 3))
}

func TestQuestionFour(t *testing.T) {
	tests := []struct {
		name string
		N    int64
		m    int64
	}{
		{
			name: "small",
			N:    4,
			m:    1,
		}, {
			name: "medium",
			N:    8,
			m:    4,
		}, {
			name: "large",
			N:    32,
			m:    24,
		}, {
			name: "huge",
			N:    645,
			m:    325,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			list := initIntModList(tt.N, tt.m)
			result := QuestionFour(list, 0)
			assert.Equal(t, tt.m, result)

			for _, i := range list {
				fmt.Print(i.value, " ", i.Get(1))
			}
		})
	}
}

func initIntModList(N, m int64) []*IntMod {
	list := []*IntMod{}
	for i := int64(0); i < N; i++ {
		if i == m {
			continue
		}
		list = append(list, &IntMod{i})
	}
	return list
}

func TestQuestionFive(t *testing.T) {
	tests := []struct {
		name string
		full []string
		sub  []string
	}{
		{
			name: "base",
			full: []string{"A", "B", "A"},
			sub:  []string{"A", "B"},
		}, {
			name: "long",
			full: []string{"A", "B", "A", "A", "B", "B", "B", "A", "B"},
			sub:  []string{"A", "B", "A", "A", "B", "B", "B", "A"},
		}, {
			name: "split",
			full: []string{"A", "B", "A", "A", "A", "A", "A", "B", "B"},
			sub:  []string{"A", "A", "B", "B"},
		}, {
			name: "nil",
			full: []string{"A", "A", "A"},
			sub:  nil,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionFive(tt.full)
			assert.Equal(t, tt.sub, result)
		})
	}
}

func TestQuestionSix(t *testing.T) {
	tests := []struct {
		name     string
		N        int
		expected int
	}{
		{
			name:     "base",
			N:        22,
			expected: 6,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionSix(tt.N)
			assert.Equal(t, tt.expected, result)
		})
	}
}
