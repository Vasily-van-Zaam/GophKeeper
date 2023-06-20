package page

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/repository"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/rivo/tview"
)

type applicationClient interface {
	Run() error
	Stop()
	Pages() *tview.Pages
	Repository() repository.Repository
	User() *core.User
}

type AppPage interface {
	Show(context.Context, bool) AppPage
	Close(bool) AppPage

	Back(func()) AppPage
	Next(func(user *core.User)) AppPage
	Reset(func()) AppPage
}
