package main

import (
	"fmt"
	_ "github.com/tangthinker/go-craft/c5-call-init/another-package" // This line is added to call the init function of another-package
)

func main() {
	fmt.Println("Hello, world!")
}

// output:
// This is another package init() function
// Hello, world!
