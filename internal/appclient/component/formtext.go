package component

import (
	"fmt"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewFormTextEditor(text *core.TextFomm, onSave func(text *core.TextFomm), onClose func()) *tview.Grid {
	metadata := text.MetaData
	textArea := tview.NewTextArea().
		SetPlaceholder("Enter text here...")
	textArea.SetText(text.Text, true)
	textArea.SetTitle("Edit Text").SetBorder(true)
	inputMetadata := tview.NewInputField().SetPlaceholder("Name").SetChangedFunc(func(text string) {
		metadata = text
	}).SetText(metadata)
	inputMetadata.SetTitle("Name")
	saveButton := tview.NewButton("✅ Save").SetSelectedFunc(func() {
		onSave(&core.TextFomm{
			Text: textArea.GetText(),

			MetaData: metadata,
		})
	}).SetStyle(tcell.StyleDefault.Attributes(tcell.AttrBold))
	closeButton := tview.NewButton("❌ Close").SetSelectedFunc(func() {
		onClose()
	}).SetStyle(tcell.StyleDefault.Attributes(tcell.AttrBold))

	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)
	pages := tview.NewPages()

	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	textArea.SetMovedFunc(updateInfos)
	updateInfos()

	mainView := tview.NewGrid().
		SetRows(0, 1, 1, 1).
		AddItem(textArea, 0, 0, 1, 3, 0, 0, true).
		AddItem(inputMetadata, 1, 0, 1, 3, 0, 0, true).
		AddItem(tview.NewBox(), 2, 0, 1, 0, 0, 0, false).
		AddItem(closeButton, 3, 0, 1, 1, 0, 0, false).
		AddItem(saveButton, 3, 1, 1, 1, 0, 0, false).
		AddItem(position, 3, 2, 1, 1, 0, 0, false)

	help1 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Navigation

[yellow]Left arrow[white]: Move left.
[yellow]Right arrow[white]: Move right.
[yellow]Down arrow[white]: Move down.
[yellow]Up arrow[white]: Move up.
[yellow]Ctrl-A, Home[white]: Move to the beginning of the current line.
[yellow]Ctrl-E, End[white]: Move to the end of the current line.
[yellow]Ctrl-F, page down[white]: Move down by one page.
[yellow]Ctrl-B, page up[white]: Move up by one page.
[yellow]Alt-Up arrow[white]: Scroll the page up.
[yellow]Alt-Down arrow[white]: Scroll the page down.
[yellow]Alt-Left arrow[white]: Scroll the page to the left.
[yellow]Alt-Right arrow[white]:  Scroll the page to the right.
[yellow]Alt-B, Ctrl-Left arrow[white]: Move back by one word.
[yellow]Alt-F, Ctrl-Right arrow[white]: Move forward by one word.

[blue]Press Enter for more help, press Escape to return.`)
	help2 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Editing[white]

Type to enter text.
[yellow]Ctrl-H, Backspace[white]: Delete the left character.
[yellow]Ctrl-D, Delete[white]: Delete the right character.
[yellow]Ctrl-K[white]: Delete until the end of the line.
[yellow]Ctrl-W[white]: Delete the rest of the word.
[yellow]Ctrl-U[white]: Delete the current line.

[blue]Press Enter for more help, press Escape to return.`)
	help3 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Selecting Text[white]

Move while holding Shift or drag the mouse.
Double-click to select a word.
[yellow]Ctrl-L[white] to select entire text.

[green]Clipboard

[yellow]Ctrl-Q[white]: Copy.
[yellow]Ctrl-X[white]: Cut.
[yellow]Ctrl-V[white]: Paste.
		
[green]Undo

[yellow]Ctrl-Z[white]: Undo.
[yellow]Ctrl-Y[white]: Redo.

[blue]Press Enter for more help, press Escape to return.`)
	help := tview.NewFrame(help1).
		SetBorders(1, 1, 0, 0, 0, 0)
	help.SetBorder(true).
		SetTitle("Help").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				pages.SwitchToPage("main")
				return nil
			} else if event.Key() == tcell.KeyEnter {
				switch {
				case help.GetPrimitive() == help1:
					help.SetPrimitive(help2)
				case help.GetPrimitive() == help2:
					help.SetPrimitive(help3)
				case help.GetPrimitive() == help3:
					help.SetPrimitive(help1)
				}
				return nil
			}
			return event
		})

	return mainView
}
