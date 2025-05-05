package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format question, answer")
	timeLimit := flag.Int("time", 30, "time to answer each question in seconds")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Unable to read input file: %s \n", *csvFileName))
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as a CSV")
	}

	problems := parseLines(records)
	correct := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, problem := range problems {
		fmt.Printf("Problem #%d = %s \n", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			answerCh <- answer // arrow is always pointing towards the way data is moving
		}()
		select {
		// if we get a message from the timer, no matter what the answer is, it is not valid.
		case <-timer.C: // Waiting for a message from this channel. Our code will block until it gets a message from this channel.
			fmt.Printf("Time exceeded! You scored %d out of %d \n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				correct += 1
			}
		}

	}

	fmt.Printf("You scored %d out of %d \n", correct, len(problems))

}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
