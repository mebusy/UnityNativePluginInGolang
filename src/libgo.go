package main

// #include "helper.h"
import "C"   // 必须紧跟c代码块，不能有空格


import (
    "fmt"
    "log"
)

func main(){ }

//export go_Add
func go_Add(a, b int) int{
    log.Printf("%d + %d = %d\n", a, b, a+ b)

    if f_callback != nil {
        C.bridge_func_callback_str_str( f_callback, C.CString(fmt.Sprintf( "calculate: %d", a+b  )) )
    }

    return a+b
}

//export go_SendMsg
func go_SendMsg( c *C.char ) int {
    fmt.Printf( "go received msg %s\n", C.GoString( c ) )
    return -1
}


// type Func_Callback func( *C.char ) *C.char

var f_callback C.Callback_S_S = nil

//export go_SetCallbackFunc
func go_SetCallbackFunc( f C.Callback_S_S ) {
    f_callback = f ;
}


