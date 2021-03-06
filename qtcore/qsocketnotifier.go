package qtcore

// /usr/include/qt/QtCore/qsocketnotifier.h
// #include <qsocketnotifier.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QSocketNotifier) InheritEvent(f func(arg0 *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QSocketNotifier struct {
	*QObject
}
type QSocketNotifier_ITF interface {
	QObject_ITF
	QSocketNotifier_PTR() *QSocketNotifier
}

func (ptr *QSocketNotifier) QSocketNotifier_PTR() *QSocketNotifier { return ptr }

func (this *QSocketNotifier) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSocketNotifier) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQSocketNotifierFromPointer(cthis unsafe.Pointer) *QSocketNotifier {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSocketNotifier{bcthis0}
}
func (*QSocketNotifier) NewFromPointer(cthis unsafe.Pointer) *QSocketNotifier {
	return NewQSocketNotifierFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:50
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QSocketNotifier) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSocketNotifier10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsocketnotifier.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSocketNotifier(qintptr, enum QSocketNotifier::Type, QObject *)
func NewQSocketNotifier(socket int64, arg1 int, parent QObject_ITF /*777 QObject **/) *QSocketNotifier {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSocketNotifierC2ExNS_4TypeEP7QObject", qtrt.FFI_TYPE_POINTER, socket, arg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSocketNotifierFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSocketNotifier")
	return gothis
}

// /usr/include/qt/QtCore/qsocketnotifier.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSocketNotifier(qintptr, enum QSocketNotifier::Type, QObject *)
func NewQSocketNotifier__(socket int64, arg1 int) *QSocketNotifier {
	// arg: 2, QObject *=Pointer, QObject=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSocketNotifierC2ExNS_4TypeEP7QObject", qtrt.FFI_TYPE_POINTER, socket, arg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSocketNotifierFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSocketNotifier")
	return gothis
}

// /usr/include/qt/QtCore/qsocketnotifier.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSocketNotifier()
func DeleteQSocketNotifier(this *QSocketNotifier) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSocketNotifierD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] qintptr socket() const
func (this *QSocketNotifier) Socket() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSocketNotifier6socketEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qsocketnotifier.h:60
// index:0
// Public Visibility=Default Availability=Available
// [4] QSocketNotifier::Type type() const
func (this *QSocketNotifier) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSocketNotifier4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const
func (this *QSocketNotifier) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSocketNotifier9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsocketnotifier.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QSocketNotifier) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSocketNotifier10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsocketnotifier.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSocketNotifier) Event(arg0 QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSocketNotifier5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QSocketNotifier__Type = int

const QSocketNotifier__Read QSocketNotifier__Type = 0
const QSocketNotifier__Write QSocketNotifier__Type = 1
const QSocketNotifier__Exception QSocketNotifier__Type = 2

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
