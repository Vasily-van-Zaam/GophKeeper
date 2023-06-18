// Application Client
// project GophKeeper Yandex Practicum
// used package Tview
// Created by Vasiliy Van-Zaam
package appclient

import (
	"context"
	"errors"
	"regexp"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
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
	Stop()
	Pages() *tview.Pages
	Repository() repository.Repository
	Config() config.Config
}

type client struct {
	pages          *tview.Pages
	privateUserKey string
	// form       *tview.Form
	app *tview.Application
	// modal      *tview.Modal
	// flex       *tview.Flex
	// textView   *tview.TextView
	// list       *tview.List
	// button     *tview.Button
	repository repository.Repository
	config     config.Config
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
		pages: tview.NewPages(),
		app:   tview.NewApplication(),
		// flex:       tview.NewFlex(),
		// textView:   tview.NewTextView(),
		// list:       tview.NewList(),
		// modal:      tview.NewModal(),
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

	accessiblePage := page.NewAccessiblePage(c, "AccessiblePage", "close")
	// accessiblePage.Show(ctx, true)
	// return nil
	loginPage.
		Next(func(puk string) {
			// log.Println(puk)

			accessiblePage.
				Show(ctx, true).
				Reset(func() {
					accessiblePage.Close(true)
					loginPage.Show(ctx, true)
				})
			loginPage.Close(true)
			accessPage.Close(true)
			c.privateUserKey = puk
			component.ModalError(errors.New(puk), "AccessiblePage", c.pages)
		}).
		Reset(func() {
			resetAccessPage.
				Back(func() {
					// c.pages.SwitchToPage("Login")
					loginPage.Show(ctx, true)
					resetAccessPage.Close(true)
				}).
				Next(func(puk string) {
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
		Next(func(puk string) {
			loginPage.Show(ctx, true)
			accessPage.Close(true)
		}).
		Back(func() {
			c.Stop()
		})
	loginPage.Show(ctx, d != nil)
	return nil
}
