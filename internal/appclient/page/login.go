package page

import (
	"context"
	"fmt"
	"strings"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
)

type loginPage struct {
	client         applicationClient
	name           string
	buttonNameBack string
	reset          func()
	back           func()
	next           func(user *core.User)
}

// Close implements AppPage.
func (l *loginPage) Close(bool) AppPage {
	l.client.Pages().RemovePage(l.name)
	return l
}

// Reset implements AppPage.
func (l *loginPage) Reset(reset func()) AppPage {
	l.reset = reset
	return l
}

// Back implements AppPage.
func (l *loginPage) Back(back func()) AppPage {
	l.back = back
	return l
}

// Next implements AppPage.
func (l *loginPage) Next(next func(user *core.User)) AppPage {
	l.next = next
	return l
}

// Show implements AppPage.
func (l *loginPage) Show(ctx context.Context, show bool) AppPage {
	if !show {
		return l
	}
	accessForms := component.NewAcessForms(l.client.Pages())
	accessForms.NewFormLogin(l.name,
		l.buttonNameBack, func(password string) {
			localUser, err := l.client.Repository().Local().GetAccessData(ctx, password)
			if err != nil {
				if strings.Contains(err.Error(), "authentication failed") {
					count, _ := l.client.Repository().Local().GetTryPasword(ctx)
					component.ModalError(
						fmt.Errorf("error password. There are %v attempts left", 3-count),
						"Login", l.client.Pages())
					if 3-count < 0 {
						l.client.Repository().Local().ResetUserData(ctx, core.DataTypeTryEnterPassword, core.DataTypeUser)
						l.client.Stop()
					}
					_ = l.client.Repository().Local().AddTryPasword(ctx, count)
					return
				}
				return
			}
			l.client.Repository().Local().ResetUserData(ctx, core.DataTypeTryEnterPassword)
			l.next(localUser)
		}, func() {
			l.reset()
		},
		func() {
			l.client.Stop()
		})
	return l
}

func NewLoginPage(
	client applicationClient,
	pageName,
	buttonNameBack string,

) AppPage {
	return &loginPage{
		client:         client,
		name:           pageName,
		buttonNameBack: buttonNameBack,
	}
}
