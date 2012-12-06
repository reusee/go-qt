#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

enum {
  CQT_CORE = 0,
  CQT_GUI = 1,
  CQT_NONE = 2,
};

typedef union StackItem {
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

void cqt_set_voidp(StackItem* si, void* p);
void cqt_set_bool(StackItem* si, bool b);
void cqt_set_char(StackItem* si, signed char c);
void cqt_set_uchar(StackItem* si, unsigned char c);
void cqt_set_short(StackItem* si, short i);
void cqt_set_ushort(StackItem* si, unsigned short i);
void cqt_set_int(StackItem* si, int i);
void cqt_set_uint(StackItem* si, unsigned int i);
void cqt_set_long(StackItem* si, long i);
void cqt_set_ulong(StackItem* si, unsigned long i);
void cqt_set_float(StackItem* si, float f);
void cqt_set_double(StackItem* si, double f);
void cqt_set_enum(StackItem* si, long e);
void cqt_set_class(StackItem* si, void* c);

void* cqt_get_voidp(StackItem* si);
bool cqt_get_bool(StackItem* si);
signed char cqt_get_char(StackItem* si);
unsigned char cqt_get_uchar(StackItem* si);
short cqt_get_short(StackItem* si);
unsigned short cqt_get_ushort(StackItem* si);
int cqt_get_int(StackItem* si);
unsigned int cqt_get_uint(StackItem* si);
long cqt_get_long(StackItem* si);
unsigned long cqt_get_ulong(StackItem* si);
float cqt_get_float(StackItem* si);
double cqt_get_double(StackItem* si);
long cqt_get_enum(StackItem* si);
void* cqt_get_class(StackItem* si);

void cqt_init();
void cqt_call(int smoke_symbol, int binding_symbol, char* klass, char* method, void* obj, Stack stack);

#ifdef __cplusplus
}
#endif
