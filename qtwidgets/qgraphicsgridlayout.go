package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h
// #include <qgraphicsgridlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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

type QGraphicsGridLayout struct {
	*QGraphicsLayout
}
type QGraphicsGridLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsGridLayout_PTR() *QGraphicsGridLayout
}

func (ptr *QGraphicsGridLayout) QGraphicsGridLayout_PTR() *QGraphicsGridLayout { return ptr }

func (this *QGraphicsGridLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayout.GetCthis()
	}
}
func (this *QGraphicsGridLayout) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayout = NewQGraphicsLayoutFromPointer(cthis)
}
func NewQGraphicsGridLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsGridLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsGridLayout{bcthis0}
}
func (*QGraphicsGridLayout) NewFromPointer(cthis unsafe.Pointer) *QGraphicsGridLayout {
	return NewQGraphicsGridLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsGridLayout(QGraphicsLayoutItem *)
func NewQGraphicsGridLayout(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsGridLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsGridLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsGridLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsGridLayout(QGraphicsLayoutItem *)
func NewQGraphicsGridLayout__() *QGraphicsGridLayout {
	// arg: 0, QGraphicsLayoutItem *=Pointer, QGraphicsLayoutItem=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsGridLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsGridLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsGridLayout()
func DeleteQGraphicsGridLayout(this *QGraphicsGridLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, int, int, Qt::Alignment)
func (this *QGraphicsGridLayout) AddItem(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int, rowSpan int, columnSpan int, alignment int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemiiii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, rowSpan, columnSpan, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, int, int, Qt::Alignment)
func (this *QGraphicsGridLayout) AddItem__(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int, rowSpan int, columnSpan int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	// arg: 5, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemiiii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, rowSpan, columnSpan, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, Qt::Alignment)
func (this *QGraphicsGridLayout) AddItem_1(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int, alignment int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *, int, int, Qt::Alignment)
func (this *QGraphicsGridLayout) AddItem_1_(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, row int, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	// arg: 3, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout7addItemEP19QGraphicsLayoutItemii6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, row, column, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalSpacing(qreal)
func (this *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout20setHorizontalSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal horizontalSpacing() const
func (this *QGraphicsGridLayout) HorizontalSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout17horizontalSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalSpacing(qreal)
func (this *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout18setVerticalSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal verticalSpacing() const
func (this *QGraphicsGridLayout) VerticalSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15verticalSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(qreal)
func (this *QGraphicsGridLayout) SetSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10setSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowSpacing(int, qreal)
func (this *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout13setRowSpacingEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowSpacing(int) const
func (this *QGraphicsGridLayout) RowSpacing(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout10rowSpacingEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnSpacing(int, qreal)
func (this *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout16setColumnSpacingEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnSpacing(int) const
func (this *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout13columnSpacingEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowStretchFactor(int, int)
func (this *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowStretchFactorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowStretchFactor(int) const
func (this *QGraphicsGridLayout) RowStretchFactor(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowStretchFactorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnStretchFactor(int, int)
func (this *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout22setColumnStretchFactorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnStretchFactor(int) const
func (this *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout19columnStretchFactorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowMinimumHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMinimumHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowMinimumHeight(int) const
func (this *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMinimumHeightEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowPreferredHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setRowPreferredHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowPreferredHeight(int) const
func (this *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18rowPreferredHeightEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowMaximumHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setRowMaximumHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rowMaximumHeight(int) const
func (this *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout16rowMaximumHeightEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowFixedHeight(int, qreal)
func (this *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout17setRowFixedHeightEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnMinimumWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMinimumWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnMinimumWidth(int) const
func (this *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMinimumWidthEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnPreferredWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout23setColumnPreferredWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnPreferredWidth(int) const
func (this *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout20columnPreferredWidthEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnMaximumWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout21setColumnMaximumWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal columnMaximumWidth(int) const
func (this *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout18columnMaximumWidthEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnFixedWidth(int, qreal)
func (this *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout19setColumnFixedWidthEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowAlignment(int, Qt::Alignment)
func (this *QGraphicsGridLayout) SetRowAlignment(row int, alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout15setRowAlignmentEi6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment rowAlignment(int) const
func (this *QGraphicsGridLayout) RowAlignment(row int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout12rowAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnAlignment(int, Qt::Alignment)
func (this *QGraphicsGridLayout) SetColumnAlignment(column int, alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout18setColumnAlignmentEi6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment columnAlignment(int) const
func (this *QGraphicsGridLayout) ColumnAlignment(column int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout15columnAlignmentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(QGraphicsLayoutItem *, Qt::Alignment)
func (this *QGraphicsGridLayout) SetAlignment(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, alignment int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout12setAlignmentEP19QGraphicsLayoutItem6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment(QGraphicsLayoutItem *) const
func (this *QGraphicsGridLayout) Alignment(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout9alignmentEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowCount() const
func (this *QGraphicsGridLayout) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount() const
func (this *QGraphicsGridLayout) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int, int) const
func (this *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:110
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int) const
func (this *QGraphicsGridLayout) ItemAt_1(index int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:109
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count() const
func (this *QGraphicsGridLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:111
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)
func (this *QGraphicsGridLayout) RemoveAt(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QGraphicsLayoutItem *)
func (this *QGraphicsGridLayout) RemoveItem(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10removeItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QGraphicsGridLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)
func (this *QGraphicsGridLayout) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsGridLayout11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const
func (this *QGraphicsGridLayout) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsgridlayout.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const
func (this *QGraphicsGridLayout) SizeHint__(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsGridLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
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
