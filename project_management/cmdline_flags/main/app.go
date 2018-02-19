/*

/tmp/test --help
Usage of /tmp/test:
  -num int
        provide an integer
  -word string
        provide a word

/tmp/test -word asd -num 11
word: asd
num 11

*/


package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "", "provide a word")
	intPtr := flag.Int("num", 0, "provide an integer")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("num", *intPtr)
}
