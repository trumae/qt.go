package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QWheelEvent struct {
	*QInputEvent
}

func (this *QWheelEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QWheelEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQWheelEventFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QWheelEvent{bcthis0}
}
func (*QWheelEvent) NewFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	return NewQWheelEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:196
// index:0
// Public virtual
// void ~QWheelEvent()
func DeleteQWheelEvent(*QWheelEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QWheelEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:199
// index:0
// Public inline
// QPoint pixelDelta()
func (this *QWheelEvent) PixelDelta() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10pixelDeltaEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:200
// index:0
// Public inline
// QPoint angleDelta()
func (this *QWheelEvent) AngleDelta() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10angleDeltaEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:202
// index:0
// Public inline
// int delta()
func (this *QWheelEvent) Delta() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent5deltaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:203
// index:0
// Public inline
// Qt::Orientation orientation()
func (this *QWheelEvent) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:206
// index:0
// Public inline
// QPoint pos()
func (this *QWheelEvent) Pos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:207
// index:0
// Public inline
// QPoint globalPos()
func (this *QWheelEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:208
// index:0
// Public inline
// int x()
func (this *QWheelEvent) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:209
// index:0
// Public inline
// int y()
func (this *QWheelEvent) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:210
// index:0
// Public inline
// int globalX()
func (this *QWheelEvent) GlobalX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:211
// index:0
// Public inline
// int globalY()
func (this *QWheelEvent) GlobalY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:213
// index:0
// Public inline
// const QPointF & posF()
func (this *QWheelEvent) PosF() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent4posFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:214
// index:0
// Public inline
// const QPointF & globalPosF()
func (this *QWheelEvent) GlobalPosF() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent10globalPosFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:216
// index:0
// Public inline
// Qt::MouseButtons buttons()
func (this *QWheelEvent) Buttons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:218
// index:0
// Public inline
// Qt::ScrollPhase phase()
func (this *QWheelEvent) Phase() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent5phaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:219
// index:0
// Public inline
// bool inverted()
func (this *QWheelEvent) Inverted() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent8invertedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:221
// index:0
// Public inline
// Qt::MouseEventSource source()
func (this *QWheelEvent) Source() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QWheelEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QWheelEvent__ = int

const QWheelEvent__DefaultDeltasPerStep QWheelEvent__ = 120

//  body block end
