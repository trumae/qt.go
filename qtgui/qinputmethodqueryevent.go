package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QInputMethodQueryEvent struct {
	*qtcore.QEvent
}

func (this *QInputMethodQueryEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QInputMethodQueryEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQInputMethodQueryEventFromPointer(cthis unsafe.Pointer) *QInputMethodQueryEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputMethodQueryEvent{bcthis0}
}
func (*QInputMethodQueryEvent) NewFromPointer(cthis unsafe.Pointer) *QInputMethodQueryEvent {
	return NewQInputMethodQueryEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:582
// index:0
// Public virtual
// void ~QInputMethodQueryEvent()
func DeleteQInputMethodQueryEvent(*QInputMethodQueryEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QInputMethodQueryEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:584
// index:0
// Public inline
// Qt::InputMethodQueries queries()
func (this *QInputMethodQueryEvent) Queries() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QInputMethodQueryEvent7queriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:586
// index:0
// Public
// void setValue(Qt::InputMethodQuery, const class QVariant &)
func (this *QInputMethodQueryEvent) SetValue(query int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QInputMethodQueryEvent8setValueEN2Qt16InputMethodQueryERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), query, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:587
// index:0
// Public
// QVariant value(Qt::InputMethodQuery)
func (this *QInputMethodQueryEvent) Value(query int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK22QInputMethodQueryEvent5valueEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), query)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
