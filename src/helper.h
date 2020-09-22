#ifndef CGO_LIB_HELPER_H
#define CGO_LIB_HELPER_H

typedef char* (*Callback_Str_Str)(const char* arg);
extern const char* bridge_func_callback_str_str(Callback_Str_Str f, char* s);

#endif
