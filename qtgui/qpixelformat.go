package qtgui

// /usr/include/qt/QtGui/qpixelformat.h
// #include <qpixelformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
type QPixelFormat struct {
	*qtrt.CObject
}

func (this *QPixelFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPixelFormat) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPixelFormatFromPointer(cthis unsafe.Pointer) *QPixelFormat {
	return &QPixelFormat{&qtrt.CObject{cthis}}
}
func (*QPixelFormat) NewFromPointer(cthis unsafe.Pointer) *QPixelFormat {
	return NewQPixelFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpixelformat.h:163
// index:0
// Public inline
// void QPixelFormat()
func NewQPixelFormat() *QPixelFormat {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixelFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:164
// index:1
// Public inline
// void QPixelFormat(enum QPixelFormat::ColorModel, uchar, uchar, uchar, uchar, uchar, uchar, enum QPixelFormat::AlphaUsage, enum QPixelFormat::AlphaPosition, enum QPixelFormat::AlphaPremultiplied, enum QPixelFormat::TypeInterpretation, enum QPixelFormat::ByteOrder, uchar)
func NewQPixelFormat_1(colorModel int, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage int, alphaPosition int, premultiplied int, typeInterpretation int, byteOrder int, subEnum byte) *QPixelFormat {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QPixelFormatC2ENS_10ColorModelEhhhhhhNS_10AlphaUsageENS_13AlphaPositionENS_18AlphaPremultipliedENS_18TypeInterpretationENS_9ByteOrderEh", ffiqt.FFI_TYPE_VOID, cthis, colorModel, firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder, subEnum)
	gopp.ErrPrint(err, rv)
	gothis := NewQPixelFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpixelformat.h:178
// index:0
// Public inline
// QPixelFormat::ColorModel colorModel()
func (this *QPixelFormat) ColorModel() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat10colorModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:179
// index:0
// Public inline
// uchar channelCount()
func (this *QPixelFormat) ChannelCount() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat12channelCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:186
// index:0
// Public inline
// uchar redSize()
func (this *QPixelFormat) RedSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat7redSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:187
// index:0
// Public inline
// uchar greenSize()
func (this *QPixelFormat) GreenSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9greenSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:188
// index:0
// Public inline
// uchar blueSize()
func (this *QPixelFormat) BlueSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat8blueSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:190
// index:0
// Public inline
// uchar cyanSize()
func (this *QPixelFormat) CyanSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat8cyanSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:191
// index:0
// Public inline
// uchar magentaSize()
func (this *QPixelFormat) MagentaSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat11magentaSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:192
// index:0
// Public inline
// uchar yellowSize()
func (this *QPixelFormat) YellowSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat10yellowSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:193
// index:0
// Public inline
// uchar blackSize()
func (this *QPixelFormat) BlackSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9blackSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:195
// index:0
// Public inline
// uchar hueSize()
func (this *QPixelFormat) HueSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat7hueSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:196
// index:0
// Public inline
// uchar saturationSize()
func (this *QPixelFormat) SaturationSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat14saturationSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:197
// index:0
// Public inline
// uchar lightnessSize()
func (this *QPixelFormat) LightnessSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat13lightnessSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:198
// index:0
// Public inline
// uchar brightnessSize()
func (this *QPixelFormat) BrightnessSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat14brightnessSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:200
// index:0
// Public inline
// uchar alphaSize()
func (this *QPixelFormat) AlphaSize() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9alphaSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:202
// index:0
// Public inline
// uchar bitsPerPixel()
func (this *QPixelFormat) BitsPerPixel() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat12bitsPerPixelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtGui/qpixelformat.h:209
// index:0
// Public inline
// QPixelFormat::AlphaUsage alphaUsage()
func (this *QPixelFormat) AlphaUsage() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat10alphaUsageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:210
// index:0
// Public inline
// QPixelFormat::AlphaPosition alphaPosition()
func (this *QPixelFormat) AlphaPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat13alphaPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:211
// index:0
// Public inline
// QPixelFormat::AlphaPremultiplied premultiplied()
func (this *QPixelFormat) Premultiplied() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat13premultipliedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:212
// index:0
// Public inline
// QPixelFormat::TypeInterpretation typeInterpretation()
func (this *QPixelFormat) TypeInterpretation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat18typeInterpretationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:213
// index:0
// Public inline
// QPixelFormat::ByteOrder byteOrder()
func (this *QPixelFormat) ByteOrder() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9byteOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:215
// index:0
// Public inline
// QPixelFormat::YUVLayout yuvLayout()
func (this *QPixelFormat) YuvLayout() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat9yuvLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpixelformat.h:216
// index:0
// Public inline
// uchar subEnum()
func (this *QPixelFormat) SubEnum() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QPixelFormat7subEnumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

type QPixelFormat__FieldWidth = int

const QPixelFormat__ModelFieldWidth QPixelFormat__FieldWidth = 4
const QPixelFormat__FirstFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__SecondFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__ThirdFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__FourthFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__FifthFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__AlphaFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__AlphaUsageFieldWidth QPixelFormat__FieldWidth = 1
const QPixelFormat__AlphaPositionFieldWidth QPixelFormat__FieldWidth = 1
const QPixelFormat__PremulFieldWidth QPixelFormat__FieldWidth = 1
const QPixelFormat__TypeInterpretationFieldWidth QPixelFormat__FieldWidth = 4
const QPixelFormat__ByteOrderFieldWidth QPixelFormat__FieldWidth = 2
const QPixelFormat__SubEnumFieldWidth QPixelFormat__FieldWidth = 6
const QPixelFormat__UnusedFieldWidth QPixelFormat__FieldWidth = 9
const QPixelFormat__TotalFieldWidthByWidths QPixelFormat__FieldWidth = 64

type QPixelFormat__Field = int

const QPixelFormat__ModelField QPixelFormat__Field = 0
const QPixelFormat__FirstField QPixelFormat__Field = 4
const QPixelFormat__SecondField QPixelFormat__Field = 10
const QPixelFormat__ThirdField QPixelFormat__Field = 16
const QPixelFormat__FourthField QPixelFormat__Field = 22
const QPixelFormat__FifthField QPixelFormat__Field = 28
const QPixelFormat__AlphaField QPixelFormat__Field = 34
const QPixelFormat__AlphaUsageField QPixelFormat__Field = 40
const QPixelFormat__AlphaPositionField QPixelFormat__Field = 41
const QPixelFormat__PremulField QPixelFormat__Field = 42
const QPixelFormat__TypeInterpretationField QPixelFormat__Field = 43
const QPixelFormat__ByteOrderField QPixelFormat__Field = 47
const QPixelFormat__SubEnumField QPixelFormat__Field = 49
const QPixelFormat__UnusedField QPixelFormat__Field = 55
const QPixelFormat__TotalFieldWidthByOffsets QPixelFormat__Field = 64

type QPixelFormat__ColorModel = int

const QPixelFormat__RGB QPixelFormat__ColorModel = 0
const QPixelFormat__BGR QPixelFormat__ColorModel = 1
const QPixelFormat__Indexed QPixelFormat__ColorModel = 2
const QPixelFormat__Grayscale QPixelFormat__ColorModel = 3
const QPixelFormat__CMYK QPixelFormat__ColorModel = 4
const QPixelFormat__HSL QPixelFormat__ColorModel = 5
const QPixelFormat__HSV QPixelFormat__ColorModel = 6
const QPixelFormat__YUV QPixelFormat__ColorModel = 7
const QPixelFormat__Alpha QPixelFormat__ColorModel = 8

type QPixelFormat__AlphaUsage = int

const QPixelFormat__UsesAlpha QPixelFormat__AlphaUsage = 0
const QPixelFormat__IgnoresAlpha QPixelFormat__AlphaUsage = 1

type QPixelFormat__AlphaPosition = int

const QPixelFormat__AtBeginning QPixelFormat__AlphaPosition = 0
const QPixelFormat__AtEnd QPixelFormat__AlphaPosition = 1

type QPixelFormat__AlphaPremultiplied = int

const QPixelFormat__NotPremultiplied QPixelFormat__AlphaPremultiplied = 0
const QPixelFormat__Premultiplied QPixelFormat__AlphaPremultiplied = 1

type QPixelFormat__TypeInterpretation = int

const QPixelFormat__UnsignedInteger QPixelFormat__TypeInterpretation = 0
const QPixelFormat__UnsignedShort QPixelFormat__TypeInterpretation = 1
const QPixelFormat__UnsignedByte QPixelFormat__TypeInterpretation = 2
const QPixelFormat__FloatingPoint QPixelFormat__TypeInterpretation = 3

type QPixelFormat__YUVLayout = int

const QPixelFormat__YUV444 QPixelFormat__YUVLayout = 0
const QPixelFormat__YUV422 QPixelFormat__YUVLayout = 1
const QPixelFormat__YUV411 QPixelFormat__YUVLayout = 2
const QPixelFormat__YUV420P QPixelFormat__YUVLayout = 3
const QPixelFormat__YUV420SP QPixelFormat__YUVLayout = 4
const QPixelFormat__YV12 QPixelFormat__YUVLayout = 5
const QPixelFormat__UYVY QPixelFormat__YUVLayout = 6
const QPixelFormat__YUYV QPixelFormat__YUVLayout = 7
const QPixelFormat__NV12 QPixelFormat__YUVLayout = 8
const QPixelFormat__NV21 QPixelFormat__YUVLayout = 9
const QPixelFormat__IMC1 QPixelFormat__YUVLayout = 10
const QPixelFormat__IMC2 QPixelFormat__YUVLayout = 11
const QPixelFormat__IMC3 QPixelFormat__YUVLayout = 12
const QPixelFormat__IMC4 QPixelFormat__YUVLayout = 13
const QPixelFormat__Y8 QPixelFormat__YUVLayout = 14
const QPixelFormat__Y16 QPixelFormat__YUVLayout = 15

type QPixelFormat__ByteOrder = int

const QPixelFormat__LittleEndian QPixelFormat__ByteOrder = 0
const QPixelFormat__BigEndian QPixelFormat__ByteOrder = 1
const QPixelFormat__CurrentSystemEndian QPixelFormat__ByteOrder = 2

//  body block end
