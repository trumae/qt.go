package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qwidgetaction.h
// dst-file: /src/widgets/qwidgetaction.go
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
  // proto:  void QWidgetAction::releaseWidget(QWidget * widget);
extern void C_ZN13QWidgetAction13releaseWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QWidgetAction::metaObject();
extern void C_ZNK13QWidgetAction10metaObjectEv(void* qthis); // 4
  // proto:  void QWidgetAction::QWidgetAction(QObject * parent);
extern void* C_ZN13QWidgetActionC2EP7QObject(void* arg0); // 3
  // proto:  QWidget * QWidgetAction::defaultWidget();
extern void* C_ZNK13QWidgetAction13defaultWidgetEv(void* qthis); // 4
  // proto:  QWidget * QWidgetAction::requestWidget(QWidget * parent);
extern void* C_ZN13QWidgetAction13requestWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QWidgetAction::~QWidgetAction();
extern void C_ZN13QWidgetActionD2Ev(void* qthis); // 4
  // proto:  void QWidgetAction::setDefaultWidget(QWidget * w);
extern void C_ZN13QWidgetAction16setDefaultWidgetEP7QWidget(void* qthis, void* arg0); // 4
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

// class sizeof(QWidgetAction)=1
type QWidgetAction struct {
  /*qbase*/ QAction;
  qclsinst unsafe.Pointer /* *C.void */;
}

// releaseWidget(class QWidget *)
func (this *QWidgetAction) Releasewidget(args ...interface{}) () {
  // releaseWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction13releaseWidgetEP7QWidget
    // invoke: void releaseWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QWidgetAction13releaseWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetAction", "releaseWidget", args)
  }

  return
}

// metaObject()
func (this *QWidgetAction) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetAction10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QWidgetAction10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetAction", "metaObject", args)
  }

  return
}

// QWidgetAction(class QObject *)
func NewQWidgetAction(args ...interface{}) *QWidgetAction {
  // QWidgetAction(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetActionC1EP7QObject
    // invoke: void QWidgetAction(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QWidgetActionC2EP7QObject(arg0)
    return &QWidgetAction{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QWidgetAction", "QWidgetAction", args)
  }

  return nil // QWidgetAction{qclsinst:qthis}
}

// defaultWidget()
func (this *QWidgetAction) Defaultwidget(args ...interface{}) (ret interface{}) {
  // defaultWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetAction13defaultWidgetEv
    // invoke: QWidget * defaultWidget()
    var ret0 = C.C_ZNK13QWidgetAction13defaultWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWidgetAction", "defaultWidget", args)
  }

  return
}

// requestWidget(class QWidget *)
func (this *QWidgetAction) Requestwidget(args ...interface{}) (ret interface{}) {
  // requestWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction13requestWidgetEP7QWidget
    // invoke: QWidget * requestWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QWidgetAction13requestWidgetEP7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QWidgetAction", "requestWidget", args)
  }

  return
}

// ~QWidgetAction()
func (this *QWidgetAction) Freeqwidgetaction(args ...interface{}) () {
  // ~QWidgetAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetActionD0Ev
    // invoke: void ~QWidgetAction()
    C.C_ZN13QWidgetActionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetAction", "~QWidgetAction", args)
  }

  return
}

// setDefaultWidget(class QWidget *)
func (this *QWidgetAction) Setdefaultwidget(args ...interface{}) () {
  // setDefaultWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction16setDefaultWidgetEP7QWidget
    // invoke: void setDefaultWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QWidgetAction16setDefaultWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetAction", "setDefaultWidget", args)
  }

  return
}

// <= body block end

