package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qrunnable.h
// dst-file: /src/core/qrunnable.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QRunnable::setAutoDelete(bool _autoDelete);
extern void C_ZN9QRunnable13setAutoDeleteEb(void* qthis, bool arg0); // 2
  // proto:  bool QRunnable::autoDelete();
extern bool C_ZNK9QRunnable10autoDeleteEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QRunnable)=16
type QRunnable struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setAutoDelete(_Bool)
func (this *QRunnable) Setautodelete(args ...interface{}) () {
  // setAutoDelete(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QRunnable13setAutoDeleteEb
    // invoke: void setAutoDelete(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QRunnable13setAutoDeleteEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QRunnable", "setAutoDelete", args)
  }

  return
}

// autoDelete()
func (this *QRunnable) Autodelete(args ...interface{}) (ret interface{}) {
  // autoDelete()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QRunnable10autoDeleteEv
    // invoke: bool autoDelete()
    var ret0 = C.C_ZNK9QRunnable10autoDeleteEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QRunnable", "autoDelete", args)
  }

  return
}

// <= body block end

