package qtgui

// /usr/include/qt/QtGui/qstatictext.h
// #include <qstatictext.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 50
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QStaticText struct {
	*qtrt.CObject
}
type QStaticText_ITF interface {
	QStaticText_PTR() *QStaticText
}

func (ptr *QStaticText) QStaticText_PTR() *QStaticText { return ptr }

func (this *QStaticText) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStaticText) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStaticTextFromPointer(cthis unsafe.Pointer) *QStaticText {
	return &QStaticText{&qtrt.CObject{cthis}}
}
func (*QStaticText) NewFromPointer(cthis unsafe.Pointer) *QStaticText {
	return NewQStaticTextFromPointer(cthis)
}

// /usr/include/qt/QtGui/qstatictext.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStaticText()
func NewQStaticText() *QStaticText {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStaticText)
	return gothis
}

// /usr/include/qt/QtGui/qstatictext.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStaticText(const QString &)
func NewQStaticText_1(text string) *QStaticText {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStaticText)
	return gothis
}

// /usr/include/qt/QtGui/qstatictext.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStaticText & operator=(QStaticText &&)
func (this *QStaticText) Operator_equal(other unsafe.Pointer /*333*/) *QStaticText {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStaticText)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:70
// index:1
// Public Visibility=Default Availability=Available
// [8] QStaticText & operator=(const QStaticText &)
func (this *QStaticText) Operator_equal_1(arg0 QStaticText_ITF) *QStaticText {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStaticText_PTR() != nil {
		convArg0 = arg0.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStaticText)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStaticText()
func DeleteQStaticText(this *QStaticText) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qstatictext.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QStaticText &)
func (this *QStaticText) Swap(other QStaticText_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStaticText_PTR() != nil {
		convArg0 = other.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QStaticText) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const
func (this *QStaticText) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstatictext.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextFormat(Qt::TextFormat)
func (this *QStaticText) SetTextFormat(textFormat int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText13setTextFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), textFormat)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat textFormat() const
func (this *QStaticText) TextFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText10textFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstatictext.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextWidth(qreal)
func (this *QStaticText) SetTextWidth(textWidth float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText12setTextWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), textWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal textWidth() const
func (this *QStaticText) TextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText9textWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qstatictext.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextOption(const QTextOption &)
func (this *QStaticText) SetTextOption(textOption QTextOption_ITF) {
	var convArg0 unsafe.Pointer
	if textOption != nil && textOption.QTextOption_PTR() != nil {
		convArg0 = textOption.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText13setTextOptionERK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:85
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextOption textOption() const
func (this *QStaticText) TextOption() *QTextOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText10textOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size() const
func (this *QStaticText) Size() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepare(const QTransform &, const QFont &)
func (this *QStaticText) Prepare(matrix QTransform_ITF, font QFont_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg1 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepare(const QTransform &, const QFont &)
func (this *QStaticText) Prepare__() {
	// arg: 0, const QTransform &=LValueReference, QTransform=Record,
	var convArg0 unsafe.Pointer
	// arg: 1, const QFont &=LValueReference, QFont=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepare(const QTransform &, const QFont &)
func (this *QStaticText) Prepare__1(matrix QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, const QFont &=LValueReference, QFont=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPerformanceHint(enum QStaticText::PerformanceHint)
func (this *QStaticText) SetPerformanceHint(performanceHint int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText18setPerformanceHintENS_15PerformanceHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), performanceHint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] QStaticText::PerformanceHint performanceHint() const
func (this *QStaticText) PerformanceHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText15performanceHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstatictext.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QStaticText &) const
func (this *QStaticText) Operator_equal_equal(arg0 QStaticText_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStaticText_PTR() != nil {
		convArg0 = arg0.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticTexteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstatictext.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QStaticText &) const
func (this *QStaticText) Operator_not_equal(arg0 QStaticText_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStaticText_PTR() != nil {
		convArg0 = arg0.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticTextneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QStaticText__PerformanceHint = int

const QStaticText__ModerateCaching QStaticText__PerformanceHint = 0
const QStaticText__AggressiveCaching QStaticText__PerformanceHint = 1

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
}

//  keep block end
