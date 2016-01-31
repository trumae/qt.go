package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qsystemtrayicon.h
// dst-file: /src/widgets/qsystemtrayicon.go
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
  // proto:  void QSystemTrayIcon::show();
extern void C_ZN15QSystemTrayIcon4showEv(void* qthis); // 2
  // proto:  void QSystemTrayIcon::QSystemTrayIcon(const QIcon & icon, QObject * parent);
extern void* C_ZN15QSystemTrayIconC2ERK5QIconP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QSystemTrayIcon::QSystemTrayIcon(QObject * parent);
extern void* C_ZN15QSystemTrayIconC2EP7QObject(void* arg0); // 3
  // proto:  void QSystemTrayIcon::hide();
extern void C_ZN15QSystemTrayIcon4hideEv(void* qthis); // 2
  // proto: static bool QSystemTrayIcon::isSystemTrayAvailable();
extern bool C_ZN15QSystemTrayIcon21isSystemTrayAvailableEv(); // 4
  // proto:  QMenu * QSystemTrayIcon::contextMenu();
extern void* C_ZNK15QSystemTrayIcon11contextMenuEv(void* qthis); // 4
  // proto:  void QSystemTrayIcon::setIcon(const QIcon & icon);
extern void C_ZN15QSystemTrayIcon7setIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  QString QSystemTrayIcon::toolTip();
extern void* C_ZNK15QSystemTrayIcon7toolTipEv(void* qthis); // 4
  // proto:  QIcon QSystemTrayIcon::icon();
extern void* C_ZNK15QSystemTrayIcon4iconEv(void* qthis); // 4
  // proto:  void QSystemTrayIcon::setContextMenu(QMenu * menu);
extern void C_ZN15QSystemTrayIcon14setContextMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  void QSystemTrayIcon::setToolTip(const QString & tip);
extern void C_ZN15QSystemTrayIcon10setToolTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QSystemTrayIcon::metaObject();
extern void C_ZNK15QSystemTrayIcon10metaObjectEv(void* qthis); // 4
  // proto:  QRect QSystemTrayIcon::geometry();
extern void* C_ZNK15QSystemTrayIcon8geometryEv(void* qthis); // 4
  // proto:  void QSystemTrayIcon::~QSystemTrayIcon();
extern void C_ZN15QSystemTrayIconD2Ev(void* qthis); // 4
  // proto:  bool QSystemTrayIcon::isVisible();
extern bool C_ZNK15QSystemTrayIcon9isVisibleEv(void* qthis); // 4
  // proto: static bool QSystemTrayIcon::supportsMessages();
extern bool C_ZN15QSystemTrayIcon16supportsMessagesEv(); // 4
  // proto:  void QSystemTrayIcon::setVisible(bool visible);
extern void C_ZN15QSystemTrayIcon10setVisibleEb(void* qthis, bool arg0); // 4
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

// class sizeof(QSystemTrayIcon)=1
type QSystemTrayIcon struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _activated QSystemTrayIcon_activated_signal;
//  _messageClicked QSystemTrayIcon_messageClicked_signal;
}

// show()
func (this *QSystemTrayIcon) Show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon4showEv
    // invoke: void show()
    C.C_ZN15QSystemTrayIcon4showEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "show", args)
  }

  return
}

// QSystemTrayIcon(const class QIcon &, class QObject *)
func NewQSystemTrayIcon(args ...interface{}) *QSystemTrayIcon {
  // QSystemTrayIcon(const class QIcon &, class QObject *)
  // QSystemTrayIcon(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIconC1ERK5QIconP7QObject
    // invoke: void QSystemTrayIcon(const class QIcon &, class QObject *)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QSystemTrayIconC2ERK5QIconP7QObject(arg0, arg1)
    return &QSystemTrayIcon{qclsinst:qthis}
  case 1:
    // invoke: _ZN15QSystemTrayIconC1EP7QObject
    // invoke: void QSystemTrayIcon(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QSystemTrayIconC2EP7QObject(arg0)
    return &QSystemTrayIcon{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "QSystemTrayIcon", args)
  }

  return nil // QSystemTrayIcon{qclsinst:qthis}
}

