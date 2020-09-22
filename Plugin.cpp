#if _MSC_VER // this is defined when compiling with Visual Studio
#define EXPORT_API __declspec(dllexport) // Visual Studio needs annotating exported functions with this
#else
#define EXPORT_API // XCode does not need annotating exported functions, so define is empty
#endif

// ------------------------------------------------------------------------
// Plugin itself

#include "libgo.h"

// Link following functions C-style (required for plugins)
extern "C"
{

// The functions we will call from Unity.
//
EXPORT_API extern GoInt lib_Add(GoInt p0, GoInt p1) {
    return go_Add(p0, p1);
}  


} // end of export C block
