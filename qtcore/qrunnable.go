package qtcore

// /usr/include/qt/QtCore/qrunnable.h
// #include <qrunnable.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
type QRunnable struct {
	*qtrt.CObject
}

func (this *QRunnable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRunnable) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQRunnableFromPointer(cthis unsafe.Pointer) *QRunnable {
	return &QRunnable{&qtrt.CObject{cthis}}
}
func (*QRunnable) NewFromPointer(cthis unsafe.Pointer) *QRunnable {
	return NewQRunnableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrunnable.h:58
// index:0
// Public pure virtual
// void run()
func (this *QRunnable) Run() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QRunnable3runEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrunnable.h:60
// index:0
// Public inline
// void QRunnable()
func NewQRunnable() *QRunnable {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QRunnableC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRunnableFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrunnable.h:61
// index:0
// Public virtual
// void ~QRunnable()
func DeleteQRunnable(*QRunnable) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QRunnableD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrunnable.h:63
// index:0
// Public inline
// bool autoDelete()
func (this *QRunnable) AutoDelete() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QRunnable10autoDeleteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qrunnable.h:64
// index:0
// Public inline
// void setAutoDelete(_Bool)
func (this *QRunnable) SetAutoDelete(_autoDelete bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QRunnable13setAutoDeleteEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), _autoDelete)
	gopp.ErrPrint(err, rv)
}

//  body block end
