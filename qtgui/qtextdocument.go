package qtgui

// /usr/include/qt/QtGui/qtextdocument.h
// #include <qtextdocument.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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

// QTextObject * createObject(const class QTextFormat &)
func (this *QTextDocument) InheritCreateObject(f func(f *QTextFormat) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createObject", f)
}

// QVariant loadResource(int, const class QUrl &)
func (this *QTextDocument) InheritLoadResource(f func(type_ int, name *qtcore.QUrl) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "loadResource", f)
}

type QTextDocument struct {
	*qtcore.QObject
}
type QTextDocument_ITF interface {
	qtcore.QObject_ITF
	QTextDocument_PTR() *QTextDocument
}

func (ptr *QTextDocument) QTextDocument_PTR() *QTextDocument { return ptr }

func (this *QTextDocument) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTextDocument) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQTextDocumentFromPointer(cthis unsafe.Pointer) *QTextDocument {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QTextDocument{bcthis0}
}
func (*QTextDocument) NewFromPointer(cthis unsafe.Pointer) *QTextDocument {
	return NewQTextDocumentFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextdocument.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QTextDocument) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(QObject *)
func NewQTextDocument(parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(QObject *)
func NewQTextDocument__() *QTextDocument {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(const QString &, QObject *)
func NewQTextDocument_1(text string, parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:120
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocument(const QString &, QObject *)
func NewQTextDocument_1_(text string) *QTextDocument {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextDocument")
	return gothis
}

// /usr/include/qt/QtGui/qtextdocument.h:121
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextDocument()
func DeleteQTextDocument(this *QTextDocument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextdocument.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * clone(QObject *) const
func (this *QTextDocument) Clone(parent qtcore.QObject_ITF /*777 QObject **/) *QTextDocument /*777 QTextDocument **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5cloneEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * clone(QObject *) const
func (this *QTextDocument) Clone__() *QTextDocument /*777 QTextDocument **/ {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5cloneEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QTextDocument) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()
func (this *QTextDocument) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUndoRedoEnabled(_Bool)
func (this *QTextDocument) SetUndoRedoEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument18setUndoRedoEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:129
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndoRedoEnabled() const
func (this *QTextDocument) IsUndoRedoEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17isUndoRedoEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:131
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUndoAvailable() const
func (this *QTextDocument) IsUndoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15isUndoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:132
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRedoAvailable() const
func (this *QTextDocument) IsRedoAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15isRedoAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int availableUndoSteps() const
func (this *QTextDocument) AvailableUndoSteps() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument18availableUndoStepsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] int availableRedoSteps() const
func (this *QTextDocument) AvailableRedoSteps() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument18availableRedoStepsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] int revision() const
func (this *QTextDocument) Revision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument8revisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentLayout(QAbstractTextDocumentLayout *)
func (this *QTextDocument) SetDocumentLayout(layout QAbstractTextDocumentLayout_ITF /*777 QAbstractTextDocumentLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QAbstractTextDocumentLayout_PTR() != nil {
		convArg0 = layout.QAbstractTextDocumentLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentLayoutEP27QAbstractTextDocumentLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:140
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractTextDocumentLayout * documentLayout() const
func (this *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout /*777 QAbstractTextDocumentLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument14documentLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractTextDocumentLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMetaInformation(enum QTextDocument::MetaInformation, const QString &)
func (this *QTextDocument) SetMetaInformation(info int, arg1 string) {
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument18setMetaInformationENS_15MetaInformationERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), info, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:147
// index:0
// Public Visibility=Default Availability=Available
// [8] QString metaInformation(enum QTextDocument::MetaInformation) const
func (this *QTextDocument) MetaInformation(info int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15metaInformationENS_15MetaInformationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), info)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml(const QByteArray &) const
func (this *QTextDocument) ToHtml(encoding qtcore.QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if encoding != nil && encoding.QByteArray_PTR() != nil {
		convArg0 = encoding.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument6toHtmlERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:150
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml(const QByteArray &) const
func (this *QTextDocument) ToHtml__() string {
	// arg: 0, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg0 = qtcore.NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument6toHtmlERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHtml(const QString &)
func (this *QTextDocument) SetHtml(html string) {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument7setHtmlERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toRawText() const
func (this *QTextDocument) ToRawText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9toRawTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toPlainText() const
func (this *QTextDocument) ToPlainText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11toPlainTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)
func (this *QTextDocument) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12setPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:158
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar characterAt(int) const
func (this *QTextDocument) CharacterAt(pos int) *qtcore.QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11characterAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQChar)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find(subString string, from int, options int) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find__(subString string) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid,
	from := int(0)
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:168
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find__1(subString string, from int) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:169
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, const QTextCursor &, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_1(subString string, cursor QTextCursor_ITF, options int) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:169
// index:1
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QString &, const QTextCursor &, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_1_(subString string, cursor QTextCursor_ITF) *QTextCursor /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(subString)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QStringRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:172
// index:2
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_2(expr qtcore.QRegExp_ITF, from int, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:172
// index:2
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_2_(expr qtcore.QRegExp_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(0)
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:172
// index:2
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_2_1(expr qtcore.QRegExp_ITF, from int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpi6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:173
// index:3
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, const QTextCursor &, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_3(expr qtcore.QRegExp_ITF, cursor QTextCursor_ITF, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:173
// index:3
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegExp &, const QTextCursor &, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_3_(expr qtcore.QRegExp_ITF, cursor QTextCursor_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegExp_PTR() != nil {
		convArg0 = expr.QRegExp_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK7QRegExpRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_4(expr qtcore.QRegularExpression_ITF, from int, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressioni6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_4_(expr qtcore.QRegularExpression_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	from := int(0)
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressioni6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:177
// index:4
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, int, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_4_1(expr qtcore.QRegularExpression_ITF, from int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressioni6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, from, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:178
// index:5
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, const QTextCursor &, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_5(expr qtcore.QRegularExpression_ITF, cursor QTextCursor_ITF, options int) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressionRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:178
// index:5
// Public Visibility=Default Availability=Available
// [8] QTextCursor find(const QRegularExpression &, const QTextCursor &, QTextDocument::FindFlags) const
func (this *QTextDocument) Find_5_(expr qtcore.QRegularExpression_ITF, cursor QTextCursor_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if expr != nil && expr.QRegularExpression_PTR() != nil {
		convArg0 = expr.QRegularExpression_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg1 = cursor.QTextCursor_PTR().GetCthis()
	}
	// arg: 2, QTextDocument::FindFlags=Typedef, QTextDocument::FindFlags=Typedef, QFlags<QTextDocument::FindFlag>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4findERK18QRegularExpressionRK11QTextCursor6QFlagsINS_8FindFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:181
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * frameAt(int) const
func (this *QTextDocument) FrameAt(pos int) *QTextFrame /*777 QTextFrame **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument7frameAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * rootFrame() const
func (this *QTextDocument) RootFrame() *QTextFrame /*777 QTextFrame **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9rootFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextObject * object(int) const
func (this *QTextDocument) Object(objectIndex int) *QTextObject /*777 QTextObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument6objectEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), objectIndex)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextObject * objectForFormat(const QTextFormat &) const
func (this *QTextDocument) ObjectForFormat(arg0 QTextFormat_ITF) *QTextObject /*777 QTextObject **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTextFormat_PTR() != nil {
		convArg0 = arg0.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument15objectForFormatERK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:187
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock findBlock(int) const
func (this *QTextDocument) FindBlock(pos int) *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9findBlockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:188
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock findBlockByNumber(int) const
func (this *QTextDocument) FindBlockByNumber(blockNumber int) *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17findBlockByNumberEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blockNumber)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:189
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock findBlockByLineNumber(int) const
func (this *QTextDocument) FindBlockByLineNumber(blockNumber int) *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument21findBlockByLineNumberEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), blockNumber)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:190
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock begin() const
func (this *QTextDocument) Begin() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:191
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock end() const
func (this *QTextDocument) End() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:193
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock firstBlock() const
func (this *QTextDocument) FirstBlock() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10firstBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:194
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlock lastBlock() const
func (this *QTextDocument) LastBlock() *QTextBlock /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9lastBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlock)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPageSize(const QSizeF &)
func (this *QTextDocument) SetPageSize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11setPageSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:197
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF pageSize() const
func (this *QTextDocument) PageSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument8pageSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultFont(const QFont &)
func (this *QTextDocument) SetDefaultFont(font QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14setDefaultFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:200
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont defaultFont() const
func (this *QTextDocument) DefaultFont() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11defaultFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:202
// index:0
// Public Visibility=Default Availability=Available
// [4] int pageCount() const
func (this *QTextDocument) PageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9pageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:204
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModified() const
func (this *QTextDocument) IsModified() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10isModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void print(QPagedPaintDevice *) const
func (this *QTextDocument) Print(printer QPagedPaintDevice_ITF /*777 QPagedPaintDevice **/) {
	var convArg0 unsafe.Pointer
	if printer != nil && printer.QPagedPaintDevice_PTR() != nil {
		convArg0 = printer.QPagedPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument5printEP17QPagedPaintDevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:216
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant resource(int, const QUrl &) const
func (this *QTextDocument) Resource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument8resourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addResource(int, const QUrl &, const QVariant &)
func (this *QTextDocument) AddResource(type_ int, name qtcore.QUrl_ITF, resource qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if resource != nil && resource.QVariant_PTR() != nil {
		convArg2 = resource.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11addResourceEiRK4QUrlRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markContentsDirty(int, int)
func (this *QTextDocument) MarkContentsDirty(from int, length int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17markContentsDirtyEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, length)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUseDesignMetrics(_Bool)
func (this *QTextDocument) SetUseDesignMetrics(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19setUseDesignMetricsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:224
// index:0
// Public Visibility=Default Availability=Available
// [1] bool useDesignMetrics() const
func (this *QTextDocument) UseDesignMetrics() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument16useDesignMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawContents(QPainter *, const QRectF &)
func (this *QTextDocument) DrawContents(painter QPainter_ITF /*777 QPainter **/, rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void drawContents(QPainter *, const QRectF &)
func (this *QTextDocument) DrawContents__(painter QPainter_ITF /*777 QPainter **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	// arg: 1, const QRectF &=LValueReference, QRectF=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12drawContentsEP8QPainterRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextWidth(qreal)
func (this *QTextDocument) SetTextWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12setTextWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal textWidth() const
func (this *QTextDocument) TextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9textWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:231
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal idealWidth() const
func (this *QTextDocument) IdealWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10idealWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal indentWidth() const
func (this *QTextDocument) IndentWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument11indentWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndentWidth(qreal)
func (this *QTextDocument) SetIndentWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14setIndentWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:236
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal documentMargin() const
func (this *QTextDocument) DocumentMargin() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument14documentMarginEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMargin(qreal)
func (this *QTextDocument) SetDocumentMargin(margin float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17setDocumentMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), margin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:239
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()
func (this *QTextDocument) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:240
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size() const
func (this *QTextDocument) Size() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:242
// index:0
// Public Visibility=Default Availability=Available
// [4] int blockCount() const
func (this *QTextDocument) BlockCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument10blockCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineCount() const
func (this *QTextDocument) LineCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument9lineCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] int characterCount() const
func (this *QTextDocument) CharacterCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument14characterCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultStyleSheet(const QString &)
func (this *QTextDocument) SetDefaultStyleSheet(sheet string) {
	var tmpArg0 = qtcore.NewQString_5(sheet)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultStyleSheetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:248
// index:0
// Public Visibility=Default Availability=Available
// [8] QString defaultStyleSheet() const
func (this *QTextDocument) DefaultStyleSheet() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17defaultStyleSheetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo(QTextCursor *)
func (this *QTextDocument) Undo(cursor QTextCursor_ITF /*777 QTextCursor **/) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4undoEP11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:286
// index:1
// Public Visibility=Default Availability=Available
// [-2] void undo()
func (this *QTextDocument) Undo_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:252
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo(QTextCursor *)
func (this *QTextDocument) Redo(cursor QTextCursor_ITF /*777 QTextCursor **/) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4redoEP11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:287
// index:1
// Public Visibility=Default Availability=Available
// [-2] void redo()
func (this *QTextDocument) Redo_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearUndoRedoStacks(enum QTextDocument::Stacks)
func (this *QTextDocument) ClearUndoRedoStacks(historyToClear int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19clearUndoRedoStacksENS_6StacksE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), historyToClear)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearUndoRedoStacks(enum QTextDocument::Stacks)
func (this *QTextDocument) ClearUndoRedoStacks__() {
	// arg: 0, QTextDocument::Stacks=Enum, QTextDocument::Stacks=Enum,
	historyToClear := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19clearUndoRedoStacksENS_6StacksE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), historyToClear)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:261
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumBlockCount() const
func (this *QTextDocument) MaximumBlockCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17maximumBlockCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextdocument.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumBlockCount(int)
func (this *QTextDocument) SetMaximumBlockCount(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument20setMaximumBlockCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:264
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextOption defaultTextOption() const
func (this *QTextDocument) DefaultTextOption() *QTextOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument17defaultTextOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTextOption(const QTextOption &)
func (this *QTextDocument) SetDefaultTextOption(option QTextOption_ITF) {
	var convArg0 unsafe.Pointer
	if option != nil && option.QTextOption_PTR() != nil {
		convArg0 = option.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument20setDefaultTextOptionERK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:267
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl baseUrl() const
func (this *QTextDocument) BaseUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument7baseUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseUrl(const QUrl &)
func (this *QTextDocument) SetBaseUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument10setBaseUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:270
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CursorMoveStyle defaultCursorMoveStyle() const
func (this *QTextDocument) DefaultCursorMoveStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QTextDocument22defaultCursorMoveStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultCursorMoveStyle(Qt::CursorMoveStyle)
func (this *QTextDocument) SetDefaultCursorMoveStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument25setDefaultCursorMoveStyleEN2Qt15CursorMoveStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentsChange(int, int, int)
func (this *QTextDocument) ContentsChange(from int, charsRemoved int, charsAdded int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14contentsChangeEiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), from, charsRemoved, charsAdded)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentsChanged()
func (this *QTextDocument) ContentsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument15contentsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoAvailable(_Bool)
func (this *QTextDocument) UndoAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument13undoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:277
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoAvailable(_Bool)
func (this *QTextDocument) RedoAvailable(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument13redoAvailableEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoCommandAdded()
func (this *QTextDocument) UndoCommandAdded() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument16undoCommandAddedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:279
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modificationChanged(_Bool)
func (this *QTextDocument) ModificationChanged(m bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument19modificationChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:280
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorPositionChanged(const QTextCursor &)
func (this *QTextDocument) CursorPositionChanged(cursor QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument21cursorPositionChangedERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void blockCountChanged(int)
func (this *QTextDocument) BlockCountChanged(newBlockCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument17blockCountChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newBlockCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void baseUrlChanged(const QUrl &)
func (this *QTextDocument) BaseUrlChanged(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14baseUrlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void documentLayoutChanged()
func (this *QTextDocument) DocumentLayoutChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument21documentLayoutChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendUndoItem(QAbstractUndoItem *)
func (this *QTextDocument) AppendUndoItem(arg0 QAbstractUndoItem_ITF /*777 QAbstractUndoItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAbstractUndoItem_PTR() != nil {
		convArg0 = arg0.QAbstractUndoItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument14appendUndoItemEP17QAbstractUndoItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModified(_Bool)
func (this *QTextDocument) SetModified(m bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11setModifiedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModified(_Bool)
func (this *QTextDocument) SetModified__() {
	// arg: 0, bool=Bool, =Invalid,
	m := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument11setModifiedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:292
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QTextObject * createObject(const QTextFormat &)
func (this *QTextDocument) CreateObject(f QTextFormat_ITF) *QTextObject /*777 QTextObject **/ {
	var convArg0 unsafe.Pointer
	if f != nil && f.QTextFormat_PTR() != nil {
		convArg0 = f.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12createObjectERK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextdocument.h:293
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant loadResource(int, const QUrl &)
func (this *QTextDocument) LoadResource(type_ int, name qtcore.QUrl_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if name != nil && name.QUrl_PTR() != nil {
		convArg1 = name.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QTextDocument12loadResourceEiRK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QTextDocument__MetaInformation = int

const QTextDocument__DocumentTitle QTextDocument__MetaInformation = 0
const QTextDocument__DocumentUrl QTextDocument__MetaInformation = 1

type QTextDocument__FindFlag = int

const QTextDocument__FindBackward QTextDocument__FindFlag = 1
const QTextDocument__FindCaseSensitively QTextDocument__FindFlag = 2
const QTextDocument__FindWholeWords QTextDocument__FindFlag = 4

type QTextDocument__ResourceType = int

const QTextDocument__HtmlResource QTextDocument__ResourceType = 1
const QTextDocument__ImageResource QTextDocument__ResourceType = 2
const QTextDocument__StyleSheetResource QTextDocument__ResourceType = 3
const QTextDocument__UserResource QTextDocument__ResourceType = 100

type QTextDocument__Stacks = int

const QTextDocument__UndoStack QTextDocument__Stacks = 1
const QTextDocument__RedoStack QTextDocument__Stacks = 2
const QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = 3

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
