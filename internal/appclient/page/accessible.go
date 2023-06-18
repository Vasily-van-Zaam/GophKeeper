package page

import (
	"context"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type accessiblePage struct {
	client         applicationClient
	name           string
	buttonNameBack string
	reset          func()
	back           func()
	next           func(puk string)
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
func (a *accessiblePage) Next(next func(puk string)) AppPage {
	a.next = next
	return a
}

// Show implements AppPage.
func (a *accessiblePage) Show(context.Context, bool) AppPage {
	button := tview.NewButton("❌").SetSelectedFunc(func() {
		a.reset()
	})
	button2 := tview.NewButton("🔄").SetSelectedFunc(func() {
		a.reset()
	})
	button2.SetBackgroundColorActivated(tcell.ColorIndianRed)

	// button.SetBorder(true).SetRect(0, 0, 22, 3)

	frame := component.NewGridAccessible(button2, button) // NewFrameAccessible(button, button, "GophKeeper v0.1")
	a.client.Pages().AddPage(a.name, frame, true, true)
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
