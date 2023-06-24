package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc/status"
)

func ModalError(err error, backPage string, pages *tview.Pages) {
	modal := tview.NewModal()
	status := status.Convert(err)
	modal.AddButtons([]string{"ok"}).
		SetText(status.Message()).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("ModalErr")
		})

	pages.AddPage("ModalErr", modal, true, true)
}
func ModalEnterPassword(title string, backPage string, pages *tview.Pages, on func(psw string)) {
	password := ""
	form := tview.NewForm().AddInputField("Password", password, 1000,
		func(textToCheck string, lastChar rune) bool {
			return true
		}, func(text string) {
			password = text
		}).AddButton("Enter Password", func() {
		on(password)
		pages.SwitchToPage(backPage)
	})
	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		AddText(title, true, 1, tcell.ColorRed).
		SetPrimitive(form).SetBorders(10, 2, 2, 2, 10, 10)
	pages.AddPage("ModalPsw", frame, true, true)
}

func ModalNewEnterPassword(title string, backPage string, pages *tview.Pages, on func(psw string)) {
	form := tview.NewForm().AddInputField("password", "", 1000,
		func(textToCheck string, lastChar rune) bool {
			return true
		}, func(text string) {}).AddButton("Enter Password", func() {
		pages.SwitchToPage(backPage)
	})
	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		AddText(title, true, 1, tcell.ColorRed).
		SetPrimitive(form).SetBorders(10, 2, 2, 2, 10, 10)

	pages.AddPage("ModalErr", frame, true, true)
}
func ModalNewUploadFile(title string, backPage string, pages *tview.Pages, on func(path string)) {
	path := ""
	form := tview.NewForm().AddInputField("path file", "", 1000,
		func(textToCheck string, lastChar rune) bool {
			return true
		}, func(text string) {
			path = text
		}).AddButton("Upload", func() {
		pages.RemovePage(title)
		on(path)
	})
	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		AddText(title, true, 1, tcell.ColorRed).
		SetPrimitive(form).SetBorders(0, 2, 2, 2, 10, 10)

	pages.AddPage(title, frame, true, true)
}
func ModalNewDownloadFile(title, fileName string, backPage string, pages *tview.Pages, on func(fileName, path string)) {
	path := ""
	name := fileName
	form := tview.NewForm().
		AddInputField("file path to save*", "", 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				path = text
			}).
		AddInputField("file name to save*", name, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				name = text
			}).
		AddButton("Save file", func() {
			pages.RemovePage(title)
			on(name, path)
		})
	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		AddText(title, true, 1, tcell.ColorRed).
		SetPrimitive(form).SetBorders(0, 2, 2, 2, 10, 10)

	pages.AddPage(title, frame, true, true)
	// form := tview.NewForm().AddInputField("path file to save", "", 1000,
	// 	func(textToCheck string, lastChar rune) bool {
	// 		return true
	// 	}, func(text string) {}).AddButton("Save", func() {
	// 	pages.SwitchToPage(backPage)
	// })
	// frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
	// 	AddText(title, true, 1, tcell.ColorRed).
	// 	SetPrimitive(form).SetBorders(10, 2, 2, 2, 10, 10)

	// pages.AddPage("ModalErr", frame, true, true)
}
