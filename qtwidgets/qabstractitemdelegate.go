package qtwidgets

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h
// #include <qabstractitemdelegate.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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

type QAbstractItemDelegate struct {
	*qtcore.QObject
}
type QAbstractItemDelegate_ITF interface {
	qtcore.QObject_ITF
	QAbstractItemDelegate_PTR() *QAbstractItemDelegate
}

func (ptr *QAbstractItemDelegate) QAbstractItemDelegate_PTR() *QAbstractItemDelegate { return ptr }

func (this *QAbstractItemDelegate) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractItemDelegate) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractItemDelegateFromPointer(cthis unsafe.Pointer) *QAbstractItemDelegate {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractItemDelegate{bcthis0}
}
func (*QAbstractItemDelegate) NewFromPointer(cthis unsafe.Pointer) *QAbstractItemDelegate {
	return NewQAbstractItemDelegateFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QAbstractItemDelegate) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemDelegate(QObject *)
func NewQAbstractItemDelegate(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractItemDelegate {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemDelegate(QObject *)
func NewQAbstractItemDelegate__() *QAbstractItemDelegate {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemDelegate")
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractItemDelegate()
func DeleteQAbstractItemDelegate(this *QAbstractItemDelegate) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegateD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QAbstractItemDelegate) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate5paintEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize sizeHint(const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QAbstractItemDelegate) SizeHint(option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg0 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate8sizeHintERK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QAbstractItemDelegate) CreateEditor(parent QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void destroyEditor(QWidget *, const QModelIndex &) const
func (this *QAbstractItemDelegate) DestroyEditor(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setEditorData(QWidget *, const QModelIndex &) const
func (this *QAbstractItemDelegate) SetEditorData(editor QWidget_ITF /*777 QWidget **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg1 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setModelData(QWidget *, QAbstractItemModel *, const QModelIndex &) const
func (this *QAbstractItemDelegate) SetModelData(editor QWidget_ITF /*777 QWidget **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateEditorGeometry(QWidget *, const QStyleOptionViewItem &, const QModelIndex &) const
func (this *QAbstractItemDelegate) UpdateEditorGeometry(editor QWidget_ITF /*777 QWidget **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg1 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg2 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool editorEvent(QEvent *, QAbstractItemModel *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) EditorEvent(event qtcore.QEvent_ITF /*777 QEvent **/, model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg1 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg2 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg3 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString elidedText(const QFontMetrics &, int, Qt::TextElideMode, const QString &)
func (this *QAbstractItemDelegate) ElidedText(fontMetrics qtgui.QFontMetrics_ITF, width int, mode int, text string) string {
	var convArg0 unsafe.Pointer
	if fontMetrics != nil && fontMetrics.QFontMetrics_PTR() != nil {
		convArg0 = fontMetrics.QFontMetrics_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10elidedTextERK12QFontMetricsiN2Qt13TextElideModeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0, width, mode, convArg3)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QAbstractItemDelegate_ElidedText(fontMetrics qtgui.QFontMetrics_ITF, width int, mode int, text string) string {
	var nilthis *QAbstractItemDelegate
	rv := nilthis.ElidedText(fontMetrics, width, mode, text)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:109
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool helpEvent(QHelpEvent *, QAbstractItemView *, const QStyleOptionViewItem &, const QModelIndex &)
func (this *QAbstractItemDelegate) HelpEvent(event qtgui.QHelpEvent_ITF /*777 QHelpEvent **/, view QAbstractItemView_ITF /*777 QAbstractItemView **/, option QStyleOptionViewItem_ITF, index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHelpEvent_PTR() != nil {
		convArg0 = event.QHelpEvent_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if view != nil && view.QAbstractItemView_PTR() != nil {
		convArg1 = view.QAbstractItemView_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if option != nil && option.QStyleOptionViewItem_PTR() != nil {
		convArg2 = option.QStyleOptionViewItem_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg3 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void commitData(QWidget *)
func (this *QAbstractItemDelegate) CommitData(editor QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate10commitDataEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemDelegate) CloseEditor(editor QWidget_ITF /*777 QWidget **/, hint int) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11closeEditorEP7QWidgetNS_11EndEditHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeEditor(QWidget *, QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemDelegate) CloseEditor__(editor QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if editor != nil && editor.QWidget_PTR() != nil {
		convArg0 = editor.QWidget_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemDelegate::EndEditHint=Elaborated, QAbstractItemDelegate::EndEditHint=Enum,
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate11closeEditorEP7QWidgetNS_11EndEditHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemdelegate.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sizeHintChanged(const QModelIndex &)
func (this *QAbstractItemDelegate) SizeHintChanged(arg0 qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QModelIndex_PTR() != nil {
		convArg0 = arg0.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractItemDelegate15sizeHintChangedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QAbstractItemDelegate__EndEditHint = int

const QAbstractItemDelegate__NoHint QAbstractItemDelegate__EndEditHint = 0
const QAbstractItemDelegate__EditNextItem QAbstractItemDelegate__EndEditHint = 1
const QAbstractItemDelegate__EditPreviousItem QAbstractItemDelegate__EndEditHint = 2
const QAbstractItemDelegate__SubmitModelCache QAbstractItemDelegate__EndEditHint = 3
const QAbstractItemDelegate__RevertModelCache QAbstractItemDelegate__EndEditHint = 4

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
