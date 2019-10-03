package main

import (
	"time"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines (lines [][]string ) []problem {
	ret := make([]problem, len(lines))
	for i,line := range lines {
		ret[i] = problem{
			q:line[0],
			a:strings.TrimSpace(line[1]),
		}
	} 
	return ret
}


func main() {

	// command line args
	probFi := flag.String("problems", "problems.csv", "A csv file with problems in column one, and answers in column two.")
	secs := flag.Int("tlimit", 30, "Time limit to take the quiz in seconds.")
	flag.Parse()

	// read in the problems quiz file
	file, err := os.Open(*probFi)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *probFi))
	}	

	// parse problem file
	r := csv.NewReader(file)
	all, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse csvfiel: %s\n", *probFi))
	}

	problems := parseLines(all)

	fmt.Println(problems)

	// initialize a timer 
	timer := time.NewTimer(time.Duration(*secs) * time.Second)
	corr := 0

	problemLoop:
		for i,p := range problems {
			fmt.Printf("Problem #%d: %s = ", i+1, p.q)
			ans := make(chan string)
			go func() {
				var answer string 
				fmt.Scanf("%s\n", &answer)
				ans <- answer
			}()

			select {
			case <- timer.C:
				fmt.Println("\nYou ran out of time!")
				break problemLoop
			case answer := <- ans:
				if answer == p.a {
					corr++
				}
			}
		}
		fmt.Printf("You got %d out of %d correct.\n", corr, len(problems))
}
