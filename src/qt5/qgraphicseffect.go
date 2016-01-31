package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qgraphicseffect.h
// dst-file: /src/widgets/qgraphicseffect.go
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
  // proto:  void QGraphicsColorizeEffect::setStrength(qreal strength);
extern void C_ZN23QGraphicsColorizeEffect11setStrengthEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsColorizeEffect::strength();
extern double C_ZNK23QGraphicsColorizeEffect8strengthEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsColorizeEffect::metaObject();
extern void C_ZNK23QGraphicsColorizeEffect10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsColorizeEffect::setColor(const QColor & c);
extern void C_ZN23QGraphicsColorizeEffect8setColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  QColor QGraphicsColorizeEffect::color();
extern void* C_ZNK23QGraphicsColorizeEffect5colorEv(void* qthis); // 4
  // proto:  void QGraphicsColorizeEffect::QGraphicsColorizeEffect(QObject * parent);
extern void* C_ZN23QGraphicsColorizeEffectC2EP7QObject(void* arg0); // 3
  // proto:  void QGraphicsColorizeEffect::~QGraphicsColorizeEffect();
extern void C_ZN23QGraphicsColorizeEffectD2Ev(void* qthis); // 4
  // proto:  QRectF QGraphicsEffect::boundingRect();
extern void* C_ZNK15QGraphicsEffect12boundingRectEv(void* qthis); // 4
  // proto:  QRectF QGraphicsEffect::boundingRectFor(const QRectF & sourceRect);
extern void* C_ZNK15QGraphicsEffect15boundingRectForERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsEffect::~QGraphicsEffect();
extern void C_ZN15QGraphicsEffectD2Ev(void* qthis); // 4
  // proto:  bool QGraphicsEffect::isEnabled();
extern bool C_ZNK15QGraphicsEffect9isEnabledEv(void* qthis); // 4
  // proto:  QGraphicsEffectSource * QGraphicsEffect::source();
extern void C_ZNK15QGraphicsEffect6sourceEv(void* qthis); // 4
  // proto:  void QGraphicsEffect::setEnabled(bool enable);
extern void C_ZN15QGraphicsEffect10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QGraphicsEffect::update();
extern void C_ZN15QGraphicsEffect6updateEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsEffect::metaObject();
extern void C_ZNK15QGraphicsEffect10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsEffect::QGraphicsEffect(QObject * parent);
extern void* C_ZN15QGraphicsEffectC2EP7QObject(void* arg0); // 3
  // proto:  void QGraphicsDropShadowEffect::setYOffset(qreal dy);
extern void C_ZN25QGraphicsDropShadowEffect10setYOffsetEd(void* qthis, double arg0); // 2
  // proto:  QRectF QGraphicsDropShadowEffect::boundingRectFor(const QRectF & rect);
extern void* C_ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsDropShadowEffect::setXOffset(qreal dx);
extern void C_ZN25QGraphicsDropShadowEffect10setXOffsetEd(void* qthis, double arg0); // 2
  // proto:  qreal QGraphicsDropShadowEffect::blurRadius();
extern double C_ZNK25QGraphicsDropShadowEffect10blurRadiusEv(void* qthis); // 4
  // proto:  void QGraphicsDropShadowEffect::setBlurRadius(qreal blurRadius);
extern void C_ZN25QGraphicsDropShadowEffect13setBlurRadiusEd(void* qthis, double arg0); // 4
  // proto:  QColor QGraphicsDropShadowEffect::color();
extern void* C_ZNK25QGraphicsDropShadowEffect5colorEv(void* qthis); // 4
  // proto:  void QGraphicsDropShadowEffect::setOffset(qreal dx, qreal dy);
extern void C_ZN25QGraphicsDropShadowEffect9setOffsetEdd(void* qthis, double arg0, double arg1); // 2
  // proto:  void QGraphicsDropShadowEffect::setOffset(qreal d);
