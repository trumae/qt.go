package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qstylepainter.h
// dst-file: /src/widgets/qstylepainter.go
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
  // proto:  void QStylePainter::drawItemPixmap(const QRect & r, int flags, const QPixmap & pixmap);
extern void C_ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap(void* qthis, void* arg0, int32_t arg1, void* arg2); // 2
  // proto:  void QStylePainter::QStylePainter(QWidget * w);
extern void* C_ZN13QStylePainterC2EP7QWidget(void* arg0); // 1
  // proto:  void QStylePainter::QStylePainter(QPaintDevice * pd, QWidget * w);
extern void* C_ZN13QStylePainterC2EP12QPaintDeviceP7QWidget(void* arg0, void* arg1); // 1
  // proto:  void QStylePainter::QStylePainter();
extern void* C_ZN13QStylePainterC2Ev(); // 1
  // proto:  bool QStylePainter::begin(QPaintDevice * pd, QWidget * w);
extern bool C_ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget(void* qthis, void* arg0, void* arg1); // 2
  // proto:  bool QStylePainter::begin(QWidget * w);
extern bool C_ZN13QStylePainter5beginEP7QWidget(void* qthis, void* arg0); // 2
  // proto:  QStyle * QStylePainter::style();
extern void* C_ZNK13QStylePainter5styleEv(void* qthis); // 2
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

// class sizeof(QStylePainter)=1
type QStylePainter struct {
  /*qbase*/ QPainter;
  qclsinst unsafe.Pointer /* *C.void */;
}

// drawItemPixmap(const class QRect &, int, const class QPixmap &)
func (this *QStylePainter) Drawitempixmap(args ...interface{}) () {
  // drawItemPixmap(const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap
    // invoke: void drawItemPixmap(const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPixmap).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStylePainter", "drawItemPixmap", args)
  }

  return
}

// QStylePainter(class QWidget *)
func NewQStylePainter(args ...interface{}) *QStylePainter {
  // QStylePainter(class QWidget *)
  // QStylePainter(class QPaintDevice *, class QWidget *)
  // QStylePainter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStylePainterC1EP7QWidget
    // invoke: void QStylePainter(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStylePainterC2EP7QWidget(arg0)
    return &QStylePainter{qclsinst:qthis}
  case 1:
    // invoke: _ZN13QStylePainterC1EP12QPaintDeviceP7QWidget
    // invoke: void QStylePainter(class QPaintDevice *, class QWidget *)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStylePainterC2EP12QPaintDeviceP7QWidget(arg0, arg1)
    return &QStylePainter{qclsinst:qthis}
  case 2:
    // invoke: _ZN13QStylePainterC1Ev
    // invoke: void QStylePainter()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStylePainterC2Ev()
    return &QStylePainter{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStylePainter", "QStylePainter", args)
  }

  return nil // QStylePainter{qclsinst:qthis}
}

// begin(class QPaintDevice *, class QWidget *)
func (this *QStylePainter) Begin(args ...interface{}) (ret interface{}) {
  // begin(class QPaintDevice *, class QWidget *)
  // begin(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget
    // invoke: bool begin(class QPaintDevice *, class QWidget *)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN13QStylePainter5beginEP7QWidget
    // invoke: bool begin(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QStylePainter5beginEP7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QStylePainter", "begin", args)
  }

  return
}

// style()
func (this *QStylePainter) Style(args ...interface{}) (ret interface{}) {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStylePainter5styleEv
    // invoke: QStyle * style()
    var ret0 = C.C_ZNK13QStylePainter5styleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QStylePainter", "style", args)
  }

  return
}

// <= body block end

