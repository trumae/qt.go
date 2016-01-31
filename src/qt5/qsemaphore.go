package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qsemaphore.h
// dst-file: /src/core/qsemaphore.go
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
  // proto:  int QSemaphore::available();
extern int32_t C_ZNK10QSemaphore9availableEv(void* qthis); // 4
  // proto:  void QSemaphore::~QSemaphore();
extern void C_ZN10QSemaphoreD2Ev(void* qthis); // 4
  // proto:  void QSemaphore::acquire(int n);
extern void C_ZN10QSemaphore7acquireEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSemaphore::tryAcquire(int n);
extern bool C_ZN10QSemaphore10tryAcquireEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QSemaphore::tryAcquire(int n, int timeout);
extern bool C_ZN10QSemaphore10tryAcquireEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QSemaphore::release(int n);
extern void C_ZN10QSemaphore7releaseEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSemaphore::QSemaphore(int n);
extern void* C_ZN10QSemaphoreC2Ei(int32_t arg0); // 3
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

// class sizeof(QSemaphore)=8
type QSemaphore struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// available()
func (this *QSemaphore) Available(args ...interface{}) (ret interface{}) {
  // available()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QSemaphore9availableEv
    // invoke: int available()
    var ret0 = C.C_ZNK10QSemaphore9availableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSemaphore", "available", args)
  }

  return
}

// ~QSemaphore()
func (this *QSemaphore) Freeqsemaphore(args ...interface{}) () {
  // ~QSemaphore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphoreD0Ev
    // invoke: void ~QSemaphore()
    C.C_ZN10QSemaphoreD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSemaphore", "~QSemaphore", args)
  }

  return
}

// acquire(int)
func (this *QSemaphore) Acquire(args ...interface{}) () {
  // acquire(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore7acquireEi
    // invoke: void acquire(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QSemaphore7acquireEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSemaphore", "acquire", args)
  }

  return
}

// tryAcquire(int)
func (this *QSemaphore) Tryacquire(args ...interface{}) (ret interface{}) {
  // tryAcquire(int)
  // tryAcquire(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore10tryAcquireEi
    // invoke: bool tryAcquire(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QSemaphore10tryAcquireEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN10QSemaphore10tryAcquireEii
    // invoke: bool tryAcquire(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QSemaphore10tryAcquireEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSemaphore", "tryAcquire", args)
  }

  return
}

// release(int)
func (this *QSemaphore) Release(args ...interface{}) () {
  // release(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphore7releaseEi
    // invoke: void release(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QSemaphore7releaseEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSemaphore", "release", args)
  }

  return
}

// QSemaphore(int)
func NewQSemaphore(args ...interface{}) *QSemaphore {
  // QSemaphore(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QSemaphoreC1Ei
    // invoke: void QSemaphore(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QSemaphoreC2Ei(arg0)
    return &QSemaphore{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSemaphore", "QSemaphore", args)
  }

  return nil // QSemaphore{qclsinst:qthis}
}

// <= body block end

