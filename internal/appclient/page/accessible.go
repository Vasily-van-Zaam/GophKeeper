package page

import (
	"context"
	"fmt"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type accessiblePage struct {
	client         applicationClient
	name           string
	buttonNameBack string
	reset          func()
	back           func()
	next           func(user *core.User)
	conf           config.Config
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

	buttonSyncServer := tview.NewButton("🔄").SetSelectedFunc(func() {

	})
	buttonSyncServer.SetBackgroundColorActivated(tcell.ColorIndianRed)
	userID := a.client.User().ID.String()

	data, err := a.client.Repository().Local().GetData(ctx, userID)

	frame := component.NewGridAccessible(data, appInfo, func(m *core.ManagerData) {
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
	}
}