extern void C_ZN25QGraphicsDropShadowEffect9setOffsetEd(void* qthis, double arg0); // 2
  // proto:  void QGraphicsDropShadowEffect::setOffset(const QPointF & ofs);
extern void C_ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsDropShadowEffect::~QGraphicsDropShadowEffect();
extern void C_ZN25QGraphicsDropShadowEffectD2Ev(void* qthis); // 4
  // proto:  void QGraphicsDropShadowEffect::QGraphicsDropShadowEffect(QObject * parent);
extern void* C_ZN25QGraphicsDropShadowEffectC2EP7QObject(void* arg0); // 3
  // proto:  QPointF QGraphicsDropShadowEffect::offset();
extern void* C_ZNK25QGraphicsDropShadowEffect6offsetEv(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsDropShadowEffect::metaObject();
extern void C_ZNK25QGraphicsDropShadowEffect10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsDropShadowEffect::setColor(const QColor & color);
extern void C_ZN25QGraphicsDropShadowEffect8setColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  qreal QGraphicsDropShadowEffect::yOffset();
extern double C_ZNK25QGraphicsDropShadowEffect7yOffsetEv(void* qthis); // 2
  // proto:  qreal QGraphicsDropShadowEffect::xOffset();
extern double C_ZNK25QGraphicsDropShadowEffect7xOffsetEv(void* qthis); // 2
  // proto:  qreal QGraphicsOpacityEffect::opacity();
extern double C_ZNK22QGraphicsOpacityEffect7opacityEv(void* qthis); // 4
  // proto:  void QGraphicsOpacityEffect::~QGraphicsOpacityEffect();
extern void C_ZN22QGraphicsOpacityEffectD2Ev(void* qthis); // 4
  // proto:  void QGraphicsOpacityEffect::setOpacity(qreal opacity);
extern void C_ZN22QGraphicsOpacityEffect10setOpacityEd(void* qthis, double arg0); // 4
  // proto:  QBrush QGraphicsOpacityEffect::opacityMask();
extern void* C_ZNK22QGraphicsOpacityEffect11opacityMaskEv(void* qthis); // 4
  // proto:  void QGraphicsOpacityEffect::QGraphicsOpacityEffect(QObject * parent);
extern void* C_ZN22QGraphicsOpacityEffectC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QGraphicsOpacityEffect::metaObject();
extern void C_ZNK22QGraphicsOpacityEffect10metaObjectEv(void* qthis); // 4
  // proto:  void QGraphicsOpacityEffect::setOpacityMask(const QBrush & mask);
extern void C_ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsBlurEffect::QGraphicsBlurEffect(QObject * parent);
extern void* C_ZN19QGraphicsBlurEffectC2EP7QObject(void* arg0); // 3
  // proto:  BlurHints QGraphicsBlurEffect::blurHints();
extern void C_ZNK19QGraphicsBlurEffect9blurHintsEv(void* qthis); // 4
  // proto:  QRectF QGraphicsBlurEffect::boundingRectFor(const QRectF & rect);
extern void* C_ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QGraphicsBlurEffect::setBlurRadius(qreal blurRadius);
extern void C_ZN19QGraphicsBlurEffect13setBlurRadiusEd(void* qthis, double arg0); // 4
  // proto:  qreal QGraphicsBlurEffect::blurRadius();
extern double C_ZNK19QGraphicsBlurEffect10blurRadiusEv(void* qthis); // 4
  // proto:  void QGraphicsBlurEffect::~QGraphicsBlurEffect();
extern void C_ZN19QGraphicsBlurEffectD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QGraphicsBlurEffect::metaObject();
extern void C_ZNK19QGraphicsBlurEffect10metaObjectEv(void* qthis); // 4
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

// class sizeof(QGraphicsColorizeEffect)=1
type QGraphicsColorizeEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _strengthChanged QGraphicsColorizeEffect_strengthChanged_signal;
//  _colorChanged QGraphicsColorizeEffect_colorChanged_signal;
}

