package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qmap.h
// dst-file: /src/core/qmap.go
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
  // proto: static void QMapDataBase::freeData(QMapDataBase * d);
extern void C_ZN12QMapDataBase8freeDataEPS_(void* arg0); // 4
  // proto:  QMapNodeBase * QMapDataBase::createNode(int size, int alignment, QMapNodeBase * parent, bool left);
extern void* C_ZN12QMapDataBase10createNodeEiiP12QMapNodeBaseb(void* qthis, int32_t arg0, int32_t arg1, void* arg2, bool arg3); // 4
  // proto:  void QMapDataBase::rotateRight(QMapNodeBase * x);
extern void C_ZN12QMapDataBase11rotateRightEP12QMapNodeBase(void* qthis, void* arg0); // 4
  // proto:  void QMapDataBase::freeTree(QMapNodeBase * root, int alignment);
extern void C_ZN12QMapDataBase8freeTreeEP12QMapNodeBasei(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QMapDataBase::recalcMostLeftNode();
extern void C_ZN12QMapDataBase18recalcMostLeftNodeEv(void* qthis); // 4
  // proto: static QMapDataBase * QMapDataBase::createData();
extern void* C_ZN12QMapDataBase10createDataEv(); // 4
  // proto:  void QMapDataBase::rotateLeft(QMapNodeBase * x);
extern void C_ZN12QMapDataBase10rotateLeftEP12QMapNodeBase(void* qthis, void* arg0); // 4
  // proto:  void QMapDataBase::rebalance(QMapNodeBase * x);
extern void C_ZN12QMapDataBase9rebalanceEP12QMapNodeBase(void* qthis, void* arg0); // 4
  // proto:  void QMapDataBase::freeNodeAndRebalance(QMapNodeBase * z);
extern void C_ZN12QMapDataBase20freeNodeAndRebalanceEP12QMapNodeBase(void* qthis, void* arg0); // 4
  // proto:  QMapNodeBase * QMapNodeBase::previousNode();
extern void* C_ZN12QMapNodeBase12previousNodeEv(void* qthis); // 2
  // proto:  QMapNodeBase * QMapNodeBase::parent();
extern void* C_ZNK12QMapNodeBase6parentEv(void* qthis); // 2
  // proto:  QMapNodeBase::Color QMapNodeBase::color();
extern void C_ZNK12QMapNodeBase5colorEv(void* qthis); // 2
  // proto:  void QMapNodeBase::setParent(QMapNodeBase * pp);
extern void C_ZN12QMapNodeBase9setParentEPS_(void* qthis, void* arg0); // 2
  // proto:  QMapNodeBase * QMapNodeBase::nextNode();
extern void* C_ZN12QMapNodeBase8nextNodeEv(void* qthis); // 2
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

// class sizeof(QMapDataBase)=1
type QMapDataBase struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMapNodeBase)=24
type QMapNodeBase struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// freeData(struct QMapDataBase *)
func (this *QMapDataBase) Freedata_S(args ...interface{}) () {
  // freeData(struct QMapDataBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapDataBase{}) // "QMapDataBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase8freeDataEPS_
    // invoke: void freeData(struct QMapDataBase *)
    var arg0 = args[0].(QMapDataBase).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QMapDataBase8freeDataEPS_(arg0)
  default:
    qtrt.ErrorResolve("QMapDataBase", "freeData", args)
  }

  return
}

// createNode(int, int, struct QMapNodeBase *, _Bool)
func (this *QMapDataBase) Createnode(args ...interface{}) (ret interface{}) {
  // createNode(int, int, struct QMapNodeBase *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
  vtys[0][3] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase10createNodeEiiP12QMapNodeBaseb
    // invoke: QMapNodeBase * createNode(int, int, struct QMapNodeBase *, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QMapNodeBase).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.bool(args[3].(bool))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN12QMapDataBase10createNodeEiiP12QMapNodeBaseb(this.qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMapDataBase", "createNode", args)
  }

  return
}

// rotateRight(struct QMapNodeBase *)
func (this *QMapDataBase) Rotateright(args ...interface{}) () {
  // rotateRight(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase11rotateRightEP12QMapNodeBase
    // invoke: void rotateRight(struct QMapNodeBase *)
    var arg0 = args[0].(QMapNodeBase).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QMapDataBase11rotateRightEP12QMapNodeBase(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMapDataBase", "rotateRight", args)
  }

  return
}

