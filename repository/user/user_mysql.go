package post

import (
	"context"
	"database/sql"
	"fmt"

	models "github.com/akamiko/entity-sample2/model"
	repo "github.com/akamiko/entity-sample2/repository"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLPostRepo(Conn *sql.DB) repo.UserRepository {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

type mysqlPostRepo struct {
	Conn *sql.DB
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.ID,
			&data.Name,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlPostRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := "Select id, name From users where id=?"

	rows, err := m.fetch(ctx, query, id)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, nil
	}

	return payload, nil
}
