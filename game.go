package main

import (
	"fmt"
	"strings"
)

func wordToLetters(word string) []string {
	letters := strings.Split(word, "")
	return letters
}

func lettersToBlanks(letters []string) []string {
	var blanks []string
	for range letters {
		blanks = append(blanks, "_")
	}
	return blanks
}

func getLetterPositions(letters []string, letter string) []int {
	positions := []int{}

	for i := range letters {
		if letters[i] == letter {
			positions = append(positions, i)
		}
	}

	return positions
}

func hasBlanks(letters []string) bool {
	for _, l := range letters {
		if l == "_" {
			return true
		}
	}
	return false
}

type Game struct {
	movesLeft      int
	letters        []string
	guessedLetters map[string]int
	answers        []string
	status         string // can be turned into enum
	word           string
}

func NewGame(word string) Game {
	letters := wordToLetters(word)
	blanks := lettersToBlanks(letters)
	return Game{
		movesLeft:      10,
		letters:        wordToLetters(word),
		guessedLetters: make(map[string]int),
		answers:        blanks,
		status:         "start",
		word:           word,
	}
}

func DisplayGameStatus(g Game) {
	fmt.Println("You have", g.movesLeft, "left, use them wisely!")
	fmt.Println(strings.Join(g.answers, " "))
}

func (g *Game) MakeMove(letter string) {
	if len(letter) != 1 {
		g.status = "invalid"
		return
	}

	_, ok := g.guessedLetters[letter]
	if ok {
		g.status = "invalid"
		return
	}

	g.guessedLetters[letter] = 1
	positions := getLetterPositions(g.letters, letter)

	if len(positions) == 0 {
		g.movesLeft = g.movesLeft - 1
		if g.movesLeft == 0 {
			g.status = "lose"
		} else {
			g.status = "wrong_move"
		}
		return
	}

	for _, n := range positions {
		g.answers[n] = letter
	}

	if hasBlanks(g.answers) {
		g.status = "right_move"
	} else {
		g.status = "win"
	}
}
