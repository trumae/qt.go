package qtcore

// /usr/include/qt/QtCore/qversionnumber.h
// #include <qversionnumber.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QVersionNumber struct {
	*qtrt.CObject
}
type QVersionNumber_ITF interface {
	QVersionNumber_PTR() *QVersionNumber
}

func (ptr *QVersionNumber) QVersionNumber_PTR() *QVersionNumber { return ptr }

func (this *QVersionNumber) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVersionNumber) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVersionNumberFromPointer(cthis unsafe.Pointer) *QVersionNumber {
	return &QVersionNumber{&qtrt.CObject{cthis}}
}
func (*QVersionNumber) NewFromPointer(cthis unsafe.Pointer) *QVersionNumber {
	return NewQVersionNumberFromPointer(cthis)
}

// /usr/include/qt/QtCore/qversionnumber.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber()
func NewQVersionNumber() *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:242
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber(int)
func NewQVersionNumber_1(maj int) *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Ei", qtrt.FFI_TYPE_POINTER, maj)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:245
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber(int, int)
func NewQVersionNumber_2(maj int, min int) *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Eii", qtrt.FFI_TYPE_POINTER, maj, min)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:248
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QVersionNumber(int, int, int)
func NewQVersionNumber_3(maj int, min int, mic int) *QVersionNumber {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberC2Eiii", qtrt.FFI_TYPE_POINTER, maj, min, mic)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVersionNumber)
	return gothis
}

// /usr/include/qt/QtCore/qversionnumber.h:251
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QVersionNumber) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qversionnumber.h:254
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNormalized() const
func (this *QVersionNumber) IsNormalized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12isNormalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qversionnumber.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int majorVersion() const
func (this *QVersionNumber) MajorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12majorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:260
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int minorVersion() const
func (this *QVersionNumber) MinorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12minorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:263
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int microVersion() const
func (this *QVersionNumber) MicroVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12microVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:266
// index:0
// Public Visibility=Default Availability=Available
// [8] QVersionNumber normalized() const
func (this *QVersionNumber) Normalized() *QVersionNumber /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber10normalizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}

// /usr/include/qt/QtCore/qversionnumber.h:270
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int segmentAt(int) const
func (this *QVersionNumber) SegmentAt(index int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber9segmentAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:273
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int segmentCount() const
func (this *QVersionNumber) SegmentCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber12segmentCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qversionnumber.h:276
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPrefixOf(const QVersionNumber &) const
func (this *QVersionNumber) IsPrefixOf(other QVersionNumber_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QVersionNumber_PTR() != nil {
		convArg0 = other.QVersionNumber_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber10isPrefixOfERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qversionnumber.h:278
// index:0
// Public static Visibility=Default Availability=Available
// [4] int compare(const QVersionNumber &, const QVersionNumber &)
func (this *QVersionNumber) Compare(v1 QVersionNumber_ITF, v2 QVersionNumber_ITF) int {
	var convArg0 unsafe.Pointer
	if v1 != nil && v1.QVersionNumber_PTR() != nil {
		convArg0 = v1.QVersionNumber_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if v2 != nil && v2.QVersionNumber_PTR() != nil {
		convArg1 = v2.QVersionNumber_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber7compareERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QVersionNumber_Compare(v1 QVersionNumber_ITF, v2 QVersionNumber_ITF) int {
	var nilthis *QVersionNumber
	rv := nilthis.Compare(v1, v2)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:280
// index:0
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber commonPrefix(const QVersionNumber &, const QVersionNumber &)
func (this *QVersionNumber) CommonPrefix(v1 QVersionNumber_ITF, v2 QVersionNumber_ITF) *QVersionNumber /*123*/ {
	var convArg0 unsafe.Pointer
	if v1 != nil && v1.QVersionNumber_PTR() != nil {
		convArg0 = v1.QVersionNumber_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if v2 != nil && v2.QVersionNumber_PTR() != nil {
		convArg1 = v2.QVersionNumber_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber12commonPrefixERKS_S1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_CommonPrefix(v1 QVersionNumber_ITF, v2 QVersionNumber_ITF) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.CommonPrefix(v1, v2)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:282
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const
func (this *QVersionNumber) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QVersionNumber8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qversionnumber.h:284
// index:0
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(const QString &, int *)
func (this *QVersionNumber) FromString(string string, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringERK7QStringPi", qtrt.FFI_TYPE_POINTER, convArg0, suffixIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_FromString(string string, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.FromString(string, suffixIndex)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:284
// index:0
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(const QString &, int *)
func (this *QVersionNumber) FromString__(string string) *QVersionNumber /*123*/ {
	var tmpArg0 = NewQString_5(string)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int *=Pointer, =Invalid,
	var suffixIndex unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringERK7QStringPi", qtrt.FFI_TYPE_POINTER, convArg0, suffixIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}

// /usr/include/qt/QtCore/qversionnumber.h:286
// index:1
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(QLatin1String, int *)
func (this *QVersionNumber) FromString_1(string QLatin1String_ITF /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var convArg0 unsafe.Pointer
	if string != nil && string.QLatin1String_PTR() != nil {
		convArg0 = string.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE13QLatin1StringPi", qtrt.FFI_TYPE_POINTER, convArg0, suffixIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_FromString_1(string QLatin1String_ITF /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.FromString_1(string, suffixIndex)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:286
// index:1
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(QLatin1String, int *)
func (this *QVersionNumber) FromString_1_(string QLatin1String_ITF /*123*/) *QVersionNumber /*123*/ {
	var convArg0 unsafe.Pointer
	if string != nil && string.QLatin1String_PTR() != nil {
		convArg0 = string.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, int *=Pointer, =Invalid,
	var suffixIndex unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE13QLatin1StringPi", qtrt.FFI_TYPE_POINTER, convArg0, suffixIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}

// /usr/include/qt/QtCore/qversionnumber.h:287
// index:2
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(QStringView, int *)
func (this *QVersionNumber) FromString_2(string QStringView_ITF /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var convArg0 unsafe.Pointer
	if string != nil && string.QStringView_PTR() != nil {
		convArg0 = string.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE11QStringViewPi", qtrt.FFI_TYPE_POINTER, convArg0, suffixIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}
func QVersionNumber_FromString_2(string QStringView_ITF /*123*/, suffixIndex unsafe.Pointer /*666*/) *QVersionNumber /*123*/ {
	var nilthis *QVersionNumber
	rv := nilthis.FromString_2(string, suffixIndex)
	return rv
}

// /usr/include/qt/QtCore/qversionnumber.h:287
// index:2
// Public static Visibility=Default Availability=Available
// [8] QVersionNumber fromString(QStringView, int *)
func (this *QVersionNumber) FromString_2_(string QStringView_ITF /*123*/) *QVersionNumber /*123*/ {
	var convArg0 unsafe.Pointer
	if string != nil && string.QStringView_PTR() != nil {
		convArg0 = string.QStringView_PTR().GetCthis()
	}
	// arg: 1, int *=Pointer, =Invalid,
	var suffixIndex unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumber10fromStringE11QStringViewPi", qtrt.FFI_TYPE_POINTER, convArg0, suffixIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVersionNumberFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVersionNumber)
	return rv2
}

func DeleteQVersionNumber(this *QVersionNumber) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QVersionNumberD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QVersionNumber__ = int

const QVersionNumber__InlineSegmentMarker QVersionNumber__ = 0
const QVersionNumber__InlineSegmentStartIdx QVersionNumber__ = 1
const QVersionNumber__InlineSegmentCount QVersionNumber__ = 7

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
}

//  keep block end
