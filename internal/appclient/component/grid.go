package component

import (
	"github.com/rivo/tview"
)

func NewGridAccessible(button2, buttonClose *tview.Button) *tview.Grid {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	menu := newPrimitive("GophKeeper v0.0.0")
	main := newPrimitive("GophKeeper v0.0.0")

	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {

		})
	// sideBar := newPrimitive("Side Bar")
	gridHeader := tview.NewGrid().
		// SetRows(0, 0, 0).
		SetColumns(8, 0, 8).
		// SetBorders(true).
		AddItem(button2, 0, 0, 2, 1, 0, 0, false).
		AddItem(main, 0, 1, 2, 1, 0, 0, false).
		AddItem(buttonClose, 0, 2, 2, 1, 0, 0, false)

	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(gridHeader, 0, 0, 1, 2, 0, 0, false) //.
		// AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	// grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
	// 	AddItem(list, 1, 0, 1, 3, 0, 0, false) //.
	// AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(list, 1, 0, 1, 1, 0, 0, false).
		AddItem(menu, 1, 1, 1, 1, 0, 0, false) //.
	// AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)
	return grid
}
