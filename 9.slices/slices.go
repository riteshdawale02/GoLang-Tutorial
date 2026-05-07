package main

import "fmt"

//slices -> dynamic
//most used construct in go
// + userful methods
func main() {
	//uninitialized slice is nil
	var nums []string

	nums = append(nums, "ritesh", "dawale")
	fmt.Println(nums == nil)
	fmt.Println(nums)
	fmt.Print(len(nums))
    

	//make([]type, initial values, max no. of element)
	var slc = make([]int, 2, 5)

	//using index
	slc[0] = 3
	slc[1] = 6
	fmt.Println(len(slc))
	fmt.Println(slc)
	
	fmt.Println(cap(slc))
	
	slc = append(slc, 1,2,3,4,5)
	fmt.Println(len(slc))
	fmt.Println((slc))


	var new = make([]int, 0, 5)
	var new1 = make([]int, len(nums))

	new = append(new, 2)

	//copy function
	copy(new1, new)

	fmt.Println(new, new1)
	 
	fmt.Println("=============")

	//slice opertor
	var first = []int {1,2,3}
	fmt.Println(first[0:2])
	fmt.Println(first[:2])
	fmt.Println(first[:])

	
}
