package page

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
)

type loginPage struct {
	client         applicationClient
	name           string
	buttonNameBack string
	reset          func()
	back           func()
	next           func(puk string)
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
func (l *loginPage) Next(next func(puk string)) AppPage {
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
				component.ModalError(err, "Login", l.client.Pages())
				return
			}
			l.next(localUser.PrivateKey)
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
