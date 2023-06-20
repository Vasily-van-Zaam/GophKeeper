package page

import (
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
	// "github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/page".

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

func (g *getAccessPage) checkValidPassword(psw string) bool {
	secure := true
	tests := []string{".{8,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, psw)
		if !t {
			secure = false
			break
		}
	}
	return secure
}

type getAccessPage struct {
	client         applicationClient
	name           string
	buttonNameBack string
	reset          func()
	back           func()
	next           func(user *core.User)
}

// Close implements AppPage.
func (g *getAccessPage) Close(bool) AppPage {
	g.client.Pages().RemovePage(g.name)
	return g
}

// Reset implements AppPage.
func (g *getAccessPage) Reset(reset func()) AppPage {
	g.reset = reset
	return g
}

// Back implements AppPage.
func (g *getAccessPage) Back(back func()) AppPage {
	g.back = back
	return g
}

// Next implements AppPage.
func (g *getAccessPage) Next(next func(user *core.User)) AppPage {
	g.next = next
	return g
}

func (g *getAccessPage) formLogin(ctx context.Context, form *core.AccessForm) []byte {
	token, errGet := g.client.Repository().Remote().GetAccess(ctx, form)
	if errGet != nil {
		component.ModalError(errGet, g.name, g.client.Pages())
		return nil
	}
	return token
}
func (g *getAccessPage) formConfirm(ctx context.Context, form *core.AccessForm) *core.User {
	user, err1 := g.client.Repository().Remote().ConfirmAccess(ctx, form)
	if err1 != nil {
		component.ModalError(err1, "Confirm", g.client.Pages())
		return nil
	}
	return user
}
func (g *getAccessPage) formCreatePsw(ctx context.Context, user *core.User, password, repeat string) {
	if password != repeat {
		component.ModalError(errors.New("password mismatch"), "CreatPassword", g.client.Pages())
		return
	} else if !g.checkValidPassword(password) {
		component.ModalError(errors.New(
			"password must be at least 8 characters long, contain"+
				" special characters, numbers, and lowercase and uppercase characters"), "CreatPassword", g.client.Pages())
		return
	}
	err := g.client.Repository().Local().ResetUserData(ctx)
	log.Println(err)
	err = g.client.Repository().Local().AddAccessData(ctx, password, user)
	if err != nil {
		component.ModalError(err, "CreatPassword", g.client.Pages())
		return
	}
	g.client.Pages().RemovePage(g.name)
	g.client.Pages().RemovePage("Confirm")
	g.client.Pages().RemovePage("CreatPassword")
	g.next(nil)
}

// Show implements AppPage.
func (g *getAccessPage) Show(ctx context.Context, show bool) AppPage {
	if !show {
		return g
	}
	accessForms := component.NewAcessForms(g.client.Pages())
	accessForms.NewFormGetAccess(g.name, g.buttonNameBack, func(form *core.AccessForm) {
		token := g.formLogin(ctx, form)
		if token == nil {
			return
		}
		accessForms.NewFormConfirmAccess("Confirm", "back", func(form *core.AccessForm) {
			form.Token = token
			user := g.formConfirm(ctx, form)
			if user != nil {
				accessForms.NewFormCreateMasterPassword("CreatPassword", "back", func(password, repeat string) {
					g.formCreatePsw(ctx, user, password, repeat)
				}, func() {
					g.client.Pages().SwitchToPage("Confirm")
				})
			}
		}, func() {
			g.client.Pages().SwitchToPage(g.name)
		})
	}, func() {
		g.back()
	})
	return g
}

// Page name is "GetAccess".
func NewGetAccessPage(
	client applicationClient,
	pageName,
	buttonNameBack string,
) AppPage {
	return &getAccessPage{
		client:         client,
		name:           pageName,
		buttonNameBack: buttonNameBack,
	}
}
