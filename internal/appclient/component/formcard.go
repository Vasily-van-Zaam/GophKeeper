package component

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type formCardBank struct {
	metaData   string
	date       string
	csv        string
	clientName string
	number     string
	closed     func()
	err        func(msg string)
	saved      func(*core.BankCardForm)
}

// HandlerSave implements ComponentFormPassword.
func (f *formCardBank) HandlerSave(h func(form *core.BankCardForm)) ComponentFormCardBank {
	f.saved = h
	return f
}

// HandlerClose implements ComponentFormPassword.
func (f *formCardBank) HandlerClose(h func()) ComponentFormCardBank {
	f.closed = h
	return f
}

// HandlerError implements ComponentFormPassword.
func (f *formCardBank) HandlerError(h func(msg string)) ComponentFormCardBank {
	f.err = h
	return f
}

// Handler implements ComponentFormPassword.
func (f *formCardBank) EditForm(frm *core.BankCardForm) *tview.Form {
	f.csv = frm.CVC
	f.date = frm.Date
	f.metaData = frm.Metadata
	f.number = frm.Number
	f.clientName = frm.ClientName
	form := tview.NewForm().
		AddInputField("Metadata", f.metaData, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.metaData = text
			}).
		AddInputField("Number", f.number, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.number = text
			}).
		AddInputField("Client Name", f.date, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.clientName = text
			}).
		AddInputField("Date", f.date, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.date = text
			}).
		AddInputField("CSV", f.clientName, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.csv = text
			}).AddButton("❌ Back", func() {
		f.closed()
	}).AddButton("✅ Save", func() {
		frm.CVC = f.csv
		frm.ClientName = f.clientName
		frm.Metadata = f.metaData
		frm.Date = f.date
		frm.Number = f.number
		f.saved(frm)
	})

	if frm.Number != "" || frm.CVC != "" {
		form.AddButton("✂️ number", func() {
			_ = clipboard.WriteAll(f.number)
		}).AddButton("✂️ date", func() {
			_ = clipboard.WriteAll(f.date)
		}).AddButton("✂️ clien name", func() {
			_ = clipboard.WriteAll(f.clientName)
		}).AddButton("✂️ csv", func() {
			_ = clipboard.WriteAll(f.csv)
		}).SetButtonStyle(tcell.StyleDefault.Attributes(tcell.AttrBold))
	}

	return form
}

type ComponentFormCardBank interface {
	EditForm(form *core.BankCardForm) *tview.Form
	HandlerError(func(msg string)) ComponentFormCardBank
	HandlerClose(func()) ComponentFormCardBank
	HandlerSave(func(form *core.BankCardForm)) ComponentFormCardBank
}

func NewFormCardBanks() ComponentFormCardBank {
	return &formCardBank{}
}
