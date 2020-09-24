// #if _MSC_VER // this is defined when compiling with Visual Studio
#ifdef _PLATFORM_WINDOWS
#define EXPORT_API __declspec(dllexport) // Visual Studio needs annotating exported functions with this
#else
#define EXPORT_API // XCode does not need annotating exported functions, so define is empty
#endif

// ------------------------------------------------------------------------
// Plugin itself

#include "libgo.h"

#include <stdio.h>


// Link following functions C-style (required for plugins)
extern "C"
{

// The functions we will call from Unity.
//
EXPORT_API extern GoInt lib_Add(GoInt p0, GoInt p1) {
	// test whether _WINDOWS_ is defined
	#ifdef _PLATFORM_WINDOWS
	p1 += 1 ;
	#endif
    return go_Add(p0, p1);
}

EXPORT_API extern GoInt lib_SendMsg(char* p0) {
    return go_SendMsg( p0 );
}

EXPORT_API extern void lib_SetCallbackFunc(Callback_S_S p0) {
    /*
    char * ret = p0( "from lib" ) ;
    printf( "%s:" , ret ) ;
    /*/
    go_SetCallbackFunc( p0 );
    //*/
}

} // end of export C block
