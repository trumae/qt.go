package qtwidgets

// /usr/include/qt/QtWidgets/qwizard.h
// #include <qwizard.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 51
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

// void setField(const class QString &, const class QVariant &)
func (this *QWizardPage) InheritSetField(f func(name string, value *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setField", f)
}

// QVariant field(const class QString &)
func (this *QWizardPage) InheritField(f func(name string) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "field", f)
}

// void registerField(const class QString &, class QWidget *, const char *, const char *)
func (this *QWizardPage) InheritRegisterField(f func(name string, widget *QWidget /*777 QWidget **/, property string, changedSignal string) /*void*/) {
	qtrt.SetAllInheritCallback(this, "registerField", f)
}

// QWizard * wizard()
func (this *QWizardPage) InheritWizard(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "wizard", f)
}

type QWizardPage struct {
	*QWidget
}
type QWizardPage_ITF interface {
	QWidget_ITF
	QWizardPage_PTR() *QWizardPage
}

func (ptr *QWizardPage) QWizardPage_PTR() *QWizardPage { return ptr }

func (this *QWizardPage) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QWizardPage) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQWizardPageFromPointer(cthis unsafe.Pointer) *QWizardPage {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QWizardPage{bcthis0}
}
func (*QWizardPage) NewFromPointer(cthis unsafe.Pointer) *QWizardPage {
	return NewQWizardPageFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwizard.h:213
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QWizardPage) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qwizard.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWizardPage(QWidget *)
func NewQWizardPage(parent QWidget_ITF /*777 QWidget **/) *QWizardPage {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPageC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWizardPage")
	return gothis
}

// /usr/include/qt/QtWidgets/qwizard.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWizardPage(QWidget *)
func NewQWizardPage__() *QWizardPage {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPageC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWizardPageFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWizardPage")
	return gothis
}

// /usr/include/qt/QtWidgets/qwizard.h:219
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWizardPage()
func DeleteQWizardPage(this *QWizardPage) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPageD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qwizard.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)
func (this *QWizardPage) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const
func (this *QWizardPage) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwizard.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSubTitle(const QString &)
func (this *QWizardPage) SetSubTitle(subTitle string) {
	var tmpArg0 = qtcore.NewQString_5(subTitle)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage11setSubTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QString subTitle() const
func (this *QWizardPage) SubTitle() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage8subTitleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwizard.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixmap(QWizard::WizardPixmap, const QPixmap &)
func (this *QWizardPage) SetPixmap(which int, pixmap qtgui.QPixmap_ITF) {
	var convArg1 unsafe.Pointer
	if pixmap != nil && pixmap.QPixmap_PTR() != nil {
		convArg1 = pixmap.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage9setPixmapEN7QWizard12WizardPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:226
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QWizard::WizardPixmap) const
func (this *QWizardPage) Pixmap(which int) *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage6pixmapEN7QWizard12WizardPixmapE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFinalPage(_Bool)
func (this *QWizardPage) SetFinalPage(finalPage bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage12setFinalPageEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), finalPage)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:228
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinalPage() const
func (this *QWizardPage) IsFinalPage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage11isFinalPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCommitPage(_Bool)
func (this *QWizardPage) SetCommitPage(commitPage bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage13setCommitPageEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), commitPage)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:230
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCommitPage() const
func (this *QWizardPage) IsCommitPage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage12isCommitPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButtonText(QWizard::WizardButton, const QString &)
func (this *QWizardPage) SetButtonText(which int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage13setButtonTextEN7QWizard12WizardButtonERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QString buttonText(QWizard::WizardButton) const
func (this *QWizardPage) ButtonText(which int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage10buttonTextEN7QWizard12WizardButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qwizard.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void initializePage()
func (this *QWizardPage) InitializePage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage14initializePageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:235
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void cleanupPage()
func (this *QWizardPage) CleanupPage() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage11cleanupPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:236
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool validatePage()
func (this *QWizardPage) ValidatePage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage12validatePageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:237
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isComplete() const
func (this *QWizardPage) IsComplete() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage10isCompleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwizard.h:238
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int nextId() const
func (this *QWizardPage) NextId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage6nextIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwizard.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void completeChanged()
func (this *QWizardPage) CompleteChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage15completeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:244
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setField(const QString &, const QVariant &)
func (this *QWizardPage) SetField(name string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage8setFieldERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:245
// index:0
// Protected Visibility=Default Availability=Available
// [16] QVariant field(const QString &) const
func (this *QWizardPage) Field(name string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage5fieldERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qwizard.h:246
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void registerField(const QString &, QWidget *, const char *, const char *)
func (this *QWizardPage) RegisterField(name string, widget QWidget_ITF /*777 QWidget **/, property string, changedSignal string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(property)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = qtrt.CString(changedSignal)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:246
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void registerField(const QString &, QWidget *, const char *, const char *)
func (this *QWizardPage) RegisterField__(name string, widget QWidget_ITF /*777 QWidget **/) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	// arg: 3, const char *=Pointer, =Invalid,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:246
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void registerField(const QString &, QWidget *, const char *, const char *)
func (this *QWizardPage) RegisterField__1(name string, widget QWidget_ITF /*777 QWidget **/, property string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	var convArg2 = qtrt.CString(property)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, const char *=Pointer, =Invalid,
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwizard.h:248
// index:0
// Protected Visibility=Default Availability=Available
// [8] QWizard * wizard() const
func (this *QWizardPage) Wizard() *QWizard /*777 QWizard **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWizardPage6wizardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWizardFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
