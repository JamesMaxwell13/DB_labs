package repository

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	_ "github.com/jmoiron/sqlx"
	"log"
	"os"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

const queriesFile = "config/queries.json"

type Query struct {
	Name string `json:"name"`
	SQL  string `json:"sql"`
}

type ColumnInfo struct {
	Name string
	Type string
}

func loadQueries() ([]Query, error) {
	data, err := os.ReadFile(queriesFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println(err)
			return []Query{}, nil
		}
		log.Println(err)
		return nil, fmt.Errorf("failed to read queries file: %w", err)
	}

	var queries []Query
	if err := json.Unmarshal(data, &queries); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to parse queries: %w", err)
	}
	return queries, nil
}

func saveQueries(queries []Query) error {
	data, err := json.MarshalIndent(queries, "", "  ")
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to marshal queries: %w", err)
	}
	if err := os.WriteFile(queriesFile, data, 0644); err != nil {
		log.Println(err)
		return fmt.Errorf("failed to write queries file: %w", err)
	}
	return nil
}

func GetQueryNames() ([]string, error) {
	queries, err := loadQueries()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	names := make([]string, len(queries))
	for i, q := range queries {
		names[i] = q.Name
	}
	return names, nil
}

func GetQueryByName(name string) (string, error) {
	queries, err := loadQueries()
	if err != nil {
		log.Println(err)
		return "", err
	}
	for _, q := range queries {
		if q.Name == name {
			return q.SQL, nil
		}
	}
	log.Printf("query '%s' not found\n", name)
	return "", fmt.Errorf("query '%s' not found", name)
}

func SaveQuery(name, sql string) error {
	queries, err := loadQueries()
	if err != nil {
		log.Println(err)
		return err
	}

	for i, q := range queries {
		if q.Name == name {
			queries[i].SQL = sql
			return saveQueries(queries)
		}
	}

	queries = append(queries, Query{Name: name, SQL: sql})
	log.Printf("saved query '%s'\n", name)
	return saveQueries(queries)
}

