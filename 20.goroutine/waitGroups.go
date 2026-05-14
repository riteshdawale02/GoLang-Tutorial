package main

import (
	"fmt"
	"sync"
)

func task1(id int, w *sync.WaitGroup) {
	defer w.Done() //defer run's after execution
	fmt.Println("doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task1(i, &wg)

	}
	wg.Wait()
}