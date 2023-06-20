package component

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/rivo/tview"
)

func (a *accessForm) changeEmail(text string) {
	a.email = text
}

func (a *accessForm) changeCode(text string) {
	a.code = text
}
func (a *accessForm) changePassword(text string) {
	a.masterPassword = text
}
func (a *accessForm) repeatPassword(text string) {
	a.repeatPasword = text
}

type AccessForms interface {
	NewFormGetAccess(pageName, buttonNameBack string, login func(form *core.AccessForm), back func())
	NewFormConfirmAccess(pageName, buttonNameBack string, login func(form *core.AccessForm), back func())
	NewFormCreateMasterPassword(pageName, buttonNameBack string, create func(password, repeat string), back func())
	NewFormLogin(pageName, buttonNameBack string, login func(password string), reset func(), back func())
}
type accessForm struct {
	email          string
	code           string
	masterPassword string
	repeatPasword  string
	pages          *tview.Pages
}

const fieldWidth = 50

// NewFormLogin implements AccessForms.
func (a *accessForm) NewFormLogin(
	pageName, buttonNameBack string, login func(password string), reset, closed func()) {
	form := tview.NewForm().SetFocus(1)
	a.masterPassword = "Psw123!@#"
	form.AddInputField("password", a.masterPassword, fieldWidth, nil, a.changePassword)
	form.AddButton(buttonNameBack, func() {
		a.pages.RemovePage(pageName)
		a.masterPassword = ""
		a.repeatPasword = ""
		closed()
	})
	form.AddButton("reset", func() {
		// a.pages.RemovePage(pageName)
		a.masterPassword = ""
		a.repeatPasword = ""
		reset()
	})
	form.AddButton("next", func() {
		login(a.masterPassword)
		form.Clear(false)
		a.email = ""
		a.masterPassword = ""
		a.repeatPasword = ""
		form.AddInputField("password", a.masterPassword, fieldWidth, nil, a.changePassword)
	})

	frame := NewFrame(form, "Enter master  password")
	a.pages.AddPage(pageName, frame, true, true)
}

// NewFormCreateMasterPassword implements AccessForms.
func (a *accessForm) NewFormCreateMasterPassword(
	pageName, buttonNameBack string, create func(password, repeat string), back func()) {
	form := tview.NewForm().SetFocus(1)

	form.AddInputField("master password", a.masterPassword, fieldWidth, nil, a.changePassword)
	form.AddInputField("repeat", a.repeatPasword, fieldWidth, nil, a.repeatPassword)

	form.AddButton(buttonNameBack, func() {
		a.pages.RemovePage(pageName)
		a.masterPassword = ""
		a.repeatPasword = ""
		back()
	})
	form.AddButton("next", func() {
		create(a.masterPassword, a.repeatPasword)
		a.email = ""
		a.masterPassword = ""
		a.repeatPasword = ""
		form.Clear(false)
		form.AddInputField("master password", a.masterPassword, fieldWidth, nil, a.changePassword)
		form.AddInputField("repeat", a.repeatPasword, fieldWidth, nil, a.repeatPassword)
	})
	frame := NewFrame(form, "Create MASTER PASSWORD")
	a.pages.AddPage(pageName, frame, true, true)
}

func NewAcessForms(pages *tview.Pages, email ...string) AccessForms {
	e := ""
	if len(email) != 0 {
		e = email[0]
	}
	return &accessForm{
		email: e,
		pages: pages,
	}
}

func (a *accessForm) NewFormGetAccess(
	pageName,
	buttonNameBack string,
	login func(form *core.AccessForm), back func()) {
	form := tview.NewForm().SetFocus(1)

	form.AddInputField("email", a.email, fieldWidth, nil, a.changeEmail)

	form.AddButton(buttonNameBack, func() {
		a.pages.RemovePage(pageName)
		back()
	})
	form.AddButton("send code", func() {
		login(&core.AccessForm{
			Email: a.email,
			Code:  a.code,
		})
	})

	frame := NewFrame(form, pageName)
	a.pages.AddPage(pageName, frame, true, true)
}

func (a *accessForm) NewFormConfirmAccess(
	pageName,
	buttonNameBack string, login func(form *core.AccessForm), back func()) {
	form := tview.NewForm().SetFocus(1)

	form.AddInputField("email", a.email, fieldWidth, nil, nil)
	form.AddInputField("code", "", fieldWidth, nil, a.changeCode)

	form.AddButton("back", func() {
		a.pages.RemovePage(pageName)
		a.code = ""
		back()
	})
	form.AddButton("login", func() {
		login(&core.AccessForm{
			Email: a.email,
			Code:  a.code,
		})
		a.code = ""
		form.Clear(false)
		form.AddInputField("email", a.email, fieldWidth, nil, nil)
		form.AddInputField("code", "", fieldWidth, nil, a.changeCode)
	})

	frame := NewFrame(form, "Confirm access")
	a.pages.AddPage(pageName, frame, true, true)
}
