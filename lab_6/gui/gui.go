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
	"sort"
)

func GUI(pg *repository.PostgresRepository) (fyne.App, fyne.Window) {
	a := app.New()
	w := a.NewWindow("Jurassic Park Database")
	w.Resize(fyne.NewSize(1200, 900))

	cardGrid := container.NewGridWithColumns(3)
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
			form.Append("", container.NewHBox(
				layout.NewSpacer(),
				container.NewVBox(
					widget.NewLabelWithStyle(col.Name, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
					entry,
				),
				layout.NewSpacer(),
			))
		}

		saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
			dialog.ShowInformation("Success", "Changes saved successfully", w)
		})
		form.Append("", container.NewVBox(saveBtn))
		return form
	}

	showEntityDetails := func(tableName string, entity map[string]interface{}) {
		content := container.NewVBox()

		keys := make([]string, 0, len(entity))
		for k := range entity {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			valueStr := fmt.Sprintf("%v", entity[key])
			content.Add(widget.NewLabelWithStyle(key, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
			content.Add(widget.NewLabel(valueStr))
			content.Add(widget.NewSeparator())
		}

		actionButtons := container.NewVBox(
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

	tables, err := repository.GetTableList(pg)
	if err != nil {
		log.Fatal(err)
	}

	var selectedTable string

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

			infoItems := []fyne.CanvasObject{}
			keys := make([]string, 0, len(e))
			for k := range e {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, key := range keys {
				label := widget.NewRichTextFromMarkdown(fmt.Sprintf("**%s**: %v", key, e[key]))
				infoItems = append(infoItems, label)
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

			actions := container.NewVBox(btnMore, btnEdit, btnDelete)

			card := container.NewVBox(
				container.NewVBox(infoItems...),
				widget.NewSeparator(),
				actions,
			)

			cardWrapped := container.New(layout.NewMaxLayout(), widget.NewCard(title, "", card))
			cardGrid.Add(cardWrapped)
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

	tableList := widget.NewList(
		func() int { return len(tables) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tables[i])
		},
	)

	tableList.OnSelected = func(id widget.ListItemID) {
		selectedTable = tables[id]
		loadTableEntities(selectedTable)
	}

	tableButtons := container.NewVBox(
		widget.NewButtonWithIcon("New Table", theme.ContentAddIcon(), func() {
			dialog.ShowInformation("New Table", "New table functionality coming soon", w)
		}),
		widget.NewButtonWithIcon("Delete Table", theme.DeleteIcon(), func() {
			if selectedTable == "" {
				dialog.ShowInformation("Error", "Please select a table first", w)
				return
			}
			dialog.ShowConfirm("Delete Table", fmt.Sprintf("Are you sure you want to delete table %s?", selectedTable), func(b bool) {
				if b {
					dialog.ShowInformation("Deleted", "Table deleted (not implemented)", w)
				}
			}, w)
		}),
		widget.NewButtonWithIcon("Export Table", theme.DownloadIcon(), func() {
			if selectedTable == "" {
				dialog.ShowInformation("Error", "Please select a table first", w)
				return
			}
			dialog.ShowInformation("Export", fmt.Sprintf("Exporting table %s to CSV", selectedTable), w)
		}),
		widget.NewButtonWithIcon("Backup All", theme.StorageIcon(), func() {
			dialog.ShowInformation("Backup", "Backing up all tables", w)
		}),
	)

	leftPanel := container.NewBorder(
		widget.NewLabelWithStyle("Tables", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		tableButtons,
		nil, nil,
		container.NewVScroll(tableList),
	)

	w.SetContent(container.NewHSplit(
		leftPanel,
		cardScroll,
	))

	return a, w
}
