package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	buf := make([]byte, 16)

	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", buf)

	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

}
