//
// similar to static linking, the resulting exe
// contain the binary of the imported package routine
//
// 0000000000467630 T github.com/powergun/goUsage/fileSystem.Exists

package main

import "fmt"
import "github.com/powergun/goUsage/fileSystem"


func main() {
	filePath := "/tmp/test"
	if fileSystem.Exists(filePath) {
		fmt.Println("file exists!")
	} else {
		fmt.Println("file does not exist")
	}
}
