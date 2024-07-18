package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), app.addHoldingsDialog),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), app.refreshPriceContent),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return toolbar
}

func (app *Config) addHoldingsDialog() {
	addAmountEntry := widget.NewEntry()
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()

	app.AddHoldingsPurchaseAmountEntry = addAmountEntry
	app.AddHoldingsPurchaseDateEntry = purchaseDateEntry
	app.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	// create a dialog
	addForm := dialog.NewForm("Add Holding", "Add", "Cancel", []*widget.FormItem{
		{Text: "Amount in toz", Widget: addAmountEntry},
		{Text: "Purchase Price", Widget: purchasePriceEntry},
		{Text: "Purchase Date", Widget: purchaseDateEntry},
	}, func(valid bool) {
		if valid {

		}
	}, app.MainWindow)

	//size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()
}
