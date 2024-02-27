package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

var points int

func main() {
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
	case <-time.After(30 * time.Second):
		fmt.Println("Time's Up!")
	}

	fmt.Printf("%d/%d Correct", points, len(records))
}

func startGame(records [][]string, gameOver chan<- bool) {
	for i, record := range records {

		fmt.Printf("Question %d\n%s\n", i+1, record[0])
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Your Answer: ")
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)

		e := evaluateAnswer(record[1], userAnswer)
		fmt.Println(e)
		fmt.Println("---------------")
	}

	gameOver <- true
}

func evaluateAnswer(testAnswer string, userAnswer string) string {
	if testAnswer == userAnswer {
		points += 1
		return "Correct"
	}

	return "Wrong"
}
