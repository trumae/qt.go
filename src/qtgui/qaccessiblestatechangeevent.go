package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QAccessibleStateChangeEvent struct {
	*QAccessibleEvent
}

func (this *QAccessibleStateChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleEvent.GetCthis()
	}
}
func NewQAccessibleStateChangeEventFromPointer(cthis unsafe.Pointer) *QAccessibleStateChangeEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleStateChangeEvent{bcthis0}
}

// /usr/include/qt/QtGui/qaccessible.h:733
// index:0
// Public virtual
// void ~QAccessibleStateChangeEvent()
func DeleteQAccessibleStateChangeEvent(*QAccessibleStateChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QAccessibleStateChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:735
// index:0
// Public inline
// QAccessible::State changedStates()
func (this *QAccessibleStateChangeEvent) ChangedStates() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QAccessibleStateChangeEvent13changedStatesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
