package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Hangman! Do you want to play? (y or n)")
	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	ans = strings.ToLower(strings.Trim(ans, "\n"))

	if ans != "y" {
		fmt.Println("Sad to see you go!")
		os.Exit(0)
	}
	word := GenerateWord()
	game := NewGame(word)
	for {
		DisplayGameStatus(game)
		guess, _ := reader.ReadString('\n')
		guess = strings.ToLower(strings.Trim(guess, "\n"))
		game.MakeMove(guess)

		if game.status == "win" {
			fmt.Println("Congratulations! You got the word.")
			fmt.Println(strings.Join(game.answers, ""))
			break
		}

		if game.status == "lose" {
			fmt.Println("You lose! Try again next time.")
			break
		}

		if game.status == "invalid" {
			fmt.Println("Hey! You can't do that, try another move.")
		}

		if game.status == "wrong_move" {
			fmt.Println("Shucks, that was a wrong guess!")
		}

		if game.status == "right_move" {
			fmt.Println("That was a correct guess!")
		}

	}

	fmt.Println("Thanks for playing hangman!")
}
