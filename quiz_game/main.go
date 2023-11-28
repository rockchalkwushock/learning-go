package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds.")
	flag.Parse()

	// Need to use *csvFilename because the "flag" package is returning a pointer to a string value.
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	// This is a label (kind of like a JUMP TO statement)
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		// Create a Channel to get the user's answer out of the go routine.
		answerCh := make(chan string)
		go func() {
			var answer string
			// Scanf is blocking our program from running so regardless if time has
			// expired the user's answer will be counted. To solve that we place it
			// inside of a go routine.
			fmt.Scanf("%s\n", &answer)
			// Send the user's answer back out to the Channel.
			answerCh <- answer
		}()

		select {
		// Wait for message from Channel.
		case <-timer.C:
			fmt.Println()
			// Instead of using a return here we can break & reference the label to break out of
			break problemLoop
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

type Problem struct {
	answer   string
	question string
}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			answer:   strings.TrimSpace(line[1]),
			question: line[0],
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
