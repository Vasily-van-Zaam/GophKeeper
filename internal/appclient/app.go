// Application Client
// project GophKeeper Yandex Practicum
// used package Tview
// Created by Vasiliy Van-Zaam
package appclient

import (
	"context"
	"log"

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
}

type client struct {
	pages *tview.Pages
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

func (c *client) startClient() error {
	ctx := context.Background()
	md := metadata.New(map[string]string{core.CtxVersionClientKey: c.config.Client().Version()})

	ctx = metadata.NewOutgoingContext(ctx, md)
	////////////////////////////////////////////////////////////////
	data, err := c.repository.Local().GetData(ctx, string(core.DataTypeUser))
	if err != nil {
		return err
	}
	if len(data) == 0 {
		accessForms := NewAcessForms(c.pages, "test@mail.ru")
		accessForms.NewFormGetAccess("GetAccess", func(form *core.AccessForm) {
			token, errGet := c.repository.Remote().GetAccess(ctx, form)
			if errGet != nil {
				ModalError(errGet, "GetAccess", c.pages)
				return
			}
			accessForms.NewFormConfirmAccess("Confirm", func(form *core.AccessForm) {
				form.Token = token

				user, err1 := c.repository.Remote().ConfirmAccess(ctx, form)
				if err1 != nil {
					ModalError(err1, "Confirm", c.pages)
					return
				}
				log.Print(user)
			}, func() {
				c.pages.SwitchToPage("GetAccess")
			})
		}, func() {
			c.Stop()
		})
	}

	return nil
}
