package qtcore

// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QTextEncoder struct {
	*qtrt.CObject
}

func (this *QTextEncoder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextEncoder) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextEncoderFromPointer(cthis unsafe.Pointer) *QTextEncoder {
	return &QTextEncoder{&qtrt.CObject{cthis}}
}
func (*QTextEncoder) NewFromPointer(cthis unsafe.Pointer) *QTextEncoder {
	return NewQTextEncoderFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextcodec.h:141
// index:0
// Public inline
// void QTextEncoder(const class QTextCodec *)
func NewQTextEncoder(codec *QTextCodec /*777 const QTextCodec **/) *QTextEncoder {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = codec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoderC2EPK10QTextCodec", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEncoderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:142
// index:1
// Public
// void QTextEncoder(const class QTextCodec *, class QTextCodec::ConversionFlags)
func NewQTextEncoder_1(codec *QTextCodec /*777 const QTextCodec **/, flags int) *QTextEncoder {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = codec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoderC2EPK10QTextCodec6QFlagsINS0_14ConversionFlagEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextEncoderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:143
// index:0
// Public
// void ~QTextEncoder()
func DeleteQTextEncoder(*QTextEncoder) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:145
// index:0
// Public
// QByteArray fromUnicode(const class QString &)
func (this *QTextEncoder) FromUnicode(str *QString) *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:147
// index:1
// Public
// QByteArray fromUnicode(class QStringView)
func (this *QTextEncoder) FromUnicode_1(str *QStringView /*123*/) *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeE11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:148
// index:2
// Public
// QByteArray fromUnicode(const class QChar *, int)
func (this *QTextEncoder) FromUnicode_2(uc *QChar /*777 const QChar **/, len int) *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = uc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextEncoder11fromUnicodeEPK5QChari", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:149
// index:0
// Public
// bool hasFailure()
func (this *QTextEncoder) HasFailure() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextEncoder10hasFailureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
