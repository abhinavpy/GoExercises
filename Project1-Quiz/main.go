package main
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// in this code basically we are going to write our code that presents the quiz problems except users
	// input and checks for correctness
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of the 'question,answer'")
	flag.Parse()
	
	// this will read the file
	// csvFilename will be a pointer to a string
	file, err := os.Open(*csvFilename)
	
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// the io package specifies the io.Reader interface, which represents the read end of
	// a stream of data. The Go standard library contains many implementations of this interface, 
	// including files, network connections, compressors, ciphers, and others.
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file: %s\n", *csvFilename))
	}

	//fmt.Println(lines)
	problems := parseLines(lines)
	
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i + 1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) // we are going to pass in a reference to answer; remember scanf trims all spaces
		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect answer.")
		}
	}
	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

func parseLines(lines [][] string) []problem{
	//Golang make() is a built-in slice function used to create a slice. 
	//The make() function takes three arguments: type, length, and capacity, and returns the slice. 
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q : line[0],
			a : line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}