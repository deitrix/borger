package borger

import (
	"database/sql"
	"fmt"
	"strings"
)

var DB *sql.DB

type Row interface {
	Scan(dest ...any) error
}

type Table[T any] struct {
	Name    string
	Columns func(*T) map[string]any
}

func (t *Table[T]) Scan(row Row, v *T, cols ...string) error {
	colDest := t.Columns(v)

	var rowDest []any
	for _, col := range cols {
		dest, ok := colDest[col]
		if !ok {
			panic("missing column: " + col)
		}
		rowDest = append(rowDest, dest)
	}

	return row.Scan(rowDest...)
}

func (t *Table[T]) Select(cols ...string) (results []T, err error) {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(cols, ", "), t.Name)

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("querying database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var v T
		if err := t.Scan(rows, &v, cols...); err != nil {
			return nil, err
		}

		results = append(results, v)
	}

	return results, nil
}
