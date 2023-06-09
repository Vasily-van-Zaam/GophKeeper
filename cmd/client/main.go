package main

import (
	"context"
	"log"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/app"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	server "github.com/Vasily-van-Zaam/GophKeeper.git/internal/transport/grpc"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var states = []string{"AK", "AL", "AR", "AZ", "CA", "CO", "CT", "DC", "DE", "FL", "GA",
	"HI", "IA", "ID", "IL", "IN", "KS", "KY", "LA", "MA", "MD", "ME",
	"MI", "MN", "MO", "MS", "MT", "NC", "ND", "NE", "NH", "NJ", "NM",
	"NV", "NY", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX",
	"UT", "VA", "VT", "WA", "WI", "WV", "WY"}

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	state       string
	business    bool
}

var contacts = make([]Contact, 0)

// Tview.
var pages = tview.NewPages()
var contactText = tview.NewTextView()
var application = tview.NewApplication()
var form = tview.NewForm()
var contactsList = tview.NewList().ShowSecondaryText(false)
var flex = tview.NewFlex()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new contactq \n(q) to quit")

func main() {
	contactsList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		setConcatText(&contacts[index])
	})
	bl := tview.NewModal()

	bl.SetRect(10, 0, 1, 1)
	// bl.SetFullScreen(true)
	b := tview.NewButton("OK")
	b.SetLabel("HELLO")
	b.SetSelectedFunc(func() {
		pages.SwitchToPage("Menu")
	})
	bl.AddButtons([]string{"ss", "sss"})
	bl.GetDrawFunc()
	b.SetRect(0, 0, 10, 10)
	b.SetBackgroundColor(tcell.ColorLawnGreen)
	b.SetBorderColor(tcell.ColorLawnGreen)
	// b.SetStyle(tcell.Style{
	// 	bg: tcell.ColorLawnGreen,
	// })

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(contactsList, 0, 1, true).
			AddItem(contactText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			application.Stop()
		} else if event.Rune() == 97 {
			form.Clear(true)

			pages.SwitchToPage("Add Contact")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add Contact", app.NewFormLogin(func(form *core.LoginForm) {
		conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		c := server.NewGrpcClient(conn)
		ctx := context.Background()
		md := metadata.New(map[string]string{"client_version": "0.1.1"})
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		got, err := c.Login(ctx, &server.LoginRequest{
			Email:    form.Email,
			Password: form.Pasword,
		})

		log.Println(got, err)
	}, func() {
		pages.SwitchToPage("Menu")
	}), true, false)

	if err := application.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addContactList() {
	contactsList.Clear()
	for index, contact := range contacts {
		contactsList.AddItem(contact.firstName+" "+contact.lastName, " ", rune(49+index), nil)
	}
}

func addContactForm() *tview.Form {
	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(firstName string) {
		contact.firstName = firstName
	})

	form.AddInputField("Last Name", "", 20, nil, func(lastName string) {
		contact.lastName = lastName
	})

	form.AddInputField("Email", "", 20, nil, func(email string) {
		contact.email = email
	})

	form.AddInputField("Phone", "", 20, nil, func(phone string) {
		contact.phoneNumber = phone
	})

	form.AddDropDown("State", states, 0, func(state string, index int) {
		contact.state = state
	})

	form.AddCheckbox("Business", false, func(business bool) {
		contact.business = business
	})

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
		addContactList()
		pages.SwitchToPage("Menu")
	})

	return form
}

func setConcatText(contact *Contact) {
	contactText.Clear()
	text := contact.firstName + " " + contact.lastName + "\n" + contact.email + "\n" + contact.phoneNumber
	contactText.SetText(text)
}
