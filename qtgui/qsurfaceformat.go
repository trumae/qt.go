package qtgui

// /usr/include/qt/QtGui/qsurfaceformat.h
// #include <qsurfaceformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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

type QSurfaceFormat struct {
	*qtrt.CObject
}
type QSurfaceFormat_ITF interface {
	QSurfaceFormat_PTR() *QSurfaceFormat
}

func (ptr *QSurfaceFormat) QSurfaceFormat_PTR() *QSurfaceFormat { return ptr }

func (this *QSurfaceFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSurfaceFormat) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSurfaceFormatFromPointer(cthis unsafe.Pointer) *QSurfaceFormat {
	return &QSurfaceFormat{&qtrt.CObject{cthis}}
}
func (*QSurfaceFormat) NewFromPointer(cthis unsafe.Pointer) *QSurfaceFormat {
	return NewQSurfaceFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSurfaceFormat()
func NewQSurfaceFormat() *QSurfaceFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSurfaceFormat)
	return gothis
}

// /usr/include/qt/QtGui/qsurfaceformat.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSurfaceFormat(QSurfaceFormat::FormatOptions)
func NewQSurfaceFormat_1(options int) *QSurfaceFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormatC2E6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, options)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSurfaceFormat)
	return gothis
}

// /usr/include/qt/QtGui/qsurfaceformat.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QSurfaceFormat & operator=(const QSurfaceFormat &)
func (this *QSurfaceFormat) Operator_equal(other QSurfaceFormat_ITF) *QSurfaceFormat {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSurfaceFormat_PTR() != nil {
		convArg0 = other.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormataSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qsurfaceformat.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSurfaceFormat()
func DeleteQSurfaceFormat(this *QSurfaceFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDepthBufferSize(int)
func (this *QSurfaceFormat) SetDepthBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat18setDepthBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int depthBufferSize() const
func (this *QSurfaceFormat) DepthBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat15depthBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStencilBufferSize(int)
func (this *QSurfaceFormat) SetStencilBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat20setStencilBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int stencilBufferSize() const
func (this *QSurfaceFormat) StencilBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat17stencilBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRedBufferSize(int)
func (this *QSurfaceFormat) SetRedBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat16setRedBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int redBufferSize() const
func (this *QSurfaceFormat) RedBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat13redBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGreenBufferSize(int)
func (this *QSurfaceFormat) SetGreenBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat18setGreenBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] int greenBufferSize() const
func (this *QSurfaceFormat) GreenBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat15greenBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlueBufferSize(int)
func (this *QSurfaceFormat) SetBlueBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat17setBlueBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] int blueBufferSize() const
func (this *QSurfaceFormat) BlueBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat14blueBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlphaBufferSize(int)
func (this *QSurfaceFormat) SetAlphaBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat18setAlphaBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int alphaBufferSize() const
func (this *QSurfaceFormat) AlphaBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat15alphaBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSamples(int)
func (this *QSurfaceFormat) SetSamples(numSamples int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setSamplesEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), numSamples)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] int samples() const
func (this *QSurfaceFormat) Samples() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat7samplesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSwapBehavior(enum QSurfaceFormat::SwapBehavior)
func (this *QSurfaceFormat) SetSwapBehavior(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapBehaviorENS_12SwapBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::SwapBehavior swapBehavior() const
func (this *QSurfaceFormat) SwapBehavior() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapBehaviorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlpha() const
func (this *QSurfaceFormat) HasAlpha() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat8hasAlphaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProfile(enum QSurfaceFormat::OpenGLContextProfile)
func (this *QSurfaceFormat) SetProfile(profile int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setProfileENS_20OpenGLContextProfileE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), profile)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::OpenGLContextProfile profile() const
func (this *QSurfaceFormat) Profile() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat7profileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderableType(enum QSurfaceFormat::RenderableType)
func (this *QSurfaceFormat) SetRenderableType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat17setRenderableTypeENS_14RenderableTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::RenderableType renderableType() const
func (this *QSurfaceFormat) RenderableType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat14renderableTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMajorVersion(int)
func (this *QSurfaceFormat) SetMajorVersion(majorVersion int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMajorVersionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), majorVersion)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] int majorVersion() const
func (this *QSurfaceFormat) MajorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12majorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinorVersion(int)
func (this *QSurfaceFormat) SetMinorVersion(minorVersion int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMinorVersionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minorVersion)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] int minorVersion() const
func (this *QSurfaceFormat) MinorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12minorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVersion(int, int)
func (this *QSurfaceFormat) SetVersion(major int, minor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setVersionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), major, minor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:138
// index:0
// Public Visibility=Default Availability=Available
// [1] bool stereo() const
func (this *QSurfaceFormat) Stereo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat6stereoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStereo(_Bool)
func (this *QSurfaceFormat) SetStereo(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setStereoEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QSurfaceFormat::FormatOptions)
func (this *QSurfaceFormat) SetOption(opt int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionE6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:147
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setOption(enum QSurfaceFormat::FormatOption, _Bool)
func (this *QSurfaceFormat) SetOption_1(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionENS_12FormatOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:147
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setOption(enum QSurfaceFormat::FormatOption, _Bool)
func (this *QSurfaceFormat) SetOption_1_(option int) {
	// arg: 1, bool=Bool, =Invalid,
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionENS_12FormatOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QSurfaceFormat::FormatOptions) const
func (this *QSurfaceFormat) TestOption(opt int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat10testOptionE6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:148
// index:1
// Public Visibility=Default Availability=Available
// [1] bool testOption(enum QSurfaceFormat::FormatOption) const
func (this *QSurfaceFormat) TestOption_1(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat10testOptionENS_12FormatOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QSurfaceFormat::FormatOptions)
func (this *QSurfaceFormat) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setOptionsE6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:149
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::FormatOptions options() const
func (this *QSurfaceFormat) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:151
// index:0
// Public Visibility=Default Availability=Available
// [4] int swapInterval() const
func (this *QSurfaceFormat) SwapInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSwapInterval(int)
func (this *QSurfaceFormat) SetSwapInterval(interval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::ColorSpace colorSpace() const
func (this *QSurfaceFormat) ColorSpace() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat10colorSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorSpace(enum QSurfaceFormat::ColorSpace)
func (this *QSurfaceFormat) SetColorSpace(colorSpace int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat13setColorSpaceENS_10ColorSpaceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), colorSpace)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:157
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultFormat(const QSurfaceFormat &)
func (this *QSurfaceFormat) SetDefaultFormat(format QSurfaceFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QSurfaceFormat_PTR() != nil {
		convArg0 = format.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat16setDefaultFormatERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSurfaceFormat_SetDefaultFormat(format QSurfaceFormat_ITF) {
	var nilthis *QSurfaceFormat
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSurfaceFormat defaultFormat()
func (this *QSurfaceFormat) DefaultFormat() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat13defaultFormatEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}
func QSurfaceFormat_DefaultFormat() *QSurfaceFormat /*123*/ {
	var nilthis *QSurfaceFormat
	rv := nilthis.DefaultFormat()
	return rv
}

type QSurfaceFormat__FormatOption = int

const QSurfaceFormat__StereoBuffers QSurfaceFormat__FormatOption = 1
const QSurfaceFormat__DebugContext QSurfaceFormat__FormatOption = 2
const QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = 4
const QSurfaceFormat__ResetNotification QSurfaceFormat__FormatOption = 8

type QSurfaceFormat__SwapBehavior = int

const QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = 0
const QSurfaceFormat__SingleBuffer QSurfaceFormat__SwapBehavior = 1
const QSurfaceFormat__DoubleBuffer QSurfaceFormat__SwapBehavior = 2
const QSurfaceFormat__TripleBuffer QSurfaceFormat__SwapBehavior = 3

type QSurfaceFormat__RenderableType = int

const QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = 0
const QSurfaceFormat__OpenGL QSurfaceFormat__RenderableType = 1
const QSurfaceFormat__OpenGLES QSurfaceFormat__RenderableType = 2
const QSurfaceFormat__OpenVG QSurfaceFormat__RenderableType = 4

type QSurfaceFormat__OpenGLContextProfile = int

const QSurfaceFormat__NoProfile QSurfaceFormat__OpenGLContextProfile = 0
const QSurfaceFormat__CoreProfile QSurfaceFormat__OpenGLContextProfile = 1
const QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = 2

type QSurfaceFormat__ColorSpace = int

const QSurfaceFormat__DefaultColorSpace QSurfaceFormat__ColorSpace = 0
const QSurfaceFormat__sRGBColorSpace QSurfaceFormat__ColorSpace = 1

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
