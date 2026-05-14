package main

import "fmt"

//[T any]
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Print(item," ")
	}
}


//specific type of elements
func printSlice2 [T string | int] (items []T) {
	for _, item := range items {
		fmt.Print(item," ")
	}
}

func printSlice3 [T comparable, V string] (items []T, name V) {
	for _, item := range items {
		fmt.Print(item," ", name)
	}
}

type stack [T any] struct {
	elements []T
}

func main() {
	nums := [] int {1,2,4}
	names := [] string {"Java", "C++", "GO"}
	printSlice(names)
	fmt.Println("\n===============")
	printSlice(nums)
	printSlice2(nums)

	myStack := stack[string] {
		elements: []string{"String", "Variables"},
		
	}

	fmt.Println(myStack)
}