package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

// isValidGameType function to check if gameType is a valid kind.
func isValidLevel(level int) bool {
	for _, b := range [7]int{1, 2, 3, 4, 5, 6, 7} {
		if b == level {
			return true
		}
	}
	return false
}

// ScoreBoard struct
type ScoreBoard struct {
	correct int8 //defaults to 0
	wrong   int8 //defaults to 0
}

func (s *ScoreBoard) printScore() {

	fmt.Printf("Correct: %d, wrong: %d\n", s.correct, s.wrong)
}

func (s *ScoreBoard) addCorrect() {
	s.correct++
}

func (s *ScoreBoard) addWrong() {
	s.wrong++
}

// FlashCard struct
type FlashCard struct {
	gameType         string
	level            int8
	expectedResponse float64
}

// NewFlashCard creates an instance of FlashCard including a math question.
func NewFlashCard(gameType string, level int8) *FlashCard {
	// TODO Ensure that gameType and level are compatible with game

	card := new(FlashCard)
	card.gameType = gameType
	card.level = level

	// Create question

	// Create expected reponse

	return card

}

func (card *FlashCard) checkResponse(reponse float64) bool {
	return almostEqual(card.expectedResponse, reponse)
}

func (card *FlashCard) printQuestion(counter int) float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	a := r1.Intn(200)
	b := r1.Intn(200)
	var operator string
	switch card.gameType {
	case "addition":
		operator = "+"
		card.expectedResponse = float64(a) + float64(b)
	case "subtraction":
		operator = "-"
		card.expectedResponse = float64(a) - float64(b)
	case "multiplication":
		operator = "*"
		card.expectedResponse = float64(a) * float64(b)
	case "division":
		operator = "/"
		card.expectedResponse = float64(a) / float64(b)
	case "mixed":
		operator = "*"
		card.expectedResponse = float64(a) * float64(b)
	}
	fmt.Printf("Question no. %d\n", counter)
	fmt.Printf("%d %s %d = ?\n", a, operator, b)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	responseInput := scanner.Text()
	response, err := strconv.Atoi(responseInput)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return float64(response)
}

// Game struct
type Game struct {
	gameType      string
	level         int8
	timeAllocated int8 //Minutes
	counter       int

	scoreBoard ScoreBoard
}

func (game *Game) initialize() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Math Flash Card Game!!!")
	fmt.Println("Practice your addition, subtraction, multiplication, and division skills!")
	fmt.Println("Press (s) for start. At anytime press (q) to exit. It is not possible to continue game.")

	fmt.Println("Choose between; addition, substraction, multiplication, division, mixed")
	scanner.Scan()
	gameTypeInput := scanner.Text()
	gameTypeInput = strings.ToLower(gameTypeInput)

	validGameType := isValidGameType(gameTypeInput)
	if validGameType == false {
		fmt.Println("Game type needs to be either; addition, substration, multiplication, division or mixed.")
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
	validLevel := isValidLevel(level)
	if validLevel == false {
		fmt.Println("Level needs to be either an integer between 1 and 7.")
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
	fmt.Printf("How many %s tasks can you solve in %d minutes on level %d?\n", game.gameType, game.timeAllocated, game.level)
}

func (game *Game) play() {
	//  repeat-until loop
	startTime := time.Now()
	for ok := true; ok; ok = time.Since(startTime).Seconds() < float64(game.timeAllocated*60) {
		game.counter++
		// Generate FlashCard
		card := NewFlashCard(game.gameType, game.level)

		// present FlashCard
		response := card.printQuestion(game.counter)
		fmt.Printf("Reponse: %d\n", int(response))

		// evaluate response
		validatedResponse := card.checkResponse(response)

		// add to score
		if validatedResponse == true {
			game.scoreBoard.addCorrect()
		} else if validatedResponse == false {
			game.scoreBoard.addWrong()
		}

	}

	fmt.Println("Time has passed.")

	game.scoreBoard.printScore()
}

// main function - game initiation and loop
func main() {
	game := Game{}
	game.initialize() // setup game
	game.play()
}
