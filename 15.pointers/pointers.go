package main

import "fmt"

//copy by value
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", num)
}

//copy by reference
func changeNumByReference(num int ) {

}

func main() {
	num := 1

	fmt.Println("Before changeNum is main",num)

	changeNum(&num)
	fmt.Println("After changeNum is main",num)
}