package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format question, answer")
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

	for i, problem := range problems {
		fmt.Printf("Problem #%d = %s \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == problem.answer {
			correct += 1
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
