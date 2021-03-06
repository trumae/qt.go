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
// extern C begin: 1
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

type QStyleOptionSlider struct {
	*QStyleOptionComplex
}
type QStyleOptionSlider_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSlider_PTR() *QStyleOptionSlider
}

func (ptr *QStyleOptionSlider) QStyleOptionSlider_PTR() *QStyleOptionSlider { return ptr }

func (this *QStyleOptionSlider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSlider) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSliderFromPointer(cthis unsafe.Pointer) *QStyleOptionSlider {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSlider{bcthis0}
}
func (*QStyleOptionSlider) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSlider {
	return NewQStyleOptionSliderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:533
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionSlider()
func NewQStyleOptionSlider() *QStyleOptionSlider {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSlider)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:537
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionSlider(int)
func NewQStyleOptionSlider_1(version int) *QStyleOptionSlider {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSlider)
	return gothis
}

func DeleteQStyleOptionSlider(this *QStyleOptionSlider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionSliderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionSlider__StyleOptionType = int

const QStyleOptionSlider__Type QStyleOptionSlider__StyleOptionType = 983041

type QStyleOptionSlider__StyleOptionVersion = int

const QStyleOptionSlider__Version QStyleOptionSlider__StyleOptionVersion = 1

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
