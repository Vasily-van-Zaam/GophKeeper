package appclient

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *accessForm) changeEmail(text string) {
	a.email = text
}

func (a *accessForm) changeCode(text string) {
	a.code = text
}

type AccessForms interface {
	NewFormGetAccess(pageName string, login func(form *core.AccessForm), back func())
	NewFormConfirmAccess(pageName string, login func(form *core.AccessForm), back func())
}
type accessForm struct {
	email string
	code  string
	pages *tview.Pages
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
	pageName string, login func(form *core.AccessForm), back func()) {
	form := tview.NewForm().SetFocus(1)

	form.AddInputField("email", a.email, 50, nil, a.changeEmail)

	form.AddButton("close", func() {
		a.pages.RemovePage(pageName)
		back()
	})
	form.AddButton("send code", func() {
		login(&core.AccessForm{
			Email: a.email,
			Code:  a.code,
		})
	})

	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(2, 2, 2, 2, 4, 4).
		SetPrimitive(form).
		AddText("GophKeeper", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Authorization", true, tview.AlignCenter, tcell.ColorWhite).
		AddText("GophKeeper", true, tview.AlignRight, tcell.ColorWhite).
		// AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
		AddText("Super GophKeeper", false, tview.AlignCenter, tcell.ColorGreen).
		AddText("GophKeeper inc", false, tview.AlignCenter, tcell.ColorGreen)

	a.pages.AddPage(pageName, frame, true, true)
}

func (a *accessForm) NewFormConfirmAccess(
	pageName string, login func(form *core.AccessForm), back func()) {
	form := tview.NewForm().SetFocus(1)

	form.AddInputField("email", a.email, 50, nil, nil)
	form.AddInputField("code", "", 50, nil, a.changeCode)

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
	})

	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(2, 2, 2, 2, 4, 4).
		SetPrimitive(form).
		AddText("GophKeeper", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Confirm access", true, tview.AlignCenter, tcell.ColorWhite).
		AddText("GophKeeper", true, tview.AlignRight, tcell.ColorWhite).
		// AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
		AddText("Super GophKeeper", false, tview.AlignCenter, tcell.ColorGreen).
		AddText("GophKeeper inc", false, tview.AlignCenter, tcell.ColorGreen)
	a.pages.AddPage(pageName, frame, true, true)
}
