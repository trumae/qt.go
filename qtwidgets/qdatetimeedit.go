package qtwidgets

// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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

// void keyPressEvent(class QKeyEvent *)
func (this *QDateTimeEdit) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QDateTimeEdit) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QDateTimeEdit) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QDateTimeEdit) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// QValidator::State validate(class QString &, int &)
func (this *QDateTimeEdit) InheritValidate(f func(input string, pos int) int) {
	qtrt.SetAllInheritCallback(this, "validate", f)
}

// void fixup(class QString &)
func (this *QDateTimeEdit) InheritFixup(f func(input string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "fixup", f)
}

// QDateTime dateTimeFromText(const class QString &)
func (this *QDateTimeEdit) InheritDateTimeFromText(f func(text string) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "dateTimeFromText", f)
}

// QString textFromDateTime(const class QDateTime &)
func (this *QDateTimeEdit) InheritTextFromDateTime(f func(dt *qtcore.QDateTime) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "textFromDateTime", f)
}

// QAbstractSpinBox::StepEnabled stepEnabled()
func (this *QDateTimeEdit) InheritStepEnabled(f func() int) {
	qtrt.SetAllInheritCallback(this, "stepEnabled", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QDateTimeEdit) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QDateTimeEdit) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void initStyleOption(class QStyleOptionSpinBox *)
