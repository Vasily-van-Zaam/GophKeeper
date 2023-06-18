package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewFrame(p tview.Primitive, title string) *tview.Frame {
	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		SetBorders(2, 2, 2, 2, 4, 4).
		SetPrimitive(p).
		AddText("GophKeeper", true, tview.AlignLeft, tcell.ColorWhite).
		AddText(title, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("GophKeeper", true, tview.AlignRight, tcell.ColorWhite).
		// AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
		AddText("Super GophKeeper", false, tview.AlignCenter, tcell.ColorGreen).
		AddText("GophKeeper inc", false, tview.AlignCenter, tcell.ColorGreen)

	return frame
}

func NewFrameAccessible(p tview.Primitive, buttonClose tview.Primitive, title string) *tview.Frame {
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("DATA"), 0, 5, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(title).SetTitleColor(tcell.ColorPaleVioletRed), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false), 0, 10, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(buttonClose, 0, 1, false).
			// AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(false).SetTitle("Bottom (5 rows)"), 0, 5, false), 0, 1, false)
		// AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 10, 1, false)

	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		// SetBorders(2, 2, 2, 2, 4, 4).
		SetPrimitive(flex)
		// AddText("GophKeeper", true, tview.AlignLeft, tcell.ColorWhite).
		// AddText(title, true, tview.AlignCenter, tcell.ColorWhite).
		// AddText("GophKeeper", true, tview.AlignRight, tcell.ColorWhite).
		// // AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
		// AddText("Super GophKeeper", false, tview.AlignCenter, tcell.ColorGreen).
		// AddText("GophKeeper inc", false, tview.AlignCenter, tcell.ColorGreen)

	return frame
}
