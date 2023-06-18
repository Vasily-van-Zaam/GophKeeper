package component

import (
	"github.com/rivo/tview"
	"google.golang.org/grpc/status"
)

func ModalError(err error, backPage string, pages *tview.Pages) {
	modal := tview.NewModal()
	status := status.Convert(err)
	//
	modal.AddButtons([]string{"ok"}).
		SetText(status.Message()).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("ModalErr")
			pages.SwitchToPage(backPage)
		})

	pages.AddPage("ModalErr", modal, true, true)
}
