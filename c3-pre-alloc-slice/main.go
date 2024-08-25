package main

import "fmt"

func main() {

	arr := make([]int, 10)
	arr = append(arr, 1)
	fmt.Println("arr :", arr)

	arr1 := make([]int, 0, 10)
	arr1 = append(arr1, 1)
	fmt.Println("arr1 :", arr1)

	// Output:
	// arr : [0 0 0 0 0 1]
	// arr1 : [1]

}