func GetTableList(pg *PostgresRepository) ([]string, error) {
	rows, err := pg.Db.Query(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
	`)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to get tables: %w", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan table name: %w", err)
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func GetTableEntities(pg *PostgresRepository, tableName string) ([]map[string]interface{}, error) {
	if strings.TrimSpace(tableName) == "" {
		log.Println("table name cannot be empty")
		return nil, fmt.Errorf("table name cannot be empty")
	}

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := pg.Db.Queryx(query)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to query table: %w", err)
	}
	defer rows.Close()

	var entities []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		entities = append(entities, formatRowValues(row))
	}
	return entities, nil
}

func formatRowValues(row map[string]interface{}) map[string]interface{} {
	for k, v := range row {
		switch val := v.(type) {
		case []byte:
			row[k] = string(val)
		case time.Time:
			if strings.Contains(strings.ToLower(k), "date") {
				row[k] = val.Format("2006-01-02")
			} else if strings.Contains(strings.ToLower(k), "time") {
				row[k] = val.Format("15:04:05")
			}
		}
	}
	return row
}

func GetTableColumns(pg *PostgresRepository, tableName string) ([]ColumnInfo, error) {
	query := `
		SELECT column_name, data_type 
		FROM information_schema.columns 
		WHERE table_name = $1
		ORDER BY ordinal_position
	`

	rows, err := pg.Db.Query(query, tableName)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		var col ColumnInfo
		if err := rows.Scan(&col.Name, &col.Type); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan column: %w", err)
		}
		columns = append(columns, col)
	}
	return columns, nil
}

func ExecuteQuery(pg *PostgresRepository, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := pg.Db.Queryx(query, args...)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan result row: %w", err)
		}
		results = append(results, formatRowValues(row))
	}
	return results, nil
}

func DeleteTable(pg *PostgresRepository, tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
	_, err := pg.Db.Exec(query)
	if err != nil {
		log.Println(err)
	}
	return err
}

func DeleteEntity(pg *PostgresRepository, tableName string, entity map[string]interface{}) error {
	var whereClauses []string
	var args []interface{}
	i := 1
	for k, v := range entity {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", k, i))
		args = append(args, v)
		i++
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, strings.Join(whereClauses, " AND "))
	_, err := pg.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateEntity(pg *PostgresRepository, tableName string, old map[string]interface{}, updated map[string]interface{}) error {
	var setClauses []string
	var args []interface{}
	i := 1
	for k, v := range updated {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", k, i))
		args = append(args, v)
		i++
	}

	var whereClauses []string
	for k, v := range old {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", k, i))
		args = append(args, v)
		i++
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(setClauses, ", "), strings.Join(whereClauses, " AND "))
	_, err := pg.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return err
}

func InsertEntity(pg *PostgresRepository, tableName string, values map[string]interface{}) error {
	var keys []string
	var params []string
	var args []interface{}
	i := 1
	for k, v := range values {
		keys = append(keys, k)
		params = append(params, fmt.Sprintf("$%d", i))
		args = append(args, v)
		i++
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(keys, ", "), strings.Join(params, ", "))
	_, err := pg.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return err
}

func BackupTable(pg *PostgresRepository, tableName string) (string, error) {
	var backup strings.Builder

	var tableDDL string
	err := pg.Db.Get(&tableDDL, fmt.Sprintf(`
		SELECT 'CREATE TABLE ' || table_name || ' (' || 
		string_agg(column_name || ' ' || data_type || 
		CASE WHEN is_nullable = 'NO' THEN ' NOT NULL' ELSE '' END, ', ') || 
		');' 
		FROM information_schema.columns 
		WHERE table_name = '%s' 
		GROUP BY table_name`, tableName))
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to get table DDL: %w", err)
	}
	backup.WriteString(tableDDL + "\n\n")

	rows, err := pg.Db.Queryx(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to query table data: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to get columns: %w", err)
	}

	for rows.Next() {
		row := make(map[string]interface{})
		err := rows.MapScan(row)
		if err != nil {
			log.Println(err)
			return "", fmt.Errorf("failed to scan row: %w", err)
		}

		var values []string
		for _, col := range columns {
			val := row[col]
			if val == nil {
				values = append(values, "NULL")
			} else {
				switch v := val.(type) {
				case []byte:
					values = append(values, fmt.Sprintf("'%s'", pg.EscapeString(string(v))))
				case string:
					values = append(values, fmt.Sprintf("'%s'", pg.EscapeString(v)))
				case time.Time:
					if strings.Contains(strings.ToLower(col), "date") {
						values = append(values, fmt.Sprintf("'%s'", v.Format("2006-01-02")))
					} else if strings.Contains(strings.ToLower(col), "time") {
						values = append(values, fmt.Sprintf("'%s'", v.Format("15:04:05")))
					} else {
						values = append(values, fmt.Sprintf("'%s'", v.Format(time.RFC3339)))
					}
				default:
					values = append(values, fmt.Sprintf("'%v'", v))
				}
			}
		}

		backup.WriteString(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);\n",
			tableName, strings.Join(columns, ", "), strings.Join(values, ", ")))
	}

	var indexes []string
	err = pg.Db.Select(&indexes, `
		SELECT indexdef 
		FROM pg_indexes 
		WHERE tablename = $1`, tableName)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to get indexes: %w", err)
	}

	for _, idx := range indexes {
		backup.WriteString(idx + ";\n")
	}

	return backup.String(), nil
}

func BackupDatabase(pg *PostgresRepository) (string, error) {
	var backup strings.Builder

	backup.WriteString(fmt.Sprintf("-- PostgreSQL database dump\n-- %s\n\n", time.Now().Format(time.RFC1123)))
	backup.WriteString("SET statement_timeout = 0;\n")
	backup.WriteString("SET lock_timeout = 0;\n")
	backup.WriteString("SET idle_in_transaction_session_timeout = 0;\n")
	backup.WriteString("SET client_encoding = 'UTF8';\n")
	backup.WriteString("SET standard_conforming_strings = on;\n")
	backup.WriteString("SET check_function_bodies = false;\n")
	backup.WriteString("SET xmloption = content;\n")
	backup.WriteString("SET client_min_messages = warning;\n")
	backup.WriteString("SET row_security = off;\n\n")

	tables, err := pg.Db.Queryx(`
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public' 
		ORDER BY table_name`)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to get tables: %w", err)
	}
	defer tables.Close()

	var tableNames []string
	for tables.Next() {
		var tableName string
		if err := tables.Scan(&tableName); err != nil {
			log.Println(err)
			return "", fmt.Errorf("failed to scan table name: %w", err)
		}
		tableNames = append(tableNames, tableName)
	}

	for _, table := range tableNames {
		backup.WriteString(fmt.Sprintf("\n--\n-- Table: %s\n--\n\n", table))
		tableBackup, err := BackupTable(pg, table)
		if err != nil {
			log.Println(err)
			return "", fmt.Errorf("failed to backup table %s: %w", table, err)
		}
		backup.WriteString(tableBackup + "\n")
	}

	var fks []string
	err = pg.Db.Select(&fks, `
		SELECT 'ALTER TABLE ' || conrelid::regclass || 
		' ADD CONSTRAINT ' || conname || 
		' ' || pg_get_constraintdef(oid) || ';'
		FROM pg_constraint 
		WHERE contype = 'f'`)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("failed to get foreign keys: %w", err)
	}

	if len(fks) > 0 {
		backup.WriteString("\n--\n-- Foreign keys\n--\n\n")
		for _, fk := range fks {
			backup.WriteString(fk + "\n")
		}
	}

	return backup.String(), nil
}

func ExportTableToExcel(pg *PostgresRepository, tableName string, writer fyne.URIWriteCloser) error {
	f := excelize.NewFile()
	defer f.Close()

	index, err := f.NewSheet(tableName)
	if err != nil {
		log.Println(err)
		return err
	}

	entities, err := GetTableEntities(pg, tableName)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(entities) == 0 {
		log.Printf("no data in table %s\n", tableName)
		return fmt.Errorf("no data in table %s", tableName)
	}

	headers := make([]string, 0, len(entities[0]))
	for k := range entities[0] {
		headers = append(headers, k)
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(tableName, cell, header)
	}

	for rowIdx, entity := range entities {
		for colIdx, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
			f.SetCellValue(tableName, cell, entity[header])
		}
	}

	f.SetActiveSheet(index)
	return f.Write(writer)
}

func ExportDatabaseToExcel(pg *PostgresRepository, writer fyne.URIWriteCloser) error {
	f := excelize.NewFile()
	defer f.Close()

	f.DeleteSheet("Sheet1")

	tables, err := GetTableList(pg)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, table := range tables {
		index, err := f.NewSheet(table)
		if err != nil {
			log.Println(err)
			return err
		}

		entities, err := GetTableEntities(pg, table)
		if err != nil {
			log.Println(err)
			return err
		}

		if len(entities) == 0 {
			continue
		}

		headers := make([]string, 0, len(entities[0]))
		for k := range entities[0] {
			headers = append(headers, k)
		}

		for i, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(i+1, 1)
			f.SetCellValue(table, cell, header)
		}

		for rowIdx, entity := range entities {
			for colIdx, header := range headers {
				cell, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
				f.SetCellValue(table, cell, entity[header])
			}
		}
		f.SetActiveSheet(index)
	}
	return f.Write(writer)
}
