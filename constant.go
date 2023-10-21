package main

import (
    "fmt"
    "math"
)
// keyword "const"
const s string = "constant"

func main() {
    // s is globally declared and accessed inside main
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d)) // type casting to long int

    fmt.Println(math.Sin(n)) // using a math function on n
}