package page

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/repository"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/rivo/tview"
)

type applicationClient interface {
	Run() error
	AppInfo() *core.AppInfo
	Stop()
	Pages() *tview.Pages
	Repository() repository.Repository
	Config() config.Config
	User() *core.User
	App() *tview.Application
	CompareDataSync(
		local []*core.ManagerData,
		remote []*core.ManagerData,
	) ([]*core.CopmareData, []*core.ManagerData, []*core.ManagerData)
}

type AppPage interface {
	Show(context.Context, bool) AppPage
	Close(bool) AppPage

	Back(func()) AppPage
	Next(func(user *core.User)) AppPage
	Reset(func()) AppPage
}
