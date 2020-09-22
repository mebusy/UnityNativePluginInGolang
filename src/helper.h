#ifndef CGO_LIB_HELPER_H
#define CGO_LIB_HELPER_H

typedef char* (*Callback_S_S)(const char* arg);
extern const char* bridge_func_callback_str_str(Callback_S_S f, char* s);

#endif
