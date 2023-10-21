package main

import "fmt"

func main() {

    fmt.Println("go" + "lang") // string is concatenated

    fmt.Println("1+1 =", 1+1) // sum 
    fmt.Println("7.0/3.0 =", 7.0/3.0) // division of floats
	fmt.Println("7/3 =", 7/3) // division of integers

	fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}