package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string , bool) {
	return "java", "c++", true
}

func processIt() func(a int)int {
	return func (a int ) int {
		return 4
	}
}
func main() {

	fmt.Println(add(3,2))
	fmt.Println("Hello")

	//if we don't want to use lang3 value -> add there _
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1,lang3, lang2)


	fn := processIt()

	result := fn(10)

	fmt.Println(result)

}