package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var points int

func main() {
	timeLimit := flag.Int("limit", 30, "the time limit of the quiz in seconds")
	flag.Parse()

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	gameOver := make(chan bool)

	go startGame(records, gameOver)

	select {
	case <-gameOver:
		// Quiz Done
	case <-time.After(time.Duration(*timeLimit) * time.Second):
		fmt.Println("Time's Up!")
	}

	fmt.Printf("%d/%d Correct", points, len(records))
}

func startGame(records [][]string, gameOver chan<- bool) {
	reader := bufio.NewReader(os.Stdin)

	for i, record := range records {
		fmt.Printf("Question %d\n%s\n", i+1, record[0])
		userAnswer := getUserInput("Your Answer: ", reader)

		e := evaluateAnswer(record[1], userAnswer)
		fmt.Println(e)
		fmt.Println("---------------")
	}

	gameOver <- true
}

func getUserInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	userAnswer, _ := reader.ReadString('\n')

	return strings.TrimSpace(userAnswer)
}

func evaluateAnswer(testAnswer string, userAnswer string) string {
	if testAnswer == userAnswer {
		points += 1
		return "Correct"
	}

	return "Wrong"
}
