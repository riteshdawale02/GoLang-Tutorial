package main

import "fmt"

func main() {

	var nums [4]int

	nums[0] = 2
	nums[1] = 3
	nums[2] = 4
	nums[3] = 5


	for i:= 0; i < len(nums); i++ {
		fmt.Print(nums[i]," ")
	}

	fmt.Print(nums)
  
	//length of array
	fmt.Print(len(nums))
	fmt.Println(nums[0])


	arr := [3]int {1,2,4,}
	fmt.Print(arr)


	//2d array
	td := [2][2]int {{1,2},{3,4}}

	fmt.Print(td)

	//fixed size
	//Memory optimization
	//constant time access
	//
}