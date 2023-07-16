package main

import (
	"fmt"
	"sync"
)

func Printsome(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("The name is %v\n", str)

}

func main() {
	var wg sync.WaitGroup
	names := []string{
		"Kunal",
		"Ram",
		"Bholu", "Rajiv", "Jethmalani",
	}

	for _, val := range names {
		go Printsome(val, &wg)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("End of program")
}
