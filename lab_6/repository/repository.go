package repository

import (
	_ "database/sql"
	"fmt"
	_ "github.com/jmoiron/sqlx"
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

func DeleteTable(pg *PostgresRepository, tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
	_, err := pg.Db.Exec(query)
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
	return err
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