// hide()
func (this *QSystemTrayIcon) Hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon4hideEv
    // invoke: void hide()
    C.C_ZN15QSystemTrayIcon4hideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "hide", args)
  }

  return
}

// isSystemTrayAvailable()
func (this *QSystemTrayIcon) Issystemtrayavailable_S(args ...interface{}) (ret interface{}) {
  // isSystemTrayAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon21isSystemTrayAvailableEv
    // invoke: bool isSystemTrayAvailable()
    var ret0 = C.C_ZN15QSystemTrayIcon21isSystemTrayAvailableEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "isSystemTrayAvailable", args)
  }

  return
}

// contextMenu()
func (this *QSystemTrayIcon) Contextmenu(args ...interface{}) (ret interface{}) {
  // contextMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon11contextMenuEv
    // invoke: QMenu * contextMenu()
    var ret0 = C.C_ZNK15QSystemTrayIcon11contextMenuEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "contextMenu", args)
  }

  return
}

// setIcon(const class QIcon &)
func (this *QSystemTrayIcon) Seticon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QSystemTrayIcon7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setIcon", args)
  }

  return
}

// toolTip()
func (this *QSystemTrayIcon) Tooltip(args ...interface{}) (ret interface{}) {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon7toolTipEv
    // invoke: QString toolTip()
    var ret0 = C.C_ZNK15QSystemTrayIcon7toolTipEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "toolTip", args)
  }

  return
}

// icon()
func (this *QSystemTrayIcon) Icon(args ...interface{}) (ret interface{}) {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon4iconEv
    // invoke: QIcon icon()
    var ret0 = C.C_ZNK15QSystemTrayIcon4iconEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "icon", args)
  }

  return
}

// setContextMenu(class QMenu *)
func (this *QSystemTrayIcon) Setcontextmenu(args ...interface{}) () {
  // setContextMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon14setContextMenuEP5QMenu
    // invoke: void setContextMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QSystemTrayIcon14setContextMenuEP5QMenu(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setContextMenu", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QSystemTrayIcon) Settooltip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QSystemTrayIcon10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setToolTip", args)
  }

  return
}

// metaObject()
func (this *QSystemTrayIcon) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QSystemTrayIcon10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "metaObject", args)
  }

  return
}

// geometry()
func (this *QSystemTrayIcon) Geometry(args ...interface{}) (ret interface{}) {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon8geometryEv
    // invoke: QRect geometry()
    var ret0 = C.C_ZNK15QSystemTrayIcon8geometryEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "geometry", args)
  }

  return
}

// ~QSystemTrayIcon()
func (this *QSystemTrayIcon) Freeqsystemtrayicon(args ...interface{}) () {
  // ~QSystemTrayIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIconD0Ev
    // invoke: void ~QSystemTrayIcon()
    C.C_ZN15QSystemTrayIconD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "~QSystemTrayIcon", args)
  }

  return
}

// isVisible()
func (this *QSystemTrayIcon) Isvisible(args ...interface{}) (ret interface{}) {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon9isVisibleEv
    // invoke: bool isVisible()
    var ret0 = C.C_ZNK15QSystemTrayIcon9isVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "isVisible", args)
  }

  return
}

// supportsMessages()
func (this *QSystemTrayIcon) Supportsmessages_S(args ...interface{}) (ret interface{}) {
  // supportsMessages()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon16supportsMessagesEv
    // invoke: bool supportsMessages()
    var ret0 = C.C_ZN15QSystemTrayIcon16supportsMessagesEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "supportsMessages", args)
  }

  return
}

// setVisible(_Bool)
func (this *QSystemTrayIcon) Setvisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QSystemTrayIcon10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setVisible", args)
  }

  return
}

// <= body block end

