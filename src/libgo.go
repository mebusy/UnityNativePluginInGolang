package main

// #include "helper.h"
// #include <stdlib.h>
import "C"   // 必须紧跟c代码块，不能有空格


import (
    "fmt"
    "log"
    "os"
    "unsafe"
)

func main(){ }

//export go_Add
func go_Add(a, b int) int{

    InfoLogger.Printf("%d + %d = %d\n", a, b, a+ b)

    if f_callback != nil {
        // cstr uses C heap, you must free it
        cstr := C.CString(fmt.Sprintf( "calculate: %d", a+b  ) )
        defer C.free( unsafe.Pointer(cstr) )
        C.bridge_func_callback_str_str( f_callback, cstr )
    }
    return a+b
}

//export go_SendMsg
func go_SendMsg( c *C.char ) int {
    InfoLogger.Printf( "go received msg %s\n", C.GoString( c ) )
    return -1
}


// type Func_Callback func( *C.char ) *C.char

var f_callback C.Callback_S_S = nil

//export go_SetCallbackFunc
func go_SetCallbackFunc( f C.Callback_S_S ) {
    InfoLogger.Println( "%+v", f )
    f_callback = f ;
}



// ====================================================

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

