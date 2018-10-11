package main

import "fmt" 

func main() {
	x := 15
	a := &x // memory arrdress

	fmt.Println("Memory", a)
	fmt.Println("Stored at a ", *a)
	*a = 5
	fmt.Println("*a = 5, so now x is: ", x)

	*a = *a * *a
	
	fmt.Println("*a = *a**a, so now x is: ", x)
}

