package qtcore

// /usr/include/qt/QtCore/qhashfunctions.h
// #include <qhashfunctions.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 93
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QHashCombine struct {
	*qtrt.CObject
}
type QHashCombine_ITF interface {
	QHashCombine_PTR() *QHashCombine
}

func (ptr *QHashCombine) QHashCombine_PTR() *QHashCombine { return ptr }

func (this *QHashCombine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHashCombine) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQHashCombineFromPointer(cthis unsafe.Pointer) *QHashCombine {
	return &QHashCombine{&qtrt.CObject{cthis}}
}
func (*QHashCombine) NewFromPointer(cthis unsafe.Pointer) *QHashCombine {
	return NewQHashCombineFromPointer(cthis)
}

func DeleteQHashCombine(this *QHashCombine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHashCombineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

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
