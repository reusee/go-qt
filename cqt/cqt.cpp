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
    bool callMethod(Smoke::Index methodId, void *obj, Smoke::Stack, bool) {
      Smoke::Method method = smoke->methods[methodId];
      string name;

      if (method.flags & Smoke::mf_protected) name += "protected ";
      if (method.flags & Smoke::mf_const) name += "const ";

      name += smoke->methodNames[method.name] + string("(");

      Smoke::Index *idx = smoke->argumentList + method.args;
      while (*idx) {
        name += smoke->types[*idx].name;
        idx++;
        if (*idx) name += ",";
      }
      name += ")";

      cout << name << endl;

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
