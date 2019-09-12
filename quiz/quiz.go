package main

import (
	"bufio"
	"time"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)


func printScore (c, i int) float32 {
	// calculate quiz score
	score := float32(c) / float32(c+i) * 100
	fmt.Printf("Score: %0.2f %%\n", score)
	return score
}


func main() {

	// command line args
	probFi := flag.String("problems", "problems.csv", "A csv file with problems in column one, and answers in column two.")
	secs := flag.Int("tlimit", 30, "Time limit to take the quiz in seconds.")
	flag.Parse()

	// read in the problems quiz file
	probs, err := os.Open(*probFi)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// init quiz reader
	rdr := csv.NewReader(bufio.NewReader(probs))

	// define stdin reader
	inp := bufio.NewReader(os.Stdin)

	// start the quiz
	fmt.Println("Let's start the Quiz!")

	ncorr := 0
	nincorr := 0

	// start the timer 
	timer := time.NewTimer(time.Duration(*secs) * time.Second)

	i := 0 
	for {
		select {
		case <-timer.C:
			fmt.Println("\nTimed out!")
			printScore(ncorr, nincorr)
			return
        default:
			line, err := rdr.Read()
			if line == nil {
				printScore(ncorr, nincorr)
				return
			}
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Problem %v: %v\n", i, line[0])

			txt, err := inp.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			txt = strings.Replace(txt, "\n", "", -1)

			if txt == line[1] {
				fmt.Printf("Yay! %v is the correct answer!\n", txt)
				ncorr++
			} else {
				fmt.Printf("%v is not the right answer\n", txt)
				nincorr++
			}
			i++
		}
	}
}
