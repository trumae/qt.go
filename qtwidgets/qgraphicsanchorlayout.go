package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h
// #include <qgraphicsanchorlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsAnchorLayout) InheritSizeHint(f func(which int, constraint *qtcore.QSizeF) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sizeHint", f)
}

type QGraphicsAnchorLayout struct {
	*QGraphicsLayout
}
type QGraphicsAnchorLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsAnchorLayout_PTR() *QGraphicsAnchorLayout
}

func (ptr *QGraphicsAnchorLayout) QGraphicsAnchorLayout_PTR() *QGraphicsAnchorLayout { return ptr }

func (this *QGraphicsAnchorLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayout.GetCthis()
	}
}
func (this *QGraphicsAnchorLayout) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayout = NewQGraphicsLayoutFromPointer(cthis)
}
func NewQGraphicsAnchorLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsAnchorLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsAnchorLayout{bcthis0}
}
func (*QGraphicsAnchorLayout) NewFromPointer(cthis unsafe.Pointer) *QGraphicsAnchorLayout {
	return NewQGraphicsAnchorLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsAnchorLayout(QGraphicsLayoutItem *)
func NewQGraphicsAnchorLayout(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsAnchorLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsAnchorLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsAnchorLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsAnchorLayout(QGraphicsLayoutItem *)
func NewQGraphicsAnchorLayout__() *QGraphicsAnchorLayout {
	// arg: 0, QGraphicsLayoutItem *=Pointer, QGraphicsLayoutItem=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsAnchorLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsAnchorLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsAnchorLayout()
func DeleteQGraphicsAnchorLayout(this *QGraphicsAnchorLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsAnchor * addAnchor(QGraphicsLayoutItem *, Qt::AnchorPoint, QGraphicsLayoutItem *, Qt::AnchorPoint)
func (this *QGraphicsAnchorLayout) AddAnchor(firstItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, firstEdge int, secondItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, secondEdge int) *QGraphicsAnchor /*777 QGraphicsAnchor **/ {
	var convArg0 unsafe.Pointer
	if firstItem != nil && firstItem.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = firstItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if secondItem != nil && secondItem.QGraphicsLayoutItem_PTR() != nil {
		convArg2 = secondItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout9addAnchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, firstEdge, convArg2, secondEdge)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsAnchorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsAnchor * anchor(QGraphicsLayoutItem *, Qt::AnchorPoint, QGraphicsLayoutItem *, Qt::AnchorPoint)
func (this *QGraphicsAnchorLayout) Anchor(firstItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, firstEdge int, secondItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, secondEdge int) *QGraphicsAnchor /*777 QGraphicsAnchor **/ {
	var convArg0 unsafe.Pointer
	if firstItem != nil && firstItem.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = firstItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if secondItem != nil && secondItem.QGraphicsLayoutItem_PTR() != nil {
		convArg2 = secondItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout6anchorEP19QGraphicsLayoutItemN2Qt11AnchorPointES1_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, firstEdge, convArg2, secondEdge)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsAnchorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addCornerAnchors(QGraphicsLayoutItem *, Qt::Corner, QGraphicsLayoutItem *, Qt::Corner)
func (this *QGraphicsAnchorLayout) AddCornerAnchors(firstItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, firstCorner int, secondItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, secondCorner int) {
	var convArg0 unsafe.Pointer
	if firstItem != nil && firstItem.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = firstItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if secondItem != nil && secondItem.QGraphicsLayoutItem_PTR() != nil {
		convArg2 = secondItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout16addCornerAnchorsEP19QGraphicsLayoutItemN2Qt6CornerES1_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, firstCorner, convArg2, secondCorner)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAnchors(QGraphicsLayoutItem *, QGraphicsLayoutItem *, Qt::Orientations)
func (this *QGraphicsAnchorLayout) AddAnchors(firstItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, secondItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, orientations int) {
	var convArg0 unsafe.Pointer
	if firstItem != nil && firstItem.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = firstItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if secondItem != nil && secondItem.QGraphicsLayoutItem_PTR() != nil {
		convArg1 = secondItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10addAnchorsEP19QGraphicsLayoutItemS1_6QFlagsIN2Qt11OrientationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, orientations)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAnchors(QGraphicsLayoutItem *, QGraphicsLayoutItem *, Qt::Orientations)
func (this *QGraphicsAnchorLayout) AddAnchors__(firstItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, secondItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if firstItem != nil && firstItem.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = firstItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if secondItem != nil && secondItem.QGraphicsLayoutItem_PTR() != nil {
		convArg1 = secondItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	// arg: 2, Qt::Orientations=Elaborated, Qt::Orientations=Typedef, QFlags<Qt::Orientation>
	orientations := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10addAnchorsEP19QGraphicsLayoutItemS1_6QFlagsIN2Qt11OrientationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, orientations)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout20setHorizontalSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout18setVerticalSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(qreal)
func (this *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10setSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal horizontalSpacing() const
func (this *QGraphicsAnchorLayout) HorizontalSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout17horizontalSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal verticalSpacing() const
func (this *QGraphicsAnchorLayout) VerticalSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout15verticalSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)
func (this *QGraphicsAnchorLayout) RemoveAt(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)
func (this *QGraphicsAnchorLayout) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:102
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count() const
func (this *QGraphicsAnchorLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int) const
func (this *QGraphicsAnchorLayout) ItemAt(index int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QGraphicsAnchorLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsAnchorLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const
func (this *QGraphicsAnchorLayout) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const
func (this *QGraphicsAnchorLayout) SizeHint__(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
