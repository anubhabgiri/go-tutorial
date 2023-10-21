package main

import "fmt"

func main() {
	/*
	* keyword "var"
	* we can declare and assign 1 or more variables in a single statement
	* 
	* We can optionally declare the type of variable explicitly
	*
	* Even without explicit type declaration, type is associated to a variable
	* after assigning value
	*
	* For each type a default value is associated
	*/

    var a = "initial"
    fmt.Println(a)

	// var a1 = a + 2 
	// this will throw invalid operation: 
	// a + 2 (mismatched types string and untyped int)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

	// this will be initialized to a default value of 0 
    var e int
    fmt.Println(e)

    f := "apple" 
	// equivalent to var f string = "apple"
	// := is shorthand for declaration and assignment
	// this works only inside functions
    fmt.Println(f)

	g := 2
	fmt.Println(g)

	// This block will throw syntax error during build
	// since we are not declaring the type or assigning a value
	// var h 
	// fmt.Println(h)
}