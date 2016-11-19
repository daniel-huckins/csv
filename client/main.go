package client

import (
	"github.com/daniel-huckins/csv/webengine"
)

// Main loads and runs the client
func Main() (err error) {
	engine := webengine.NewWebView(800, 600)
	html, err := HTML()
	if err != nil {
		log.WithError(err).Error("creating html")
		return
	}
	engine.LoadHTML(html)

	engine.Show()

	return
}
