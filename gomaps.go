package main

import ("fmt")

func main() {
	fmt.Println("")

	//var grades map[string] float32
	// or
	grades := make(map[string] float32)

	grades["Alpha"] = 42
	grades["Beta"] = 52
	grades["Gamma"] = 62
	grades["Epsilon"] = 72

	fmt.Println(grades)
	TimsGrade := grades["Alpha"]
	fmt.Println(TimsGrade)

	delete(grades, "Alpha")

	fmt.Println(grades)

	for key, value :=  range grades {
		fmt.Println(key,":",value)

	}
}

