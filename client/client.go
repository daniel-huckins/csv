/*Package client -
tried to use https://github.com/sourcegraph/go-webkit2
but couldn't get it to work

trying to edit it a little bit
*/
package client

/*
#cgo pkg-config: webkit2gtk-4.0
#include "./client.c"
*/
import "C"

// a lot is taken from the tutorial
// https://wiki.gnome.org/Projects/WebKitGtk/ProgrammingGuide/Tutorial

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daniel-huckins/csv"
)

var log = csv.Logger()

// WebView sets up the gtk WebKitWebView
type WebView struct {
	view *C.struct__WebKitWebView
	win  *C.struct__GtkWidget
}

func init() {
	initGtk()
}

// NewWebView constructs a new webview
func NewWebView(width, height int) *WebView {
	view := C.newWebView()
	win := C.newGtkWindow(view, C.int(width), C.int(height))
	return &WebView{
		view: view,
		win:  win,
	}
}

// LoadURI - loads html from a url
func (w *WebView) LoadURI(uri string) error {
	cstr := (*C.gchar)(C.CString(uri))
	C.webkit_web_view_load_uri(w.view, cstr)
	return nil
}

// LoadHTMLFile - loads html from a file
// can i just pass in nil here for NULL?
func (w *WebView) LoadHTMLFile(file *os.File) error {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.WithError(err).Error("reading file")
		return err
	}
	html := fmt.Sprintf("%s", data)
	content := (*C.gchar)(C.CString(html))
	C.webkit_web_view_load_html(w.view, content, nil)
	return nil
}

// Show makes the window appear
func (w *WebView) Show() {
	C.showWindow(w.view, w.win)
	C.gtk_main()
}

func initGtk() {
	C.init_gtk(C.int(0))
}
