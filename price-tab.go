package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (app *Config) pricesTab() *fyne.Container {

}

func (app *Config) getChart() *canvas.Image {
	return nil
}

func (app *Config) downloadFile(URL, fileName string) error {
	// get the response bytes from calling a url
	resp, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("received wrong response code when downloading image")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
