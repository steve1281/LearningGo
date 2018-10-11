package main

import ("fmt"
		"sync")


var wg sync.WaitGroup

func foo(c chan int, someValue int){
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	fmt.Println(".......")

	fooVal := make(chan int, 100)  // add a buffer.

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}

	// could range over the channel
	// how do we know fooVal is done?
	wg.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}



}

