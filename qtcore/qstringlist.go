package qtcore

// /usr/include/qt/QtCore/qstringlist.h
// #include <qstringlist.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QStringList struct {
	*qtrt.CObject
}
type QStringList_ITF interface {
	QStringList_PTR() *QStringList
}

func (ptr *QStringList) QStringList_PTR() *QStringList { return ptr }

func (this *QStringList) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringList) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringListFromPointer(cthis unsafe.Pointer) *QStringList {
	return &QStringList{&qtrt.CObject{cthis}}
}
func (*QStringList) NewFromPointer(cthis unsafe.Pointer) *QStringList {
	return NewQStringListFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringlist.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStringList()
func NewQStringList() *QStringList {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringList)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:106
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QStringList(const QString &)
func NewQStringList_1(i string) *QStringList {
	var tmpArg0 = NewQString_5(i)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringList)
	return gothis
}

// /usr/include/qt/QtCore/qstringlist.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const
func (this *QStringList) Contains(str string, cs int) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool contains(const QString &, Qt::CaseSensitivity) const
func (this *QStringList) Contains__(str string) bool {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsERK7QStringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:123
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringList) Contains_1(str QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:123
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool contains(QLatin1String, Qt::CaseSensitivity) const
func (this *QStringList) Contains_1_(str QLatin1String_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if str != nil && str.QLatin1String_PTR() != nil {
		convArg0 = str.QLatin1String_PTR().GetCthis()
	}
	// arg: 1, Qt::CaseSensitivity=Elaborated, Qt::CaseSensitivity=Enum,
	cs := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList8containsE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlist.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringList operator+(const QStringList &) const
func (this *QStringList) Operator_add(other QStringList_ITF) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStringList_PTR() != nil {
		convArg0 = other.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringListplERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstringlist.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringList & operator<<(const QString &)
func (this *QStringList) Operator_left_shift(str string) *QStringList {
	var tmpArg0 = NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListlsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstringlist.h:129
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QStringList & operator<<(const QStringList &)
func (this *QStringList) Operator_left_shift_1(l QStringList_ITF) *QStringList {
	var convArg0 unsafe.Pointer
	if l != nil && l.QStringList_PTR() != nil {
		convArg0 = l.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListlsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstringlist.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegExp &, int) const
func (this *QStringList) IndexOf(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegExp &, int) const
func (this *QStringList) IndexOf__(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:137
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(QRegExp &, int) const
func (this *QStringList) IndexOf_1(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:137
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(QRegExp &, int) const
func (this *QStringList) IndexOf_1_(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:143
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int) const
func (this *QStringList) IndexOf_2(re QRegularExpression_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:143
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int indexOf(const QRegularExpression &, int) const
func (this *QStringList) IndexOf_2_(re QRegularExpression_ITF) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList7indexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegExp &, int) const
func (this *QStringList) LastIndexOf(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegExp &, int) const
func (this *QStringList) LastIndexOf__(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:138
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(QRegExp &, int) const
func (this *QStringList) LastIndexOf_1(rx QRegExp_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:138
// index:1
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(QRegExp &, int) const
func (this *QStringList) LastIndexOf_1_(rx QRegExp_ITF) int {
	var convArg0 unsafe.Pointer
	if rx != nil && rx.QRegExp_PTR() != nil {
		convArg0 = rx.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfER7QRegExpi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:144
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int) const
func (this *QStringList) LastIndexOf_2(re QRegularExpression_ITF, from int) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlist.h:144
// index:2
// Public inline Visibility=Default Availability=Available
// [4] int lastIndexOf(const QRegularExpression &, int) const
func (this *QStringList) LastIndexOf_2_(re QRegularExpression_ITF) int {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQStringList(this *QStringList) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringListD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
}

//  keep block end
