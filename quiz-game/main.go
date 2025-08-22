package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var csvPath string
var timeLimit int

func init() {

	flag.StringVar(
		&csvPath,
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer' (default \"problems.csv\")",
	)

	flag.IntVar(&timeLimit, "limit", 30, "the time limit for the quiz in seconds (default 30)")

	flag.Parse()
}

func main() {

	file, err := os.Open(csvPath)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %v\n", err))
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to read the CSV file: %v\n", err))
	}

	problems := parseLines(records)

	questionsSize := len(records)
	correct := 0

	fmt.Println("Press enter to start the quiz")
	fmt.Scanln()
	fmt.Println("Starting the quiz...")

problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerCh := make(chan string, 1)
		timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

		go func() {

			var userAnswer string
			_, err := fmt.Fscan(os.Stdin, &userAnswer)

			if err != nil {
				exit(fmt.Sprintf("Failed to read user answer: %v\n", err))
			}

			select {
			case answerCh <- userAnswer:
			default:
			}

		}()

		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			break problemLoop
		case answer := <-answerCh:
			timer.Stop()
			if formatAnswer(answer) == formatAnswer(p.answer) {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, questionsSize)
}

func formatAnswer(answer string) string {
	return strings.TrimSpace(strings.ToLower(answer))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}
