package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qsurface.h
// dst-file: /src/gui/qsurface.go
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
  // proto:  void QSurface::~QSurface();
extern void C_ZN8QSurfaceD2Ev(void* qthis); // 4
  // proto:  QSurface::SurfaceClass QSurface::surfaceClass();
extern void C_ZNK8QSurface12surfaceClassEv(void* qthis); // 4
  // proto:  bool QSurface::supportsOpenGL();
extern bool C_ZNK8QSurface14supportsOpenGLEv(void* qthis); // 4
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

// class sizeof(QSurface)=24
type QSurface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QSurface()
func (this *QSurface) Freeqsurface(args ...interface{}) () {
  // ~QSurface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSurfaceD0Ev
    // invoke: void ~QSurface()
    C.C_ZN8QSurfaceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurface", "~QSurface", args)
  }

  return
}

// surfaceClass()
func (this *QSurface) Surfaceclass(args ...interface{}) () {
  // surfaceClass()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSurface12surfaceClassEv
    // invoke: QSurface::SurfaceClass surfaceClass()
    C.C_ZNK8QSurface12surfaceClassEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurface", "surfaceClass", args)
  }

  return
}

// supportsOpenGL()
func (this *QSurface) Supportsopengl(args ...interface{}) (ret interface{}) {
  // supportsOpenGL()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSurface14supportsOpenGLEv
    // invoke: bool supportsOpenGL()
    var ret0 = C.C_ZNK8QSurface14supportsOpenGLEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSurface", "supportsOpenGL", args)
  }

  return
}

// <= body block end

