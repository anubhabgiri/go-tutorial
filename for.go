package main

import "fmt"

func main() {

	// different styles of for loop
    i := 1 // short hand for var i int = 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1 // increment within the loop body
    }

	// increment in the loop declaration
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

	// loop without any condition
	// basically equivalent to while(true) in CPP
    for {
        fmt.Println("loop")
        break // "return" is also used
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}