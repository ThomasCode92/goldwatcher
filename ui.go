package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// get the current price of gold
	openPrice, currentPrice, priceChange := app.GetPriceText()

	// put price information into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
	app.PriceContainer = priceContent

	// add container to the window
	finalContent := container.NewVBox(priceContent)
	app.MainWindow.SetContent(finalContent)
}
