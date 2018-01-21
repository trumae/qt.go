package qtgui

// /usr/include/qt/QtGui/qpagedpaintdevice.h
// #include <qpagedpaintdevice.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QPagedPaintDevice struct {
	*QPaintDevice
}

func (this *QPagedPaintDevice) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func NewQPagedPaintDeviceFromPointer(cthis unsafe.Pointer) *QPagedPaintDevice {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPagedPaintDevice{bcthis0}
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:58
// index:0
// Public
// void QPagedPaintDevice()
func NewQPagedPaintDevice() *QPagedPaintDevice {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDeviceC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPagedPaintDeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:59
// index:0
// Public virtual
// void ~QPagedPaintDevice()
func DeleteQPagedPaintDevice(*QPagedPaintDevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:61
// index:0
// Public pure virtual
// bool newPage()
func (this *QPagedPaintDevice) NewPage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice7newPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:219
// index:0
// Public
// bool setPageLayout(const class QPageLayout &)
func (this *QPagedPaintDevice) SetPageLayout(pageLayout *QPageLayout) bool {
	var convArg0 = pageLayout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageLayoutERK11QPageLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:220
// index:0
// Public
// bool setPageSize(const class QPageSize &)
func (this *QPagedPaintDevice) SetPageSize(pageSize *QPageSize) bool {
	var convArg0 = pageSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeERK9QPageSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:226
// index:1
// Public virtual
// void setPageSize(enum QPagedPaintDevice::PageSize)
func (this *QPagedPaintDevice) SetPageSize_1(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeENS_8PageSizeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:221
// index:0
// Public
// bool setPageOrientation(class QPageLayout::Orientation)
func (this *QPagedPaintDevice) SetPageOrientation(orientation int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice18setPageOrientationEN11QPageLayout11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:222
// index:0
// Public
// bool setPageMargins(const class QMarginsF &)
func (this *QPagedPaintDevice) SetPageMargins(margins *qtcore.QMarginsF) bool {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:223
// index:1
// Public
// bool setPageMargins(const class QMarginsF &, class QPageLayout::Unit)
func (this *QPagedPaintDevice) SetPageMargins_1(margins *qtcore.QMarginsF, units int) bool {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsFN11QPageLayout4UnitE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &units)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:224
// index:0
// Public
// QPageLayout pageLayout()
func (this *QPagedPaintDevice) PageLayout() *QPageLayout /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:227
// index:0
// Public
// QPagedPaintDevice::PageSize pageSize()
func (this *QPagedPaintDevice) PageSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice8pageSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:229
// index:0
// Public virtual
// void setPageSizeMM(const class QSizeF &)
func (this *QPagedPaintDevice) SetPageSizeMM(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageSizeMMERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:230
// index:0
// Public
// QSizeF pageSizeMM()
func (this *QPagedPaintDevice) PageSizeMM() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageSizeMMEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:241
// index:0
// Public
// QPagedPaintDevice::Margins margins()
func (this *QPagedPaintDevice) Margins() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice7marginsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:246
// index:0
// Protected
// QPageLayout devicePageLayout()
func (this *QPagedPaintDevice) DevicePageLayout() *QPageLayout /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice16devicePageLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:247
// index:1
// Protected
// QPageLayout & devicePageLayout()
func (this *QPagedPaintDevice) DevicePageLayout_1() *QPageLayout {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice16devicePageLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

//  body block end
