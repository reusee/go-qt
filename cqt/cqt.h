#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

enum {
  CQT_CORE = 0,
  CQT_GUI = 1,
  CQT_NONE = 2,
};

typedef union {
  void* s_voidp;
  bool s_bool;
  signed char s_char;
  unsigned char s_uchar;
  short s_short;
  unsigned short s_ushort;
  int s_int;
  unsigned int s_uint;
  long s_long;
  unsigned long s_ulong;
  float s_float;
  double s_double;
  long s_enum;
  void* s_class;
} StackItem;
typedef StackItem* Stack;

void cqt_init();
void cqt_call(int smoke_symbol, int binding_symbol, char* klass, char* method, void* obj, Stack stack);

#ifdef __cplusplus
}
#endif
