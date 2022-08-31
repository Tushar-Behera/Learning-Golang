package main

import "fmt"

func main() {

//_______ Variable Manupulation______

//1st method ------>

	var a int // static declaration of variable
	a = 50    // initialization of variable

	fmt.Println("value of a =", a)

// 2nd method ----->

	var b float32 = 40.5 // declartion with initialization
	//or
	var c = "Tushar" // declaraton with initialization

	fmt.Println("value of b =", b)
	fmt.Println("value of c =", c)

// 3rd type ----->

	d := [3]string{"Austin", "jackob", "markus"} // direct initialization or dynamic [this is an array but type is according to the value or type dclared]

	fmt.Println("Value of d =", d)

// To check type of a Variable use  "printf" and [ %T ] only for type and [ %t ] for type with value

	fmt.Printf("\nType of [a] is %T\n", a)
	fmt.Printf("Type of [b] is %T\n", b)
	fmt.Printf("Type of [c] is %T\n", c)
	fmt.Printf("Type of [d] is %T\n", d)

	fmt.Printf("\n\nType of [a] is %t\n", a)
	fmt.Printf("Type of [b] is %t\n", b)
	fmt.Printf("Type of [c] is %t\n", c)
	fmt.Printf("Type of [d] is %t\n", d)

}
