package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daniel-huckins/csv"
)

var (
	log          = csv.Logger()
	manifestPath = fromPkgPath("app/build/asset-manifest.json")
)

func fromPkgPath(ext string) string {
	return fmt.Sprintf("%s/src/github.com/daniel-huckins/csv/client/%s", os.Getenv("GOPATH"), ext)
}

// the assetManifest tells us where the built files are relative to "app/build"
// https://github.com/danethurber/webpack-manifest-plugin
type assetManifest struct {
	MainCSS string `json:"main.css"`
	MainJS  string `json:"main.js"`
}

type assets struct {
	CSS string
	JS  string
}

func readManifest(fpath string) (*assetManifest, error) {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.WithError(err).Error("reading asset manifest file")
		return nil, err
	}
	manifest := new(assetManifest)
	err = json.Unmarshal(data, manifest)
	if err != nil {
		log.WithError(err).Error("unmarshalling manifest")
		return nil, err
	}
	return manifest, nil
}

func readBytes2String(fpath string) (string, error) {
	fpath = fromPkgPath(fmt.Sprintf("app/build/%s", fpath))
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.WithError(err).Errorf("error reading file '%s'", fpath)
		return "", err
	}
	return fmt.Sprintf("%s", data), nil
}

// HTML gets the complete html/js/css file to render
func HTML() (string, error) {
	manifest, err := readManifest(manifestPath)
	if err != nil {
		log.WithError(err).Error("reading manifest")
		return "", err
	}

	js, err := readBytes2String(manifest.MainJS)
	if err != nil {
		log.WithError(err).Error("reading javascript")
		return "", err
	}
	css, err := readBytes2String(manifest.MainCSS)
	if err != nil {
		log.WithError(err).Error("reading css")
		return "", err
	}

	return fmt.Sprintf(indexHTML, css, js), nil
}
