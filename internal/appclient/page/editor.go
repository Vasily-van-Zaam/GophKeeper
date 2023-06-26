package page

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/appclient/component"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type editorPage struct {
	data   core.Manager
	client applicationClient
	name   string
	reset  func()
	back   func()
	next   func(user *core.User)
	conf   config.Config
}

// Back implements AppPage.
func (e *editorPage) Back(back func()) AppPage {
	e.back = back
	return e
}

// Close implements AppPage.
func (e *editorPage) Close(bool) AppPage {
	e.client.Pages().RemovePage(e.name)
	return e
}

// Next implements AppPage.
func (e *editorPage) Next(next func(user *core.User)) AppPage {
	e.next = next
	return e
}

// Reset implements AppPage.
func (e *editorPage) Reset(reset func()) AppPage {
	e.reset = reset
	return e
}

func (e *editorPage) redraw(form tview.Primitive) {
	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
		AddText(e.name, true, 1, tcell.ColorRed).
		SetPrimitive(form).SetBorders(0, 2, 0, 2, 1, 1)

	e.client.Pages().AddPage(e.name, frame, true, true)
}

// Show implements AppPage.
func (e *editorPage) Show(ctx context.Context, show bool) AppPage {
	if !show {
		return e
	}
	var formEdit tview.Primitive

	switch e.data.Get().InfoData().DataType {
	case string(core.DataTypeCard):
		var fCard = core.BankCardForm{}
		id := e.data.Get().InfoData().ID

		if id != nil {
			password, err := e.data.AddEncription(e.conf.Encryptor()).Get().Data(e.client.User().PrivateKey)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return e
			}
			_ = json.Unmarshal(password, &fCard)
		}
		formCard := component.NewFormCardBanks()
		formEdit = formCard.EditForm(&fCard)
		formCard.HandlerClose(func() {
			e.back()
		})
		formCard.HandlerSave(func(form *core.BankCardForm) {
			e.data.AddEncription(e.conf.Encryptor())
			err := e.data.Set().BankCard(e.client.User().PrivateKey, form)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			data, err := e.data.ToData()
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			if id == nil {
				_, err = e.client.Repository().Local().AddData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			} else {
				_, err = e.client.Repository().Local().ChangeData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			}
			e.back()
		})
		formCard.HandlerError(func(msg string) {
			e.back()
		})
	case string(core.DataTypePassword):
		var fPsw = core.PasswordForm{}
		id := e.data.Get().InfoData().ID

		if id != nil {
			password, err := e.data.AddEncription(e.conf.Encryptor()).Get().Data(e.client.User().PrivateKey)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return e
			}
			_ = json.Unmarshal(password, &fPsw)
		}
		formPsw := component.NewFormPassword()
		formEdit = formPsw.EditForm(&fPsw)
		formPsw.HandlerClose(func() {
			e.back()
		})
		formPsw.HandlerSave(func(form *core.PasswordForm) {
			e.data.AddEncription(e.conf.Encryptor())
			err := e.data.Set().Password(e.client.User().PrivateKey, form)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			data, err := e.data.ToData()
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			if id == nil {
				_, err = e.client.Repository().Local().AddData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			} else {
				_, err = e.client.Repository().Local().ChangeData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			}
			e.back()
		})
		formPsw.HandlerError(func(msg string) {
			e.back()
		})
	case string(core.DataTypeText):
		var fText = core.TextFomm{}
		id := e.data.Get().InfoData().ID

		textData := e.data.AddEncription(e.conf.Encryptor())

		if id != nil {
			tb, err := textData.Get().Data(e.client.User().PrivateKey)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return e
			}
			_ = json.Unmarshal(tb, &fText)
		}
		formEdit = component.NewFormTextEditor(&fText, func(text *core.TextFomm) {
			err := textData.Set().Text(e.client.User().PrivateKey, text)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			data, err := textData.ToData()
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			if id == nil {
				_, err = e.client.Repository().Local().AddData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			} else {
				_, err = e.client.Repository().Local().ChangeData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			}
			e.back()
		}, func() {
			e.back()
		})
	case string(core.DataTypeFile):
		var fFile = core.FileForm{}
		id := e.data.Get().InfoData().ID

		if id != nil {
			password, err := e.data.AddEncription(e.conf.Encryptor()).Get().Data(e.client.User().PrivateKey)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return e
			}
			_ = json.Unmarshal(password, &fFile)
		}

		formFile := component.NewFormFile()
		formEdit = formFile.EditForm(&fFile)
		formFile.HandlerClose(func() {
			e.back()
		})

		formFile.HandlerSave(func(form *core.FileForm) {
			e.data.AddEncription(e.conf.Encryptor())
			err := e.data.Set().File(e.client.User().PrivateKey, form)
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			data, err := e.data.ToData()
			if err != nil {
				component.ModalError(err, e.name, e.client.Pages())
				return
			}
			if id == nil {
				_, err = e.client.Repository().Local().AddData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			} else {
				_, err = e.client.Repository().Local().ChangeData(ctx, data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
			}

			e.back()
		})

		formFile.HandlerUploadFile(func(form *core.FileForm) {
			var (
				file *os.File
				err  error
			)
			component.ModalNewUploadFile("Upload new file", e.name, e.client.Pages(), func(path string) {
				file, err = os.Open(path)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
				defer file.Close()

				// Get the file size
				stat, errSt := file.Stat()
				if errSt != nil {
					component.ModalError(errSt, e.name, e.client.Pages())
					return
				}
				size := stat.Size()

				if size > 2560000 {
					component.ModalError(errors.New("File cannot exceed 2.56MB"), e.name, e.client.Pages())
					return
				}
				// Read the file into a byte slice
				bs := make([]byte, size)
				_, err = bufio.NewReader(file).Read(bs)
				if err != nil && err != io.EOF {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
				form.Name = file.Name()
				form.Data = bs
				form.Size = size
				e.client.Pages().RemovePage(e.name)
				formEdit = formFile.EditForm(form)
				e.redraw(formEdit)
			})
		})
		formFile.HandlerDownloadFile(func(form *core.FileForm) {
			component.ModalNewDownloadFile("Save file", form.Name, e.name, e.client.Pages(), func(name, path string) {
				file, err := os.Create(path + "/" + name)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
				}
				defer file.Close()
				_, err = file.Write(form.Data)
				if err != nil {
					component.ModalError(err, e.name, e.client.Pages())
					return
				}
				component.ModalError(errors.New("✅ OK - fale saved"), e.name, e.client.Pages())
			})
		})

	default:
		formEdit = tview.NewForm()
	}

	e.redraw(formEdit)
	return e
}

func NewEditorPage(
	conf config.Config,
	client applicationClient,
	data core.Manager,
	pageName string,
) AppPage {
	return &editorPage{
		data:   data,
		client: client,
		name:   pageName,
		conf:   conf,
	}
}
