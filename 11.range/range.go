package main

import "fmt"

// iterating over data structures
func main() {

	nums := []int {5,7,9}

	var sum = 0
	for i, num := range nums {
		sum += num
		fmt.Println(num, i)
	}
	fmt.Println(sum)
	


	//map range

	m := map[string]string{"fname":"ritesh", "lname":"dawale"}

	for k, v := range m {
		fmt.Println(k,"-",v)
	}


	// for i := 0; i < len(nums); i++ {
	// 	fmt.Print(nums[i]," ")
	// }

	//unicode code
	//starting byte of rune
	//255
	for i , c := range "GOLANG" {
		fmt.Println(i,"",c)
	}

}