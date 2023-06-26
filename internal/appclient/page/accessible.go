package page

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc/metadata"
)

type accessiblePage struct {
	client         applicationClient
	name           string
	buttonNameBack string
	reset          func()
	back           func()
	next           func(user *core.User)
	conf           config.Config
	auth           core.Auth
}

// Adding secreet info fore a remote request.
func (a *accessiblePage) addAuthContext(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{core.CtxVersionClientKey: a.conf.Client().Version()})
	ctx = metadata.NewIncomingContext(ctx, md)
	userID, _ := a.auth.EncryptData(ctx, a.client.User().ID)
	md.Append(core.CtxTokenKey, hex.EncodeToString(userID))
	return metadata.NewOutgoingContext(ctx, md)
}

// Close implements AppPage.
func (a *accessiblePage) Close(bool) AppPage {
	a.client.Pages().RemovePage(a.name)
	return a
}

// Reset implements AppPage.
func (a *accessiblePage) Reset(reset func()) AppPage {
	a.reset = reset
	return a
}

// Back implements AppPage.
func (a *accessiblePage) Back(back func()) AppPage {
	a.back = back
	return a
}

// Next implements AppPage.
func (a *accessiblePage) Next(next func(user *core.User)) AppPage {
	a.next = next
	return a
}

// Show implements AppPage.
func (a *accessiblePage) Show(ctx context.Context, show bool) AppPage {
	if !show {
		return a
	}
	button := tview.NewButton("❌").SetSelectedFunc(func() {
		a.reset()
	})

	appInfo := a.client.AppInfo()
	userID := a.client.User().ID.String()
	localData, err := a.client.Repository().Local().GetData(ctx, userID)

	if err != nil {
		component.ModalError(err, a.name, a.client.Pages())
	}

	buttonSyncServer := tview.NewButton("🔄").SetSelectedFunc(func() {
		component.ModalInfo("Synchronize data with the server?", a.name, a.client.Pages(), func() {
			remoteData, errGetDataInfo := a.client.Repository().Remote().GetData(a.addAuthContext(ctx))
			if errGetDataInfo != nil {
				component.ModalError(errGetDataInfo, a.name, a.client.Pages())
				return
			}
			compare, newRemote, newLocal := a.client.CompareDataSync(localData, remoteData)

			a.client.Pages().RemovePage("ModalInfo")

			component.ModalInfo(
				fmt.Sprintf("differences %v, remote %v local %v, update?", len(compare),
					len(newRemote), len(newLocal)), a.name, a.client.Pages(), func() {
					if len(newRemote) != 0 {
						list, err := a.client.Repository().Local().AddData(ctx, newRemote...)
						if err != nil {
							component.ModalError(err, a.name, a.client.Pages())
							return
						}
						localData = append(localData, list...)
						a.client.Pages().RemovePage(a.name)
						a.Show(ctx, true)
					}
					if len(newLocal) != 0 {
						_, err = a.client.Repository().Remote().AddData(ctx, newLocal...)
						if err != nil {
							component.ModalError(err, a.name, a.client.Pages())
							return
						}
						a.client.Pages().RemovePage(a.name)
						a.Show(ctx, true)
					}
				})
		})
	})
	buttonSyncServer.SetBackgroundColorActivated(tcell.ColorIndianRed)

	frame := component.NewGridAccessible(localData, appInfo, func(m *core.ManagerData) {
		editorPageName := fmt.Sprintf("Edit %v", m.InfoData.DataType)
		a.client.Pages().RemovePage(a.name)
		edit := NewEditorPage(
			a.conf,
			a.client,
			core.NewManagerFromData(m),
			editorPageName,
		).Back(func() {
			a.client.Pages().RemovePage(editorPageName)
			a.Show(ctx, true)
		})
		edit.Show(ctx, true)
	}, func(dataType core.DataType) {
		additorPageName := fmt.Sprintf("Add new %v", dataType)
		a.client.Pages().RemovePage(a.name)
		edit := NewEditorPage(
			a.conf,
			a.client,
			core.NewManager(a.client.User().ID, dataType),
			additorPageName,
		).Back(func() {
			a.client.Pages().RemovePage(additorPageName)
			a.Show(ctx, true)
		})
		edit.Show(ctx, true)
	}, buttonSyncServer, button)
	a.client.Pages().AddPage(a.name, frame, true, true)
	if err != nil {
		component.ModalError(err, "AccessiblePage", a.client.Pages())
		// return a
	}

	return a
}

func NewAccessiblePage(
	conf config.Config,
	client applicationClient,
	pageName,
	buttonNameBack string,
) AppPage {
	return &accessiblePage{
		client:         client,
		name:           pageName,
		buttonNameBack: buttonNameBack,
		conf:           conf,
		auth:           core.NewAuth(conf),
	}
}
