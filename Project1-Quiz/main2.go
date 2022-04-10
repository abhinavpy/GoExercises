package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/csv"
)

/*
In part 2 we are going to basically add a time limit to this quiz.
Create a program that will accept the CSV filepath and a time limit (in seconds) as flags
and will then run the quiz reading each problem in order and stopping the quiz as soon as
the time limit has been exceeded.

Users should be able to press enter (or some other key) before the timer starts, and then
the questions should be printed out to the screen one at a time until the user provides an 
answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should output the total number of questions there were in total.
Questions given invalid answers or unanswered answers are considered incorrect.

*/
/*
	Go by example Timers
	We often want to execute Go code at some point in the future, or repeatedly at some 
	interval. Go's built in timer and ticker features make both of these tasks easy. 
	We'll look first at timers then at tickers.
	package main
	import (
		"fmt"
		"time"
	)

	func main() {
		timer1 := time.NewTimer(2 * time.Second)

		<-timer1.C
		fmt.Println("Timer 1 fired")
		timer2 := time.NewTimer(time.Second)
		go func() {
			<-timer2.C
			fmt.Println("Timer 2 fired")
		}()
		stop2 := timer2.Stop()
		if stop2 {
			fmt.Println("Timer 2 stopped")
		}
		time.Sleep(2 * time.Second)
	}
	Timers represent a single event in the future. You tell the timer how long you want to 
	wait, and it provides a channel that will be notified at that time. This timer will wait 
	2 seconds. 
	The <-timer1.C blocks on the timer's channel C until it sends a value indicating that the timer
	fired.
	If you just wanted to wait, you could have used time.Sleep. One reason a timer may be useful is 
	that you can cancel a timer before it fires. Here's an example of that.
	Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
	The first timer will fire ~2s after we start the program, but the second should be stopped before
	it has a chance to fire.

	Go by example Tickers
	Timers are for when you want to do something in the future - tickers are for when you want
	to do something repeatedly at regular intervals. Here's an example of a ticker that ticks 
	periodically until we stop it.
	Tickers use a similar mechanism to timers: a channel that is sent values. Here we'll use the
	select builtin on the channel to await the values as they arrive every 500ms.
	package main
	import(
		"fmt"
		"time"
	)
	func main() {
		ticker := time.NewTicker(500 * time.Millisecond)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <- done:
					return
				case t := <-ticker.C
					fmt.Println("Tick at", t)
				}
			}
		}()

		time.Sleep(1600 * time.Millisecond)
		ticker.Stop()
		done <- true
		fmt.Println("Ticker stopped")
	}
*/

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of the 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprinf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	// look up timer package on the web

}