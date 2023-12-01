package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	conn *sql.DB
}

func New(dsn string) (*MysqlDB, error) {
	conn, err := connect(dsn)
	if err != nil {
		return nil, err
	}
	return &MysqlDB{
		conn: conn,
	}, nil
}

func (db *MysqlDB) CreateTables() error {
	tableQuery := []string{
		table_employee_query,
		table_menu_items_query,
		table_orders_query,
		table_orders_item_query,
		table_payment_query,
	}

	for _, query := range tableQuery {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := db.conn.ExecContext(ctx, query)
		if err != nil {
			return fmt.Errorf("error creating table: \n%s\ncause: %w", query, err)
		}
	}
	return nil
}
