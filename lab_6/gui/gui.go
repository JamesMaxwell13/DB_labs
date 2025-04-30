package gui

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"lab_6/repository"
)

const (
	maxTextLength = 18
	maxCardFields = 4
)

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

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case []byte:
		if strings.HasPrefix(string(v), "$") {
			return string(v)
		}
		return string(v)
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	default:
		return fmt.Sprintf("%v", v)
	}
}

func CreateEntityForm(pg *repository.PostgresRepository, tableName string, entity map[string]interface{}, isNew bool, onSuccess func()) fyne.CanvasObject {
	form := container.NewVBox()
	fields := make(map[string]*widget.Entry)
	columns, err := repository.GetTableColumns(pg, tableName)
	if err != nil {
		log.Println(err)
		return form
	}

	for _, col := range columns {
		entry := widget.NewEntry()
		if !isNew {
			if val, ok := entity[col.Name]; ok {
				entry.SetText(formatValue(val))
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
		if err != nil {
			dialog.ShowError(err, nil)
			log.Println(err)
		} else if onSuccess != nil {
			log.Println("entity saved successfully")
			onSuccess()
		}
	})

	return container.NewBorder(nil, saveBtn, nil, nil, container.NewVScroll(form))
}

func ShowEntityDetails(pg *repository.PostgresRepository, tableName string, entity map[string]interface{}, onUpdate func()) fyne.Window {
	title := getEntityTitle(entity)
	w := fyne.CurrentApp().NewWindow(fmt.Sprintf("%s: %s", tableName, title))
	icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
	w.SetIcon(icon)
	l
	content := container.NewVBox()
	for _, key := range sortEntityKeys(entity) {
		content.Add(widget.NewLabelWithStyle(key, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
		content.Add(widget.NewLabel(formatValue(entity[key])))
		content.Add(widget.NewSeparator())
	}

	scrollContent := container.NewVScroll(content)

	actions := container.NewVBox(
		widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
			editW := fyne.CurrentApp().NewWindow("Edit Entity")
			editW.SetContent(CreateEntityForm(pg, tableName, entity, false, func() {
				onUpdate()
				editW.Close()
				if w := fyne.CurrentApp().Driver().AllWindows()[0]; w != nil {
					dialog.ShowInformation("Success", "Entity updated successfully", w)
				}
			}))
			icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
			editW.SetIcon(icon)
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
						if parent := fyne.CurrentApp().Driver().AllWindows()[0]; parent != nil {
							log.Println("entity deleted successfully")
							dialog.ShowInformation("Deleted", "Entity deleted successfully", parent)
						}
					} else {
						dialog.ShowError(err, w)
						log.Println(err)
					}
				}
			}, w)
		}),
	)

	mainContent := container.NewBorder(
		container.NewVBox(
			widget.NewLabelWithStyle("Entity Details", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewSeparator(),
		),
		actions,
		nil, nil,
		scrollContent,
	)

	w.SetContent(mainContent)
	w.Resize(fyne.NewSize(400, 600))
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
		typeSelect := widget.NewSelect([]string{"INTEGER", "REAL", "TEXT", "BOOLEAN", "DATE", "TIMESTAMP", "MONEY"}, nil)
		typeSelect.PlaceHolder = "Select type"

		row := container.NewGridWithColumns(2,
			container.NewVBox(
				widget.NewLabel("Column Name"),
				fieldNameEntry,
			),
			container.NewVBox(
				widget.NewLabel("Data Type"),
				typeSelect,
			),
		)
		fieldRows.Add(row)
	}
	addField()
	form.Add(fieldRows)

	buttons := container.NewVBox(
		widget.NewButtonWithIcon("Add Column", theme.ContentAddIcon(), addField),
		widget.NewButtonWithIcon("Create Table", theme.DocumentSaveIcon(), func() {
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
					dialog.ShowError(fmt.Errorf("all columns must have name and type"), w)
					log.Println("all columns must have name and type")
					return
				}
				fieldDefs = append(fieldDefs, fmt.Sprintf("%s %s", nameEntry.Text, typeSelect.Selected))
			}

			if len(fieldDefs) == 0 {
				dialog.ShowError(fmt.Errorf("at least one column is required"), w)
				log.Println("at least one column is required")
				return
			}

			query := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(fieldDefs, ", "))
			_, err := pg.Db.Exec(query)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to create table: %w", err), w)
				log.Println("failed to create table")
				return
			}

			dialog.ShowInformation("Success", fmt.Sprintf("Table %s created", tableName), w)
			w.Close()
		}),
	)

	return container.NewBorder(
		nil,                        // Верх
		buttons,                    // Низ (две кнопки)
		nil,                        // Лево
		nil,                        // Право
		container.NewVScroll(form), // Центр (форма с прокруткой)
	)
}

