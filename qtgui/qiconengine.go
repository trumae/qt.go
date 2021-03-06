package qtgui

// /usr/include/qt/QtGui/qiconengine.h
// #include <qiconengine.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
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

type QIconEngine struct {
	*qtrt.CObject
}
type QIconEngine_ITF interface {
	QIconEngine_PTR() *QIconEngine
}

func (ptr *QIconEngine) QIconEngine_PTR() *QIconEngine { return ptr }

func (this *QIconEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIconEngine) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQIconEngineFromPointer(cthis unsafe.Pointer) *QIconEngine {
	return &QIconEngine{&qtrt.CObject{cthis}}
}
func (*QIconEngine) NewFromPointer(cthis unsafe.Pointer) *QIconEngine {
	return NewQIconEngineFromPointer(cthis)
}

// /usr/include/qt/QtGui/qiconengine.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIconEngine()
func NewQIconEngine() *QIconEngine {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngineC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIconEngine)
	return gothis
}

// /usr/include/qt/QtGui/qiconengine.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIconEngine()
func DeleteQIconEngine(this *QIconEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qiconengine.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QRect &, QIcon::Mode, QIcon::State)
func (this *QIconEngine) Paint(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRect_ITF, mode int, state int) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg1 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine5paintEP8QPainterRK5QRectN5QIcon4ModeENS5_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize actualSize(const QSize &, QIcon::Mode, QIcon::State)
func (this *QIconEngine) ActualSize(size qtcore.QSize_ITF, mode int, state int) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine10actualSizeERK5QSizeN5QIcon4ModeENS3_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap pixmap(const QSize &, QIcon::Mode, QIcon::State)
func (this *QIconEngine) Pixmap(size qtcore.QSize_ITF, mode int, state int) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine6pixmapERK5QSizeN5QIcon4ModeENS3_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addPixmap(const QPixmap &, QIcon::Mode, QIcon::State)
func (this *QIconEngine) AddPixmap(pixmap QPixmap_ITF, mode int, state int) {
	var convArg0 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg0 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine9addPixmapERK7QPixmapN5QIcon4ModeENS3_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addFile(const QString &, const QSize &, QIcon::Mode, QIcon::State)
func (this *QIconEngine) AddFile(fileName string, size qtcore.QSize_ITF, mode int, state int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine7addFileERK7QStringRK5QSizeN5QIcon4ModeENS6_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qiconengine.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString key() const
func (this *QIconEngine) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QIconEngine3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qiconengine.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QIconEngine * clone() const
func (this *QIconEngine) Clone() *QIconEngine /*777 QIconEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QIconEngine5cloneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIconEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qiconengine.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool read(QDataStream &)
func (this *QIconEngine) Read(in qtcore.QDataStream_ITF) bool {
	var convArg0 unsafe.Pointer
	if in != nil && in.QDataStream_PTR() != nil {
		convArg0 = in.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine4readER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qiconengine.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool write(QDataStream &) const
func (this *QIconEngine) Write(out qtcore.QDataStream_ITF) bool {
	var convArg0 unsafe.Pointer
	if out != nil && out.QDataStream_PTR() != nil {
		convArg0 = out.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QIconEngine5writeER11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qiconengine.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString iconName() const
func (this *QIconEngine) IconName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QIconEngine8iconNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qiconengine.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const
func (this *QIconEngine) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QIconEngine6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qiconengine.h:82
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap scaledPixmap(const QSize &, QIcon::Mode, QIcon::State, qreal)
func (this *QIconEngine) ScaledPixmap(size qtcore.QSize_ITF, mode int, state int, scale float64) *QPixmap /*123*/ {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine12scaledPixmapERK5QSizeN5QIcon4ModeENS3_5StateEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state, scale)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qiconengine.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void virtual_hook(int, void *)
func (this *QIconEngine) Virtual_hook(id int, data unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QIconEngine12virtual_hookEiPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, data)
	qtrt.ErrPrint(err, rv)
}

type QIconEngine__IconEngineHook = int

const QIconEngine__AvailableSizesHook QIconEngine__IconEngineHook = 1
const QIconEngine__IconNameHook QIconEngine__IconEngineHook = 2
const QIconEngine__IsNullHook QIconEngine__IconEngineHook = 3
const QIconEngine__ScaledPixmapHook QIconEngine__IconEngineHook = 4

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
