package main

import "fmt"

func main() {

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	/* you can leave stuff out. for example:
	for {
		fmt.Println("forever...")
	}
	*/

	// you can break the for loop
	x := 5
	for  {
		fmt.Println("Do stuff", x)
		x +=3
		if x> 25 {
			break;
		}
	}


}

