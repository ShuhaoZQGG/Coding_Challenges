package main

import "fmt"

func gethelpMessage() string {
	return fmt.Sprintln("To get help: ./wc help or ./wc -help or ./wc --help") +
		fmt.Sprintln("Commands available:") +
		fmt.Sprintln("./wc				please enter your messagess, and enter ctrl + c to terminate it to see results") +
		fmt.Sprintln("./wc -l filename          	get number of lines of a file") +
		fmt.Sprintln("./wc -c filename           	get number of bytes of a file") +
		fmt.Sprintln("./wc -w filename           	get number of words of a file") +
		fmt.Sprintln("./wc -m filename           	get number of characters of a file") +
		fmt.Sprintln("cat filename | ./wc -l          get number of characters from the stdin") +
		fmt.Sprintln("cat filename | ./wc -c          get number of bytes from the stdin") +
		fmt.Sprintln("cat filename | ./wc -w          get number of words from the stdin") +
		fmt.Sprintln("cat filename | ./wc -m          get number of characters from the stdin")
}
