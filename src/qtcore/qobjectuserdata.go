package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QObjectUserData struct {
	*qtrt.CObject
}

func (this *QObjectUserData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQObjectUserDataFromPointer(cthis unsafe.Pointer) *QObjectUserData {
	return &QObjectUserData{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qobject.h:478
// index:0
// Public virtual
// void ~QObjectUserData()
func DeleteQObjectUserData(*QObjectUserData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QObjectUserDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
