package app

import (
	"github.com/rivo/tview"
)

func acceptEmail(textToCheck string, lastChar rune) bool {
	// log.Println(textToCheck, lastChar)
	return true
}

func changeEmail(text string) {
	// log.Println(text)
}

func NewFormLogin() *tview.Form {
	form := tview.NewForm()

	form.AddInputField("email", "", 200, acceptEmail, changeEmail)
	form.AddButton("login", func() {})
	return form
}
