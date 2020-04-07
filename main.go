package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// timer

// Constants

const (
	numberOfLevels           int8    = 7
	defaultTimeMinutes       int8    = 2
	float64EqualityThreshold float64 = 1e-9
)

// almostEqual function to check if two float64 are almost equal
func almostEqual(a, b float64) bool {
	// https://floating-point-gui.de/errors/comparison/#look-out-for-edge-cases
	return math.Abs(a-b) <= float64EqualityThreshold
}

// isValidGameType function to check if gameType is a valid kind.
func isValidGameType(gameType string) bool {
	switch gameType {
	case
		"addition",
		"subtraction",
		"multiplication",
		"division",
		"mixed":
		return true
	}
	return false
}

// ScoreBoard struct
type ScoreBoard struct {
	correct int8 //defaults to 0
	wrong   int8 //defaults to 0
}

func (s *ScoreBoard) printScore() {

	fmt.Printf("Correct: %d, wrong: %d", s.correct, s.wrong)
}

// Game struct
type Game struct {
	gameType      string
	level         int8
	timeAllocated int8 //Minutes

	scoreBoard ScoreBoard
}

func (game *Game) initialize() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Choose between; addition, substraction, multiplication, division, mixed")
	scanner.Scan()
	gameTypeInput := scanner.Text()
	gameTypeInput = strings.ToLower(gameTypeInput)

	validGameType := isValidGameType(gameTypeInput)
	if validGameType == false {
		fmt.Println(" Game type needs to be either; addition, substration, multiplication, division or mixed.")
		os.Exit(2)
	}
	game.gameType = gameTypeInput

	// TODO: Validate input
	fmt.Println("Choose level; 1 to 7")
	scanner.Scan()
	levelInput := scanner.Text()
	level, err := strconv.Atoi(levelInput)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	game.level = int8(level)

	// TODO: Validate input
	// TODO: Handle default values
	fmt.Println("Choose total time allocated in minutes (default)")
	scanner.Scan()
	timeAllocatedInput := scanner.Text()
	timeAllocated, err := strconv.Atoi(timeAllocatedInput)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	game.timeAllocated = int8(timeAllocated)

	fmt.Println("Starting game!!!")
	fmt.Printf("How many %s tasks can you solve in %d minutes on level %d?", game.gameType, game.timeAllocated, game.level)
}

func (game *Game) play() {
	for {

	}

	game.scoreBoard.printScore()
}

// FlashCard struct
type FlashCard struct {
	gameType         string
	expectedResponse float64
}

func (card *FlashCard) checkResponse(reponse float64) bool {
	return almostEqual(card.expectedResponse, reponse)
}

// main function - game initiation and loop
func main() {
	fmt.Println("Welcome to the Math Flash Card Game!!!")
	fmt.Println("Practice your addition, subtraction, multiplication, and division skills!")
	fmt.Println("Press (s) for start. At anytime press (q) to exit. It is not possible to continue game.")

	game := Game{}
	game.initialize() // setup game
	game.play()
}
