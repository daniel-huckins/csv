package main

import (
	"github.com/daniel-huckins/csv"
	"github.com/daniel-huckins/csv/client"
)

var log = csv.Logger()

func main() {
	var err error
	err = client.Main()
	if err != nil {
		log.WithError(err).Fatal("")
	}
}
