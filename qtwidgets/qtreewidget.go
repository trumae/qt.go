package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 69
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
func (this *QTreeWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// QStringList mimeTypes()
func (this *QTreeWidget) InheritMimeTypes(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "mimeTypes", f)
}

// bool dropMimeData(class QTreeWidgetItem *, int, const class QMimeData *, Qt::DropAction)
func (this *QTreeWidget) InheritDropMimeData(f func(parent *QTreeWidgetItem /*777 QTreeWidgetItem **/, index int, data *qtcore.QMimeData /*777 const QMimeData **/, action int) bool) {
	qtrt.SetAllInheritCallback(this, "dropMimeData", f)
}

// Qt::DropActions supportedDropActions()
func (this *QTreeWidget) InheritSupportedDropActions(f func() int) {
	qtrt.SetAllInheritCallback(this, "supportedDropActions", f)
}

// QModelIndex indexFromItem(const class QTreeWidgetItem *, int)
func (this *QTreeWidget) InheritIndexFromItem(f func(item *QTreeWidgetItem /*777 const QTreeWidgetItem **/, column int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "indexFromItem", f)
}

// QTreeWidgetItem * itemFromIndex(const class QModelIndex &)
func (this *QTreeWidget) InheritItemFromIndex(f func(index *qtcore.QModelIndex) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "itemFromIndex", f)
}

// void dropEvent(class QDropEvent *)
func (this *QTreeWidget) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

type QTreeWidget struct {
	*QTreeView
}
type QTreeWidget_ITF interface {
	QTreeView_ITF
	QTreeWidget_PTR() *QTreeWidget
}

func (ptr *QTreeWidget) QTreeWidget_PTR() *QTreeWidget { return ptr }

