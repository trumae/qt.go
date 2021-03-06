package qtcore

// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QReadWriteLock struct {
	*qtrt.CObject
}
type QReadWriteLock_ITF interface {
	QReadWriteLock_PTR() *QReadWriteLock
}

func (ptr *QReadWriteLock) QReadWriteLock_PTR() *QReadWriteLock { return ptr }

func (this *QReadWriteLock) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QReadWriteLock) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQReadWriteLockFromPointer(cthis unsafe.Pointer) *QReadWriteLock {
	return &QReadWriteLock{&qtrt.CObject{cthis}}
}
func (*QReadWriteLock) NewFromPointer(cthis unsafe.Pointer) *QReadWriteLock {
	return NewQReadWriteLockFromPointer(cthis)
}

// /usr/include/qt/QtCore/qreadwritelock.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QReadWriteLock(enum QReadWriteLock::RecursionMode)
func NewQReadWriteLock(recursionMode int) *QReadWriteLock {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLockC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, recursionMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQReadWriteLock)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QReadWriteLock(enum QReadWriteLock::RecursionMode)
func NewQReadWriteLock__() *QReadWriteLock {
	// arg: 0, QReadWriteLock::RecursionMode=Enum, QReadWriteLock::RecursionMode=Enum,
	recursionMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLockC2ENS_13RecursionModeE", qtrt.FFI_TYPE_POINTER, recursionMode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQReadWriteLock)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QReadWriteLock()
func DeleteQReadWriteLock(this *QReadWriteLock) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLockD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qreadwritelock.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockForRead()
func (this *QReadWriteLock) LockForRead() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock11lockForReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLockForRead()
func (this *QReadWriteLock) TryLockForRead() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock14tryLockForReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:62
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryLockForRead(int)
func (this *QReadWriteLock) TryLockForRead_1(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock14tryLockForReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lockForWrite()
func (this *QReadWriteLock) LockForWrite() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock12lockForWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLockForWrite()
func (this *QReadWriteLock) TryLockForWrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock15tryLockForWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:66
// index:1
// Public Visibility=Default Availability=Available
// [1] bool tryLockForWrite(int)
func (this *QReadWriteLock) TryLockForWrite_1(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock15tryLockForWriteEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qreadwritelock.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QReadWriteLock) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QReadWriteLock6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QReadWriteLock__RecursionMode = int

const QReadWriteLock__NonRecursive QReadWriteLock__RecursionMode = 0
const QReadWriteLock__Recursive QReadWriteLock__RecursionMode = 1

type QReadWriteLock__StateForWaitCondition = int

const QReadWriteLock__LockedForRead QReadWriteLock__StateForWaitCondition = 0
const QReadWriteLock__LockedForWrite QReadWriteLock__StateForWaitCondition = 1
const QReadWriteLock__Unlocked QReadWriteLock__StateForWaitCondition = 2
const QReadWriteLock__RecursivelyLocked QReadWriteLock__StateForWaitCondition = 3

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