// class sizeof(QGraphicsEffect)=1
type QGraphicsEffect struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _enabledChanged QGraphicsEffect_enabledChanged_signal;
}

// class sizeof(QGraphicsDropShadowEffect)=1
type QGraphicsDropShadowEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _colorChanged QGraphicsDropShadowEffect_colorChanged_signal;
//  _offsetChanged QGraphicsDropShadowEffect_offsetChanged_signal;
//  _blurRadiusChanged QGraphicsDropShadowEffect_blurRadiusChanged_signal;
}

// class sizeof(QGraphicsOpacityEffect)=1
type QGraphicsOpacityEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _opacityMaskChanged QGraphicsOpacityEffect_opacityMaskChanged_signal;
//  _opacityChanged QGraphicsOpacityEffect_opacityChanged_signal;
}

// class sizeof(QGraphicsBlurEffect)=1
type QGraphicsBlurEffect struct {
  /*qbase*/ QGraphicsEffect;
  qclsinst unsafe.Pointer /* *C.void */;
//  _blurHintsChanged QGraphicsBlurEffect_blurHintsChanged_signal;
//  _blurRadiusChanged QGraphicsBlurEffect_blurRadiusChanged_signal;
}

// setStrength(qreal)
func (this *QGraphicsColorizeEffect) Setstrength(args ...interface{}) () {
  // setStrength(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffect11setStrengthEd
    // invoke: void setStrength(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsColorizeEffect11setStrengthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "setStrength", args)
  }

  return
}

// strength()
func (this *QGraphicsColorizeEffect) Strength(args ...interface{}) (ret interface{}) {
  // strength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect8strengthEv
    // invoke: qreal strength()
    var ret0 = C.C_ZNK23QGraphicsColorizeEffect8strengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "strength", args)
  }

  return
}

// metaObject()
func (this *QGraphicsColorizeEffect) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK23QGraphicsColorizeEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "metaObject", args)
  }

  return
}

// setColor(const class QColor &)
func (this *QGraphicsColorizeEffect) Setcolor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffect8setColorERK6QColor
    // invoke: void setColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN23QGraphicsColorizeEffect8setColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "setColor", args)
  }

  return
}

// color()
func (this *QGraphicsColorizeEffect) Color(args ...interface{}) (ret interface{}) {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK23QGraphicsColorizeEffect5colorEv
    // invoke: QColor color()
    var ret0 = C.C_ZNK23QGraphicsColorizeEffect5colorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "color", args)
  }

  return
}

// QGraphicsColorizeEffect(class QObject *)
func NewQGraphicsColorizeEffect(args ...interface{}) *QGraphicsColorizeEffect {
  // QGraphicsColorizeEffect(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffectC1EP7QObject
    // invoke: void QGraphicsColorizeEffect(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QGraphicsColorizeEffectC2EP7QObject(arg0)
    return &QGraphicsColorizeEffect{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "QGraphicsColorizeEffect", args)
  }

  return nil // QGraphicsColorizeEffect{qclsinst:qthis}
}

// ~QGraphicsColorizeEffect()
func (this *QGraphicsColorizeEffect) Freeqgraphicscolorizeeffect(args ...interface{}) () {
  // ~QGraphicsColorizeEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QGraphicsColorizeEffectD0Ev
    // invoke: void ~QGraphicsColorizeEffect()
    C.C_ZN23QGraphicsColorizeEffectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsColorizeEffect", "~QGraphicsColorizeEffect", args)
  }

  return
}

// boundingRect()
func (this *QGraphicsEffect) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect12boundingRectEv
    // invoke: QRectF boundingRect()
    var ret0 = C.C_ZNK15QGraphicsEffect12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "boundingRect", args)
  }

  return
}

// boundingRectFor(const class QRectF &)
func (this *QGraphicsEffect) Boundingrectfor(args ...interface{}) (ret interface{}) {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect15boundingRectForERK6QRectF
    // invoke: QRectF boundingRectFor(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QGraphicsEffect15boundingRectForERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "boundingRectFor", args)
  }

  return
}

