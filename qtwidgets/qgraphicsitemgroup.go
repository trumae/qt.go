package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

type QGraphicsItemGroup struct {
	*QGraphicsItem
}
type QGraphicsItemGroup_ITF interface {
	QGraphicsItem_ITF
	QGraphicsItemGroup_PTR() *QGraphicsItemGroup
}

func (ptr *QGraphicsItemGroup) QGraphicsItemGroup_PTR() *QGraphicsItemGroup { return ptr }

func (this *QGraphicsItemGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsItem.GetCthis()
	}
}
func (this *QGraphicsItemGroup) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsItem = NewQGraphicsItemFromPointer(cthis)
}
func NewQGraphicsItemGroupFromPointer(cthis unsafe.Pointer) *QGraphicsItemGroup {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QGraphicsItemGroup{bcthis0}
}
func (*QGraphicsItemGroup) NewFromPointer(cthis unsafe.Pointer) *QGraphicsItemGroup {
	return NewQGraphicsItemGroupFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1004
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsItemGroup(QGraphicsItem *)
func NewQGraphicsItemGroup(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsItemGroup {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsItemGroupC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsItemGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsItemGroup)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1004
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsItemGroup(QGraphicsItem *)
func NewQGraphicsItemGroup__() *QGraphicsItemGroup {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsItemGroupC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsItemGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsItemGroup)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1005
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsItemGroup()
func DeleteQGraphicsItemGroup(this *QGraphicsItemGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsItemGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1007
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addToGroup(QGraphicsItem *)
func (this *QGraphicsItemGroup) AddToGroup(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1008
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeFromGroup(QGraphicsItem *)
func (this *QGraphicsItemGroup) RemoveFromGroup(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1010
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect() const
func (this *QGraphicsItemGroup) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1011
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsItemGroup) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1011
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsItemGroup) Paint__(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1013
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *) const
func (this *QGraphicsItemGroup) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1014
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea() const
func (this *QGraphicsItemGroup) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1017
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const
func (this *QGraphicsItemGroup) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QGraphicsItemGroup__ = int

const QGraphicsItemGroup__Type QGraphicsItemGroup__ = 10

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
