package main

import ("fmt")

func add(x ,y float64) float64 {
	return x+y
}

func multiple( a, b string) (string, string) {
	return a,b
}

func main() {
	// quite a few types
	//var num1 float64 = 5.6
	//var num2 float64 = 9.5

	//var num1,num2 float64 = 5.6, 9.5

	num1,num2 := 5.6, 9.5 // go will make this a float64.

	fmt.Println(add(num1, num2));

	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1, w2));

	var a int = 62;
	var b float64 = float64(a);

	x := a //  will be type int.

	fmt.Println(b,x);



}

