package qtcore

// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
type QTextDecoder struct {
	*qtrt.CObject
}

func (this *QTextDecoder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQTextDecoderFromPointer(cthis unsafe.Pointer) *QTextDecoder {
	return &QTextDecoder{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qtextcodec.h:158
// index:0
// Public inline
// void QTextDecoder(const class QTextCodec *)
func NewQTextDecoder(codec *QTextCodec /*444 const QTextCodec **/) *QTextDecoder {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = codec.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoderC2EPK10QTextCodec", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDecoderFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:160
// index:0
// Public
// void ~QTextDecoder()
func DeleteQTextDecoder(*QTextDecoder) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoderD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:161
// index:0
// Public
// QString toUnicode(const char *, int)
func (this *QTextDecoder) ToUnicode(chars string, len int) *QString /*123*/ {
	var convArg0 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:162
// index:1
// Public
// QString toUnicode(const class QByteArray &)
func (this *QTextDecoder) ToUnicode_1(ba *QByteArray) *QString /*123*/ {
	var convArg0 = ba.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:163
// index:2
// Public
// void toUnicode(class QString *, const char *, int)
func (this *QTextDecoder) ToUnicode_2(target *QString /*444 QString **/, chars string, len int) {
	var convArg0 = target.GetCthis()
	var convArg1 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTextDecoder9toUnicodeEP7QStringPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextcodec.h:164
// index:0
// Public
// bool hasFailure()
func (this *QTextDecoder) HasFailure() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTextDecoder10hasFailureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
