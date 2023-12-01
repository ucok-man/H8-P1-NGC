package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type MenuItem struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Category    string
}

type MenuItemModel struct {
	conn *sql.DB
}

func (m *MenuItemModel) Insert(param *MenuItem) (*MenuItem, error) {
	query := `
		insert into menu_items(mi_name, mi_description, mi_price, mi_category)
		values (?, ?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.conn.ExecContext(ctx, query, param.Name, param.Description, param.Price, param.Category)
	if err != nil {
		return nil, fmt.Errorf("error inserting employee: %w", err)
	}

	param.ID, err = result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("last insert id error: %v", err)
	}
	return param, nil
}
