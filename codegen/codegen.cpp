#include <smoke.h>
#include <smoke/qtcore_smoke.h>
#include <smoke/qtgui_smoke.h>

#include <iostream>
#include <stdio.h>
using namespace std;

int main(int argc, char *argv[]) {
  init_qtcore_Smoke();
  init_qtgui_Smoke();

  Smoke* smokes[] = {
    qtcore_Smoke,
    qtgui_Smoke,
  };

  for (auto i = begin(smokes); i != end(smokes); i++) {
    Smoke* smoke = *i;

    for (int i = 0; i < smoke->numClasses; ++i) {
      Smoke::Class klass = smoke->classes[i];
      if (klass.className) {
        cout << klass.className << endl;
      }
    }

    for (int i = 0; i < smoke->numMethods; ++i) {
      Smoke::Method method = smoke->methods[i];
      Smoke::Class klass = smoke->classes[method.classId];
      if (klass.className) {
        cout << klass.className << " ";
      }
      cout << smoke->methodNames[method.name] << " ";
      for (int i = 0; i < method.numArgs; ++i) {
        Smoke::Type arg_type = smoke->types[smoke->argumentList[method.args + i]];
        cout << arg_type.name << " ";
      }
      if (method.ret > 0) {
        Smoke::Type ret_type = smoke->types[method.ret];
        cout << "->" << ret_type.name << " ";
      }
      if (method.flags & Smoke::mf_enum) {
        cout << "ENUM ";
      }
      if (method.flags & Smoke::mf_signal) {
        cout << "SIGNAL ";
      }
      cout << endl;
    }

  }

  return 0;
}