func CreateQueryForm(pg *repository.PostgresRepository, w fyne.Window) fyne.CanvasObject {
	queryNames, err := repository.GetQueryNames()
	if err != nil {
		err := fmt.Errorf("error getting query names: %w", err)
		dialog.ShowError(err, w)
		log.Println(err)
		return widget.NewLabel("Failed to load queries")
	}

	querySelect := widget.NewSelect(queryNames, nil)
	queryEntry := widget.NewMultiLineEntry()
	queryEntry.SetPlaceHolder("Enter SQL query here...")
	queryEntry.Wrapping = fyne.TextWrapWord

	querySelect.OnChanged = func(name string) {
		if sql, err := repository.GetQueryByName(name); err == nil {
			queryEntry.SetText(sql)
		}
	}

	saveBtn := widget.NewButton("Save Query", func() {
		if queryEntry.Text == "" {
			err = fmt.Errorf("no query to save")
			dialog.ShowInformation("Error", err.Error(), w)
			log.Println(err)
			return
		}
		nameEntry := widget.NewEntry()
		nameEntry.SetPlaceHolder("Query name")
		dialog.ShowForm("Save Query", "Save", "Cancel", []*widget.FormItem{
			widget.NewFormItem("Name", nameEntry),
		}, func(b bool) {
			if b && nameEntry.Text != "" {
				if err := repository.SaveQuery(nameEntry.Text, queryEntry.Text); err != nil {
					dialog.ShowError(err, w)
					log.Println(err)
				} else {
					dialog.ShowInformation("Success", "Query saved", w)
					if names, err := repository.GetQueryNames(); err == nil {
						querySelect.Options = names
						querySelect.Refresh()
					}
				}
			}
		}, w)
	})

	deleteBtn := widget.NewButton("Delete Query", func() {
		if querySelect.Selected == "" {
			dialog.ShowInformation("Error", "No query selected to delete", w)
			return
		}

		dialog.ShowConfirm("Delete Query",
			fmt.Sprintf("Are you sure you want to delete query '%s'?", querySelect.Selected),
			func(b bool) {
				if b {
					err := repository.DeleteQuery(querySelect.Selected)
					if err != nil {
						dialog.ShowError(err, w)
						log.Println(err)
						return
					}

					if names, err := repository.GetQueryNames(); err == nil {
						querySelect.Options = names
						querySelect.Selected = ""
						querySelect.Refresh()
						queryEntry.SetText("")
					}
					dialog.ShowInformation("Success", "Query deleted", w)
				}
			}, w)
	})

	runBtn := widget.NewButton("Run Query", func() {
		if queryEntry.Text == "" {
			err := fmt.Errorf("no query to save")
			dialog.ShowInformation("Error", err.Error(), w)
			log.Println(err)
			return
		}
		results, err := repository.ExecuteQuery(pg, queryEntry.Text)
		if err != nil {
			dialog.ShowError(err, w)
			log.Println(err)
			return
		}
		resultWindow := fyne.CurrentApp().NewWindow("Query Results")
		cardGrid := container.NewGridWithColumns(3)
		cardScroll := container.NewVScroll(cardGrid)

		exportBtn := widget.NewButtonWithIcon("Export to Excel", theme.DownloadIcon(), func() {
			dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					log.Println(err)
					return
				}
				if writer == nil {
					return
				}
				defer writer.Close()

				if err := repository.ExportQueryResultsToExcel(results, writer); err != nil {
					dialog.ShowError(err, w)
					log.Println(err)
					return
				}

				dialog.ShowInformation("Success", "Results exported successfully", w)
			}, resultWindow)
		})

		resultWindow.SetContent(container.NewBorder(
			exportBtn,
			nil,
			nil,
			nil,
			cardScroll,
		))

		icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
		resultWindow.SetIcon(icon)
		resultWindow.Resize(fyne.NewSize(1200, 800))

		if len(results) > 0 {
			for _, entity := range results {
				e := entity
				title := getEntityTitle(e)
				info := container.NewVBox()
				keys := sortEntityKeys(e)
				for _, key := range keys {
					text := fmt.Sprintf("%s: %v", key, formatValue(e[key]))
					label := widget.NewLabel(text)
					label.Wrapping = fyne.TextTruncate
					info.Add(label)
				}
				card := widget.NewCard(title, "", info)
				cardGrid.Add(card)
			}
		} else {
			cardGrid.Add(widget.NewLabel("No results returned"))
		}

		resultWindow.Show()
	})

	buttons := container.NewGridWithColumns(3,
		saveBtn,
		deleteBtn,
		runBtn,
	)

	content := container.NewBorder(
		container.NewVBox(
			widget.NewLabel("Saved Queries:"),
			querySelect,
			widget.NewSeparator(),
			widget.NewLabel("Query:"),
		),
		buttons,
		nil,
		nil,
		container.NewMax(
			container.NewBorder(
				nil,
				nil,
				nil,
				nil,
				queryEntry,
			),
		),
	)
	return content
}

