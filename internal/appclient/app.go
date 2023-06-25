// Application Client
// project GophKeeper Yandex Practicum
// used package Tview
// Created by Vasiliy Van-Zaam
package appclient

import (
	"context"
	"regexp"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/page"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/repository"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/storage/localstore"
	"github.com/rivo/tview"
	"google.golang.org/grpc/metadata"
)

type ApplicationClient interface {
	Run() error
	AppInfo() *core.AppInfo
	Stop()
	Pages() *tview.Pages
	Repository() repository.Repository
	Config() config.Config
	User() *core.User
	App() *tview.Application
	CompareDataSync(local []*core.ManagerData, remote []*core.ManagerData) ([]*core.CopmareData, []*core.ManagerData, []*core.ManagerData)
}

type client struct {
	pages       *tview.Pages
	user        *core.User
	managerData []*core.ManagerData
	app         *tview.Application
	repository  repository.Repository
	config      config.Config
}

// AppInfo implements ApplicationClient.
func (c *client) AppInfo() *core.AppInfo {
	info := core.NewAppInfo(c.config.Client(), c.repository.Store())
	return info
}

// App implements ApplicationClient.
func (c *client) App() *tview.Application {
	return c.app
}

// User implements ApplicationClient.
func (c *client) User() *core.User {
	return c.user
}

// Config implements ApplicationClient.
func (c *client) Config() config.Config {
	return c.config
}

// Pages implements ApplicationClient.
func (c *client) Pages() *tview.Pages {
	return c.pages
}

// Repository implements ApplicationClient.
func (c *client) Repository() repository.Repository {
	return c.repository
}

// Stop implements ApplicationClient.
func (c *client) Stop() {
	c.app.Stop()
}

func (c *client) Run() error {
	err := c.startClient()
	if err != nil {
		return err
	}
	return c.app.SetRoot(c.pages, true).EnableMouse(true).Run()
}

// Create new Client.
func New(conf config.Config) (ApplicationClient, error) {
	lstore, err := localstore.New(conf)
	if err != nil {
		return nil, err
	}

	return &client{
		pages:      tview.NewPages(),
		app:        tview.NewApplication(),
		repository: repository.New(conf, lstore),
		config:     conf,
	}, nil
}

func (c *client) checkValidPassword(psw string) bool {
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
func (c *client) startClient() error {
	ctx := context.Background()
	md := metadata.New(map[string]string{core.CtxVersionClientKey: c.config.Client().Version()})

	ctx = metadata.NewOutgoingContext(ctx, md)
	////////////////////////////////////////////////////////////////

	accessPage := page.NewGetAccessPage(c, "GetAccess", "close")
	resetAccessPage := page.NewGetAccessPage(c, "GetAccess", "back")
	loginPage := page.NewLoginPage(c, "Login", "close")

	accessiblePage := page.NewAccessiblePage(c.config, c, "AccessiblePage", "close")

	loginPage.
		Next(func(user *core.User) {
			// log.Println(puk)
			c.user = user
			accessiblePage.
				Show(ctx, true).
				Reset(func() {
					accessiblePage.Close(true)
					loginPage.Show(ctx, true)
				})
			loginPage.Close(true)
			accessPage.Close(true)
		}).
		Reset(func() {
			resetAccessPage.
				Back(func() {
					// c.pages.SwitchToPage("Login")
					loginPage.Show(ctx, true)
					resetAccessPage.Close(true)
				}).
				Next(func(user *core.User) {
					loginPage.Show(ctx, true)
					resetAccessPage.Close(true)
				}).
				Show(ctx, true)
		}).
		Back(func() {
			c.Stop()
		})

	d, errAccess := c.repository.Store().GetAccessData(ctx)
	accessPage.
		Show(ctx, errAccess != nil).
		Next(func(user *core.User) {
			loginPage.Show(ctx, true)
			accessPage.Close(true)
		}).
		Back(func() {
			c.Stop()
		})
	loginPage.Show(ctx, d != nil)
	return nil
}

// Compare data sync by hash.
func (*client) CompareDataSync(
	local []*core.ManagerData,
	remote []*core.ManagerData,
) ([]*core.CopmareData, []*core.ManagerData, []*core.ManagerData) {
	comapare := make([]*core.CopmareData, 0)

	findNewData := func(d *core.ManagerData, data []*core.ManagerData) []*core.ManagerData {
		list := make([]*core.ManagerData, 0)

		for _, check := range data {
			if d.InfoData.ID.String() != check.InfoData.ID.String() {
				list = append(list, check)
			}
		}

		return list
	}
	newRemoteData := remote
	newLocalData := local
	for _, ld := range local {
		for _, rd := range newRemoteData {
			lID := ld.InfoData.ID.String()
			rID := rd.InfoData.ID.String()
			lHash := ld.InfoData.Hash
			rHash := rd.InfoData.Hash
			check := lID == rID && lHash != rHash
			if check {
				comapare = append(comapare, &core.CopmareData{
					Local:  ld,
					Remote: rd,
				})
			}
		}
		newRemoteData = findNewData(ld, newRemoteData)
	}
	for _, rd := range remote {
		newLocalData = findNewData(rd, newLocalData)
	}

	return comapare, newRemoteData, newLocalData
}
