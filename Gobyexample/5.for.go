package main

import "fmt"

/*
	Branching with if and else in Go is strightforward.
	Here's a basic example.
	You can have an if statement without an else.
	A statement can precede conditionals; any variables declared 
	in this statement are available in all branches.
	Note that you don't need parentheses around conditionals in Go,
	but that the braces are requried.
	There is no ternary if in Go, so you'll need to use a full if 
	statement even for basic conditions.
*/

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}