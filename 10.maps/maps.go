package main

import (
	"fmt"
	"maps"
)

func main() {

	//creating map
	m := make(map[string]string)

	//setting an elements

	m["name"] = "golang"
	m["service"] = "backend"

	//get a element
	fmt.Println(m["name"])
	fmt.Println(m["service"])

	fmt.Println(m["name"],"helllo")

	fmt.Println(len(m))

	//delete elements
	delete(m, "service")

	fmt.Println(m)

	//creating map without make() function

	mp := map[string]int{"price":1000, "phones": 2}
	
	fmt.Println(mp)
	
	clear(mp)

	m1 := map[string]int{"price":1000, "phones": 2}
	m2 := map[string]int{"price":1000, "phones": 2}

	fmt.Println(maps.Equal(m1,m2))

	v, ok := m1["price"]

	fmt.Println("=================")
	fmt.Println(v)

	if ok {
		fmt.Println("All ok")
	} else {
		fmt.Println("Not ok")
	}
}
