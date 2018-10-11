package main

import ("fmt"
		"time")


func say( s string) {
	for  i:=0; i<3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	fmt.Println("go routines...")
	// linear approach
	say("Hey")
	say("There")
	fmt.Println(".......")
	// make one a go routine (light weight thread)
	go say("Hey")
	say("There")
	fmt.Println(".......")
	// but if both are go, the program exists before output
	go say("Hey")
	go say("There")
	// could fix with a delay..
	// time.Sleep(time.Second)
	fmt.Println(".......")
}

