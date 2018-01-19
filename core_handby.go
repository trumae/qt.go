package main

/*
#include <stdlib.h>
#include <stdint.h>

extern void ffi_call_ex(void*fn, int retype, uint64_t *rc, int argc, uint8_t* argtys, uint64_t* argvals);

*/
import "C"
import (
	"unsafe"
)

type QWidget5 struct {
	o unsafe.Pointer
}

func NewQWidget5() *QWidget5 {
	o := &QWidget5{}

	t := C.calloc(1, 128)
	o.o = t

	argtys := make([]byte, 20)
	argvals := make([]uint64, 20)

	argtys[0] = FFI_TYPE_POINTER
	argvals[0] = uint64(uintptr(t))

	var p *C.char
	argtys[1] = FFI_TYPE_POINTER
	argvals[1] = (uint64)(uintptr(unsafe.Pointer(p)))

	var cf C.int = 0
	argtys[2] = FFI_TYPE_INT
	argvals[2] = (uint64)(cf)
	var f int = 0

	// InvokeQtFunc5("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE",
	//	FFI_TYPE_VOID, 3, argtys, argvals)

	InvokeQtFunc6("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE",
		FFI_TYPE_VOID, t, unsafe.Pointer(p), f)
	// InvokeQtFunc("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", 0, nil)
	return o
}

func (w *QWidget5) Show() {

	argtys := make([]byte, 20)
	argvals := make([]uint64, 20)

	argtys[0] = FFI_TYPE_POINTER
	argvals[0] = uint64(uintptr(w.o))

	// InvokeQtFunc5("_ZN7QWidget4showEv", FFI_TYPE_VOID, 1, argtys, argvals)
	InvokeQtFunc6("_ZN7QWidget4showEv", FFI_TYPE_VOID, w.o)
}

func (w *QWidget5) SetObjectName() {
	// _ZN7QObject13setObjectNameERK7QString
}

type QApplication5 struct {
	o unsafe.Pointer
}

func NewQApplication50() *QApplication5 {
	o := &QApplication5{}
	t := C.calloc(1, 128)
	// args := C.CString("testprog")
	InvokeQtFunc("_ZN12QApplicationC2ERiPPci", 0,
		// InvokeQtFunc("_ZN3AbcC2ERiPPci", nil,
		// InvokeQtFunc("_ZN3AbcC2Eiiii", nil,
		[]byte{0, 0},
		uint64(uintptr(t)), 0)

	return o
}

func NewQApplication5() *QApplication5 {
	o := &QApplication5{}

	argtys := make([]byte, 20)
	argvals := make([]uint64, 20)

	t := C.calloc(1, 256)

	argtys[0] = FFI_TYPE_POINTER
	argvals[0] = uint64(uintptr(t))

	var argc C.int = 2
	argtys[1] = FFI_TYPE_POINTER
	argvals[1] = uint64(uintptr(unsafe.Pointer(&argc)))

	argtys[2] = FFI_TYPE_POINTER
	var twostr = make([]unsafe.Pointer, 2)
	twostr[0] = unsafe.Pointer(C.CString("testprog"))
	twostr[1] = unsafe.Pointer(C.CString("ikoklll"))
	argvals[2] = uint64(uintptr(unsafe.Pointer(&twostr[0])))

	var flag C.int = 0
	argtys[3] = FFI_TYPE_INT
	argvals[3] = uint64(uintptr(unsafe.Pointer(&flag)))

	// InvokeQtFunc5("_ZN12QApplicationC2ERiPPci", FFI_TYPE_VOID, argtys, argvals)
	InvokeQtFunc5("_ZN12QApplicationC2ERiPPci", FFI_TYPE_VOID, 4, argtys, argvals)
	o.o = unsafe.Pointer(t)
	return o
}

func (app *QApplication5) Exec5() {
	argtys := make([]byte, 20)
	argvals := make([]uint64, 20)

	argtys[0] = FFI_TYPE_POINTER
	argvals[0] = uint64(uintptr(app.o))

	InvokeQtFunc6("_ZN12QApplication4execEv", FFI_TYPE_INT, app.o)
	// retval, err := InvokeQtFunc5("_ZN12QApplication4execEv", FFI_TYPE_INT, argtys, argvals)
	// log.Println(err, retval)
}

////
type QString5 struct {
	cptr unsafe.Pointer
}

func NewQString5() *QString5 {
	o := &QString5{}

	return o
}
