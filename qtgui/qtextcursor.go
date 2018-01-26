package qtgui

// /usr/include/qt/QtGui/qtextcursor.h
// #include <qtextcursor.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QTextCursor struct {
	*qtrt.CObject
}

func (this *QTextCursor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextCursor) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextCursorFromPointer(cthis unsafe.Pointer) *QTextCursor {
	return &QTextCursor{&qtrt.CObject{cthis}}
}
func (*QTextCursor) NewFromPointer(cthis unsafe.Pointer) *QTextCursor {
	return NewQTextCursorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextcursor.h:69
// index:0
// Public
// void QTextCursor()
func NewQTextCursor() *QTextCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:70
// index:1
// Public
// void QTextCursor(class QTextDocument *)
func NewQTextCursor_1(document *QTextDocument /*777 QTextDocument **/) *QTextCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = document.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:73
// index:2
// Public
// void QTextCursor(class QTextFrame *)
func NewQTextCursor_2(frame *QTextFrame /*777 QTextFrame **/) *QTextCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = frame.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2EP10QTextFrame", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:74
// index:3
// Public
// void QTextCursor(const class QTextBlock &)
func NewQTextCursor_3(block *QTextBlock) *QTextCursor {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorC2ERK10QTextBlock", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCursorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextcursor.h:80
// index:0
// Public
// void ~QTextCursor()
func DeleteQTextCursor(*QTextCursor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:82
// index:0
// Public inline
// void swap(class QTextCursor &)
func (this *QTextCursor) Swap(other *QTextCursor) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:84
// index:0
// Public
// bool isNull()
func (this *QTextCursor) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:91
// index:0
// Public
// void setPosition(int, enum QTextCursor::MoveMode)
func (this *QTextCursor) SetPosition(pos int, mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11setPositionEiNS_8MoveModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:92
// index:0
// Public
// int position()
func (this *QTextCursor) Position() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:93
// index:0
// Public
// int positionInBlock()
func (this *QTextCursor) PositionInBlock() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor15positionInBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:95
// index:0
// Public
// int anchor()
func (this *QTextCursor) Anchor() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor6anchorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:97
// index:0
// Public
// void insertText(const class QString &)
func (this *QTextCursor) InsertText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:98
// index:1
// Public
// void insertText(const class QString &, const class QTextCharFormat &)
func (this *QTextCursor) InsertText_1(text *qtcore.QString, format *QTextCharFormat) {
	var convArg0 = text.GetCthis()
	var convArg1 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertTextERK7QStringRK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:131
// index:0
// Public
// bool movePosition(enum QTextCursor::MoveOperation, enum QTextCursor::MoveMode, int)
func (this *QTextCursor) MovePosition(op int, arg1 int, n int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor12movePositionENS_13MoveOperationENS_8MoveModeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), op, arg1, n)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:133
// index:0
// Public
// bool visualNavigation()
func (this *QTextCursor) VisualNavigation() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor16visualNavigationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:134
// index:0
// Public
// void setVisualNavigation(_Bool)
func (this *QTextCursor) SetVisualNavigation(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor19setVisualNavigationEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:136
// index:0
// Public
// void setVerticalMovementX(int)
func (this *QTextCursor) SetVerticalMovementX(x int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor20setVerticalMovementXEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:137
// index:0
// Public
// int verticalMovementX()
func (this *QTextCursor) VerticalMovementX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor17verticalMovementXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:139
// index:0
// Public
// void setKeepPositionOnInsert(_Bool)
func (this *QTextCursor) SetKeepPositionOnInsert(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor23setKeepPositionOnInsertEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:140
// index:0
// Public
// bool keepPositionOnInsert()
func (this *QTextCursor) KeepPositionOnInsert() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor20keepPositionOnInsertEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:142
// index:0
// Public
// void deleteChar()
func (this *QTextCursor) DeleteChar() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10deleteCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:143
// index:0
// Public
// void deletePreviousChar()
func (this *QTextCursor) DeletePreviousChar() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor18deletePreviousCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:151
// index:0
// Public
// void select(enum QTextCursor::SelectionType)
func (this *QTextCursor) Select(selection int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor6selectENS_13SelectionTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:153
// index:0
// Public
// bool hasSelection()
func (this *QTextCursor) HasSelection() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12hasSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:154
// index:0
// Public
// bool hasComplexSelection()
func (this *QTextCursor) HasComplexSelection() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor19hasComplexSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:155
// index:0
// Public
// void removeSelectedText()
func (this *QTextCursor) RemoveSelectedText() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor18removeSelectedTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:156
// index:0
// Public
// void clearSelection()
func (this *QTextCursor) ClearSelection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14clearSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:157
// index:0
// Public
// int selectionStart()
func (this *QTextCursor) SelectionStart() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor14selectionStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:158
// index:0
// Public
// int selectionEnd()
func (this *QTextCursor) SelectionEnd() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12selectionEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:160
// index:0
// Public
// QString selectedText()
func (this *QTextCursor) SelectedText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12selectedTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:161
// index:0
// Public
// QTextDocumentFragment selection()
func (this *QTextCursor) Selection() *QTextDocumentFragment /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor9selectionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:162
// index:0
// Public
// void selectedTableCells(int *, int *, int *, int *)
func (this *QTextCursor) SelectedTableCells(firstRow unsafe.Pointer /*666*/, numRows unsafe.Pointer /*666*/, firstColumn unsafe.Pointer /*666*/, numColumns unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor18selectedTableCellsEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &firstRow, &numRows, &firstColumn, &numColumns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:164
// index:0
// Public
// QTextBlock block()
func (this *QTextCursor) Block() *QTextBlock /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor5blockEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:166
// index:0
// Public
// QTextCharFormat charFormat()
func (this *QTextCursor) CharFormat() *QTextCharFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor10charFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:167
// index:0
// Public
// void setCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) SetCharFormat(format *QTextCharFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor13setCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:168
// index:0
// Public
// void mergeCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) MergeCharFormat(modifier *QTextCharFormat) {
	var convArg0 = modifier.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor15mergeCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:170
// index:0
// Public
// QTextBlockFormat blockFormat()
func (this *QTextCursor) BlockFormat() *QTextBlockFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor11blockFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:171
// index:0
// Public
// void setBlockFormat(const class QTextBlockFormat &)
func (this *QTextCursor) SetBlockFormat(format *QTextBlockFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14setBlockFormatERK16QTextBlockFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:172
// index:0
// Public
// void mergeBlockFormat(const class QTextBlockFormat &)
func (this *QTextCursor) MergeBlockFormat(modifier *QTextBlockFormat) {
	var convArg0 = modifier.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor16mergeBlockFormatERK16QTextBlockFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:174
// index:0
// Public
// QTextCharFormat blockCharFormat()
func (this *QTextCursor) BlockCharFormat() *QTextCharFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor15blockCharFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:175
// index:0
// Public
// void setBlockCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) SetBlockCharFormat(format *QTextCharFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor18setBlockCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:176
// index:0
// Public
// void mergeBlockCharFormat(const class QTextCharFormat &)
func (this *QTextCursor) MergeBlockCharFormat(modifier *QTextCharFormat) {
	var convArg0 = modifier.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor20mergeBlockCharFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:178
// index:0
// Public
// bool atBlockStart()
func (this *QTextCursor) AtBlockStart() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12atBlockStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:179
// index:0
// Public
// bool atBlockEnd()
func (this *QTextCursor) AtBlockEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor10atBlockEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:180
// index:0
// Public
// bool atStart()
func (this *QTextCursor) AtStart() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor7atStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:181
// index:0
// Public
// bool atEnd()
func (this *QTextCursor) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:183
// index:0
// Public
// void insertBlock()
func (this *QTextCursor) InsertBlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:184
// index:1
// Public
// void insertBlock(const class QTextBlockFormat &)
func (this *QTextCursor) InsertBlock_1(format *QTextBlockFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockERK16QTextBlockFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:185
// index:2
// Public
// void insertBlock(const class QTextBlockFormat &, const class QTextCharFormat &)
func (this *QTextCursor) InsertBlock_2(format *QTextBlockFormat, charFormat *QTextCharFormat) {
	var convArg0 = format.GetCthis()
	var convArg1 = charFormat.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertBlockERK16QTextBlockFormatRK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:187
// index:0
// Public
// QTextList * insertList(const class QTextListFormat &)
func (this *QTextCursor) InsertList(format *QTextListFormat) *QTextList /*777 QTextList **/ {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertListERK15QTextListFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:188
// index:1
// Public
// QTextList * insertList(class QTextListFormat::Style)
func (this *QTextCursor) InsertList_1(style int) *QTextList /*777 QTextList **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertListEN15QTextListFormat5StyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:190
// index:0
// Public
// QTextList * createList(const class QTextListFormat &)
func (this *QTextCursor) CreateList(format *QTextListFormat) *QTextList /*777 QTextList **/ {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10createListERK15QTextListFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:191
// index:1
// Public
// QTextList * createList(class QTextListFormat::Style)
func (this *QTextCursor) CreateList_1(style int) *QTextList /*777 QTextList **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10createListEN15QTextListFormat5StyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), style)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:192
// index:0
// Public
// QTextList * currentList()
func (this *QTextCursor) CurrentList() *QTextList /*777 QTextList **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor11currentListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:194
// index:0
// Public
// QTextTable * insertTable(int, int, const class QTextTableFormat &)
func (this *QTextCursor) InsertTable(rows int, cols int, format *QTextTableFormat) *QTextTable /*777 QTextTable **/ {
	var convArg2 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertTableEiiRK16QTextTableFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rows, cols, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:195
// index:1
// Public
// QTextTable * insertTable(int, int)
func (this *QTextCursor) InsertTable_1(rows int, cols int) *QTextTable /*777 QTextTable **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertTableEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rows, cols)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:196
// index:0
// Public
// QTextTable * currentTable()
func (this *QTextCursor) CurrentTable() *QTextTable /*777 QTextTable **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12currentTableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:198
// index:0
// Public
// QTextFrame * insertFrame(const class QTextFrameFormat &)
func (this *QTextCursor) InsertFrame(format *QTextFrameFormat) *QTextFrame /*777 QTextFrame **/ {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertFrameERK16QTextFrameFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:199
// index:0
// Public
// QTextFrame * currentFrame()
func (this *QTextCursor) CurrentFrame() *QTextFrame /*777 QTextFrame **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12currentFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextcursor.h:201
// index:0
// Public
// void insertFragment(const class QTextDocumentFragment &)
func (this *QTextCursor) InsertFragment(fragment *QTextDocumentFragment) {
	var convArg0 = fragment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14insertFragmentERK21QTextDocumentFragment", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:204
// index:0
// Public
// void insertHtml(const class QString &)
func (this *QTextCursor) InsertHtml(html *qtcore.QString) {
	var convArg0 = html.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor10insertHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:207
// index:0
// Public
// void insertImage(const class QTextImageFormat &, class QTextFrameFormat::Position)
func (this *QTextCursor) InsertImage(format *QTextImageFormat, alignment int) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK16QTextImageFormatN16QTextFrameFormat8PositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:208
// index:1
// Public
// void insertImage(const class QTextImageFormat &)
func (this *QTextCursor) InsertImage_1(format *QTextImageFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK16QTextImageFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:209
// index:2
// Public
// void insertImage(const class QString &)
func (this *QTextCursor) InsertImage_2(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:210
// index:3
// Public
// void insertImage(const class QImage &, const class QString &)
func (this *QTextCursor) InsertImage_3(image *QImage, name *qtcore.QString) {
	var convArg0 = image.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor11insertImageERK6QImageRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:212
// index:0
// Public
// void beginEditBlock()
func (this *QTextCursor) BeginEditBlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor14beginEditBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:213
// index:0
// Public
// void joinPreviousEditBlock()
func (this *QTextCursor) JoinPreviousEditBlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor21joinPreviousEditBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:214
// index:0
// Public
// void endEditBlock()
func (this *QTextCursor) EndEditBlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextCursor12endEditBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:223
// index:0
// Public
// bool isCopyOf(const class QTextCursor &)
func (this *QTextCursor) IsCopyOf(other *QTextCursor) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor8isCopyOfERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextcursor.h:225
// index:0
// Public
// int blockNumber()
func (this *QTextCursor) BlockNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor11blockNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:226
// index:0
// Public
// int columnNumber()
func (this *QTextCursor) ColumnNumber() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor12columnNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextcursor.h:228
// index:0
// Public
// QTextDocument * document()
func (this *QTextCursor) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextCursor8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QTextCursor__MoveMode = int

const QTextCursor__MoveAnchor QTextCursor__MoveMode = 0
const QTextCursor__KeepAnchor QTextCursor__MoveMode = 1

type QTextCursor__MoveOperation = int

const QTextCursor__NoMove QTextCursor__MoveOperation = 0
const QTextCursor__Start QTextCursor__MoveOperation = 1
const QTextCursor__Up QTextCursor__MoveOperation = 2
const QTextCursor__StartOfLine QTextCursor__MoveOperation = 3
const QTextCursor__StartOfBlock QTextCursor__MoveOperation = 4
const QTextCursor__StartOfWord QTextCursor__MoveOperation = 5
const QTextCursor__PreviousBlock QTextCursor__MoveOperation = 6
const QTextCursor__PreviousCharacter QTextCursor__MoveOperation = 7
const QTextCursor__PreviousWord QTextCursor__MoveOperation = 8
const QTextCursor__Left QTextCursor__MoveOperation = 9
const QTextCursor__WordLeft QTextCursor__MoveOperation = 10
const QTextCursor__End QTextCursor__MoveOperation = 11
const QTextCursor__Down QTextCursor__MoveOperation = 12
const QTextCursor__EndOfLine QTextCursor__MoveOperation = 13
const QTextCursor__EndOfWord QTextCursor__MoveOperation = 14
const QTextCursor__EndOfBlock QTextCursor__MoveOperation = 15
const QTextCursor__NextBlock QTextCursor__MoveOperation = 16
const QTextCursor__NextCharacter QTextCursor__MoveOperation = 17
const QTextCursor__NextWord QTextCursor__MoveOperation = 18
const QTextCursor__Right QTextCursor__MoveOperation = 19
const QTextCursor__WordRight QTextCursor__MoveOperation = 20
const QTextCursor__NextCell QTextCursor__MoveOperation = 21
const QTextCursor__PreviousCell QTextCursor__MoveOperation = 22
const QTextCursor__NextRow QTextCursor__MoveOperation = 23
const QTextCursor__PreviousRow QTextCursor__MoveOperation = 24

type QTextCursor__SelectionType = int

const QTextCursor__WordUnderCursor QTextCursor__SelectionType = 0
const QTextCursor__LineUnderCursor QTextCursor__SelectionType = 1
const QTextCursor__BlockUnderCursor QTextCursor__SelectionType = 2
const QTextCursor__Document QTextCursor__SelectionType = 3

//  body block end
