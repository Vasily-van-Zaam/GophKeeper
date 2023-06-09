package app

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/rivo/tview"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func acceptEmail(textToCheck string, lastChar rune) bool {
	// log.Println(textToCheck, lastChar)
	return true
}

func changeEmail(text string) {
	email = text
}
func changePasword(text string) {
	password = text
}

var (
	email    string = ""
	password string = ""
)

func NewFormLogin(login func(form *core.LoginForm), back func()) *tview.Form {
	form := tview.NewForm()

	form.AddInputField("email", "", 200, acceptEmail, changeEmail)
	form.AddInputField("password", "", 200, acceptEmail, changePasword)

	form.AddButton("back", back)
	form.AddButton("login", func() {
		login(&core.LoginForm{
			Email:   email,
			Pasword: password,
		})
	})
	return form
}
