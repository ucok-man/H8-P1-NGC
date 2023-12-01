package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Employee struct {
	ID        int64
	FirstName string
	LastName  string
	Position  string
}

type EmployeeModel struct {
	conn *sql.DB
}

func (m *EmployeeModel) Insert(param *Employee) (*Employee, error) {
	query := `
		insert into employees(emp_firstname, emp_lastname, emp_position)
		values (?, ?, ?);	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := m.conn.ExecContext(ctx, query, param.FirstName, param.LastName, param.Position)
	if err != nil {
		return nil, fmt.Errorf("error inserting employee: %v", err)
	}

	param.ID, err = result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("last insert id error: %v", err)
	}
	return param, nil
}
