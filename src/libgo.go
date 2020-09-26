package main

// #cgo windows CPPFLAGS: -D_PLATFORM_WINDOWS=1
// #cgo android CPPFLAGS: -march=armv7-a -mfpu=neon -mfloat-abi=hard
// #cgo android LDFLAGS:  -Wl,--no-warn-mismatch
// #include "helper.h"
// #include <stdlib.h>
import "C"   // 必须紧跟c代码块，不能有空格


import (
    "fmt"
    "libgo/libmisc"
    "unsafe"
)

func main(){ }

//export go_Add
func go_Add(a, b int) int{
    libmisc.Log("I", "%d + %d = %d\n", a, b, a+ b)

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
    libmisc.Log( "I", "go received msg %s\n", C.GoString( c ) )
    return -1
}


// type Func_Callback func( *C.char ) *C.char

var f_callback C.Callback_S_S = nil

//export go_SetCallbackFunc
func go_SetCallbackFunc( f C.Callback_S_S ) {
    libmisc.Log( "I", "%+v", f )
    f_callback = f ;
}



// ====================================================

