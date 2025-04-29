package gui

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"lab_6/repository"
)

func ShowEntityDetails(pg *repository.PostgresRepository, tableName string, entity map[string]interface{}, onUpdate func()) fyne.Window {
	title := getEntityTitle(entity)
	w := fyne.CurrentApp().NewWindow(fmt.Sprintf("%s: %s", tableName, title))
	content := container.NewVBox()
	for _, key := range sortEntityKeys(entity) {
		content.Add(widget.NewLabelWithStyle(key, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		content.Add(widget.NewLabel(fmt.Sprintf("%v", entity[key])))
		content.Add(widget.NewSeparator())
	}
	actions := container.NewVBox(
		widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
			editW := fyne.CurrentApp().NewWindow("Edit Entity")
			editW.SetContent(CreateEntityForm(pg, tableName, entity, false, func() {
				onUpdate()
				editW.Close()
			}))
			editW.Resize(fyne.NewSize(600, 400))
			editW.Show()
		}),
		widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
			dialog.ShowConfirm("Delete", "Are you sure you want to delete this entity?", func(b bool) {
				if b {
					err := repository.DeleteEntity(pg, tableName, entity)
					if err == nil {
						onUpdate()
						w.Close()
						dialog.ShowInformation("Deleted", "Entity deleted successfully", nil)
					}
				}
			}, w)
		}),
	)
	w.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Entity Details", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		content,
		actions,
	))
	w.Resize(fyne.NewSize(400, w.Canvas().Size().Height))
	return w
}

func CreateTableForm(pg *repository.PostgresRepository, w fyne.Window) fyne.CanvasObject {
	form := container.NewVBox()
	tableNameEntry := widget.NewEntry()
	tableNameEntry.SetPlaceHolder("Table name")
	form.Add(container.NewVBox(
		widget.NewLabelWithStyle("Table Name", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		tableNameEntry,
	))

	fieldRows := container.NewVBox()
	addField := func() {
		fieldNameEntry := widget.NewEntry()
		typeSelect := widget.NewSelect([]string{"INTEGER", "TEXT", "BOOLEAN", "DATE", "SERIAL"}, nil)
		row := container.NewGridWithColumns(2,
			container.NewVBox(widget.NewLabel("Name of column"), fieldNameEntry),
			container.NewVBox(widget.NewLabel("Type"), typeSelect),
		)
		fieldRows.Add(row)
	}
	addField()
	form.Add(fieldRows)
	form.Add(widget.NewButtonWithIcon("Add Field", theme.ContentAddIcon(), addField))

	form.Add(widget.NewButtonWithIcon("Create Table", theme.DocumentSaveIcon(), func() {
		tableName := tableNameEntry.Text
		if tableName == "" {
			dialog.ShowError(fmt.Errorf("table name cannot be empty"), w)
			return
		}
		var fieldDefs []string
		for _, obj := range fieldRows.Objects {
			row := obj.(*fyne.Container)
			nameEntry := row.Objects[0].(*fyne.Container).Objects[1].(*widget.Entry)
			typeSelect := row.Objects[1].(*fyne.Container).Objects[1].(*widget.Select)
			if nameEntry.Text == "" || typeSelect.Selected == "" {
				dialog.ShowError(fmt.Errorf("all fields must have name and type"), w)
				return
			}
			fieldDefs = append(fieldDefs, fmt.Sprintf("%s %s", nameEntry.Text, typeSelect.Selected))
		}
		if len(fieldDefs) == 0 {
			dialog.ShowError(fmt.Errorf("at least one field is required"), w)
			return
		}
		query := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(fieldDefs, ", "))
		_, err := pg.Db.Query(query)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation("Success", fmt.Sprintf("Table %s created", tableName), w)
	}))
	return container.NewVBox(
		widget.NewLabelWithStyle("Create New Table", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
	)
}

// GUI — основная точка входа
const maxTextLength = 18
const maxCardFields = 4

func truncateText(text string) string {
	runes := []rune(text)
	if len(runes) > maxTextLength {
		return string(runes[:maxTextLength-1]) + "…"
	}
	return text
}

func sortEntityKeys(entity map[string]interface{}) []string {
	keys := make([]string, 0, len(entity))
	for k := range entity {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	priority := map[string]int{"id": -2, "name": -1}
	sort.SliceStable(keys, func(i, j int) bool {
		pi, pj := priority[keys[i]], priority[keys[j]]
		if pi != pj {
			return pi < pj
		}
		return keys[i] < keys[j]
	})
	return keys
}

func getEntityTitle(e map[string]interface{}) string {
	if name, ok := e["name"]; ok {
		return truncateText(fmt.Sprintf("%v", name))
	} else if id, ok := e["id"]; ok {
		return fmt.Sprintf("%v", id)
	}
	return "Entity"
}

func getEntitySortKey(e map[string]interface{}) string {
	if name, ok := e["name"]; ok {
		return fmt.Sprintf("%v", name)
	} else if id, ok := e["id"]; ok {
		if num, err := strconv.Atoi(fmt.Sprintf("%v", id)); err == nil {
			return fmt.Sprintf("%020d", num)
		}
		return fmt.Sprintf("%v", id)
	}
	return ""
}

func CreateEntityForm(pg *repository.PostgresRepository, tableName string, entity map[string]interface{}, isNew bool, onSuccess func()) fyne.CanvasObject {
	form := container.NewVBox()
	fields := make(map[string]*widget.Entry)
	columns, err := repository.GetTableColumns(pg, tableName)
	if err != nil {
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
		form.Add(container.NewVBox(
			widget.NewLabelWithStyle(col.Name, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			entry,
		))
	}
	saveBtn := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		values := make(map[string]interface{})
		for col, entry := range fields {
			values[col] = entry.Text
		}
		var err error
		if isNew {
			err = repository.InsertEntity(pg, tableName, values)
		} else {
			err = repository.UpdateEntity(pg, tableName, entity, values)
		}
		if err == nil && onSuccess != nil {
			onSuccess()
			dialog.ShowInformation("Success", "Changes saved successfully", nil)
		}
	})
	return container.NewBorder(nil, saveBtn, nil, nil, container.NewVScroll(form))
}

func LoadTableEntities(pg *repository.PostgresRepository, tableName string, cardGrid *fyne.Container, w fyne.Window, a fyne.App, onEdit func(func()), onDelete func(func())) {
	entities, err := repository.GetTableEntities(pg, tableName)
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	cardGrid.Objects = nil
	sort.SliceStable(entities, func(i, j int) bool {
		return getEntitySortKey(entities[i]) < getEntitySortKey(entities[j])
	})
	for _, entity := range entities {
		e := entity
		title := getEntityTitle(e)
		info := container.NewVBox()
		keys := sortEntityKeys(e)
		for i, key := range keys {
			if i >= maxCardFields {
				break
			}
			text := fmt.Sprintf("%s: %v", key, e[key])
			label := widget.NewLabel(text)
			label.Wrapping = fyne.TextTruncate
			info.Add(label)
		}
		btnMore := widget.NewButtonWithIcon("More", theme.MoreHorizontalIcon(), func() {
			ShowEntityDetails(pg, tableName, e, func() {
				LoadTableEntities(pg, tableName, cardGrid, w, a, onEdit, onDelete)
			}).Show()
		})
		btnEdit := widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
			editW := a.NewWindow("Edit Entity")
			editW.SetContent(CreateEntityForm(pg, tableName, e, false, func() {
				LoadTableEntities(pg, tableName, cardGrid, w, a, onEdit, onDelete)
				editW.Close()
			}))
			editW.Resize(fyne.NewSize(600, 400))
			editW.Show()
		})
		btnDelete := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
			dialog.ShowConfirm("Delete", "Are you sure you want to delete this entity?", func(b bool) {
				if b {
					err := repository.DeleteEntity(pg, tableName, e)
					if err == nil {
						LoadTableEntities(pg, tableName, cardGrid, w, a, onEdit, onDelete)
						dialog.ShowInformation("Deleted", "Entity deleted successfully", w)
					}
				}
			}, w)
		})
		actions := container.NewVBox(btnMore, btnEdit, btnDelete)
		card := widget.NewCard(title, "", container.NewVBox(info, widget.NewSeparator(), actions))
		cardGrid.Add(card)
	}
	cardGrid.Refresh()
}

