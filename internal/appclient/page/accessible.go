package page

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
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
	// if !show {
	// 	return a
	// }
	button := tview.NewButton("❌").SetSelectedFunc(func() {
		a.reset()
	})
	button2 := tview.NewButton("🔄").SetSelectedFunc(func() {
		a.reset()
	})
	button2.SetBackgroundColorActivated(tcell.ColorIndianRed)
	userID := a.client.User().ID.String()
	// button.SetBorder(true).SetRect(0, 0, 22, 3)

	data, err := a.client.Repository().Local().GetData(ctx, userID)
	frame := component.NewGridAccessible(data, button2, button)
	a.client.Pages().AddPage(a.name, frame, true, true)
	if err != nil {
		component.ModalError(err, "AccessiblePage", a.client.Pages())
		// return a
	}
	// NewFrameAccessible(button, button, "GophKeeper v0.1")

	return a
}

func NewAccessiblePage(
	client applicationClient,
	pageName,
	buttonNameBack string,
) AppPage {
	return &accessiblePage{
		client:         client,
		name:           pageName,
		buttonNameBack: buttonNameBack,
	}
}
