package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	// read all the problem form the quix.csv
	//1. Open the file
	if fObj, err := os.Open(fileName); err == nil {
		//2. we will create a new reader
		csvR := csv.NewReader(fObj)
		//3. it will read all the lines
		if cLines, err := csvR.ReadAll(); err == nil {
			//4. call the parseProblem function
			return parseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("error in reading data in csv"+"format from %s file; %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("error in opening %s file; %s", fileName, err.Error())
	}

}

func main() {
	// 1. Input the name of the file
	fName := flag.String("f", "quiz.csv", "Path of csv file")
	// 2. Set the duration of the time
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()
	// 3. Pull the problem from the file (calling our problem puller func)
	problems, err := problemPuller(*fName)
	// 4. Handle the error
	if err != nil {
		exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}
	// 5. Create a variable to count our correct answers
	correctAns := 0
	// 6. Using the duration of the timer, we want to initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string) //go channel uses here
	// 7.loop through the problems, print the questions, we'll accept the answers
problemLoop:
	for i, p := range problems {

		var answer string
		fmt.Printf("Problem %d: %s = ", i+1, p.q)

		go func() { //go routine
			fmt.Scanf("%s", &answer)
			ansC <- answer

		}() //go routine coll itself
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}
	// 8. we'll calculate and print out the result
	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
	fmt.Printf("Press enter to exit")
	<-ansC
}

func parseProblem(lines [][]string) []problem {
	// go over the lines and parse them, with problem struct
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