// ~QGraphicsEffect()
func (this *QGraphicsEffect) Freeqgraphicseffect(args ...interface{}) () {
  // ~QGraphicsEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffectD0Ev
    // invoke: void ~QGraphicsEffect()
    C.C_ZN15QGraphicsEffectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "~QGraphicsEffect", args)
  }

  return
}

// isEnabled()
func (this *QGraphicsEffect) Isenabled(args ...interface{}) (ret interface{}) {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect9isEnabledEv
    // invoke: bool isEnabled()
    var ret0 = C.C_ZNK15QGraphicsEffect9isEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "isEnabled", args)
  }

  return
}

// source()
func (this *QGraphicsEffect) Source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect6sourceEv
    // invoke: QGraphicsEffectSource * source()
    C.C_ZNK15QGraphicsEffect6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "source", args)
  }

  return
}

// setEnabled(_Bool)
func (this *QGraphicsEffect) Setenabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffect10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QGraphicsEffect10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "setEnabled", args)
  }

  return
}

// update()
func (this *QGraphicsEffect) Update(args ...interface{}) () {
  // update()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffect6updateEv
    // invoke: void update()
    C.C_ZN15QGraphicsEffect6updateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "update", args)
  }

  return
}

// metaObject()
func (this *QGraphicsEffect) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGraphicsEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QGraphicsEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "metaObject", args)
  }

  return
}

// QGraphicsEffect(class QObject *)
func NewQGraphicsEffect(args ...interface{}) *QGraphicsEffect {
  // QGraphicsEffect(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGraphicsEffectC1EP7QObject
    // invoke: void QGraphicsEffect(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QGraphicsEffectC2EP7QObject(arg0)
    return &QGraphicsEffect{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsEffect", "QGraphicsEffect", args)
  }

  return nil // QGraphicsEffect{qclsinst:qthis}
}

// setYOffset(qreal)
func (this *QGraphicsDropShadowEffect) Setyoffset(args ...interface{}) () {
  // setYOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect10setYOffsetEd
    // invoke: void setYOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsDropShadowEffect10setYOffsetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setYOffset", args)
  }

  return
}

// boundingRectFor(const class QRectF &)
func (this *QGraphicsDropShadowEffect) Boundingrectfor(args ...interface{}) (ret interface{}) {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF
    // invoke: QRectF boundingRectFor(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK25QGraphicsDropShadowEffect15boundingRectForERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "boundingRectFor", args)
  }

  return
}

// setXOffset(qreal)
func (this *QGraphicsDropShadowEffect) Setxoffset(args ...interface{}) () {
  // setXOffset(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect10setXOffsetEd
    // invoke: void setXOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsDropShadowEffect10setXOffsetEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setXOffset", args)
  }

  return
}

// blurRadius()
func (this *QGraphicsDropShadowEffect) Blurradius(args ...interface{}) (ret interface{}) {
  // blurRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect10blurRadiusEv
    // invoke: qreal blurRadius()
    var ret0 = C.C_ZNK25QGraphicsDropShadowEffect10blurRadiusEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "blurRadius", args)
  }

  return
}

// setBlurRadius(qreal)
func (this *QGraphicsDropShadowEffect) Setblurradius(args ...interface{}) () {
  // setBlurRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect13setBlurRadiusEd
    // invoke: void setBlurRadius(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsDropShadowEffect13setBlurRadiusEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setBlurRadius", args)
  }

  return
}

// color()
func (this *QGraphicsDropShadowEffect) Color(args ...interface{}) (ret interface{}) {
  // color()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect5colorEv
    // invoke: QColor color()
    var ret0 = C.C_ZNK25QGraphicsDropShadowEffect5colorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "color", args)
  }

  return
}

