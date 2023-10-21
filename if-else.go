package main

import "fmt"

func main() {

	// else is not necessary with if
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

	// brackets are not required for conditions but
	// braces are required
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

	// CANNOT add more than 1 else if
	// A statement can precede conditionals; 
	// any variables declared in this statement are
	// available in the current and all subsequent branches.

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}