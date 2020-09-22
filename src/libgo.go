package main

import (
    "fmt"
    "C"
)

func main(){}

//export Add
func Add(a, b int) int{
    fmt.Printf("%d + %d = %d\n", a, b, a+ b)
    return a+b
}
