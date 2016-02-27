package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qkeyeventtransition.h
// dst-file: /src/widgets/qkeyeventtransition.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QKeyEventTransition::metaObject();
extern void C_ZNK19QKeyEventTransition10metaObjectEv(void* qthis); // 4
  // proto:  int QKeyEventTransition::key();
extern int32_t C_ZNK19QKeyEventTransition3keyEv(void* qthis); // 4
  // proto:  Qt::KeyboardModifiers QKeyEventTransition::modifierMask();
extern void C_ZNK19QKeyEventTransition12modifierMaskEv(void* qthis); // 4
  // proto:  void QKeyEventTransition::setKey(int key);
extern void C_ZN19QKeyEventTransition6setKeyEi(void* qthis, int32_t arg0); // 4
  // proto:  void QKeyEventTransition::~QKeyEventTransition();
extern void C_ZN19QKeyEventTransitionD2Ev(void* qthis); // 4
  // proto:  void QKeyEventTransition::QKeyEventTransition(QState * sourceState);
extern void* C_ZN19QKeyEventTransitionC2EP6QState(void* arg0); // 3
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QKeyEventTransition)=1
type QKeyEventTransition struct {
  /*qbase*/ qtcore.QEventTransition;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QKeyEventTransition) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QKeyEventTransition10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "metaObject", args)
  }

  return
}

// key()
func (this *QKeyEventTransition) Key(args ...interface{}) (ret interface{}) {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition3keyEv
    // invoke: int key()
    var ret0 = C.C_ZNK19QKeyEventTransition3keyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "key", args)
  }

  return
}

// modifierMask()
func (this *QKeyEventTransition) ModifierMask(args ...interface{}) () {
  // modifierMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QKeyEventTransition12modifierMaskEv
    // invoke: Qt::KeyboardModifiers modifierMask()
    C.C_ZNK19QKeyEventTransition12modifierMaskEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "modifierMask", args)
  }

  return
}

// setKey(int)
func (this *QKeyEventTransition) SetKey(args ...interface{}) () {
  // setKey(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QKeyEventTransition6setKeyEi
    // invoke: void setKey(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN19QKeyEventTransition6setKeyEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "setKey", args)
  }

  return
}

// ~QKeyEventTransition()
func (this *QKeyEventTransition) Free(args ...interface{}) () {
  // ~QKeyEventTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QKeyEventTransitionD0Ev
    // invoke: void ~QKeyEventTransition()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN19QKeyEventTransitionD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "~QKeyEventTransition", args)
  }

  return
}

// QKeyEventTransition(class QState *)
func GcfreeQKeyEventTransition(this *QKeyEventTransition) {
  qtrt.UniverseFree(this)
}
func NewQKeyEventTransition(args ...interface{}) *QKeyEventTransition {
  // QKeyEventTransition(class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QKeyEventTransitionC1EP6QState
    // invoke: void QKeyEventTransition(class QState *)
    var arg0 = args[0].(*qtcore.QState).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QKeyEventTransitionC2EP6QState(arg0)
    this := &QKeyEventTransition{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQKeyEventTransition)
    return this
  default:
    qtrt.ErrorResolve("QKeyEventTransition", "QKeyEventTransition", args)
  }

  return nil // QKeyEventTransition{Qclsinst:qthis}
}

// <= body block end
