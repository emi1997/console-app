package main

import (
	"fmt"
)

func main(){

	//Declaration:
	letters := []string{"a", "b", "c", "d"}
	//Type: []string
	//No specified length
	//Can be made with the make function:
	var s[]byte = make([]byte, 5, 5)
	//make allocates an array, and returns a slice referring to that array
	//Arguments: 1. type, 2. length, 3. capacity (optional)
	length := len(s) == 5
	//inspecting the length of a slice
	capacity := cap(s) == 5
	//inspecting the capacity of a slice
	b := []string{"g", "o", "l", "a", "n", "g"}
	var slice []string = b[:2]
	//b[:2] == []string{"g", "o"}
	fmt.Println(letters, length, capacity, b[1:4], slice)
	//Increase the cpacity of a slice:
	//1. make a bigger slice and copy the data from the previous slice
	t := make([]byte, len(s), (cap(s)+1)*2)//+1 in case s == 0
	copy(t, s)
	fmt.Println(t, s)
	//2. Append data to the end of a slice:
	a := make([]int, 1)
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)
}
