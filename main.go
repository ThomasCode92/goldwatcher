package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var myApp Config

func main() {
	// create a fyne application
	fyneApp := app.NewWithID("dev.thomascode92.goldwatcher")
	myApp.App = fyneApp

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database

	// create a database repository

	// create and size a fyne window
	win := fyneApp.NewWindow("GoldWatcher")

	// show and run the application
	win.ShowAndRun()
}
