package main

import (
    "fmt"
    "slices"
)

func main() {

	/* 
	To some extent slices are like vectors in CPP
	Unlike arrays, slices are typed only by the elements they contain (not the number of elements). 
	An uninitialized slice equals to nil and has length 0.
	*/

    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)


	/* 
	To create an empty slice with non-zero length, use the builtin make. 
	Here we make a slice of strings of length 3 (initially zero-valued). 
	By default a new slice’s capacity is equal to its length; if we know the slice is going to grow ahead of time, 
	it’s possible to pass a capacity explicitly as an additional parameter to make.
	*/
    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// We can set and get just like with arrays.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

	/*
	Slices support a “slice” operator with the syntax slice[low:high]. 
	For example, this gets a slice of the elements s[2], s[3], and s[4]. 
	*/
    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }

	/*Slices can be composed into multi-dimensional data structures. 
	The length of the inner slices can vary, unlike with multi-dimensional arrays. */
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}