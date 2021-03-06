package qtgui

// /usr/include/qt/QtGui/qimagewriter.h
// #include <qimagewriter.h>
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

type QImageWriter struct {
	*qtrt.CObject
}
type QImageWriter_ITF interface {
	QImageWriter_PTR() *QImageWriter
}

func (ptr *QImageWriter) QImageWriter_PTR() *QImageWriter { return ptr }

func (this *QImageWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageWriter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQImageWriterFromPointer(cthis unsafe.Pointer) *QImageWriter {
	return &QImageWriter{&qtrt.CObject{cthis}}
}
func (*QImageWriter) NewFromPointer(cthis unsafe.Pointer) *QImageWriter {
	return NewQImageWriterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimagewriter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter()
func NewQImageWriter() *QImageWriter {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter(QIODevice *, const QByteArray &)
func NewQImageWriter_1(device qtcore.QIODevice_ITF /*777 QIODevice **/, format qtcore.QByteArray_ITF) *QImageWriter {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter(const QString &, const QByteArray &)
func NewQImageWriter_2(fileName string, format qtcore.QByteArray_ITF) *QImageWriter {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg1 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2ERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QImageWriter(const QString &, const QByteArray &)
func NewQImageWriter_2_(fileName string) *QImageWriter {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg1 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterC2ERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQImageWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageWriter)
	return gothis
}

// /usr/include/qt/QtGui/qimagewriter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QImageWriter()
func DeleteQImageWriter(this *QImageWriter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimagewriter.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)
func (this *QImageWriter) SetFormat(format qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QByteArray_PTR() != nil {
		convArg0 = format.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format() const
func (this *QImageWriter) Format() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QImageWriter) SetDevice(device qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const
func (this *QImageWriter) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qimagewriter.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QImageWriter) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const
func (this *QImageWriter) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagewriter.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(int)
func (this *QImageWriter) SetQuality(quality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter10setQualityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), quality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int quality() const
func (this *QImageWriter) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagewriter.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompression(int)
func (this *QImageWriter) SetCompression(compression int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter14setCompressionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), compression)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int compression() const
func (this *QImageWriter) Compression() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter11compressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimagewriter.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGamma(float)
func (this *QImageWriter) SetGamma(gamma float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter8setGammaEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gamma)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] float gamma() const
func (this *QImageWriter) Gamma() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter5gammaEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qimagewriter.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSubType(const QByteArray &)
func (this *QImageWriter) SetSubType(type_ qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if type_ != nil && type_.QByteArray_PTR() != nil {
		convArg0 = type_.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter10setSubTypeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray subType() const
func (this *QImageWriter) SubType() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter7subTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimagewriter.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptimizedWrite(_Bool)
func (this *QImageWriter) SetOptimizedWrite(optimize bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter17setOptimizedWriteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), optimize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool optimizedWrite() const
func (this *QImageWriter) OptimizedWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter14optimizedWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProgressiveScanWrite(_Bool)
func (this *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter23setProgressiveScanWriteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progressive)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool progressiveScanWrite() const
func (this *QImageWriter) ProgressiveScanWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter20progressiveScanWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] QImageIOHandler::Transformations transformation() const
func (this *QImageWriter) Transformation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter14transformationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformation(QImageIOHandler::Transformations)
func (this *QImageWriter) SetTransformation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter17setTransformationE6QFlagsIN15QImageIOHandler14TransformationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)
func (this *QImageWriter) SetDescription(description string) {
	var tmpArg0 = qtcore.NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description() const
func (this *QImageWriter) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagewriter.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &, const QString &)
func (this *QImageWriter) SetText(key string, text string) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter7setTextERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canWrite() const
func (this *QImageWriter) CanWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter8canWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write(const QImage &)
func (this *QImageWriter) Write(image QImage_ITF) bool {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QImageWriter5writeERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qimagewriter.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] QImageWriter::ImageWriterError error() const
func (this *QImageWriter) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qimagewriter.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const
func (this *QImageWriter) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qimagewriter.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOption(QImageIOHandler::ImageOption) const
func (this *QImageWriter) SupportsOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QImageWriter14supportsOptionEN15QImageIOHandler11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QImageWriter__ImageWriterError = int

const QImageWriter__UnknownError QImageWriter__ImageWriterError = 0
const QImageWriter__DeviceError QImageWriter__ImageWriterError = 1
const QImageWriter__UnsupportedFormatError QImageWriter__ImageWriterError = 2
const QImageWriter__InvalidImageError QImageWriter__ImageWriterError = 3

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