func GUI(pg *repository.PostgresRepository) (fyne.App, fyne.Window) {
	a := app.New()
	w := a.NewWindow("Jurassic Park Database")
	w.Resize(fyne.NewSize(1200, 900))

	cardGrid := container.NewGridWithColumns(3)
	cardScroll := container.NewVScroll(cardGrid)

	var selectedTable string
	tables, err := repository.GetTableList(pg)
	if err != nil {
		dialog.ShowError(err, w)
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
		LoadTableEntities(pg, selectedTable, cardGrid, w, a, nil, nil)
	}

	tableButtons := container.NewVBox(
		widget.NewButtonWithIcon("New Table", theme.ContentAddIcon(), func() {
			newTableWindow := a.NewWindow("Create Table")
			newTableWindow.SetContent(CreateTableForm(pg, newTableWindow))
			newTableWindow.SetOnClosed(func() {
				tables, _ = repository.GetTableList(pg)
				tableList.Refresh()
			})
			newTableWindow.Show()
		}),
		widget.NewButtonWithIcon("Delete Table", theme.DeleteIcon(), func() {
			if selectedTable == "" {
				dialog.ShowInformation("Error", "Please select a table first", w)
				return
			}
			dialog.ShowConfirm("Delete Table", fmt.Sprintf("Are you sure you want to delete table %s?", selectedTable), func(b bool) {
				if b {
					err := repository.DeleteTable(pg, selectedTable)
					if err == nil {
						selectedTable = ""
						tables, _ = repository.GetTableList(pg)
						tableList.Refresh()
						cardGrid.Objects = nil
						cardGrid.Refresh()
						dialog.ShowInformation("Deleted", "Table deleted", w)
					}
				}
			}, w)
		}),
		widget.NewButton("Export Table", func() {
			dialog.ShowInformation("Not implemented", "Export Table clicked", w)
		}),
		widget.NewButton("Export Database", func() {
			dialog.ShowInformation("Not implemented", "Export DB clicked", w)
		}),
		widget.NewButton("Backup Database", func() {
			dialog.ShowInformation("Not implemented", "Backup DB clicked", w)
		}),
		widget.NewButton("Run Query", func() {
			dialog.ShowInformation("Not implemented", "Query clicked", w)
		}),
	)

	leftPanel := container.NewBorder(
		widget.NewLabelWithStyle("Tables", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewVBox(tableButtons), nil, nil,
		container.NewVScroll(tableList),
	)

	split := container.NewHSplit(leftPanel, cardScroll)
	split.Offset = 0.25
	w.SetContent(split)
	return a, w
}