func LoadTableEntities(pg *repository.PostgresRepository, tableName string, cardGrid *fyne.Container, w fyne.Window, a fyne.App, onEdit func(func()), onDelete func(func())) {
	entities, err := repository.GetTableEntities(pg, tableName)
	if err != nil {
		dialog.ShowError(err, w)
		log.Println(err)
		return
	}
	cardGrid.Objects = nil

	if len(entities) == 0 {
		cardGrid.Add(widget.NewLabel("No entities found"))
		return
	}

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
			text := fmt.Sprintf("%s: %v", key, formatValue(e[key]))
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
				if parent := fyne.CurrentApp().Driver().AllWindows()[0]; parent != nil {
					dialog.ShowInformation("Success", "Entity updated successfully", parent)
				}
			}))
			icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
			editW.SetIcon(icon)
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
						log.Println("entity deleted successfully")
					} else {
						dialog.ShowError(err, w)
						log.Println(err)
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
	os.Setenv("LANGUAGE", "en_US")
	os.Setenv("LANG", "en_US.UTF-8")
	os.Setenv("LC_ALL", "en_US.UTF-8")

	a := app.NewWithID("com.jurassicpark.database")
	w := a.NewWindow("Jurassic Park Database")
	w.Resize(fyne.NewSize(1200, 900))
	a.Preferences().SetString("language", "en")
	icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
	w.SetIcon(icon)

	cardGrid := container.NewGridWithColumns(3)
	cardScroll := container.NewVScroll(cardGrid)

	var selectedTable string
	tables, err := repository.GetTableList(pg)
	if err != nil {
		dialog.ShowError(err, w)
		log.Println(err)
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

	backupDatabase := func(isTableBackup bool) {
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				log.Println(err)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			var backupData string
			var errBackup error
			if isTableBackup && selectedTable != "" {
				backupData, errBackup = repository.BackupTable(pg, selectedTable)
			} else {
				backupData, errBackup = repository.BackupDatabase(pg)
			}

			if errBackup != nil {
				dialog.ShowError(errBackup, w)
				log.Println(errBackup)
				return
			}

			_, err = writer.Write([]byte(backupData))
			if err != nil {
				dialog.ShowError(err, w)
				log.Println(err)
				return
			}

			dialog.ShowInformation("Success", "Backup completed successfully", w)
		}, w)
	}

	exportDatabase := func(isTableExport bool) {
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				log.Println(err)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			var errExport error
			if isTableExport && selectedTable != "" {
				errExport = repository.ExportTableToExcel(pg, selectedTable, writer)
			} else {
				errExport = repository.ExportDatabaseToExcel(pg, writer)
			}

			if errExport != nil {
				dialog.ShowError(errExport, w)
				log.Println(errExport)
				return
			}

			dialog.ShowInformation("Success", "Export completed successfully", w)
		}, w)
	}

	newEntityBtn := widget.NewButtonWithIcon(" New Entity ", theme.ContentAddIcon(), func() {
		if selectedTable == "" {
			dialog.ShowInformation("Error", "Please select a table first", w)
			return
		}
		newEntityWindow := a.NewWindow("New Entity")
		icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
		newEntityWindow.SetIcon(icon)
		newEntityWindow.SetContent(CreateEntityForm(pg, selectedTable, make(map[string]interface{}), true, func() {
			LoadTableEntities(pg, selectedTable, cardGrid, w, a, nil, nil)
			newEntityWindow.Close()
			dialog.ShowInformation("Success", "Entity created successfully", w)
		}))
		newEntityWindow.Resize(fyne.NewSize(600, 400))
		newEntityWindow.Show()
	})
	newEntityContainer := container.NewBorder(
		nil, nil, nil, nil,
		container.NewMax(newEntityBtn),
	)

	tableButtons := container.NewVBox(
		widget.NewButtonWithIcon("New Table", theme.ContentAddIcon(), func() {
			newTableWindow := a.NewWindow("Create Table")
			newTableWindow.SetContent(CreateTableForm(pg, newTableWindow))
			icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
			newTableWindow.SetIcon(icon)
			newTableWindow.Resize(fyne.NewSize(600, 400))
			newTableWindow.SetOnClosed(func() {
				tables, _ = repository.GetTableList(pg)
				tableList.Refresh()
			})
			newTableWindow.Show()
		}),
		widget.NewButtonWithIcon("Delete Table", theme.DeleteIcon(), func() {
			if selectedTable == "" {
				err := fmt.Errorf("no table selected")
				dialog.ShowInformation("Error", err.Error(), w)
				log.Println(err)
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
						log.Println("table deleted successfully")
					} else {
						dialog.ShowError(err, w)
						log.Println(err)
					}
				}
			}, w)
		}),
		widget.NewButtonWithIcon("Export Table", theme.DownloadIcon(), func() {
			if selectedTable == "" {
				err := fmt.Errorf("no table selected")
				dialog.ShowInformation("Error", err.Error(), w)
				log.Println(err)
				return
			}
			exportDatabase(true)
		}),
		widget.NewButtonWithIcon("Export Database", theme.DownloadIcon(), func() {
			exportDatabase(false)
		}),
		widget.NewButtonWithIcon("Backup Table", theme.StorageIcon(), func() {
			if selectedTable == "" {
				err := fmt.Errorf("no table selected")
				dialog.ShowInformation("Error", err.Error(), w)
				log.Println(err)
				return
			}
			backupDatabase(true)
		}),
		widget.NewButtonWithIcon("Backup Database", theme.StorageIcon(), func() {
			backupDatabase(false)
		}),
		widget.NewButtonWithIcon("Run Query", theme.DocumentIcon(), func() {
			queryWindow := a.NewWindow("Run Query")
			queryWindow.SetContent(CreateQueryForm(pg, queryWindow))
			icon, _ := fyne.LoadResourceFromPath("C:\\BSUIR\\sem_6\\DB\\lab_6\\dino.png")
			queryWindow.SetIcon(icon)
			queryWindow.Resize(fyne.NewSize(600, 400))
			queryWindow.Show()
		}),
	)

	leftPanel := container.NewBorder(
		widget.NewLabelWithStyle("Tables", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewVBox(tableButtons), nil, nil,
		container.NewVScroll(tableList),
	)

	rightContent := container.NewBorder(
		newEntityContainer,
		nil,
		nil,
		nil,
		cardScroll,
	)

	split := container.NewHSplit(leftPanel, rightContent)
	split.Offset = 0.25
	w.SetContent(split)
	return a, w
}
