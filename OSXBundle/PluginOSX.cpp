
#include "libgo.h"

// Link following functions C-style (required for plugins)
// for building OSX bundle Only
extern "C"
{

GoInt lib_SendMsg(char* p0) {
    return go_SendMsg( p0 );
}

} // end of export C block
