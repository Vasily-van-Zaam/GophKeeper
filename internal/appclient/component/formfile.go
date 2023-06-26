package component

import (
	"strconv"

	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type formFile struct {
	metaData string
	data     []byte
	name     string
	size     int64
	closed   func()
	err      func(msg string)
	saved    func(*core.FileForm)
	upload   func(form *core.FileForm)
	download func(form *core.FileForm)
}

// HandlerDownloadFile implements ComponentFormFile.
func (f *formFile) HandlerDownloadFile(h func(form *core.FileForm)) componentFormFile {
	f.download = h
	return f
}

// HandlerUploadFile implements ComponentFormFile.
func (f *formFile) HandlerUploadFile(h func(form *core.FileForm)) componentFormFile {
	f.upload = h
	return f
}

// HandlerSave implements ComponentFormPassword.
func (f *formFile) HandlerSave(h func(form *core.FileForm)) componentFormFile {
	f.saved = h
	return f
}

// HandlerClose implements ComponentFormPassword.
func (f *formFile) HandlerClose(h func()) componentFormFile {
	f.closed = h
	return f
}

// HandlerError implements ComponentFormPassword.
func (f *formFile) HandlerError(h func(msg string)) componentFormFile {
	f.err = h
	return f
}

// Handler implements ComponentFormPassword.
func (f *formFile) EditForm(frm *core.FileForm) *tview.Form {
	f.data = frm.Data
	f.metaData = frm.MetaData
	f.name = frm.Name
	f.size = frm.Size
	form := tview.NewForm().
		AddInputField("Metadata", f.metaData, 1000,
			func(textToCheck string, lastChar rune) bool {
				return true
			}, func(text string) {
				f.metaData = text
			}).
		AddInputField("name", f.name, 1000,
			func(textToCheck string, lastChar rune) bool {
				return false
			}, func(text string) {
				f.name = text
			}).
		AddInputField("size file", core.SizeToSize(f.size), 1000,
			func(textToCheck string, lastChar rune) bool {
				return false
			}, func(text string) {
				d, _ := strconv.ParseInt(text, 10, 64)
				f.size = d
			}).
		AddButton("❌ Back", func() { f.closed() }).
		AddButton("📎 Upload new file", func() {
			frm.MetaData = f.metaData
			f.upload(frm)
		}).
		AddButton("✅ Save", func() {
			frm.MetaData = f.metaData
			frm.Name = f.name
			frm.Data = f.data
			f.saved(frm)
		}).SetButtonStyle(tcell.StyleDefault.Attributes(tcell.AttrBold))
	if frm.Name != "" {
		form.AddButton("💾 Download current file", func() {
			frm.MetaData = f.metaData
			frm.Name = f.name
			f.download(frm)
		})
	}

	return form
}

type componentFormFile interface {
	EditForm(form *core.FileForm) *tview.Form
	HandlerError(func(msg string)) componentFormFile
	HandlerClose(func()) componentFormFile
	HandlerSave(func(form *core.FileForm)) componentFormFile
	HandlerUploadFile(func(form *core.FileForm)) componentFormFile
	HandlerDownloadFile(func(form *core.FileForm)) componentFormFile
}

func NewFormFile() componentFormFile {
	return &formFile{}
}
