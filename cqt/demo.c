#include "cqt.h"
#include <stdio.h>

int main(int argc, char** argv) {
  cqt_init();

  StackItem stack[3];
  stack[1].s_voidp = &argc;
  stack[2].s_voidp = argv;

  cqt_call(CQT_CORE, CQT_GUI, "QApplication", "QApplication$?", 0, stack);
  void* app = stack[0].s_voidp;

  cqt_call(CQT_CORE, CQT_GUI, "QWidget", "QWidget", 0, stack);
  void* widget = stack[0].s_voidp;

  cqt_call(CQT_CORE, CQT_NONE, "QWidget", "show", widget, 0);

  cqt_call(CQT_GUI, CQT_NONE, "QApplication", "exec", 0, stack);

  return 0;
}
