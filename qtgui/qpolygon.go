package qtgui

// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 77
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

type QPolygon struct {
	*qtrt.CObject
}
type QPolygon_ITF interface {
	QPolygon_PTR() *QPolygon
}

func (ptr *QPolygon) QPolygon_PTR() *QPolygon { return ptr }

func (this *QPolygon) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPolygon) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPolygonFromPointer(cthis unsafe.Pointer) *QPolygon {
	return &QPolygon{&qtrt.CObject{cthis}}
}
func (*QPolygon) NewFromPointer(cthis unsafe.Pointer) *QPolygon {
	return NewQPolygonFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpolygon.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPolygon()
func NewQPolygon() *QPolygon {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygon)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:61
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QPolygon(int)
func NewQPolygon_1(size int) *QPolygon {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonC2Ei", qtrt.FFI_TYPE_POINTER, size)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygon)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPolygon(const QRect &, _Bool)
func NewQPolygon_2(r qtcore.QRect_ITF, closed bool) *QPolygon {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonC2ERK5QRectb", qtrt.FFI_TYPE_POINTER, convArg0, closed)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygon)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPolygon(const QRect &, _Bool)
func NewQPolygon_2_(r qtcore.QRect_ITF) *QPolygon {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid,
	closed := false
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonC2ERK5QRectb", qtrt.FFI_TYPE_POINTER, convArg0, closed)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygon)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:67
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPolygon(int, const int *)
func NewQPolygon_3(nPoints int, points unsafe.Pointer /*666*/) *QPolygon {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonC2EiPKi", qtrt.FFI_TYPE_POINTER, nPoints, points)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPolygon)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QPolygon()
func DeleteQPolygon(this *QPolygon) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpolygon.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPolygon & operator=(QPolygon &&)
func (this *QPolygon) Operator_equal(other unsafe.Pointer /*333*/) *QPolygon {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPolygon & operator=(const QPolygon &)
func (this *QPolygon) Operator_equal_1(other QPolygon_ITF) *QPolygon {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPolygon_PTR() != nil {
		convArg0 = other.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygonaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPolygon &)
func (this *QPolygon) Swap(other QPolygon_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPolygon_PTR() != nil {
		convArg0 = other.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void translate(int, int)
func (this *QPolygon) Translate(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon9translateEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:79
// index:1
// Public Visibility=Default Availability=Available
// [-2] void translate(const QPoint &)
func (this *QPolygon) Translate_1(offset qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg0 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon9translateERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon translated(int, int) const
func (this *QPolygon) Translated(dx int, dy int) *QPolygon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon10translatedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPolygon translated(const QPoint &) const
func (this *QPolygon) Translated_1(offset qtcore.QPoint_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg0 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon10translatedERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:84
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect() const
func (this *QPolygon) BoundingRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void point(int, int *, int *) const
func (this *QPolygon) Point(i int, x unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon5pointEiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i, x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:87
// index:1
// Public Visibility=Default Availability=Available
// [8] QPoint point(int) const
func (this *QPolygon) Point_1(i int) *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon5pointEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPoint(int, int, int)
func (this *QPolygon) SetPoint(index int, x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon8setPointEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPoint(int, const QPoint &)
func (this *QPolygon) SetPoint_1(index int, p qtcore.QPoint_ITF) {
	var convArg1 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg1 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon8setPointEiRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPoints(int, const int *)
func (this *QPolygon) SetPoints(nPoints int, points unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon9setPointsEiPKi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nPoints, points)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void putPoints(int, int, const int *)
func (this *QPolygon) PutPoints(index int, nPoints int, points unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiPKi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, nPoints, points)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void putPoints(int, int, const QPolygon &, int)
func (this *QPolygon) PutPoints_1(index int, nPoints int, from QPolygon_ITF, fromIndex int) {
	var convArg2 unsafe.Pointer
	if from != nil && from.QPolygon_PTR() != nil {
		convArg2 = from.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiRKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, nPoints, convArg2, fromIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void putPoints(int, int, const QPolygon &, int)
func (this *QPolygon) PutPoints_1_(index int, nPoints int, from QPolygon_ITF) {
	var convArg2 unsafe.Pointer
	if from != nil && from.QPolygon_PTR() != nil {
		convArg2 = from.QPolygon_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid,
	fromIndex := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiRKS_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, nPoints, convArg2, fromIndex)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool containsPoint(const QPoint &, Qt::FillRule) const
func (this *QPolygon) ContainsPoint(pt qtcore.QPoint_ITF, fillRule int) bool {
	var convArg0 unsafe.Pointer
	if pt != nil && pt.QPoint_PTR() != nil {
		convArg0 = pt.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon13containsPointERK6QPointN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, fillRule)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpolygon.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon united(const QPolygon &) const
func (this *QPolygon) United(r QPolygon_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygon_PTR() != nil {
		convArg0 = r.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon6unitedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon intersected(const QPolygon &) const
func (this *QPolygon) Intersected(r QPolygon_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygon_PTR() != nil {
		convArg0 = r.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon11intersectedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygon subtracted(const QPolygon &) const
func (this *QPolygon) Subtracted(r QPolygon_ITF) *QPolygon /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygon_PTR() != nil {
		convArg0 = r.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon10subtractedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPolygonFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPolygon)
	return rv2
}

// /usr/include/qt/QtGui/qpolygon.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool intersects(const QPolygon &) const
func (this *QPolygon) Intersects(r QPolygon_ITF) bool {
	var convArg0 unsafe.Pointer
	if r != nil && r.QPolygon_PTR() != nil {
		convArg0 = r.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPolygon10intersectsERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
}

//  keep block end
