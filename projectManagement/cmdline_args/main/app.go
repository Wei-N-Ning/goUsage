//
// NOTE:
// to build and run the executable
//
// cd to this_directory:
// /work/dev/golang/src/github.com/powergun/goUsage/projectManagement/cmdLineArgs/main
// go build -o /tmp/test app.go

package main

import (
	"fmt"
	"os"
)

// example:
// go build -o /tmp/test app.go
// /tmp/test a b a=1
//
// arguments:
// [/tmp/test a b a=1]
func main() {
	fmt.Println("arguments:")
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		panic("Insufficient arguments")
	}
}
