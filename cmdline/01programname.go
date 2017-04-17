package main

import "fmt"
import "os"
import "path/filepath"

func main() {
	program := filepath.Base(os.Args[0])
	fmt.Println(program)
}
