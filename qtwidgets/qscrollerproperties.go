package qtwidgets

// /usr/include/qt/QtWidgets/qscrollerproperties.h
// #include <qscrollerproperties.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QScrollerProperties struct {
	*qtrt.CObject
}

func (this *QScrollerProperties) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QScrollerProperties) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQScrollerPropertiesFromPointer(cthis unsafe.Pointer) *QScrollerProperties {
	return &QScrollerProperties{&qtrt.CObject{cthis}}
}
func (*QScrollerProperties) NewFromPointer(cthis unsafe.Pointer) *QScrollerProperties {
	return NewQScrollerPropertiesFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:60
// index:0
// Public
// void QScrollerProperties()
func NewQScrollerProperties() *QScrollerProperties {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerPropertiesC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollerPropertiesFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:63
// index:0
// Public virtual
// void ~QScrollerProperties()
func DeleteQScrollerProperties(*QScrollerProperties) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerPropertiesD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:68
// index:0
// Public static
// void setDefaultScrollerProperties(const class QScrollerProperties &)
func (this *QScrollerProperties) SetDefaultScrollerProperties(sp *QScrollerProperties) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerProperties28setDefaultScrollerPropertiesERKS_", ffiqt.FFI_TYPE_POINTER, sp)
	gopp.ErrPrint(err, rv)
}
func QScrollerProperties_SetDefaultScrollerProperties(sp *QScrollerProperties) {
	var nilthis *QScrollerProperties
	nilthis.SetDefaultScrollerProperties(sp)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:69
// index:0
// Public static
// void unsetDefaultScrollerProperties()
func (this *QScrollerProperties) UnsetDefaultScrollerProperties() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerProperties30unsetDefaultScrollerPropertiesEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QScrollerProperties_UnsetDefaultScrollerProperties() {
	var nilthis *QScrollerProperties
	nilthis.UnsetDefaultScrollerProperties()
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:117
// index:0
// Public
// QVariant scrollMetric(enum QScrollerProperties::ScrollMetric)
func (this *QScrollerProperties) ScrollMetric(metric int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QScrollerProperties12scrollMetricENS_12ScrollMetricE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), metric)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:118
// index:0
// Public
// void setScrollMetric(enum QScrollerProperties::ScrollMetric, const class QVariant &)
func (this *QScrollerProperties) SetScrollMetric(metric int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QScrollerProperties15setScrollMetricENS_12ScrollMetricERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), metric, convArg1)
	gopp.ErrPrint(err, rv)
}

type QScrollerProperties__OvershootPolicy = int

const QScrollerProperties__OvershootWhenScrollable QScrollerProperties__OvershootPolicy = 0
const QScrollerProperties__OvershootAlwaysOff QScrollerProperties__OvershootPolicy = 1
const QScrollerProperties__OvershootAlwaysOn QScrollerProperties__OvershootPolicy = 2

type QScrollerProperties__FrameRates = int

const QScrollerProperties__Standard QScrollerProperties__FrameRates = 0
const QScrollerProperties__Fps60 QScrollerProperties__FrameRates = 1
const QScrollerProperties__Fps30 QScrollerProperties__FrameRates = 2
const QScrollerProperties__Fps20 QScrollerProperties__FrameRates = 3

type QScrollerProperties__ScrollMetric = int

const QScrollerProperties__MousePressEventDelay QScrollerProperties__ScrollMetric = 0
const QScrollerProperties__DragStartDistance QScrollerProperties__ScrollMetric = 1
const QScrollerProperties__DragVelocitySmoothingFactor QScrollerProperties__ScrollMetric = 2
const QScrollerProperties__AxisLockThreshold QScrollerProperties__ScrollMetric = 3
const QScrollerProperties__ScrollingCurve QScrollerProperties__ScrollMetric = 4
const QScrollerProperties__DecelerationFactor QScrollerProperties__ScrollMetric = 5
const QScrollerProperties__MinimumVelocity QScrollerProperties__ScrollMetric = 6
const QScrollerProperties__MaximumVelocity QScrollerProperties__ScrollMetric = 7
const QScrollerProperties__MaximumClickThroughVelocity QScrollerProperties__ScrollMetric = 8
const QScrollerProperties__AcceleratingFlickMaximumTime QScrollerProperties__ScrollMetric = 9
const QScrollerProperties__AcceleratingFlickSpeedupFactor QScrollerProperties__ScrollMetric = 10
const QScrollerProperties__SnapPositionRatio QScrollerProperties__ScrollMetric = 11
const QScrollerProperties__SnapTime QScrollerProperties__ScrollMetric = 12
const QScrollerProperties__OvershootDragResistanceFactor QScrollerProperties__ScrollMetric = 13
const QScrollerProperties__OvershootDragDistanceFactor QScrollerProperties__ScrollMetric = 14
const QScrollerProperties__OvershootScrollDistanceFactor QScrollerProperties__ScrollMetric = 15
const QScrollerProperties__OvershootScrollTime QScrollerProperties__ScrollMetric = 16
const QScrollerProperties__HorizontalOvershootPolicy QScrollerProperties__ScrollMetric = 17
const QScrollerProperties__VerticalOvershootPolicy QScrollerProperties__ScrollMetric = 18
const QScrollerProperties__FrameRate QScrollerProperties__ScrollMetric = 19
const QScrollerProperties__ScrollMetricCount QScrollerProperties__ScrollMetric = 20

//  body block end