package component

import (
	"encoding/json"
	"fmt"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewGridAccessible(data []*core.ManagerData, button2, buttonClose *tview.Button) *tview.Grid {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	menu := newPrimitive("GophKeeper v0.0.0")
	main := newPrimitive("GophKeeper v0.0.0")
	// var list *tview.List
	// list = NewList(data, func(items []*core.ManagerData) {
	// 	list.Clear()
	// 	list = tview.NewList().
	// 		AddItem("List item 1", "Some explanatory text", 'a', nil).
	// 		AddItem("List item 2", "Some explanatory text", 'b', nil).
	// 		AddItem("List item 3", "Some explanatory text", 'c', nil).
	// 		AddItem("List item 4", "Some explanatory text", 'd', nil).
	// 		AddItem("Quit", "Press to exit", 'q', func() {

	// 		})
	// })
	// list = allList

	rootDir := "/"
	root := tview.NewTreeNode("...").
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
		// list := make([]*tview.TreeNode, 10)
		for i, v := range data {
			var c int
			_ = json.Unmarshal(v.Data, &c)
			node := tview.NewTreeNode(fmt.Sprintln(i, v.InfoData.DataType, c))

			target.AddChild(node)
		}

		// files, err := ioutil.ReadDir(path)
		// if err != nil {
		// 	panic(err)
		// }
		// for _, file := range files {
		// 	node := tview.NewTreeNode(file.Name()).
		// 		SetReference(filepath.Join(path, file.Name())).
		// 		SetSelectable(file.IsDir())
		// 	if file.IsDir() {
		// 		node.SetColor(tcell.ColorGreen)
		// 	}
		// 	target.AddChild(node)
		// }
	}

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	// 	})
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
		SetColumns(20, 0).
		SetBorders(true).
		AddItem(gridHeader, 0, 0, 1, 2, 0, 0, false) //.
		// AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	// grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
	// 	AddItem(list, 1, 0, 1, 3, 0, 0, false) //.
	// AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(tree, 1, 0, 1, 1, 0, 0, false).
		AddItem(menu, 1, 1, 1, 1, 0, 0, false) //.
	// AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)
	return grid
}
