package qtnetwork

// /usr/include/qt/QtNetwork/qudpsocket.h
// #include <qudpsocket.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

type QUdpSocket struct {
	*QAbstractSocket
}
type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func (ptr *QUdpSocket) QUdpSocket_PTR() *QUdpSocket { return ptr }

func (this *QUdpSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSocket.GetCthis()
	}
}
func (this *QUdpSocket) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSocket = NewQAbstractSocketFromPointer(cthis)
}
func NewQUdpSocketFromPointer(cthis unsafe.Pointer) *QUdpSocket {
	bcthis0 := NewQAbstractSocketFromPointer(cthis)
	return &QUdpSocket{bcthis0}
}
func (*QUdpSocket) NewFromPointer(cthis unsafe.Pointer) *QUdpSocket {
	return NewQUdpSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QUdpSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qudpsocket.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUdpSocket(QObject *)
func NewQUdpSocket(parent qtcore.QObject_ITF /*777 QObject **/) *QUdpSocket {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUdpSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUdpSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qudpsocket.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUdpSocket(QObject *)
func NewQUdpSocket__() *QUdpSocket {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUdpSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUdpSocket")
	return gothis
}

// /usr/include/qt/QtNetwork/qudpsocket.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUdpSocket()
func DeleteQUdpSocket(this *QUdpSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool joinMulticastGroup(const QHostAddress &)
func (this *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket18joinMulticastGroupERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:65
// index:1
// Public Visibility=Default Availability=Available
// [1] bool joinMulticastGroup(const QHostAddress &, const QNetworkInterface &)
func (this *QUdpSocket) JoinMulticastGroup_1(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if iface != nil && iface.QNetworkInterface_PTR() != nil {
		convArg1 = iface.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket18joinMulticastGroupERK12QHostAddressRK17QNetworkInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool leaveMulticastGroup(const QHostAddress &)
func (this *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket19leaveMulticastGroupERK12QHostAddress", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:68
// index:1
// Public Visibility=Default Availability=Available
// [1] bool leaveMulticastGroup(const QHostAddress &, const QNetworkInterface &)
func (this *QUdpSocket) LeaveMulticastGroup_1(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	var convArg0 unsafe.Pointer
	if groupAddress != nil && groupAddress.QHostAddress_PTR() != nil {
		convArg0 = groupAddress.QHostAddress_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if iface != nil && iface.QNetworkInterface_PTR() != nil {
		convArg1 = iface.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket19leaveMulticastGroupERK12QHostAddressRK17QNetworkInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkInterface multicastInterface() const
func (this *QUdpSocket) MulticastInterface() *QNetworkInterface /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket18multicastInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkInterface)
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMulticastInterface(const QNetworkInterface &)
func (this *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QNetworkInterface_PTR() != nil {
		convArg0 = iface.QNetworkInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket21setMulticastInterfaceERK17QNetworkInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qudpsocket.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasPendingDatagrams() const
func (this *QUdpSocket) HasPendingDatagrams() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket19hasPendingDatagramsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qudpsocket.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 pendingDatagramSize() const
func (this *QUdpSocket) PendingDatagramSize() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUdpSocket19pendingDatagramSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkDatagram receiveDatagram(qint64)
func (this *QUdpSocket) ReceiveDatagram(maxSize int64) *QNetworkDatagram /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket15receiveDatagramEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QNetworkDatagram receiveDatagram(qint64)
func (this *QUdpSocket) ReceiveDatagram__() *QNetworkDatagram /*123*/ {
	// arg: 0, qint64=Typedef, qint64=Typedef, long long
	maxSize := int64(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket15receiveDatagramEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxSize)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQNetworkDatagramFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQNetworkDatagram)
	return rv2
}

// /usr/include/qt/QtNetwork/qudpsocket.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readDatagram(char *, qint64, QHostAddress *, quint16 *)
func (this *QUdpSocket) ReadDatagram(data string, maxlen int64, host QHostAddress_ITF /*777 QHostAddress **/, port unsafe.Pointer /*666*/) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg2 = host.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket12readDatagramEPcxP12QHostAddressPt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readDatagram(char *, qint64, QHostAddress *, quint16 *)
func (this *QUdpSocket) ReadDatagram__(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	// arg: 2, QHostAddress *=Pointer, QHostAddress=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, quint16 *=Pointer, quint16=Typedef, unsigned short
	port := unsafe.Pointer(nil)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket12readDatagramEPcxP12QHostAddressPt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 readDatagram(char *, qint64, QHostAddress *, quint16 *)
func (this *QUdpSocket) ReadDatagram__1(data string, maxlen int64, host QHostAddress_ITF /*777 QHostAddress **/) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg2 = host.QHostAddress_PTR().GetCthis()
	}
	// arg: 3, quint16 *=Pointer, quint16=Typedef, unsigned short
	port := unsafe.Pointer(nil)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket12readDatagramEPcxP12QHostAddressPt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const QNetworkDatagram &)
func (this *QUdpSocket) WriteDatagram(datagram QNetworkDatagram_ITF) int64 {
	var convArg0 unsafe.Pointer
	if datagram != nil && datagram.QNetworkDatagram_PTR() != nil {
		convArg0 = datagram.QNetworkDatagram_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramERK16QNetworkDatagram", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:81
// index:1
// Public Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const char *, qint64, const QHostAddress &, quint16)
func (this *QUdpSocket) WriteDatagram_1(data string, len_ int64, host QHostAddress_ITF, port uint16) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg2 = host.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramEPKcxRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_, convArg2, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtNetwork/qudpsocket.h:82
// index:2
// Public inline Visibility=Default Availability=Available
// [8] qint64 writeDatagram(const QByteArray &, const QHostAddress &, quint16)
func (this *QUdpSocket) WriteDatagram_2(datagram qtcore.QByteArray_ITF, host QHostAddress_ITF, port uint16) int64 {
	var convArg0 unsafe.Pointer
	if datagram != nil && datagram.QByteArray_PTR() != nil {
		convArg0 = datagram.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if host != nil && host.QHostAddress_PTR() != nil {
		convArg1 = host.QHostAddress_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUdpSocket13writeDatagramERK10QByteArrayRK12QHostAddresst", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, port)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
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
