package main

import "fmt"

func main() {
	fmt.Print("without line!\n")
	fmt.Println("with line !")
	fmt.Printf("Hello, %s!\n", "formatting")
	// hey bro this is a comment

	fmt.Println("------------------variable-------------------")

	var a int = 10
	var b int = 20
	var name string = "suraj vishwakarma"

	fmt.Printf("Value of a is %d\n", a)
	fmt.Printf("Value of b is %d\n", b)
	fmt.Printf("Value of name is %s\n", name)

	x := 100
	y := 200
	fmt.Printf("Value of X is %d  and Y is %d \n", x, y)

	fmt.Println("------------------constant-------------------")
	const pi float64 = 3.14
	fmt.Printf("Value of pi is %f\n", pi)
}
