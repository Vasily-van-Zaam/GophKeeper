package component

import (
	"fmt"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewGridAccessible(data []*core.ManagerData, appInfo *core.AppInfo,
	onSelect func(m *core.ManagerData),
	onAdd func(dataType core.DataType),
	syncButton, buttonClose *tview.Button) *tview.Grid {
	selectedList := tview.NewList()
	searchInput := tview.NewInputField().
		SetFieldWidth(40).SetLabel(" Search")
	searchInput.SetPlaceholder("Select group for search")
	root := tview.NewTreeNode("📁").SetSelectable(true)
	// SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	allNodes := []string{"🔑passwod", "💳card", "📝text", "📁file"}
	gridHeader := tview.NewGrid()
	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode) {
		for _, name := range allNodes {
			node := tview.NewTreeNode(fmt.Sprintln(name))

			node.SetSelectedFunc(func() {
				nodeName := node.GetText()
				list := make([]*core.ManagerData, 0)
				emojy := ""
				var dataType core.DataType

				switch nodeName {
				case "🔑passwod\n":
					dataType = core.DataTypePassword
					emojy = "🔑"
					for _, v := range data {
						if v.DataType == string(core.DataTypePassword) {
							list = append(list, v)
						}
					}
				case "💳card\n":
					dataType = core.DataTypeCard
					emojy = "💳"
					for _, v := range data {
						if v.DataType == string(core.DataTypeCard) {
							list = append(list, v)
						}
					}
				case "📝text\n":
					dataType = core.DataTypeText
					emojy = "📝"
					for _, v := range data {
						if v.DataType == string(core.DataTypeText) {
							list = append(list, v)
						}
					}
				case "📁file\n":
					dataType = core.DataTypeFile
					emojy = "📁"
					for _, v := range data {
						if v.DataType == string(core.DataTypeFile) {
							list = append(list, v)
						}
					}
				default:
					emojy = ""
				}

				if len(node.GetChildren()) == 0 {
					node.AddChild(tview.NewTreeNode(emojy + "➕").SetSelectedFunc(func() {
						onAdd(dataType)
					}).SetColor(tcell.ColorGreen))
					selectedList.Clear()
					for i, n := range list {
						child := tview.NewTreeNode(fmt.Sprintln(i+1, emojy, n.MetaData))
						node.AddChild(child)

						dataM := *n

						selectedList.AddItem(
							n.MetaData,
							"✏️ "+n.UpdatedAt.Local().Format("2006-01-02 15:04"),
							rune(0),
							func() {
								onSelect(&dataM)
							},
						).SetHighlightFullLine(true)
						child.SetSelectedFunc(func() {
							onSelect(&dataM)
						})
					}
					searchInput.SetText("")
					searchInput.SetPlaceholder("Selected search in " + string(dataType))
					searchInput.SetChangedFunc(func(text string) {
						indexes := selectedList.FindItems(text, "", true, true)
						if len(indexes) > 0 {
							selectedList.SetCurrentItem(indexes[0]).FindItems("", "", true, true)
						}
					})
				} else {
					node.ClearChildren()
				}
			})
			target.AddChild(node)
		}
	}
	add(root)

	f := tview.NewInputField().SetFieldWidth(100).
		SetLabel(" Version ").SetText(fmt.Sprint(appInfo.Version))
	f2 := tview.NewInputField().SetFieldWidth(100).
		SetLabel(" store size  ").SetText(fmt.Sprint(appInfo.SizeStore, " "))
	f3 := tview.NewInputField().SetFieldWidth(100).
		SetLabel(" last sync ").SetText(fmt.Sprint(appInfo.LastSync, " "))
	gridHeader.
		// SetRows(0, 0, 0).
		SetColumns(8, 30, 0, 8).
		// SetBorders(true).
		AddItem(syncButton, 0, 0, 3, 1, 0, 0, false).
		AddItem(f, 0, 1, 1, 1, 0, 0, false).
		AddItem(f2, 1, 1, 1, 1, 0, 0, false).
		AddItem(f3, 2, 1, 1, 1, 0, 0, false).
		AddItem(searchInput, 1, 2, 1, 1, 0, 0, false).
		AddItem(buttonClose, 0, 3, 3, 1, 0, 0, false).
		SetBorder(false).SetBorderStyle(tcell.StyleDefault.Background(tcell.Color102)) // .SetBorderPadding(5, 5, 5, 5)

	grid := tview.NewGrid().
		SetRows(3, 0).
		SetColumns(40, 0).
		SetBorders(true).
		AddItem(gridHeader, 0, 0, 1, 2, 0, 0, false) //.

	grid.AddItem(tree, 1, 0, 1, 1, 0, 0, true).
		AddItem(selectedList, 1, 1, 1, 1, 0, 0, false) //.

	return grid
}
