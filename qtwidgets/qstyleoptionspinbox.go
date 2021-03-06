package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QStyleOptionSpinBox struct {
	*QStyleOptionComplex
}
type QStyleOptionSpinBox_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSpinBox_PTR() *QStyleOptionSpinBox
}

func (ptr *QStyleOptionSpinBox) QStyleOptionSpinBox_PTR() *QStyleOptionSpinBox { return ptr }

func (this *QStyleOptionSpinBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSpinBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionSpinBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSpinBox{bcthis0}
}
func (*QStyleOptionSpinBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSpinBox {
	return NewQStyleOptionSpinBoxFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:552
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionSpinBox()
func NewQStyleOptionSpinBox() *QStyleOptionSpinBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionSpinBoxC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSpinBox)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:556
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionSpinBox(int)
func NewQStyleOptionSpinBox_1(version int) *QStyleOptionSpinBox {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionSpinBoxC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSpinBoxFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSpinBox)
	return gothis
}

func DeleteQStyleOptionSpinBox(this *QStyleOptionSpinBox) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionSpinBoxD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionSpinBox__StyleOptionType = int

const QStyleOptionSpinBox__Type QStyleOptionSpinBox__StyleOptionType = 983042

type QStyleOptionSpinBox__StyleOptionVersion = int

const QStyleOptionSpinBox__Version QStyleOptionSpinBox__StyleOptionVersion = 1

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