func (this *QTreeWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTreeView.GetCthis()
	}
}
func (this *QTreeWidget) SetCthis(cthis unsafe.Pointer) {
	this.QTreeView = NewQTreeViewFromPointer(cthis)
}
func NewQTreeWidgetFromPointer(cthis unsafe.Pointer) *QTreeWidget {
	bcthis0 := NewQTreeViewFromPointer(cthis)
	return &QTreeWidget{bcthis0}
}
func (*QTreeWidget) NewFromPointer(cthis unsafe.Pointer) *QTreeWidget {
	return NewQTreeWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:257
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QTreeWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidget(QWidget *)
func NewQTreeWidget(parent QWidget_ITF /*777 QWidget **/) *QTreeWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTreeWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTreeWidget(QWidget *)
func NewQTreeWidget__() *QTreeWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTreeWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTreeWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:265
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTreeWidget()
func DeleteQTreeWidget(this *QTreeWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnCount() const
func (this *QTreeWidget) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)
func (this *QTreeWidget) SetColumnCount(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setColumnCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:270
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * invisibleRootItem() const
func (this *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget17invisibleRootItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * topLevelItem(int) const
func (this *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget12topLevelItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:272
// index:0
// Public Visibility=Default Availability=Available
// [4] int topLevelItemCount() const
func (this *QTreeWidget) TopLevelItemCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget17topLevelItemCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:273
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertTopLevelItem(int, QTreeWidgetItem *)
func (this *QTreeWidget) InsertTopLevelItem(index int, item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg1 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addTopLevelItem(QTreeWidgetItem *)
func (this *QTreeWidget) AddTopLevelItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:275
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * takeTopLevelItem(int)
func (this *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget16takeTopLevelItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:276
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfTopLevelItem(QTreeWidgetItem *) const
func (this *QTreeWidget) IndexOfTopLevelItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:281
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * headerItem() const
func (this *QTreeWidget) HeaderItem() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10headerItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderItem(QTreeWidgetItem *)
func (this *QTreeWidget) SetHeaderItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeaderLabels(const QStringList &)
func (this *QTreeWidget) SetHeaderLabels(labels qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg0 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15setHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:284
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeaderLabel(const QString &)
func (this *QTreeWidget) SetHeaderLabel(label string) {
	var tmpArg0 = qtcore.NewQString_5(label)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setHeaderLabelERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * currentItem() const
func (this *QTreeWidget) CurrentItem() *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget11currentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:287
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentColumn() const
func (this *QTreeWidget) CurrentColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13currentColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTreeWidgetItem *)
func (this *QTreeWidget) SetCurrentItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:289
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTreeWidgetItem *, int)
func (this *QTreeWidget) SetCurrentItem_1(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:290
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setCurrentItem(QTreeWidgetItem *, int, QItemSelectionModel::SelectionFlags)
func (this *QTreeWidget) SetCurrentItem_2(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int, command int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi6QFlagsIN19QItemSelectionModel13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:292
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemAt(const QPoint &) const
func (this *QTreeWidget) ItemAt(p qtcore.QPoint_ITF) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPoint_PTR() != nil {
		convArg0 = p.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:293
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemAt(int, int) const
func (this *QTreeWidget) ItemAt_1(x int, y int) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget6itemAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:294
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect visualItemRect(const QTreeWidgetItem *) const
func (this *QTreeWidget) VisualItemRect(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:296
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortColumn() const
func (this *QTreeWidget) SortColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10sortColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtreewidget.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sortItems(int, Qt::SortOrder)
func (this *QTreeWidget) SortItems(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget9sortItemsEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QTreeWidgetItem *, int)
func (this *QTreeWidget) EditItem(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editItem(QTreeWidgetItem *, int)
func (this *QTreeWidget) EditItem__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QTreeWidgetItem *, int)
func (this *QTreeWidget) OpenPersistentEditor(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void openPersistentEditor(QTreeWidgetItem *, int)
func (this *QTreeWidget) OpenPersistentEditor__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QTreeWidgetItem *, int)
func (this *QTreeWidget) ClosePersistentEditor(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closePersistentEditor(QTreeWidgetItem *, int)
func (this *QTreeWidget) ClosePersistentEditor__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(QTreeWidgetItem *, int) const
func (this *QTreeWidget) IsPersistentEditorOpen(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget22isPersistentEditorOpenEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentEditorOpen(QTreeWidgetItem *, int) const
func (this *QTreeWidget) IsPersistentEditorOpen__(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget22isPersistentEditorOpenEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:305
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * itemWidget(QTreeWidgetItem *, int) const
func (this *QTreeWidget) ItemWidget(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) *QWidget /*777 QWidget **/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:306
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemWidget(QTreeWidgetItem *, int, QWidget *)
func (this *QTreeWidget) SetItemWidget(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:307
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void removeItemWidget(QTreeWidgetItem *, int)
func (this *QTreeWidget) RemoveItemWidget(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:309
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemSelected(const QTreeWidgetItem *) const
func (this *QTreeWidget) IsItemSelected(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:310
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSelected(const QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemSelected(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, select_ bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, select_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:315
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemHidden(const QTreeWidgetItem *) const
func (this *QTreeWidget) IsItemHidden(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemHidden(const QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemHidden(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, hide bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hide)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:318
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isItemExpanded(const QTreeWidgetItem *) const
func (this *QTreeWidget) IsItemExpanded(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemExpanded(const QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetItemExpanded(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, expand bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, expand)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:321
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFirstItemColumnSpanned(const QTreeWidgetItem *) const
func (this *QTreeWidget) IsFirstItemColumnSpanned(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFirstItemColumnSpanned(const QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) SetFirstItemColumnSpanned(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, span bool) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, span)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:324
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemAbove(const QTreeWidgetItem *) const
func (this *QTreeWidget) ItemAbove(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:325
// index:0
// Public Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemBelow(const QTreeWidgetItem *) const
func (this *QTreeWidget) ItemBelow(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:327
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSelectionModel(QItemSelectionModel *)
func (this *QTreeWidget) SetSelectionModel(selectionModel qtcore.QItemSelectionModel_ITF /*777 QItemSelectionModel **/) {
	var convArg0 unsafe.Pointer
	if selectionModel != nil && selectionModel.QItemSelectionModel_PTR() != nil {
		convArg0 = selectionModel.QItemSelectionModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QTreeWidgetItem *, QAbstractItemView::ScrollHint)
func (this *QTreeWidget) ScrollToItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, hint int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12scrollToItemEPK15QTreeWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scrollToItem(const QTreeWidgetItem *, QAbstractItemView::ScrollHint)
func (this *QTreeWidget) ScrollToItem__(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, QAbstractItemView::ScrollHint=Elaborated, QAbstractItemView::ScrollHint=Enum,
	hint := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12scrollToItemEPK15QTreeWidgetItemN17QAbstractItemView10ScrollHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, hint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void expandItem(const QTreeWidgetItem *)
func (this *QTreeWidget) ExpandItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:333
// index:0
// Public Visibility=Default Availability=Available
// [-2] void collapseItem(const QTreeWidgetItem *)
func (this *QTreeWidget) CollapseItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:334
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QTreeWidget) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:337
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemPressed(QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemPressed(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemPressedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemClicked(QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemClicked(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemClickedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:339
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemDoubleClicked(QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemDoubleClicked(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget17itemDoubleClickedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemActivated(QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemActivated(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13itemActivatedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:341
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemEntered(QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemEntered(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemEnteredEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:342
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QTreeWidgetItem *, int)
func (this *QTreeWidget) ItemChanged(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget11itemChangedEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:343
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemExpanded(QTreeWidgetItem *)
func (this *QTreeWidget) ItemExpanded(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12itemExpandedEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:344
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemCollapsed(QTreeWidgetItem *)
func (this *QTreeWidget) ItemCollapsed(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget13itemCollapsedEP15QTreeWidgetItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:345
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentItemChanged(QTreeWidgetItem *, QTreeWidgetItem *)
func (this *QTreeWidget) CurrentItemChanged(current QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, previous QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QTreeWidgetItem_PTR() != nil {
		convArg0 = current.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QTreeWidgetItem_PTR() != nil {
		convArg1 = previous.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget18currentItemChangedEP15QTreeWidgetItemS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:346
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemSelectionChanged()
func (this *QTreeWidget) ItemSelectionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget20itemSelectionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:349
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QTreeWidget) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:350
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const
func (this *QTreeWidget) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:356
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(QTreeWidgetItem *, int, const QMimeData *, Qt::DropAction)
func (this *QTreeWidget) DropMimeData(parent QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, index int, data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QTreeWidgetItem_PTR() != nil {
		convArg0 = parent.QTreeWidgetItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg2 = data.QMimeData_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, index, convArg2, action)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtreewidget.h:358
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const
func (this *QTreeWidget) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:367
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(const QTreeWidgetItem *, int) const
func (this *QTreeWidget) IndexFromItem(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/, column int) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:367
// index:0
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(const QTreeWidgetItem *, int) const
func (this *QTreeWidget) IndexFromItem__(item QTreeWidgetItem_ITF /*777 const QTreeWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:368
// index:1
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QTreeWidgetItem *, int) const
func (this *QTreeWidget) IndexFromItem_1(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/, column int) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:368
// index:1
// Protected Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(QTreeWidgetItem *, int) const
func (this *QTreeWidget) IndexFromItem_1_(item QTreeWidgetItem_ITF /*777 QTreeWidgetItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QTreeWidgetItem_PTR() != nil {
		convArg0 = item.QTreeWidgetItem_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qtreewidget.h:369
// index:0
// Protected Visibility=Default Availability=Available
// [8] QTreeWidgetItem * itemFromIndex(const QModelIndex &) const
func (this *QTreeWidget) ItemFromIndex(index qtcore.QModelIndex_ITF) *QTreeWidgetItem /*777 QTreeWidgetItem **/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTreeWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qtreewidget.h:373
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QTreeWidget) DropEvent(event qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QDropEvent_PTR() != nil {
		convArg0 = event.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTreeWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
