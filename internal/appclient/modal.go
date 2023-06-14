package appclient

import "github.com/rivo/tview"

func ModalError(err error, backPage string, pages *tview.Pages) {
	modal := tview.NewModal()
	//
	modal.AddButtons([]string{"ok"}).
		SetText(err.Error()).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("ModalErr")
			pages.SwitchToPage(backPage)
		})

	pages.AddPage("ModalErr", modal, true, true)
}
