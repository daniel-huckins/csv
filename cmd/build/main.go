package main

import (
	"github.com/daniel-huckins/csv"
	"github.com/daniel-huckins/csv/client"
	"github.com/daniel-huckins/csv/webengine"
)

var log = csv.Logger()

func main() {
	var err error
	view := webengine.NewWebView(800, 600)
	html, err := client.HTML()
	if err != nil {
		log.WithError(err).Fatal("loading html")
	}
	err = view.LoadHTML(html)
	if err != nil {
		log.WithError(err).Fatal("loading html")
	}
	view.Show()
}
