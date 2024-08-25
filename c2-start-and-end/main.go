package main

import (
	"log"
)

func main() {
	defer AspectExe()()

	log.Println("main")
}

func AspectExe() func() {

	// before exec
	log.Println("before exec")

	// after exec
	return func() {
		log.Println("after exec")
	}

}
