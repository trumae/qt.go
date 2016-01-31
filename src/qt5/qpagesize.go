package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qpagesize.h
// dst-file: /src/gui/qpagesize.go
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
  // proto:  int QPageSize::windowsId();
extern int32_t C_ZNK9QPageSize9windowsIdEv(void* qthis); // 4
  // proto:  QSize QPageSize::sizePixels(int resolution);
extern void* C_ZNK9QPageSize10sizePixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QPageSize::rectPixels(int resolution);
extern void* C_ZNK9QPageSize10rectPixelsEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QPageSize::name();
extern void* C_ZNK9QPageSize4nameEv(void* qthis); // 4
  // proto:  QPageSize::Unit QPageSize::definitionUnits();
extern void C_ZNK9QPageSize15definitionUnitsEv(void* qthis); // 4
  // proto:  bool QPageSize::isValid();
extern bool C_ZNK9QPageSize7isValidEv(void* qthis); // 4
  // proto:  bool QPageSize::isEquivalentTo(const QPageSize & other);
extern bool C_ZNK9QPageSize14isEquivalentToERKS_(void* qthis, void* arg0); // 4
  // proto:  void QPageSize::~QPageSize();
extern void C_ZN9QPageSizeD2Ev(void* qthis); // 4
  // proto:  QSize QPageSize::sizePoints();
extern void* C_ZNK9QPageSize10sizePointsEv(void* qthis); // 4
  // proto:  void QPageSize::swap(QPageSize & other);
extern void C_ZN9QPageSize4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QString QPageSize::key();
extern void* C_ZNK9QPageSize3keyEv(void* qthis); // 4
  // proto:  QRect QPageSize::rectPoints();
extern void* C_ZNK9QPageSize10rectPointsEv(void* qthis); // 4
  // proto:  void QPageSize::QPageSize(const QPageSize & other);
extern void* C_ZN9QPageSizeC2ERKS_(void* arg0); // 3
  // proto:  void QPageSize::QPageSize();
extern void* C_ZN9QPageSizeC2Ev(); // 3
  // proto: static QPageSize::PageSizeId QPageSize::id(int windowsId);
extern void C_ZN9QPageSize2idEi(int32_t arg0); // 4
  // proto:  QPageSize::PageSizeId QPageSize::id();
extern void C_ZNK9QPageSize2idEv(void* qthis); // 4
  // proto:  QSizeF QPageSize::definitionSize();
extern void* C_ZNK9QPageSize14definitionSizeEv(void* qthis); // 4
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

// class sizeof(QPageSize)=1
type QPageSize struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// windowsId()
func (this *QPageSize) Windowsid(args ...interface{}) (ret interface{}) {
  // windowsId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize9windowsIdEv
    // invoke: int windowsId()
    var ret0 = C.C_ZNK9QPageSize9windowsIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "windowsId", args)
  }

  return
}

// sizePixels(int)
func (this *QPageSize) Sizepixels(args ...interface{}) (ret interface{}) {
  // sizePixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10sizePixelsEi
    // invoke: QSize sizePixels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QPageSize10sizePixelsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "sizePixels", args)
  }

  return
}

// rectPixels(int)
func (this *QPageSize) Rectpixels(args ...interface{}) (ret interface{}) {
  // rectPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10rectPixelsEi
    // invoke: QRect rectPixels(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QPageSize10rectPixelsEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "rectPixels", args)
  }

  return
}

// name()
func (this *QPageSize) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize4nameEv
    // invoke: QString name()
    var ret0 = C.C_ZNK9QPageSize4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "name", args)
  }

  return
}

// definitionUnits()
func (this *QPageSize) Definitionunits(args ...interface{}) () {
  // definitionUnits()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize15definitionUnitsEv
    // invoke: QPageSize::Unit definitionUnits()
    C.C_ZNK9QPageSize15definitionUnitsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageSize", "definitionUnits", args)
  }

  return
}

// isValid()
func (this *QPageSize) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK9QPageSize7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "isValid", args)
  }

  return
}

// isEquivalentTo(const class QPageSize &)
func (this *QPageSize) Isequivalentto(args ...interface{}) (ret interface{}) {
  // isEquivalentTo(const class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize14isEquivalentToERKS_
    // invoke: bool isEquivalentTo(const class QPageSize &)
    var arg0 = args[0].(QPageSize).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QPageSize14isEquivalentToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "isEquivalentTo", args)
  }

  return
}

// ~QPageSize()
func (this *QPageSize) Freeqpagesize(args ...interface{}) () {
  // ~QPageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPageSizeD0Ev
    // invoke: void ~QPageSize()
    C.C_ZN9QPageSizeD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageSize", "~QPageSize", args)
  }

  return
}

// sizePoints()
func (this *QPageSize) Sizepoints(args ...interface{}) (ret interface{}) {
  // sizePoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10sizePointsEv
    // invoke: QSize sizePoints()
    var ret0 = C.C_ZNK9QPageSize10sizePointsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "sizePoints", args)
  }

  return
}

// swap(class QPageSize &)
func (this *QPageSize) Swap(args ...interface{}) () {
  // swap(class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize4swapERS_
    // invoke: void swap(class QPageSize &)
    var arg0 = args[0].(QPageSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QPageSize4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPageSize", "swap", args)
  }

  return
}

// key()
func (this *QPageSize) Key(args ...interface{}) (ret interface{}) {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize3keyEv
    // invoke: QString key()
    var ret0 = C.C_ZNK9QPageSize3keyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "key", args)
  }

  return
}

// rectPoints()
func (this *QPageSize) Rectpoints(args ...interface{}) (ret interface{}) {
  // rectPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10rectPointsEv
    // invoke: QRect rectPoints()
    var ret0 = C.C_ZNK9QPageSize10rectPointsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "rectPoints", args)
  }

  return
}

// QPageSize(const class QPageSize &)
func NewQPageSize(args ...interface{}) *QPageSize {
  // QPageSize(const class QPageSize &)
  // QPageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPageSizeC1ERKS_
    // invoke: void QPageSize(const class QPageSize &)
    var arg0 = args[0].(QPageSize).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QPageSizeC2ERKS_(arg0)
    return &QPageSize{qclsinst:qthis}
  case 1:
    // invoke: _ZN9QPageSizeC1Ev
    // invoke: void QPageSize()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QPageSizeC2Ev()
    return &QPageSize{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPageSize", "QPageSize", args)
  }

  return nil // QPageSize{qclsinst:qthis}
}

// id(int)
func (this *QPageSize) Id_S(args ...interface{}) () {
  // id(int)
  // id()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize2idEi
    // invoke: QPageSize::PageSizeId id(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QPageSize2idEi(arg0)
  case 1:
    // invoke: _ZNK9QPageSize2idEv
    // invoke: QPageSize::PageSizeId id()
    C.C_ZNK9QPageSize2idEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPageSize", "id", args)
  }

  return
}

// definitionSize()
func (this *QPageSize) Definitionsize(args ...interface{}) (ret interface{}) {
  // definitionSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize14definitionSizeEv
    // invoke: QSizeF definitionSize()
    var ret0 = C.C_ZNK9QPageSize14definitionSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPageSize", "definitionSize", args)
  }

  return
}

// <= body block end

