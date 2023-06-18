package page

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/repository"
	"github.com/rivo/tview"
)

type applicationClient interface {
	Run() error
	Stop()
	Pages() *tview.Pages
	Repository() repository.Repository
}

type AppPage interface {
	Show(context.Context, bool) AppPage
	Close(bool) AppPage

	Back(func()) AppPage
	Next(func(puk string)) AppPage
	Reset(func()) AppPage
}
