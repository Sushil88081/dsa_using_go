package main

import "fmt"

func main() {
	// Create a new array of integer
	//find smallest in an array
	arr := [5]int{4, 5, -12, 7, 1}
	var smallest int

	for i := 0; i < len(arr); i++ {

		if arr[i] < smallest {
			smallest = arr[i]

		}
	}
	fmt.Println(smallest)
}