// freeTree(struct QMapNodeBase *, int)
func (this *QMapDataBase) Freetree(args ...interface{}) () {
  // freeTree(struct QMapNodeBase *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase8freeTreeEP12QMapNodeBasei
    // invoke: void freeTree(struct QMapNodeBase *, int)
    var arg0 = args[0].(QMapNodeBase).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QMapDataBase8freeTreeEP12QMapNodeBasei(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMapDataBase", "freeTree", args)
  }

  return
}

// recalcMostLeftNode()
func (this *QMapDataBase) Recalcmostleftnode(args ...interface{}) () {
  // recalcMostLeftNode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase18recalcMostLeftNodeEv
    // invoke: void recalcMostLeftNode()
    C.C_ZN12QMapDataBase18recalcMostLeftNodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMapDataBase", "recalcMostLeftNode", args)
  }

  return
}

// createData()
func (this *QMapDataBase) Createdata_S(args ...interface{}) (ret interface{}) {
  // createData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase10createDataEv
    // invoke: QMapDataBase * createData()
    var ret0 = C.C_ZN12QMapDataBase10createDataEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMapDataBase{}) // "QMapDataBase *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMapDataBase", "createData", args)
  }

  return
}

// rotateLeft(struct QMapNodeBase *)
func (this *QMapDataBase) Rotateleft(args ...interface{}) () {
  // rotateLeft(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase10rotateLeftEP12QMapNodeBase
    // invoke: void rotateLeft(struct QMapNodeBase *)
    var arg0 = args[0].(QMapNodeBase).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QMapDataBase10rotateLeftEP12QMapNodeBase(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMapDataBase", "rotateLeft", args)
  }

  return
}

// rebalance(struct QMapNodeBase *)
func (this *QMapDataBase) Rebalance(args ...interface{}) () {
  // rebalance(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase9rebalanceEP12QMapNodeBase
    // invoke: void rebalance(struct QMapNodeBase *)
    var arg0 = args[0].(QMapNodeBase).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QMapDataBase9rebalanceEP12QMapNodeBase(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMapDataBase", "rebalance", args)
  }

  return
}

// freeNodeAndRebalance(struct QMapNodeBase *)
func (this *QMapDataBase) Freenodeandrebalance(args ...interface{}) () {
  // freeNodeAndRebalance(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapDataBase20freeNodeAndRebalanceEP12QMapNodeBase
    // invoke: void freeNodeAndRebalance(struct QMapNodeBase *)
    var arg0 = args[0].(QMapNodeBase).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QMapDataBase20freeNodeAndRebalanceEP12QMapNodeBase(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMapDataBase", "freeNodeAndRebalance", args)
  }

  return
}

// previousNode()
func (this *QMapNodeBase) Previousnode(args ...interface{}) (ret interface{}) {
  // previousNode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapNodeBase12previousNodeEv
    // invoke: QMapNodeBase * previousNode()
    var ret0 = C.C_ZN12QMapNodeBase12previousNodeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMapNodeBase", "previousNode", args)
  }

  return
}

// parent()
func (this *QMapNodeBase) Parent(args ...interface{}) (ret interface{}) {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QMapNodeBase6parentEv
    // invoke: QMapNodeBase * parent()
    var ret0 = C.C_ZNK12QMapNodeBase6parentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMapNodeBase", "parent", args)
  }

  return
}

// color()
func (this *QMapNodeBase) Color(args ...interface{}) () {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QMapNodeBase5colorEv
    // invoke: QMapNodeBase::Color color()
    C.C_ZNK12QMapNodeBase5colorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMapNodeBase", "color", args)
  }

  return
}

// setParent(struct QMapNodeBase *)
func (this *QMapNodeBase) Setparent(args ...interface{}) () {
  // setParent(struct QMapNodeBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapNodeBase9setParentEPS_
    // invoke: void setParent(struct QMapNodeBase *)
    var arg0 = args[0].(QMapNodeBase).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QMapNodeBase9setParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMapNodeBase", "setParent", args)
  }

  return
}

// nextNode()
func (this *QMapNodeBase) Nextnode(args ...interface{}) (ret interface{}) {
  // nextNode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QMapNodeBase8nextNodeEv
    // invoke: QMapNodeBase * nextNode()
    var ret0 = C.C_ZN12QMapNodeBase8nextNodeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMapNodeBase{}) // "QMapNodeBase *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMapNodeBase", "nextNode", args)
  }

  return
}

// <= body block end

