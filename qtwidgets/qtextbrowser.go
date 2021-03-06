package qtwidgets

// /usr/include/qt/QtWidgets/qtextbrowser.h
// #include <qtextbrowser.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 77
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QTextBrowser) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QTextBrowser) InheritKeyPressEvent(f func(ev *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QTextBrowser) InheritMouseMoveEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QTextBrowser) InheritMousePressEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QTextBrowser) InheritMouseReleaseEvent(f func(ev *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QTextBrowser) InheritFocusOutEvent(f func(ev *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QTextBrowser) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QTextBrowser) InheritPaintEvent(f func(e *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

type QTextBrowser struct {
	*QTextEdit
}
type QTextBrowser_ITF interface {
	QTextEdit_ITF
	QTextBrowser_PTR() *QTextBrowser
}

func (ptr *QTextBrowser) QTextBrowser_PTR() *QTextBrowser { return ptr }

func (this *QTextBrowser) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextEdit.GetCthis()
	}
}
func (this *QTextBrowser) SetCthis(cthis unsafe.Pointer) {
	this.QTextEdit = NewQTextEditFromPointer(cthis)
}
func NewQTextBrowserFromPointer(cthis unsafe.Pointer) *QTextBrowser {
	bcthis0 := NewQTextEditFromPointer(cthis)
	return &QTextBrowser{bcthis0}
}
func (*QTextBrowser) NewFromPointer(cthis unsafe.Pointer) *QTextBrowser {
	return NewQTextBrowserFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QTextBrowser) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextBrowser(QWidget *)
func NewQTextBrowser(parent QWidget_ITF /*777 QWidget **/) *QTextBrowser {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowserC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBrowserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextBrowser")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextBrowser(QWidget *)
func NewQTextBrowser__() *QTextBrowser {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowserC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBrowserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextBrowser")
	return gothis
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextBrowser()
func DeleteQTextBrowser(this *QTextBrowser) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowserD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl source() const
func (this *QTextBrowser) Source() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList searchPaths() const
func (this *QTextBrowser) SearchPaths() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser11searchPathsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSearchPaths(const QStringList &)
func (this *QTextBrowser) SetSearchPaths(paths qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg0 = paths.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser14setSearchPathsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant loadResource(int, const QUrl &)
func (this *QTextBrowser) LoadResource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser12loadResourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBackwardAvailable() const
func (this *QTextBrowser) IsBackwardAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser19isBackwardAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isForwardAvailable() const
func (this *QTextBrowser) IsForwardAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser18isForwardAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearHistory()
func (this *QTextBrowser) ClearHistory() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser12clearHistoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString historyTitle(int) const
func (this *QTextBrowser) HistoryTitle(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser12historyTitleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl historyUrl(int) const
func (this *QTextBrowser) HistoryUrl(arg0 int) *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser10historyUrlEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int backwardHistoryCount() const
func (this *QTextBrowser) BackwardHistoryCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser20backwardHistoryCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int forwardHistoryCount() const
func (this *QTextBrowser) ForwardHistoryCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser19forwardHistoryCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openExternalLinks() const
func (this *QTextBrowser) OpenExternalLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser17openExternalLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenExternalLinks(_Bool)
func (this *QTextBrowser) SetOpenExternalLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser20setOpenExternalLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool openLinks() const
func (this *QTextBrowser) OpenLinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QTextBrowser9openLinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpenLinks(_Bool)
func (this *QTextBrowser) SetOpenLinks(open bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser12setOpenLinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), open)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSource(const QUrl &)
func (this *QTextBrowser) SetSource(name qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg0 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser9setSourceERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void backward()
func (this *QTextBrowser) Backward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser8backwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void forward()
func (this *QTextBrowser) Forward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser7forwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void home()
func (this *QTextBrowser) Home() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser4homeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reload()
func (this *QTextBrowser) Reload() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser6reloadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void backwardAvailable(_Bool)
func (this *QTextBrowser) BackwardAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser17backwardAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void forwardAvailable(_Bool)
func (this *QTextBrowser) ForwardAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser16forwardAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void historyChanged()
func (this *QTextBrowser) HistoryChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser14historyChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sourceChanged(const QUrl &)
func (this *QTextBrowser) SourceChanged(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13sourceChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QUrl &)
func (this *QTextBrowser) Highlighted(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QString &)
func (this *QTextBrowser) Highlighted_1(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser11highlightedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void anchorClicked(const QUrl &)
func (this *QTextBrowser) AnchorClicked(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13anchorClickedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QTextBrowser) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:108
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QTextBrowser) KeyPressEvent(ev qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QKeyEvent_PTR() != nil {
		convArg0 = ev.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QTextBrowser) MouseMoveEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QTextBrowser) MousePressEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QTextBrowser) MouseReleaseEvent(ev qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QMouseEvent_PTR() != nil {
		convArg0 = ev.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QTextBrowser) FocusOutEvent(ev qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if ev != nil && ev.QFocusEvent_PTR() != nil {
		convArg0 = ev.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QTextBrowser) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtextbrowser.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QTextBrowser) PaintEvent(e qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QPaintEvent_PTR() != nil {
		convArg0 = e.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QTextBrowser10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
