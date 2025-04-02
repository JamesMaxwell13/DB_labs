package data

import "database/sql"

var Db *sql.DB

func GetTableList() ([]string, error) {
	rows, err := Db.Query(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
		AND table_type = 'BASE TABLE'
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}
