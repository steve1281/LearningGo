package main

import ("fmt"
		"time"
		"sync")

var wg sync.WaitGroup

func say( s string) {
	defer wg.Done() // runs when function ends, or "panics"
	for  i:=0; i<3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
	fmt.Println(".......")
}

