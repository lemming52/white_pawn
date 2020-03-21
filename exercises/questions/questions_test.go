package questions

import (
	"fmt"
	"reflect"
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
		}, {
			name:     "other",
			N:        25,
			expected: 9,
		}, {
			name:     "zero",
			N:        1,
			expected: 0,
		}, {
			name:     "large",
			N:        1000,
			expected: 300,
		}, {
			name:     "odd",
			N:        245,
			expected: 101,
		}, {
			name:     "big",
			N:        2456,
			expected: 1253,
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

func TestQuestionSeven(t *testing.T) {
	tests := []struct {
		name            string
		frequencies     map[string]int
		synonyms        [][]string
		expected        map[string]int
		expectedEntries int
	}{
		{
			name: "base",
			frequencies: map[string]int{
				"A": 1,
				"B": 2,
				"a": 3,
			},
			synonyms: [][]string{
				[]string{"A", "a"},
			},
			expected: map[string]int{
				"A": 4,
				"B": 2,
				"a": 4,
			},
			expectedEntries: 2,
		}, {
			name: "names",
			frequencies: map[string]int{
				"John":        15,
				"Jon":         12,
				"Chris":       13,
				"Kris":        4,
				"Christopher": 19,
			},
			synonyms: [][]string{
				[]string{"John", "Jon"},
				[]string{"John", "Johnny"},
				[]string{"Chris", "Kris"},
				[]string{"Kris", "Christopher"},
			},
			expected: map[string]int{
				"John":        27,
				"Johnny":      27,
				"Jon":         27,
				"Kris":        36,
				"Chris":       36,
				"Christopher": 36,
			},
			expectedEntries: 2,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionSeven(tt.frequencies, tt.synonyms)
			assert.Equal(t, tt.expectedEntries, len(result))
			for k, v := range result {
				entry, ok := tt.expected[k]
				assert.Equal(t, ok, true)
				assert.Equal(t, entry, v)
			}
		})
	}
}

func TestQuestionEight(t *testing.T) {
	tests := []struct {
		name  string
		staff []*CircusPerson
		tower []*CircusPerson
	}{
		{
			name: "base",
			staff: []*CircusPerson{
				&CircusPerson{65, 100},
				&CircusPerson{70, 150},
				&CircusPerson{56, 90},
				&CircusPerson{75, 190},
				&CircusPerson{60, 95},
				&CircusPerson{68, 110},
			},
			tower: []*CircusPerson{
				&CircusPerson{56, 90},
				&CircusPerson{60, 95},
				&CircusPerson{65, 100},
				&CircusPerson{68, 110},
				&CircusPerson{70, 150},
				&CircusPerson{75, 190},
			},
		}, {
			name: "breaking",
			staff: []*CircusPerson{
				&CircusPerson{1, 100},
				&CircusPerson{2, 34},
				&CircusPerson{3, 35},
				&CircusPerson{4, 36},
			},
			tower: []*CircusPerson{
				&CircusPerson{2, 34},
				&CircusPerson{3, 35},
				&CircusPerson{4, 36},
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionEight(tt.staff)
			assert.Equal(t, tt.tower, result)
		})
	}
}

func TestQuestionNine(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		expected []int
	}{
		{
			name:     "base",
			k:        7,
			expected: []int{1, 3, 5, 7, 9, 15, 21},
		}, {
			name:     "longer",
			k:        10,
			expected: []int{1, 3, 5, 7, 9, 15, 21, 25, 27, 35},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionNine(tt.k)
			ok := reflect.DeepEqual(tt.expected, result)
			if !ok {
				t.Error("result array should equal", result, tt.expected)
			}
		})
	}
}

func TestQuestionTen(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		array    []int
	}{
		{
			name:     "base",
			expected: 5,
			array:    []int{2, 3, 5, 5, 5, 5, 8},
		}, {
			name:     "book",
			expected: 7,
			array:    []int{3, 1, 7, 1, 1, 7, 7, 3, 7, 7, 7},
		}, {
			name:     "no majority",
			expected: -1,
			array:    []int{1, 2, 3, 4},
		}, {
			name:     "book other",
			expected: 5,
			array:    []int{1, 2, 5, 9, 5, 9, 5, 5, 5},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionTen(tt.array)
			ok := reflect.DeepEqual(tt.expected, result)
			if !ok {
				t.Error("result array should equal", result, tt.expected)
			}
		})
	}
}

func TestQuestionEleven(t *testing.T) {
	tests := []struct {
		name       string
		expected   int
		array      []string
		candidate1 string
		candidate2 string
	}{
		{
			name:       "base",
			expected:   3,
			array:      []string{"word1", "a", "b", "word2"},
			candidate1: "word1",
			candidate2: "word2",
		}, {
			name:       "multiple",
			expected:   2,
			array:      []string{"word1", "a", "b", "word2", "a", "word1"},
			candidate1: "word1",
			candidate2: "word2",
		}, {
			name:       "multiple",
			expected:   1,
			array:      []string{"word1", "a", "b", "word2", "a", "A", "word2", "b", "word1", "word2"},
			candidate1: "word1",
			candidate2: "word2",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionEleven(tt.array, tt.candidate1, tt.candidate2)
			ok := reflect.DeepEqual(tt.expected, result)
			if !ok {
				t.Error("result array should equal", result, tt.expected)
			}
		})
	}
}

func TestQuestionTwelve(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		initial  string
		expected string
	}{
		{
			name:     "base",
			array:    []int{1, 2, 3, 4},
			expected: "1 2 3 4",
		}, {
			name:     "out of order",
			array:    []int{3, 2, 4, 1},
			expected: "1 2 3 4",
		}, {
			name:     "long",
			array:    []int{3, 2, 4, 7, 5, 9, 8, 6, 1},
			expected: "1 2 3 4 5 6 7 8 9",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tree := buildTree(tt.array)
			head := QuestionTwelve(tree)
			assert.Equal(t, tt.expected, head.print())
		})
	}
}

func buildTree(array []int) *BiNode {
	head := newNode(array[0])
	for i := 1; i < len(array); i++ {
		fmt.Println(array[i])
		head.insert(newNode(array[i]))
	}
	return head
}

func TestQuestionThirteen(t *testing.T) {
	tests := []struct {
		name            string
		text            string
		dict            map[string]bool
		expectedInvalid int
		expectedText    string
	}{
		{
			name: "base",
			text: "weare",
			dict: map[string]bool{
				"we":  true,
				"are": true,
			},
			expectedInvalid: 0,
			expectedText:    "we are",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, invalid := QuestionThirteen(tt.text, tt.dict)
			assert.Equal(t, tt.expectedText, result)
			assert.Equal(t, tt.expectedInvalid, invalid)
		})
	}
}

func TestQuestionFourteen(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		k        int
		expected []int
	}{
		{
			name:     "base",
			array:    []int{4, 7, 1, 2},
			k:        2,
			expected: []int{1, 2},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionFourteen(tt.array, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}
