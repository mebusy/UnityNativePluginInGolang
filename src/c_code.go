package main

/*
#include "../helper.h"

const char* bridge_func_callback_str_str(Callback_S_S f, const char* s) {
    return f(s);
}
*/
import "C"

// Using //export in a file places a restriction on the preamble: since it is copied into two different C output files, it must not contain any definitions, only declarations. If a file contains both definitions and declarations, then the two output files will produce duplicate symbols and the linker will fail. To avoid this, definitions must be placed in preambles in other files, or in C source files.

