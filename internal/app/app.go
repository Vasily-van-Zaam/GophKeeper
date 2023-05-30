// Application Client
// project GophKeeper Yandex Practicum
// used package Tview
// Created by Vasiliy Van-Zaam
package app

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Tview variables.
var (
	pages        = tview.NewPages()
	contactText  = tview.NewTextView()
	app          = tview.NewApplication()
	form         = tview.NewForm()
	contactsList = tview.NewList().ShowSecondaryText(false)
	flex         = tview.NewFlex()
	text         = tview.NewTextView().
			SetTextColor(tcell.ColorGreen).
			SetText("(a) to add a new contactq \n(q) to quit")
)

type client struct {
	Pages *tview.Pages
	App   *tview.Application
}

// Create new Client.
func NewClient() error {
	// f := form.AddButton("hello", func() {})
	// p1 := pages.AddPage("h", f, false, false)
	_ = core.User{}
	return nil
}
