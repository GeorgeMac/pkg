package psql

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

var (
	Statement = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

// TruncateAll takes a database connection, lists all the tables which
// aren't tracking schema_migrations and issues a cascading truncate
// across each of them.
func TruncateAll(db *sql.DB) error {
	rows, err := Statement.
		Select("tablename").
		From("pg_tables").
		Where(sq.Eq{"schemaname": "public"}).
		Where(sq.NotEq{"tablename": "schema_migrations"}).
		RunWith(db).
		Query()
	if err != nil {
		return err
	}

	var tables []string
	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			return err
		}

		tables = append(tables, tablename)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	for _, table := range tables {
		if _, err := db.Exec(fmt.Sprintf(`TRUNCATE TABLE %q CASCADE;`, table)); err != nil {
			return err
		}
	}

	return nil
}