func (this *QDateTimeEdit) InheritInitStyleOption(f func(option *QStyleOptionSpinBox /*777 QStyleOptionSpinBox **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QDateTimeEdit struct {
	*QAbstractSpinBox
}
type QDateTimeEdit_ITF interface {
	QAbstractSpinBox_ITF
	QDateTimeEdit_PTR() *QDateTimeEdit
}

func (ptr *QDateTimeEdit) QDateTimeEdit_PTR() *QDateTimeEdit { return ptr }

func (this *QDateTimeEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSpinBox.GetCthis()
	}
}
func (this *QDateTimeEdit) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSpinBox = NewQAbstractSpinBoxFromPointer(cthis)
}
func NewQDateTimeEditFromPointer(cthis unsafe.Pointer) *QDateTimeEdit {
	bcthis0 := NewQAbstractSpinBoxFromPointer(cthis)
	return &QDateTimeEdit{bcthis0}
}
func (*QDateTimeEdit) NewFromPointer(cthis unsafe.Pointer) *QDateTimeEdit {
	return NewQDateTimeEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QDateTimeEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(QWidget *)
func NewQDateTimeEdit(parent QWidget_ITF /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(QWidget *)
func NewQDateTimeEdit__() *QDateTimeEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QDateTime &, QWidget *)
func NewQDateTimeEdit_1(dt qtcore.QDateTime_ITF, parent QWidget_ITF /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if dt != nil && dt.QDateTime_PTR() != nil {
		convArg0 = dt.QDateTime_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK9QDateTimeP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QDateTime &, QWidget *)
func NewQDateTimeEdit_1_(dt qtcore.QDateTime_ITF) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if dt != nil && dt.QDateTime_PTR() != nil {
		convArg0 = dt.QDateTime_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK9QDateTimeP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:97
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QDate &, QWidget *)
func NewQDateTimeEdit_2(d qtcore.QDate_ITF, parent QWidget_ITF /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if d != nil && d.QDate_PTR() != nil {
		convArg0 = d.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QDateP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:97
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QDate &, QWidget *)
func NewQDateTimeEdit_2_(d qtcore.QDate_ITF) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if d != nil && d.QDate_PTR() != nil {
		convArg0 = d.QDate_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QDateP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:98
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QTime &, QWidget *)
func NewQDateTimeEdit_3(t qtcore.QTime_ITF, parent QWidget_ITF /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if t != nil && t.QTime_PTR() != nil {
		convArg0 = t.QTime_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QTimeP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:98
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QTime &, QWidget *)
func NewQDateTimeEdit_3_(t qtcore.QTime_ITF) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if t != nil && t.QTime_PTR() != nil {
		convArg0 = t.QTime_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK5QTimeP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:194
// index:4
// Protected Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QVariant &, QVariant::Type, QWidget *)
func NewQDateTimeEdit_4(val qtcore.QVariant_ITF, parserType int, parent QWidget_ITF /*777 QWidget **/) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if val != nil && val.QVariant_PTR() != nil {
		convArg0 = val.QVariant_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg2 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK8QVariantNS0_4TypeEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, parserType, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:194
// index:4
// Protected Visibility=Default Availability=Available
// [-2] void QDateTimeEdit(const QVariant &, QVariant::Type, QWidget *)
func NewQDateTimeEdit_4_(val qtcore.QVariant_ITF, parserType int) *QDateTimeEdit {
	var convArg0 unsafe.Pointer
	if val != nil && val.QVariant_PTR() != nil {
		convArg0 = val.QVariant_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditC2ERK8QVariantNS0_4TypeEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, parserType, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateTimeEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDateTimeEdit()
func DeleteQDateTimeEdit(this *QDateTimeEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime dateTime() const
func (this *QDateTimeEdit) DateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8dateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate date() const
func (this *QDateTimeEdit) Date() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit4dateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime time() const
func (this *QDateTimeEdit) Time() *qtcore.QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit4timeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime minimumDateTime() const
func (this *QDateTimeEdit) MinimumDateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit15minimumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMinimumDateTime()
func (this *QDateTimeEdit) ClearMinimumDateTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMinimumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDateTime(const QDateTime &)
func (this *QDateTimeEdit) SetMinimumDateTime(dt qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if dt != nil && dt.QDateTime_PTR() != nil {
		convArg0 = dt.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMinimumDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime maximumDateTime() const
func (this *QDateTimeEdit) MaximumDateTime() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit15maximumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMaximumDateTime()
func (this *QDateTimeEdit) ClearMaximumDateTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit20clearMaximumDateTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumDateTime(const QDateTime &)
func (this *QDateTimeEdit) SetMaximumDateTime(dt qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if dt != nil && dt.QDateTime_PTR() != nil {
		convArg0 = dt.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18setMaximumDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateTimeRange(const QDateTime &, const QDateTime &)
func (this *QDateTimeEdit) SetDateTimeRange(min qtcore.QDateTime_ITF, max qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if min != nil && min.QDateTime_PTR() != nil {
		convArg0 = min.QDateTime_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if max != nil && max.QDateTime_PTR() != nil {
		convArg1 = max.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDateTimeRangeERK9QDateTimeS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate minimumDate() const
func (this *QDateTimeEdit) MinimumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumDate(const QDate &)
func (this *QDateTimeEdit) SetMinimumDate(min qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if min != nil && min.QDate_PTR() != nil {
		convArg0 = min.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMinimumDate()
func (this *QDateTimeEdit) ClearMinimumDate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QDate maximumDate() const
func (this *QDateTimeEdit) MaximumDate() *qtcore.QDate /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDate)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumDate(const QDate &)
func (this *QDateTimeEdit) SetMaximumDate(max qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if max != nil && max.QDate_PTR() != nil {
		convArg0 = max.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMaximumDate()
func (this *QDateTimeEdit) ClearMaximumDate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumDateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateRange(const QDate &, const QDate &)
func (this *QDateTimeEdit) SetDateRange(min qtcore.QDate_ITF, max qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if min != nil && min.QDate_PTR() != nil {
		convArg0 = min.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if max != nil && max.QDate_PTR() != nil {
		convArg1 = max.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit12setDateRangeERK5QDateS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime minimumTime() const
func (this *QDateTimeEdit) MinimumTime() *qtcore.QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11minimumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumTime(const QTime &)
func (this *QDateTimeEdit) SetMinimumTime(min qtcore.QTime_ITF) {
	var convArg0 unsafe.Pointer
	if min != nil && min.QTime_PTR() != nil {
		convArg0 = min.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMinimumTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMinimumTime()
func (this *QDateTimeEdit) ClearMinimumTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMinimumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:129
// index:0
// Public Visibility=Default Availability=Available
// [4] QTime maximumTime() const
func (this *QDateTimeEdit) MaximumTime() *qtcore.QTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11maximumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumTime(const QTime &)
func (this *QDateTimeEdit) SetMaximumTime(max qtcore.QTime_ITF) {
	var convArg0 unsafe.Pointer
	if max != nil && max.QTime_PTR() != nil {
		convArg0 = max.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit14setMaximumTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMaximumTime()
func (this *QDateTimeEdit) ClearMaximumTime() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16clearMaximumTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeRange(const QTime &, const QTime &)
func (this *QDateTimeEdit) SetTimeRange(min qtcore.QTime_ITF, max qtcore.QTime_ITF) {
	var convArg0 unsafe.Pointer
	if min != nil && min.QTime_PTR() != nil {
		convArg0 = min.QTime_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if max != nil && max.QTime_PTR() != nil {
		convArg1 = max.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit12setTimeRangeERK5QTimeS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] QDateTimeEdit::Sections displayedSections() const
func (this *QDateTimeEdit) DisplayedSections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit17displayedSectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] QDateTimeEdit::Section currentSection() const
func (this *QDateTimeEdit) CurrentSection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit14currentSectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] QDateTimeEdit::Section sectionAt(int) const
func (this *QDateTimeEdit) SectionAt(index int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit9sectionAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetCurrentSection(section int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCurrentSectionENS_7SectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentSectionIndex() const
func (this *QDateTimeEdit) CurrentSectionIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit19currentSectionIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentSectionIndex(int)
func (this *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit22setCurrentSectionIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QCalendarWidget * calendarWidget() const
func (this *QDateTimeEdit) CalendarWidget() *QCalendarWidget /*777 QCalendarWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit14calendarWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCalendarWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCalendarWidget(QCalendarWidget *)
func (this *QDateTimeEdit) SetCalendarWidget(calendarWidget QCalendarWidget_ITF /*777 QCalendarWidget **/) {
	var convArg0 unsafe.Pointer
	if calendarWidget != nil && calendarWidget.QCalendarWidget_PTR() != nil {
		convArg0 = calendarWidget.QCalendarWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit17setCalendarWidgetEP15QCalendarWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] int sectionCount() const
func (this *QDateTimeEdit) SectionCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit12sectionCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelectedSection(enum QDateTimeEdit::Section)
func (this *QDateTimeEdit) SetSelectedSection(section int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18setSelectedSectionENS_7SectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sectionText(enum QDateTimeEdit::Section) const
func (this *QDateTimeEdit) SectionText(section int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11sectionTextENS_7SectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:152
// index:0
// Public Visibility=Default Availability=Available
// [8] QString displayFormat() const
func (this *QDateTimeEdit) DisplayFormat() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit13displayFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDisplayFormat(const QString &)
func (this *QDateTimeEdit) SetDisplayFormat(format string) {
	var tmpArg0 = qtcore.NewQString_5(format)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16setDisplayFormatERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:155
// index:0
// Public Visibility=Default Availability=Available
// [1] bool calendarPopup() const
func (this *QDateTimeEdit) CalendarPopup() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit13calendarPopupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCalendarPopup(_Bool)
func (this *QDateTimeEdit) SetCalendarPopup(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit16setCalendarPopupEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:158
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TimeSpec timeSpec() const
func (this *QDateTimeEdit) TimeSpec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8timeSpecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimeSpec(Qt::TimeSpec)
func (this *QDateTimeEdit) SetTimeSpec(spec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11setTimeSpecEN2Qt8TimeSpecE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:161
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const
func (this *QDateTimeEdit) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:163
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()
func (this *QDateTimeEdit) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:164
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void stepBy(int)
func (this *QDateTimeEdit) StepBy(steps int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit6stepByEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), steps)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QDateTimeEdit) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dateTimeChanged(const QDateTime &)
func (this *QDateTimeEdit) DateTimeChanged(dateTime qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if dateTime != nil && dateTime.QDateTime_PTR() != nil {
		convArg0 = dateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit15dateTimeChangedERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void timeChanged(const QTime &)
func (this *QDateTimeEdit) TimeChanged(time qtcore.QTime_ITF) {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11timeChangedERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dateChanged(const QDate &)
func (this *QDateTimeEdit) DateChanged(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11dateChangedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDateTime(const QDateTime &)
func (this *QDateTimeEdit) SetDateTime(dateTime qtcore.QDateTime_ITF) {
	var convArg0 unsafe.Pointer
	if dateTime != nil && dateTime.QDateTime_PTR() != nil {
		convArg0 = dateTime.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit11setDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDate(const QDate &)
func (this *QDateTimeEdit) SetDate(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit7setDateERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTime(const QTime &)
func (this *QDateTimeEdit) SetTime(time qtcore.QTime_ITF) {
	var convArg0 unsafe.Pointer
	if time != nil && time.QTime_PTR() != nil {
		convArg0 = time.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit7setTimeERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:178
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QDateTimeEdit) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:180
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QDateTimeEdit) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:182
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QDateTimeEdit) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:183
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QDateTimeEdit) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:184
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QValidator::State validate(QString &, int &) const
func (this *QDateTimeEdit) Validate(input string, pos int) int {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit8validateER7QStringRi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pos)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:185
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void fixup(QString &) const
func (this *QDateTimeEdit) Fixup(input string) {
	var tmpArg0 = qtcore.NewQString_5(input)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit5fixupER7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:187
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QDateTime dateTimeFromText(const QString &) const
func (this *QDateTimeEdit) DateTimeFromText(text string) *qtcore.QDateTime /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:188
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QString textFromDateTime(const QDateTime &) const
func (this *QDateTimeEdit) TextFromDateTime(dt qtcore.QDateTime_ITF) string {
	var convArg0 unsafe.Pointer
	if dt != nil && dt.QDateTime_PTR() != nil {
		convArg0 = dt.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:189
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] QAbstractSpinBox::StepEnabled stepEnabled() const
func (this *QDateTimeEdit) StepEnabled() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit11stepEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:190
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QDateTimeEdit) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:191
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QDateTimeEdit) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QDateTimeEdit10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:192
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionSpinBox *) const
func (this *QDateTimeEdit) InitStyleOption(option QStyleOptionSpinBox_ITF /*777 QStyleOptionSpinBox **/) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionSpinBox_PTR() != nil {
		convArg0 = option.QStyleOptionSpinBox_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QDateTimeEdit__Section = int

const QDateTimeEdit__NoSection QDateTimeEdit__Section = 0
const QDateTimeEdit__AmPmSection QDateTimeEdit__Section = 1
const QDateTimeEdit__MSecSection QDateTimeEdit__Section = 2
const QDateTimeEdit__SecondSection QDateTimeEdit__Section = 4
const QDateTimeEdit__MinuteSection QDateTimeEdit__Section = 8
const QDateTimeEdit__HourSection QDateTimeEdit__Section = 16
const QDateTimeEdit__DaySection QDateTimeEdit__Section = 256
const QDateTimeEdit__MonthSection QDateTimeEdit__Section = 512
const QDateTimeEdit__YearSection QDateTimeEdit__Section = 1024
const QDateTimeEdit__TimeSections_Mask QDateTimeEdit__Section = 31
const QDateTimeEdit__DateSections_Mask QDateTimeEdit__Section = 1792

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
