package main

import (
    "fmt"
    "C"
)

func main(){ }

//export go_Add
func go_Add(a, b int) int{
    fmt.Printf("%d + %d = %d\n", a, b, a+ b)

    if f_callback != nil {
        ret := f_callback( fmt.Sprintf( "result is %d", a+b  )  )
        fmt.Printf("callback resp: %s\n", ret )
    }

    return a+b
}

//export go_SendMsg
func go_SendMsg( msg string ) int {
    fmt.Printf( "go received msg %s\n", msg )
    return -1
}


type Func_Callback func( string ) string

var f_callback Func_Callback = nil

//export go_SetCallbackFunc
func go_SetCallbackFunc( f Func_Callback ) {
    f_callback = f ;
}
