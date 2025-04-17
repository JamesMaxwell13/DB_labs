package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"lab_6/repository"
	"log"
)

func GUI(pg *repository.PostgresRepository) (fyne.App, fyne.Window) {
	a := app.New()
	w := a.NewWindow("Jurassic Park Database")
	w.Resize(fyne.NewSize(1200, 900))

	cardGrid := container.NewGridWithColumns(3)
	cardGrid.Layout = layout.NewGridWrapLayout(fyne.NewSize(380, 300))
	cardScroll := container.NewVScroll(cardGrid)

	createEntityForm := func(tableName string, entity map[string]interface{}, isNew bool) *widget.Form {
		form := &widget.Form{}
		fields := make(map[string]*widget.Entry)

		columns, err := repository.GetTableColumns(pg, tableName)
		if err != nil {
			dialog.ShowError(err, w)
			return form
		}

		for _, col := range columns {
			entry := widget.NewEntry()
			if !isNew {
				if val, ok := entity[col.Name]; ok {
					entry.SetText(fmt.Sprintf("%v", val))
				}
			}
			fields[col.Name] = entry
			form.Append(col.Name, entry)
		}

		form.OnSubmit = func() {
			dialog.ShowInformation("Success", "Changes saved successfully", w)
		}

		return form
	}

	showEntityDetails := func(tableName string, entity map[string]interface{}) {
		content := container.NewVBox()

		for key, value := range entity {
			valueStr := fmt.Sprintf("%v", value)
			content.Add(widget.NewLabelWithStyle(key, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
			content.Add(widget.NewLabel(valueStr))
			content.Add(widget.NewSeparator())
		}

		actionButtons := container.NewHBox(
			widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
				editWindow := a.NewWindow("Edit Entity")
				editWindow.Resize(fyne.NewSize(600, 400))
				editWindow.SetContent(createEntityForm(tableName, entity, false))
				editWindow.Show()
			}),
			widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
				dialog.ShowConfirm("Delete", "Are you sure you want to delete this entity?", func(b bool) {
					if b {
						dialog.ShowInformation("Deleted", "Entity deleted successfully", w)
					}
				}, w)
			}),
		)

		mainContent := container.NewVBox(
			widget.NewLabelWithStyle("Entity Details", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewSeparator(),
			content,
			actionButtons,
		)

		detailsWindow := a.NewWindow("Entity Details")
		detailsWindow.SetContent(mainContent)
		detailsWindow.Show()
	}

	loadTableEntities := func(tableName string) {
		entities, err := repository.GetTableEntities(pg, tableName)
		if err != nil {
			log.Printf("Error getting entities from table %s: %v", tableName, err)
			return
		}

		cardGrid.Objects = nil

		for _, entity := range entities {
			e := entity
			title := ""
			if name, ok := e["name"]; ok {
				title = fmt.Sprintf("%v", name)
			} else if id, ok := e["id"]; ok {
				title = fmt.Sprintf("ID: %v", id)
			} else {
				title = "Entity"
			}

			info := container.NewVBox()
			for key, val := range e {
				info.Add(widget.NewLabel(fmt.Sprintf("%s: %v", key, val)))
			}

			btnMore := widget.NewButtonWithIcon("More", theme.MoreHorizontalIcon(), func() {
				showEntityDetails(tableName, e)
			})
			btnEdit := widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
				editWindow := a.NewWindow("Edit Entity")
				editWindow.SetContent(createEntityForm(tableName, e, false))
				editWindow.Resize(fyne.NewSize(600, 400))
				editWindow.Show()
			})
			btnDelete := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
				dialog.ShowConfirm("Delete", "Are you sure you want to delete this entity?", func(b bool) {
					if b {
						dialog.ShowInformation("Deleted", "Entity deleted (placeholder)", w)
					}
				}, w)
			})

			card := widget.NewCard(title, "", container.NewVBox(
				info,
				container.NewHBox(btnMore, btnEdit, btnDelete),
			))

			cardGrid.Add(card)
		}

		cardGrid.Add(widget.NewCard("", "",
			container.NewCenter(widget.NewButtonWithIcon("New Entity", theme.ContentAddIcon(), func() {
				newEntityWindow := a.NewWindow("New Entity")
				newEntityWindow.SetContent(createEntityForm(tableName, nil, true))
				newEntityWindow.Resize(fyne.NewSize(600, 400))
				newEntityWindow.Show()
			})),
		))

		cardGrid.Refresh()
	}

	tables, err := repository.GetTableList(pg)
	if err != nil {
		log.Fatal(err)
	}

	tableList := widget.NewList(
		func() int { return len(tables) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tables[i])
		},
	)
	tableScroll := container.NewVScroll(tableList)

	var selectedTable string
	tableList.OnSelected = func(id widget.ListItemID) {
		selectedTable = tables[id]
		loadTableEntities(selectedTable)
	}

	tableButtons := container.NewVBox()
	buttons := []struct {
		label  string
		icon   fyne.Resource
		action func()
	}{
		{"Query", theme.DocumentIcon(), func() {
			queryWindow := a.NewWindow("SQL Query")
			queryWindow.Resize(fyne.NewSize(600, 400))
			queryEntry := widget.NewMultiLineEntry()
			queryEntry.SetPlaceHolder("Enter your SQL query here...")
			executeBtn := widget.NewButton("Execute", func() {
				dialog.ShowInformation("Query", "Query execution will be implemented here", queryWindow)
			})
			queryWindow.SetContent(container.NewVBox(widget.NewLabel("SQL Query:"), queryEntry, executeBtn))
			queryWindow.Show()
		}},
		{"New Table", theme.ContentAddIcon(), func() {
			dialog.ShowInformation("New Table", "New table functionality will be implemented here", w)
		}},
		{"Delete Table", theme.DeleteIcon(), func() {
			if selectedTable == "" {
				dialog.ShowInformation("Error", "Please select a table first", w)
				return
			}
			dialog.ShowConfirm("Delete Table", fmt.Sprintf("Are you sure you want to delete table %s?", selectedTable), func(b bool) {
				if b {
					dialog.ShowInformation("Deleted", "Table deleted (not really, just a mock)", w)
				}
			}, w)
		}},
		{"Export Table", theme.DownloadIcon(), func() {
			if selectedTable == "" {
				dialog.ShowInformation("Error", "Please select a table first", w)
				return
			}
			dialog.ShowInformation("Export", fmt.Sprintf("Exporting table %s to CSV", selectedTable), w)
		}},
		{"Backup Table", theme.StorageIcon(), func() {
			if selectedTable == "" {
				dialog.ShowInformation("Error", "Please select a table first", w)
				return
			}
			dialog.ShowInformation("Backup", fmt.Sprintf("Backing up table %s", selectedTable), w)
		}},
		{"Backup All", theme.StorageIcon(), func() {
			dialog.ShowInformation("Backup All", "Backing up all tables", w)
		}},
	}

	for _, btn := range buttons {
		button := widget.NewButtonWithIcon(btn.label, btn.icon, btn.action)
		button.Alignment = widget.ButtonAlignLeading
		tableButtons.Add(button)
	}

	leftPanel := container.NewBorder(
		widget.NewLabelWithStyle("Tables", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		tableButtons,
		nil, nil,
		tableScroll,
	)

	split := container.NewHSplit(leftPanel, cardScroll)
	split.Offset = 0.25

	w.SetContent(split)
	return a, w
}
