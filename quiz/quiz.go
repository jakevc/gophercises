package main

import (
	"strings"
	"fmt"
	"log"
	"io"
	"bufio"
	"os"
	"encoding/csv"
	"flag"
)

func main() {

	// command line args
	probFi := flag.String("problems", "problems.csv", "A csv file with problems in column one, and answers in column two.")
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
	
	i := 1
	for {
		line, err := rdr.Read()
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
			ncorr ++
		} else {
			fmt.Printf("%v is not the right answer\n", txt)
			nincorr ++
		}
		i ++ 
	}
	
	// calculate quiz score 
	score := float32(ncorr)/float32(ncorr+nincorr) * 100
	fmt.Println("\nYou finished the quiz!")
	fmt.Printf("Correct: %v\n", ncorr)
	fmt.Printf("Incorrect: %v\n", nincorr)
	fmt.Printf("Score: %0.2f %%\n", score)
}
