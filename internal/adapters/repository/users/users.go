package users

import (
	"context"
	"database/sql"

	"users-api/internal/core/domain/users"
)

type PostgresUser struct {
	db *sql.DB
}

func NewPostgresUser(db *sql.DB) *PostgresUser {
	user := &PostgresUser{
		db: db,
	}
	return user
}

func (u *PostgresUser) Find(ctx context.Context) ([]users.User, error) {
	var list = make([]users.User, 0)
	rows, err := u.db.QueryContext(ctx, "select id, name, age from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		list = append(list, user)
	}
	return list, nil
}

func (u *PostgresUser) GetById(ctx context.Context, userId string) (*users.User, error) {
	rows := u.db.QueryRowContext(ctx, "select id, name, age from users where id = $1", userId)

	var user users.User
	if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *PostgresUser) Save(ctx context.Context, user *users.User) error {
	if err := u.db.QueryRowContext(ctx,
		`insert into users ("name", "age", "email", "password") values ($1, $2, $3, $4)`,
		user.Name, user.Age, user.Email, user.Password,
	).Scan(); err != nil {
		if err.Error() != "sql: no rows in result set" {
			return err
		}
	}
	return nil
}