// setOffset(qreal, qreal)
func (this *QGraphicsDropShadowEffect) Setoffset(args ...interface{}) () {
  // setOffset(qreal, qreal)
  // setOffset(qreal)
  // setOffset(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetEdd
    // invoke: void setOffset(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    C.C_ZN25QGraphicsDropShadowEffect9setOffsetEdd(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetEd
    // invoke: void setOffset(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsDropShadowEffect9setOffsetEd(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF
    // invoke: void setOffset(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsDropShadowEffect9setOffsetERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setOffset", args)
  }

  return
}

// ~QGraphicsDropShadowEffect()
func (this *QGraphicsDropShadowEffect) Freeqgraphicsdropshadoweffect(args ...interface{}) () {
  // ~QGraphicsDropShadowEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffectD0Ev
    // invoke: void ~QGraphicsDropShadowEffect()
    C.C_ZN25QGraphicsDropShadowEffectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "~QGraphicsDropShadowEffect", args)
  }

  return
}

// QGraphicsDropShadowEffect(class QObject *)
func NewQGraphicsDropShadowEffect(args ...interface{}) *QGraphicsDropShadowEffect {
  // QGraphicsDropShadowEffect(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffectC1EP7QObject
    // invoke: void QGraphicsDropShadowEffect(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN25QGraphicsDropShadowEffectC2EP7QObject(arg0)
    return &QGraphicsDropShadowEffect{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "QGraphicsDropShadowEffect", args)
  }

  return nil // QGraphicsDropShadowEffect{qclsinst:qthis}
}

// offset()
func (this *QGraphicsDropShadowEffect) Offset(args ...interface{}) (ret interface{}) {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect6offsetEv
    // invoke: QPointF offset()
    var ret0 = C.C_ZNK25QGraphicsDropShadowEffect6offsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPointF{}) // "QPointF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "offset", args)
  }

  return
}

// metaObject()
func (this *QGraphicsDropShadowEffect) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK25QGraphicsDropShadowEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "metaObject", args)
  }

  return
}

// setColor(const class QColor &)
func (this *QGraphicsDropShadowEffect) Setcolor(args ...interface{}) () {
  // setColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QGraphicsDropShadowEffect8setColorERK6QColor
    // invoke: void setColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN25QGraphicsDropShadowEffect8setColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "setColor", args)
  }

  return
}

// yOffset()
func (this *QGraphicsDropShadowEffect) Yoffset(args ...interface{}) (ret interface{}) {
  // yOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect7yOffsetEv
    // invoke: qreal yOffset()
    var ret0 = C.C_ZNK25QGraphicsDropShadowEffect7yOffsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "yOffset", args)
  }

  return
}

// xOffset()
func (this *QGraphicsDropShadowEffect) Xoffset(args ...interface{}) (ret interface{}) {
  // xOffset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QGraphicsDropShadowEffect7xOffsetEv
    // invoke: qreal xOffset()
    var ret0 = C.C_ZNK25QGraphicsDropShadowEffect7xOffsetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsDropShadowEffect", "xOffset", args)
  }

  return
}

// opacity()
func (this *QGraphicsOpacityEffect) Opacity(args ...interface{}) (ret interface{}) {
  // opacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect7opacityEv
    // invoke: qreal opacity()
    var ret0 = C.C_ZNK22QGraphicsOpacityEffect7opacityEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "opacity", args)
  }

  return
}

// ~QGraphicsOpacityEffect()
func (this *QGraphicsOpacityEffect) Freeqgraphicsopacityeffect(args ...interface{}) () {
  // ~QGraphicsOpacityEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffectD0Ev
    // invoke: void ~QGraphicsOpacityEffect()
    C.C_ZN22QGraphicsOpacityEffectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "~QGraphicsOpacityEffect", args)
  }

  return
}

// setOpacity(qreal)
func (this *QGraphicsOpacityEffect) Setopacity(args ...interface{}) () {
  // setOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffect10setOpacityEd
    // invoke: void setOpacity(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN22QGraphicsOpacityEffect10setOpacityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "setOpacity", args)
  }

  return
}

