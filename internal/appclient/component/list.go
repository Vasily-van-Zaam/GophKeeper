package component

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/rivo/tview"
)

func NewList(items []*core.ManagerData, selected func(items []*core.ManagerData)) *tview.List {
	list := tview.NewList().
		AddItem("Passwords", "List passwords", 'p', func() {
			// tview.NewList().AddItem("Cards", "List ban cards", 'c', nil)
			selected(items)
		}).
		AddItem("Cards", "List ban cards", 'c', nil).
		AddItem("Texts", "List texts", 't', nil).
		AddItem("Files", "List files", 'f', nil) //.
		// AddItem("Quit", "Press to exit", 'q', func() {

		// })

	// for _, item := range items {
	// 	list.AddItem(item.DataType, item.MetaData, 11, nil)
	// }
	return list
}
