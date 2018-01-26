package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleOptionTitleBar struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionTitleBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionTitleBar) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionTitleBarFromPointer(cthis unsafe.Pointer) *QStyleOptionTitleBar {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionTitleBar{bcthis0}
}
func (*QStyleOptionTitleBar) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTitleBar {
	return NewQStyleOptionTitleBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:619
// index:0
// Public
// void QStyleOptionTitleBar()
func NewQStyleOptionTitleBar() *QStyleOptionTitleBar {
	cthis := qtrt.Calloc(1, 256) // 96
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionTitleBarC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTitleBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:623
// index:1
// Protected
// void QStyleOptionTitleBar(int)
func NewQStyleOptionTitleBar_1(version int) *QStyleOptionTitleBar {
	cthis := qtrt.Calloc(1, 256) // 96
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionTitleBarC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionTitleBarFromPointer(cthis)
	return gothis
}

type QStyleOptionTitleBar__StyleOptionType = int

const QStyleOptionTitleBar__Type QStyleOptionTitleBar__StyleOptionType = 983045

type QStyleOptionTitleBar__StyleOptionVersion = int

const QStyleOptionTitleBar__Version QStyleOptionTitleBar__StyleOptionVersion = 1

//  body block end
