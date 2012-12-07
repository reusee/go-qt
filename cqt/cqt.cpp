#include "cqt.h"

#include <smoke.h>
#include <smoke/qtcore_smoke.h>
#include <smoke/qtgui_smoke.h>

#include <iostream>
using namespace std;

class Binding : public SmokeBinding {
  public:
    Binding(Smoke *s): SmokeBinding(s) {}
    void deleted(Smoke::Index classId, void *obj) {}
    bool callMethod(Smoke::Index methodId, void *obj, Smoke::Stack args, bool isAbstract) {
      return false;
    }
    char* className(Smoke::Index classId) {
      return (char*)smoke->classes[classId].className;
    }
};

Binding *bindings[2];

void cqt_init() {
  init_qtcore_Smoke();
  init_qtgui_Smoke();
  bindings[CQT_CORE] = new Binding(qtcore_Smoke);
  bindings[CQT_GUI] = new Binding(qtgui_Smoke);
}

void cqt_call(int smoke_symbol, int binding_symbol, char* klass_name, char* method_name, void* obj, Stack stack) {
  Smoke *smoke;
  Binding *binding = bindings[binding_symbol];
  switch (smoke_symbol) {
    case CQT_CORE: 
      smoke = qtcore_Smoke;
      break;
    case CQT_GUI: 
      smoke = qtgui_Smoke;
      break;
  }
  Smoke::ModuleIndex classId = smoke->findClass(klass_name);
  Smoke::ModuleIndex methodId = classId.smoke->findMethod(klass_name, method_name);
  Smoke::Class klass = classId.smoke->classes[classId.index];
  Smoke::Method method = methodId.smoke->methods[methodId.smoke->methodMaps[methodId.index].method];
  (*klass.classFn)(method.method, obj, (Smoke::Stack)stack);

  if (binding_symbol != CQT_NONE) {
    Smoke::StackItem bindStack[2];
    obj = stack[0].s_voidp;
    bindStack[1].s_voidp = binding;
    (*klass.classFn)(0, obj, bindStack);
  }
}

void cqt_set_voidp(StackItem* si, void* p) {
  si->s_voidp = p;
}
void cqt_set_bool(StackItem* si, bool b) {
  si->s_bool = b;
}
void cqt_set_char(StackItem* si, signed char c) {
  si->s_char = c;
}
void cqt_set_uchar(StackItem* si, unsigned char c) {
  si->s_uchar = c;
}
void cqt_set_short(StackItem* si, short i) {
  si->s_short = i;
}
void cqt_set_ushort(StackItem* si, unsigned short i) {
  si->s_ushort = i;
}
void cqt_set_int(StackItem* si, int i) {
  si->s_int = i;
}
void cqt_set_uint(StackItem* si, unsigned int i) {
  si->s_uint = i;
}
void cqt_set_long(StackItem* si, long i) {
  si->s_long = i;
}
void cqt_set_ulong(StackItem* si, unsigned long i) {
  si->s_ulong = i;
}
void cqt_set_float(StackItem* si, float f) {
  si->s_float = f;
}
void cqt_set_double(StackItem* si, double f) {
  si->s_double = f;
}
void cqt_set_enum(StackItem* si, long e) {
  si->s_enum = e;
}
void cqt_set_class(StackItem* si, void* c) {
  si->s_class = c;
}

void* cqt_get_voidp(StackItem* si) {
  return si->s_voidp;
}
bool cqt_get_bool(StackItem* si) {
  return si->s_bool;
}
signed char cqt_get_char(StackItem* si) {
  return si->s_char;
}
unsigned char cqt_get_uchar(StackItem* si) {
  return si->s_uchar;
}
short cqt_get_short(StackItem* si) {
  return si->s_short;
}
unsigned short cqt_get_ushort(StackItem* si) {
  return si->s_ushort;
}
int cqt_get_int(StackItem* si) {
  return si->s_int;
}
unsigned int cqt_get_uint(StackItem* si) {
  return si->s_uint;
}
long cqt_get_long(StackItem* si) {
  return si->s_long;
}
unsigned long cqt_get_ulong(StackItem* si) {
  return si->s_ulong;
}
float cqt_get_float(StackItem* si) {
  return si->s_float;
}
double cqt_get_double(StackItem* si) {
  return si->s_double;
}
long cqt_get_enum(StackItem* si) {
  return si->s_enum;
}
void* cqt_get_class(StackItem* si) {
  return si->s_class;
}
