package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QWidgetItemV2 struct {
	*QWidgetItem
}

func (this *QWidgetItemV2) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidgetItem.GetCthis()
	}
}
func NewQWidgetItemV2FromPointer(cthis unsafe.Pointer) *QWidgetItemV2 {
	bcthis0 := NewQWidgetItemFromPointer(cthis)
	return &QWidgetItemV2{bcthis0}
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:148
// index:0
// Public
// void QWidgetItemV2(class QWidget *)
func NewQWidgetItemV2(widget *QWidget /*444 QWidget **/) *QWidgetItemV2 {
	cthis := qtrt.Calloc(1, 256) // 88
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetItemV2C2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWidgetItemV2FromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:149
// index:0
// Public virtual
// void ~QWidgetItemV2()
func DeleteQWidgetItemV2(*QWidgetItemV2) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetItemV2D2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:151
// index:0
// Public virtual
// QSize sizeHint()
func (this *QWidgetItemV2) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetItemV28sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:152
// index:0
// Public virtual
// QSize minimumSize()
func (this *QWidgetItemV2) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetItemV211minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:153
// index:0
// Public virtual
// QSize maximumSize()
func (this *QWidgetItemV2) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetItemV211maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlayoutitem.h:154
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QWidgetItemV2) HeightForWidth(width int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetItemV214heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
