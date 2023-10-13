package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
)

var (
	minNum, maxNum          int = 1, 1000
	varboseMode, maxGuesses     = true, 10
	guess, tries, randNum   int
	name                    string
	saveFileName            string = "leaderboard.txt"
	num                     int    = 1
)

func main() {
	playGame()
}

func playGame() {
	randNum = rand.Intn(maxNum-1) + 1

	fmt.Println("The Random Number is: ", randNum)
	fmt.Println("Enter your Name: ")
	fmt.Scanln(&name)

	fmt.Println("Top 5 Scores: ")

	exists := fileExists(saveFileName)

	if exists {
		readFileLines(saveFileName)
		startGame()
	} else {
		startGame()
	}
}

func startGame() {
	fmt.Println("Guess the random number between 1 - 1000: ")

	for guess != randNum {

		fmt.Scanln(&guess)
		tries++

		if guess == randNum {
			fmt.Println("Winner!!")
			fmt.Println("Number of tries: ", tries)
			replayGame()
		}

		if guess < randNum {
			fmt.Println("Too Low!  Try again:")
		} else if guess > randNum {
			fmt.Println("Too high! Try again:")
		}
	}
	createFile(saveFileName)
}

func replayGame() {
	var playAgain string
	fmt.Println("Would you like to play again? (y/n)")
	fmt.Scanln(&playAgain)

	if playAgain == "y" || playAgain == "Y" {
		fmt.Println("You choose to play again!")
		playGame()
	} else if playAgain == "n" || playAgain == "N" {
		fmt.Println("Exiting App")
		gameOver()
	} else {
		fmt.Println("Invalid entry")
		replayGame()
	}
}

func gameOver() {
	fmt.Println("Thanks for playing " + name)
}

func readFileLines(fileName string) {
	file, err := os.Open(fileName)

	checkErr(err)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string // will hold each line

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		fileLines = append(fileLines, line)
	}

	file.Close()
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	checkErr(err)

	fmt.Println("Data:", data)
}

func createFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)
	defer file.Close()

	result := strconv.Itoa(num) + "_ " + name + " - " + strconv.Itoa(tries) + " Guesses\n"

	_, err = file.WriteString(result)
	checkErr(err)

	num++
}

func fileExists(fileName string) bool {
	fileInfo, err := os.Stat(fileName)

	if fileInfo != nil && err == nil {
		return true
	} else {
		return false
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
