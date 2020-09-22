package main

import (
    "fmt"
    "C"
)

func main(){}

//export go_Add
func go_Add(a, b int) int{
    fmt.Printf("%d + %d = %d\n", a, b, a+ b)
    return a+b
}

