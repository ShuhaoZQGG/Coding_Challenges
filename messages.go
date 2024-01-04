package main

import "fmt"

func gethelpMessage() {
	fmt.Println("To get help: ./wc help or ./wc -help or ./wc --help")
	fmt.Println("Commands available:")
	fmt.Println("./wc				please enter your messagess, and enter ctrl + c to terminate it to see results")
	fmt.Println("./wc -l filename          	get number of lines of a file")
	fmt.Println("./wc -c filename           	get number of bytes of a file")
	fmt.Println("./wc -w filename           	get number of words of a file")
	fmt.Println("./wc -m filename           	get number of characters of a file")
	fmt.Println("cat filename | ./wc -l          get number of characters from the stdin")
	fmt.Println("cat filename | ./wc -c          get number of bytes from the stdin")
	fmt.Println("cat filename | ./wc -w          get number of words from the stdin")
	fmt.Println("cat filename | ./wc -m          get number of characters from the stdin")
}
