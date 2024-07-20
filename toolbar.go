package main

import (
	"goldwatcher/repository"
	"strconv"
	"time"

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

	dateValidator := func(s string) error {
		if _, err := time.Parse("2006-01-02", s); err != nil {
			return err
		}
		return nil
	}
	purchaseDateEntry.Validator = dateValidator

	isIntValidator := func(s string) error {
		if _, err := strconv.Atoi(s); err != nil {
			return err
		}
		return nil
	}
	addAmountEntry.Validator = isIntValidator

	isFloatValidator := func(s string) error {
		if _, err := strconv.ParseFloat(s, 32); err != nil {
			return err
		}
		return nil
	}
	purchasePriceEntry.Validator = isFloatValidator
	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	// create a dialog
	addForm := dialog.NewForm("Add Holding", "Add", "Cancel", []*widget.FormItem{
		{Text: "Amount in toz", Widget: addAmountEntry},
		{Text: "Purchase Price", Widget: purchasePriceEntry},
		{Text: "Purchase Date", Widget: purchaseDateEntry},
	}, func(valid bool) {
		if valid {
			amount, _ := strconv.Atoi(addAmountEntry.Text)
			purchaseDate, _ := time.Parse("2006-01-02", purchaseDateEntry.Text)
			purchasePrice, _ := strconv.ParseFloat(purchasePriceEntry.Text, 32)

			purchasePrice = purchasePrice * 100

			holding := repository.Holdings{Amount: amount, PurchaseDate: purchaseDate, PurchasePrice: int(purchasePrice)}

			_, err := app.DB.InsertHolding(holding)
			if err != nil {
				app.ErrorLog.Println(err)
			}

			app.refreshHoldingsTable()
		}
	}, app.MainWindow)

	//size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()
}
