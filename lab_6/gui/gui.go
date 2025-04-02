package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"lab_6/data"
	"log"
)

func GUI() (fyne.App, fyne.Window) {
	a := app.New()
	w := a.NewWindow("Jurassic Park Database")

	tables, err := data.GetTableList()
	if err != nil {
		log.Fatal(err)
	}

	tableList := widget.NewList(
		func() int {
			return len(tables)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tables[i])
		},
	)

	// Create a hamburger menu
	menu := container.NewVBox(
		widget.NewLabel("Tables"),
		tableList,
	)

	// Set the content of the window
	w.SetContent(container.NewBorder(nil, nil, menu, nil, widget.NewLabel("Select a table from the menu")))

	// Show the window
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()

	return a, w
}