// opacityMask()
func (this *QGraphicsOpacityEffect) Opacitymask(args ...interface{}) (ret interface{}) {
  // opacityMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect11opacityMaskEv
    // invoke: QBrush opacityMask()
    var ret0 = C.C_ZNK22QGraphicsOpacityEffect11opacityMaskEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "opacityMask", args)
  }

  return
}

// QGraphicsOpacityEffect(class QObject *)
func NewQGraphicsOpacityEffect(args ...interface{}) *QGraphicsOpacityEffect {
  // QGraphicsOpacityEffect(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffectC1EP7QObject
    // invoke: void QGraphicsOpacityEffect(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QGraphicsOpacityEffectC2EP7QObject(arg0)
    return &QGraphicsOpacityEffect{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "QGraphicsOpacityEffect", args)
  }

  return nil // QGraphicsOpacityEffect{qclsinst:qthis}
}

// metaObject()
func (this *QGraphicsOpacityEffect) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QGraphicsOpacityEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK22QGraphicsOpacityEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "metaObject", args)
  }

  return
}

// setOpacityMask(const class QBrush &)
func (this *QGraphicsOpacityEffect) Setopacitymask(args ...interface{}) () {
  // setOpacityMask(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush
    // invoke: void setOpacityMask(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsOpacityEffect", "setOpacityMask", args)
  }

  return
}

// QGraphicsBlurEffect(class QObject *)
func NewQGraphicsBlurEffect(args ...interface{}) *QGraphicsBlurEffect {
  // QGraphicsBlurEffect(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsBlurEffectC1EP7QObject
    // invoke: void QGraphicsBlurEffect(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QGraphicsBlurEffectC2EP7QObject(arg0)
    return &QGraphicsBlurEffect{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "QGraphicsBlurEffect", args)
  }

  return nil // QGraphicsBlurEffect{qclsinst:qthis}
}

// blurHints()
func (this *QGraphicsBlurEffect) Blurhints(args ...interface{}) () {
  // blurHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect9blurHintsEv
    // invoke: BlurHints blurHints()
    C.C_ZNK19QGraphicsBlurEffect9blurHintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "blurHints", args)
  }

  return
}

// boundingRectFor(const class QRectF &)
func (this *QGraphicsBlurEffect) Boundingrectfor(args ...interface{}) (ret interface{}) {
  // boundingRectFor(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF
    // invoke: QRectF boundingRectFor(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "boundingRectFor", args)
  }

  return
}

// setBlurRadius(qreal)
func (this *QGraphicsBlurEffect) Setblurradius(args ...interface{}) () {
  // setBlurRadius(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsBlurEffect13setBlurRadiusEd
    // invoke: void setBlurRadius(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN19QGraphicsBlurEffect13setBlurRadiusEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "setBlurRadius", args)
  }

  return
}

// blurRadius()
func (this *QGraphicsBlurEffect) Blurradius(args ...interface{}) (ret interface{}) {
  // blurRadius()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect10blurRadiusEv
    // invoke: qreal blurRadius()
    var ret0 = C.C_ZNK19QGraphicsBlurEffect10blurRadiusEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "blurRadius", args)
  }

  return
}

// ~QGraphicsBlurEffect()
func (this *QGraphicsBlurEffect) Freeqgraphicsblureffect(args ...interface{}) () {
  // ~QGraphicsBlurEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QGraphicsBlurEffectD0Ev
    // invoke: void ~QGraphicsBlurEffect()
    C.C_ZN19QGraphicsBlurEffectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "~QGraphicsBlurEffect", args)
  }

  return
}

// metaObject()
func (this *QGraphicsBlurEffect) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QGraphicsBlurEffect10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QGraphicsBlurEffect10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGraphicsBlurEffect", "metaObject", args)
  }

  return
}

// <= body block end

