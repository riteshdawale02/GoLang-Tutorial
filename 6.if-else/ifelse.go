package main

import "fmt"

func main() {
	age :=  29

	if age < 18 {
		fmt.Print("Minor")
	} else if age <= 21 {
		fmt.Println("Teen")	
	}else {
		fmt.Print("Adult\n")
	}

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("Yes")
	}

	//IMPORTANT
	if score := 93; score >= 90 {
		fmt.Println("Class A")
	} else if (score >= 35){
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}



}