package main

import (
	"database/sql"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"lab_6/data"
	"lab_6/gui"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	data.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Db.Close()

	err = data.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := data.Db.Query("SELECT datname FROM pg_database WHERE datistemplate = false")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Databases:")
	for rows.Next() {
		var dbName string
		err := rows.Scan(&dbName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dbName)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	a := app.New()
	w := a.NewWindow("Jurassic Park Database")
	a, w = gui.GUI()
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

//
//var db *sql.DB
//
//func main() {
//	myApp := app.New()
//	myWindow := myApp.NewWindow("Jurassic Park Database Manager")
//	myWindow.Resize(fyne.NewSize(800, 600))
//
//	// Подключение к базе данных
//	connectBtn := widget.NewButton("Connect to Database", func() {
//		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//			host, port, user, password, dbname)
//		var err error
//		db, err = sql.Open("postgres", psqlInfo)
//		if err != nil {
//			dialog.ShowError(err, myWindow)
//			return
//		}
//		err = db.Ping()
//		if err != nil {
//			dialog.ShowError(err, myWindow)
//			return
//		}
//		dialog.ShowInformation("Success", "Connected to database!", myWindow)
//	})
//
//	// Таблицы
//	tablesBtn := widget.NewButton("Manage Tables", func() {
//		showTablesWindow(myWindow)
//	})
//
//	// Запросы
//	queriesBtn := widget.NewButton("Run Queries", func() {
//		showQueriesWindow(myWindow)
//	})
//
//	// Резервные копии
//	backupBtn := widget.NewButton("Backup/Restore", func() {
//		showBackupWindow(myWindow)
//	})
//
//	// Экспорт
//	exportBtn := widget.NewButton("Export Data", func() {
//		showExportWindow(myWindow)
//	})
//
//	mainMenu := container.NewVBox(
//		widget.NewLabel("Jurassic Park Database Manager"),
//		connectBtn,
//		tablesBtn,
//		queriesBtn,
//		backupBtn,
//		exportBtn,
//	)
//
//	myWindow.SetContent(mainMenu)
//	myWindow.ShowAndRun()
//}
//
//func showTablesWindow(parent fyne.Window) {
//	window := fyne.CurrentApp().NewWindow("Manage Tables")
//	window.Resize(fyne.NewSize(600, 400))
//
//	tables, err := getTableList()
//	if err != nil {
//		dialog.ShowError(err, parent)
//		return
//	}
//
//	var selectedTable string
//	tableList := widget.NewList(
//		func() int { return len(tables) },
//		func() fyne.CanvasObject { return widget.NewLabel("") },
//		func(i int, o fyne.CanvasObject) {
//			o.(*widget.Label).SetText(tables[i])
//		},
//	)
//
//	// Добавляем обработчик выбора
//	tableList.OnSelected = func(id int) {
//		selectedTable = tables[id]
//	}
//
//	createTableBtn := widget.NewButton("Create New Table", func() {
//		showCreateTableWindow(window)
//	})
//
//	deleteTableBtn := widget.NewButton("Delete Selected Table", func() {
//		if selectedTable == "" {
//			dialog.ShowInformation("Error", "Please select a table first", window)
//			return
//		}
//		showDeleteTableWindow(window, selectedTable)
//	})
//
//	editTableBtn := widget.NewButton("Edit Selected Table", func() {
//		if selectedTable == "" {
//			dialog.ShowInformation("Error", "Please select a table first", window)
//			return
//		}
//		showEditTableWindow(window, selectedTable)
//	})
//
//	content := container.NewBorder(
//		nil,
//		container.NewHBox(createTableBtn, deleteTableBtn, editTableBtn),
//		nil,
//		nil,
//		tableList,
//	)
//
//	window.SetContent(content)
//	window.Show()
//}
//
//func showQueriesWindow(parent fyne.Window) {
//	window := fyne.CurrentApp().NewWindow("Run Queries")
//	window.Resize(fyne.NewSize(800, 600))
//
//	queryEntry := widget.NewMultiLineEntry()
//	queryEntry.SetPlaceHolder("Enter your SQL query here...")
//
//	resultTable := widget.NewTable(
//		func() (int, int) { return 0, 0 },
//		func() fyne.CanvasObject { return widget.NewLabel("") },
//		func(i widget.TableCellID, o fyne.CanvasObject) { o.(*widget.Label).SetText("") },
//	)
//
//	executeBtn := widget.NewButton("Execute", func() {
//		query := queryEntry.Text
//		if query == "" {
//			dialog.ShowInformation("Error", "Please enter a query", window)
//			return
//		}
//
//		rows, err := db.Query(query)
//		if err != nil {
//			dialog.ShowError(err, window)
//			return
//		}
//		defer rows.Close()
//
//		columns, err := rows.Columns()
//		if err != nil {
//			dialog.ShowError(err, window)
//			return
//		}
//
//		var results [][]string
//		results = append(results, columns)
//
//		values := make([]interface{}, len(columns))
//		valuePtrs := make([]interface{}, len(columns))
//		for i := range columns {
//			valuePtrs[i] = &values[i]
//		}
//
//		for rows.Next() {
//			err := rows.Scan(valuePtrs...)
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//
//			var row []string
//			for _, v := range values {
//				if v == nil {
//					row = append(row, "NULL")
//				} else {
//					row = append(row, fmt.Sprintf("%v", v))
//				}
//			}
//			results = append(results, row)
//		}
//
//		resultTable.Length = func() (int, int) {
//			return len(results), len(columns)
//		}
//		resultTable.CreateCell = func() fyne.CanvasObject {
//			return widget.NewLabel("")
//		}
//		resultTable.UpdateCell = func(i widget.TableCellID, o fyne.CanvasObject) {
//			o.(*widget.Label).SetText(results[i.Row][i.Col])
//		}
//	})
//
//	saveQueryBtn := widget.NewButton("Save Query", func() {
//		query := queryEntry.Text
//		if query == "" {
//			dialog.ShowInformation("Error", "No query to save", window)
//			return
//		}
//
//		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//			if writer == nil {
//				return
//			}
//			defer writer.Close()
//
//			_, err = writer.Write([]byte(query))
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//			dialog.ShowInformation("Success", "Query saved successfully", window)
//		}, window)
//	})
//
//	content := container.NewBorder(
//		container.NewVBox(
//			widget.NewLabel("Enter SQL Query:"),
//			queryEntry,
//			container.NewHBox(executeBtn, saveQueryBtn),
//		),
//		nil,
//		nil,
//		nil,
//		resultTable,
//	)
//
//	window.SetContent(content)
//	window.Show()
//}
//
//func showBackupWindow(parent fyne.Window) {
//	window := fyne.CurrentApp().NewWindow("Backup/Restore")
//	window.Resize(fyne.NewSize(600, 400))
//
//	backupBtn := widget.NewButton("Create Backup", func() {
//		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//			if writer == nil {
//				return
//			}
//			defer writer.Close()
//
//			// Здесь должна быть логика создания резервной копии
//			// Это упрощенный пример
//			tables, err := getTableList()
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//
//			var backupData strings.Builder
//			for _, table := range tables {
//				backupData.WriteString(fmt.Sprintf("-- Backup for table: %s\n", table))
//				rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))
//				if err != nil {
//					dialog.ShowError(err, window)
//					return
//				}
//				defer rows.Close()
//
//				columns, err := rows.Columns()
//				if err != nil {
//					dialog.ShowError(err, window)
//					return
//				}
//
//				for rows.Next() {
//					values := make([]interface{}, len(columns))
//					valuePtrs := make([]interface{}, len(columns))
//					for i := range columns {
//						valuePtrs[i] = &values[i]
//					}
//
//					err := rows.Scan(valuePtrs...)
//					if err != nil {
//						dialog.ShowError(err, window)
//						return
//					}
//
//					var row strings.Builder
//					row.WriteString(fmt.Sprintf("INSERT INTO %s VALUES (", table))
//					for i, v := range values {
//						if i > 0 {
//							row.WriteString(", ")
//						}
//						if v == nil {
//							row.WriteString("NULL")
//						} else {
//							row.WriteString(fmt.Sprintf("'%v'", v))
//						}
//					}
//					row.WriteString(");\n")
//					backupData.WriteString(row.String())
//				}
//			}
//
//			_, err = writer.Write([]byte(backupData.String()))
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//			dialog.ShowInformation("Success", "Backup created successfully", window)
//		}, window)
//	})
//
//	restoreBtn := widget.NewButton("Restore Backup", func() {
//		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//			if reader == nil {
//				return
//			}
//			defer reader.Close()
//
//			data := make([]byte, 1024*1024) // 1MB buffer
//			n, err := reader.Read(data)
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//
//			// Здесь должна быть логика восстановления из резервной копии
//			// Это упрощенный пример
//			queries := strings.Split(string(data[:n]), ";")
//			for _, query := range queries {
//				query = strings.TrimSpace(query)
//				if query == "" {
//					continue
//				}
//				_, err := db.Exec(query)
//				if err != nil {
//					dialog.ShowError(fmt.Errorf("error executing query: %v\nQuery: %s", err, query), window)
//					return
//				}
//			}
//
//			dialog.ShowInformation("Success", "Backup restored successfully", window)
//		}, window)
//	})
//
//	content := container.NewVBox(
//		widget.NewLabel("Database Backup and Restore"),
//		backupBtn,
//		restoreBtn,
//	)
//
//	window.SetContent(content)
//	window.Show()
//}
//
//func showExportWindow(parent fyne.Window) {
//	window := fyne.CurrentApp().NewWindow("Export Data")
//	window.Resize(fyne.NewSize(600, 400))
//
//	tables, err := getTableList()
//	if err != nil {
//		dialog.ShowError(err, parent)
//		return
//	}
//
//	tableSelect := widget.NewSelect(tables, func(s string) {})
//	queryEntry := widget.NewMultiLineEntry()
//	queryEntry.SetPlaceHolder("Or enter custom query for export...")
//
//	exportBtn := widget.NewButton("Export to CSV", func() {
//		var query string
//		if queryEntry.Text != "" {
//			query = queryEntry.Text
//		} else if tableSelect.Selected != "" {
//			query = fmt.Sprintf("SELECT * FROM %s", tableSelect.Selected)
//		} else {
//			dialog.ShowInformation("Error", "Please select a table or enter a query", window)
//			return
//		}
//
//		rows, err := db.Query(query)
//		if err != nil {
//			dialog.ShowError(err, window)
//			return
//		}
//		defer rows.Close()
//
//		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//			if writer == nil {
//				return
//			}
//			defer writer.Close()
//
//			csvWriter := csv.NewWriter(writer)
//			defer csvWriter.Flush()
//
//			columns, err := rows.Columns()
//			if err != nil {
//				dialog.ShowError(err, window)
//				return
//			}
//			csvWriter.Write(columns)
//
//			values := make([]interface{}, len(columns))
//			valuePtrs := make([]interface{}, len(columns))
//			for i := range columns {
//				valuePtrs[i] = &values[i]
//			}
//
//			for rows.Next() {
//				err := rows.Scan(valuePtrs...)
//				if err != nil {
//					dialog.ShowError(err, window)
//					return
//				}
//
//				var row []string
//				for _, v := range values {
//					if v == nil {
//						row = append(row, "")
//					} else {
//						row = append(row, fmt.Sprintf("%v", v))
//					}
//				}
//				csvWriter.Write(row)
//			}
//
//			dialog.ShowInformation("Success", "Data exported successfully", window)
//		}, window)
//	})
//
//	content := container.NewVBox(
//		widget.NewLabel("Export Data to CSV"),
//		widget.NewLabel("Select table:"),
//		tableSelect,
//		widget.NewLabel("Or enter custom query:"),
//		queryEntry,
//		exportBtn,
//	)
//
//	window.SetContent(content)
//	window.Show()
//}
//
//// Вспомогательные функции
//
//func getTableList() ([]string, error) {
//	rows, err := db.Query(`
//		SELECT table_name
//		FROM information_schema.tables
//		WHERE table_schema = 'public'
//		AND table_type = 'BASE TABLE'
//	`)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var tables []string
//	for rows.Next() {
//		var table string
//		err := rows.Scan(&table)
//		if err != nil {
//			return nil, err
//		}
//		tables = append(tables, table)
//	}
//	return tables, nil
//}
//
//func showCreateTableWindow(parent fyne.Window) {
//	// Реализация окна создания таблицы
//}
//
//func showDeleteTableWindow(parent fyne.Window, tableName string) {
//	// Реализация окна удаления таблицы
//}
//
//func showEditTableWindow(parent fyne.Window, tableName string) {
//	// Реализация окна редактирования таблицы
//}
