package main

import (
	"fmt"
	"os"

	"github.com/Yakiyo/monkey/repl"
)

func main() {
	fmt.Println("Initializing Monkey repl")
	fmt.Println()
	repl.Start(os.Stdin, os.Stdout)
	
}
