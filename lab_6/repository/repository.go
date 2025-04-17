package repository

import (
	_ "database/sql"
	"fmt"
	"strings"
)

func GetTableList(pg *PostgresRepository) ([]string, error) {
	rows, err := pg.Db.Query(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
	`)
	if err != nil {
		return nil, err
	}
	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func GetTableEntities(pg *PostgresRepository, tableName string) ([]map[string]interface{}, error) {
	if strings.TrimSpace(tableName) == "" {
		return nil, fmt.Errorf("имя таблицы не может быть пустым")
	}

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := pg.Db.Queryx(query)
	if err != nil {
		return nil, err
	}
	var entities []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			return nil, err
		}
		entities = append(entities, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entities, nil
}

// ColumnInfo содержит информацию о колонке таблицы
type ColumnInfo struct {
	Name string
	Type string
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
		return nil, err
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		var col ColumnInfo
		if err := rows.Scan(&col.Name, &col.Type); err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}

	return columns, nil
}
