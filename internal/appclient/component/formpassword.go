package component

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type formPassword struct {
	metaData string
	password string
	login    string
	resource string
	closed   func()
	err      func(msg string)
	saved    func(*core.PasswordForm)
}

// HandlerSave implements ComponentFormPassword.
func (f *formPassword) HandlerSave(h func(form *core.PasswordForm)) ComponentFormPassword {
	f.saved = h
	return f
}

// HandlerClose implements ComponentFormPassword.
func (f *formPassword) HandlerClose(h func()) ComponentFormPassword {
	f.closed = h
	return f
}

// HandlerError implements ComponentFormPassword.
func (f *formPassword) HandlerError(h func(msg string)) ComponentFormPassword {
	f.err = h
	return f
}

// Handler implements ComponentFormPassword.
func (f *formPassword) EditForm(frm *core.PasswordForm) *tview.Form {
	f.login = frm.Login
	f.metaData = frm.MetaData
	f.resource = frm.Resource
	f.password = frm.Password
	form := tview.NewForm().
		AddInputField("Metadata", f.metaData, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.metaData = text
			}).
		AddInputField("Login", f.login, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.login = text
			}).
		AddInputField("Password", f.password, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.password = text
			}).
		AddInputField("Resource", f.resource, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.resource = text
			}).AddButton("❌ Back", func() {
		f.closed()
	}).AddButton("✅ Save", func() {
		frm.Login = f.login
		frm.Password = f.password
		frm.MetaData = f.metaData
		frm.Resource = f.resource
		f.saved(frm)
	})

	if frm.Password != "" || frm.Login != "" {
		form.AddButton("✂️ login", func() {
			_ = clipboard.WriteAll(f.login)
		}).AddButton("✂️ resouce", func() {
			_ = clipboard.WriteAll(f.resource)
		}).AddButton("✂️ password", func() {
			_ = clipboard.WriteAll(f.password)
		}).SetButtonStyle(tcell.StyleDefault.Attributes(tcell.AttrBold))
	}

	return form
}

type ComponentFormPassword interface {
	EditForm(form *core.PasswordForm) *tview.Form
	HandlerError(func(msg string)) ComponentFormPassword
	HandlerClose(func()) ComponentFormPassword
	HandlerSave(func(form *core.PasswordForm)) ComponentFormPassword
}

func NewFormPassword() ComponentFormPassword {
	return &formPassword{}
}
