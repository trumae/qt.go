package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QStyleOption struct {
	*qtrt.CObject
}

func (this *QStyleOption) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStyleOption) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQStyleOptionFromPointer(cthis unsafe.Pointer) *QStyleOption {
	return &QStyleOption{&qtrt.CObject{cthis}}
}
func (*QStyleOption) NewFromPointer(cthis unsafe.Pointer) *QStyleOption {
	return NewQStyleOptionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:102
// index:0
// Public
// void QStyleOption(int, int)
func NewQStyleOption(version int, type_ int) *QStyleOption {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOptionC2Eii", ffiqt.FFI_TYPE_VOID, cthis, version, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:104
// index:0
// Public
// void ~QStyleOption()
func DeleteQStyleOption(*QStyleOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOptionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:106
// index:0
// Public
// void init(const class QWidget *)
func (this *QStyleOption) Init(w *QWidget /*777 const QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOption4initEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:107
// index:0
// Public inline
// void initFrom(const class QWidget *)
func (this *QStyleOption) InitFrom(w *QWidget /*777 const QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOption8initFromEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QStyleOption__OptionType = int

const QStyleOption__SO_Default QStyleOption__OptionType = 0
const QStyleOption__SO_FocusRect QStyleOption__OptionType = 1
const QStyleOption__SO_Button QStyleOption__OptionType = 2
const QStyleOption__SO_Tab QStyleOption__OptionType = 3
const QStyleOption__SO_MenuItem QStyleOption__OptionType = 4
const QStyleOption__SO_Frame QStyleOption__OptionType = 5
const QStyleOption__SO_ProgressBar QStyleOption__OptionType = 6
const QStyleOption__SO_ToolBox QStyleOption__OptionType = 7
const QStyleOption__SO_Header QStyleOption__OptionType = 8
const QStyleOption__SO_DockWidget QStyleOption__OptionType = 9
const QStyleOption__SO_ViewItem QStyleOption__OptionType = 10
const QStyleOption__SO_TabWidgetFrame QStyleOption__OptionType = 11
const QStyleOption__SO_TabBarBase QStyleOption__OptionType = 12
const QStyleOption__SO_RubberBand QStyleOption__OptionType = 13
const QStyleOption__SO_ToolBar QStyleOption__OptionType = 14
const QStyleOption__SO_GraphicsItem QStyleOption__OptionType = 15
const QStyleOption__SO_Complex QStyleOption__OptionType = 983040
const QStyleOption__SO_Slider QStyleOption__OptionType = 983041
const QStyleOption__SO_SpinBox QStyleOption__OptionType = 983042
const QStyleOption__SO_ToolButton QStyleOption__OptionType = 983043
const QStyleOption__SO_ComboBox QStyleOption__OptionType = 983044
const QStyleOption__SO_TitleBar QStyleOption__OptionType = 983045
const QStyleOption__SO_GroupBox QStyleOption__OptionType = 983046
const QStyleOption__SO_SizeGrip QStyleOption__OptionType = 983047
const QStyleOption__SO_CustomBase QStyleOption__OptionType = 3840
const QStyleOption__SO_ComplexCustomBase QStyleOption__OptionType = 251658240

type QStyleOption__StyleOptionType = int

const QStyleOption__Type QStyleOption__StyleOptionType = 0

type QStyleOption__StyleOptionVersion = int

const QStyleOption__Version QStyleOption__StyleOptionVersion = 1

//  body block end
