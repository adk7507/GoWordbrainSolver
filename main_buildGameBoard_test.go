package main

import(
	"testing"
  "github.com/stretchr/testify/assert"
)
    // Returns a board and a list of integers when given valid inputs.
func TestValidInputs(t *testing.T) {
  wlInput := []string{"4,4,4,4"}
  gbInput := []string{"abcd\nefgh\nijkl\nmnop"}
  expectedBoard := &board{
    size:       4,
    length:     16,
    characters: []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'},
    neighbors: [][]int{
      {1, 4, 5},
      {0, 2, 4, 5, 6},
      {1, 3, 5, 6, 7},
      {2, 6, 7},
      {0, 1, 5, 8, 9},
      {0, 1, 2, 4, 6, 8, 9, 10},
      {1, 2, 3, 5, 7, 9, 10, 11},
      {2, 3, 6, 10, 11},
      {4, 5, 9, 12, 13},
      {4, 5, 6, 8, 10, 12, 13, 14},
      {5, 6, 7, 9, 11, 13, 14, 15},
      {6, 7, 10, 14, 15},
      {8, 9, 13},
      {8, 9, 10, 12, 14},
      {9, 10, 11, 13, 15},
      {10, 11, 14},
    },
  }
  expectedWlInts := []int{4,4,4,4}

  board, wlInts := initGameFromUserInput(wlInput, gbInput)

  assert.Equal(t, expectedBoard, board)
  assert.Equal(t, expectedWlInts, wlInts)
}

    // Handles uppercase characters in the game board input.
func TestUppercaseCharacters(t *testing.T) {
  wlInput := []string{"4,4,4,4"}
  gbInput := []string{"ABCD\nEFGH\nIJKL\nMNOP"}
  expectedBoard := &board{
    size:       4,
    length:     16,
    characters: []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'},
    neighbors: [][]int{
      {1, 4, 5},
      {0, 2, 4, 5, 6},
      {1, 3, 5, 6, 7},
      {2, 6, 7},
      {0, 1, 5, 8, 9},
      {0, 1, 2, 4, 6, 8, 9, 10},
      {1, 2, 3, 5, 7, 9, 10, 11},
      {2, 3, 6, 10, 11},
      {4, 5, 9, 12, 13},
      {4, 5, 6, 8, 10, 12, 13, 14},
      {5, 6, 7, 9, 11, 13, 14, 15},
      {6, 7, 10, 14, 15},
      {8, 9, 13},
      {8, 9, 10, 12, 14},
      {9, 10, 11, 13, 15},
      {10, 11, 14},
    },
  }
  expectedWlInts := []int{4,4,4,4}

  board, wlInts := initGameFromUserInput(wlInput, gbInput)

  assert.Equal(t, expectedBoard, board)
  assert.Equal(t, expectedWlInts, wlInts)
}

    // Handles special characters in the game board input.
func TestSpecialCharacters(t *testing.T) {
  wlInput := []string{"1"}
  gbInput := []string{"!@#$\n%^&*\n()_+\n{}[]"}
  expectedBoard := (*board)(nil)
  expectedWlInts := []int(nil)

  board, wlInts := initGameFromUserInput(wlInput, gbInput)

  assert.Equal(t, expectedBoard, board)
  assert.Equal(t, expectedWlInts, wlInts)
}

    // Returns None when given an empty word list input.
func TestEmptyWordLenghts(t *testing.T) {
  wlInput := []string{}
  gbInput := []string{"abcd\nefgh\nijkl\nmnop"}

  board, wlInts := initGameFromUserInput(wlInput, gbInput)

  assert.Nil(t, board)
  assert.Nil(t, wlInts)
}

    // Returns None when given an empty game board input.
func TestEmptyGameBoard(t *testing.T) {
  wlInput := []string{"1,2,3,4"}
  gbInput := []string{}

  board, wlInts := initGameFromUserInput(wlInput, gbInput)

  assert.Nil(t, board)
  assert.Nil(t, wlInts)
}

    // Returns None when given a word list input with a non-integer value.
func TestNonIntegerWordLengths(t *testing.T) {
  wlInput := []string{"1,2,3,a"}
  gbInput := []string{"abcd\nefgh\nijkl\nmnop"}

  board, wlInts := initGameFromUserInput(wlInput, gbInput)

  assert.Nil(t, board)
  assert.Nil(t, wlInts)
}
