// Application Client
// project GophKeeper Yandex Practicum
// used package Tview
// Created by Vasiliy Van-Zaam
package appclient

import (
	"context"
	"errors"
	"log"
	"regexp"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/repository"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/storage/localstore"
	"github.com/gdamore/tcell/v2"
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
	var (
		err       error
		localUser *core.User
	)

	localUser, err = c.repository.Local().GetAccessData(ctx)
	log.Print("-------", localUser, err, c.repository.Local().GetAccessData)
	if err != nil {
		if err.Error() != "not found" {
			return err
		}
	}

	if localUser == nil {
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

				if user != nil {
					accessForms.NewFormCreateMasterPassword("CreatPassword", func(password, repeat string) {
						if password != repeat {
							ModalError(errors.New("password mismatch"), "CreatPassword", c.pages)
							return
						} else if !c.checkValidPassword(password) {
							ModalError(errors.New(
								"password must be at least 8 characters long, contain"+
									" special characters, numbers, and lowercase and uppercase characters"), "CreatPassword", c.pages)
						}
						err = c.repository.Local().AddAccessData(ctx, password, user)
						if err != nil {
							ModalError(err, "CreatPassword", c.pages)
						}
					}, func() {
						c.pages.SwitchToPage("Confirm")
					})
				}
			}, func() {
				c.pages.SwitchToPage("GetAccess")
			})
		}, func() {
			c.Stop()
		})
	} else {
		frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
			SetBorders(2, 2, 2, 2, 4, 4).
			AddText("GophKeeper", true, tview.AlignLeft, tcell.ColorWhite).
			AddText("Create MASTER PASSWORD", true, tview.AlignCenter, tcell.ColorWhite).
			AddText("GophKeeper", true, tview.AlignRight, tcell.ColorWhite).
			// AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
			AddText("Super GophKeeper", false, tview.AlignCenter, tcell.ColorGreen).
			AddText("GophKeeper inc", false, tview.AlignCenter, tcell.ColorGreen)
		c.pages.AddPage("Main", frame, true, true)
	}

	return nil
}
