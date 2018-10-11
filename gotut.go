package main

import ("fmt"
        "math"
        "math/rand")

func foo() {
	fmt.Println("Welcome to go, sqrt of 4 is ", math.Sqrt(4))
}

func bar() {
	fmt.Println("A number from 0-100", rand.Intn(100))
}

func main() {
	foo();
	bar();
}

