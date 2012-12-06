package qt

// #cgo LDFLAGS: -lcqt
// #include <cqt.h>
// #include <stdlib.h>
import "C"
import (
  "unsafe"
)

func TestHelloWorld() {
  C.cqt_init()
  var stack [3]C.StackItem;
  var argc C.int;
  C.cqt_set_voidp(&stack[1], unsafe.Pointer(&argc))
  C.cqt_set_voidp(&stack[2], nil)
  klass := C.CString("QApplication")
  method := C.CString("QApplication$?")
  C.cqt_call(C.CQT_CORE, C.CQT_GUI, klass, method, nil, &stack[0])
  C.free(unsafe.Pointer(klass))
  C.free(unsafe.Pointer(method))

  app := C.cqt_get_voidp(&stack[0])
  _ = app

  klass = C.CString("QWidget")
  C.cqt_call(C.CQT_CORE, C.CQT_GUI, klass, klass, nil, &stack[0])
  widget := C.cqt_get_voidp(&stack[0])

  method = C.CString("show")
  C.cqt_call(C.CQT_CORE, C.CQT_NONE, klass, method, widget, nil)
  C.free(unsafe.Pointer(klass))
  C.free(unsafe.Pointer(method))

  klass = C.CString("QApplication")
  method = C.CString("exec")
  C.cqt_call(C.CQT_GUI, C.CQT_NONE, klass, method, nil, &stack[0])
  C.free(unsafe.Pointer(klass))
  C.free(unsafe.Pointer(method))
}
