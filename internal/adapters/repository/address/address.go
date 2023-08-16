package address

import (
	"context"
	"database/sql"
	"users-api/internal/core/domain/address"
)

type PostgresAddress struct {
	db *sql.DB
}

func NewPostgresAddress(db *sql.DB) *PostgresAddress {
	address := &PostgresAddress{
		db: db,
	}
	return address
}

func (u *PostgresAddress) Find(ctx context.Context, userId string) ([]address.Address, error) {
	var list = make([]address.Address, 0)
	rows, err := u.db.QueryContext(ctx, "select id, district, street, number, zipcode from address where user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var address address.Address
		if err := rows.Scan(&address.Id, &address.District, &address.Street, &address.Number, &address.ZipCode); err != nil {
			return nil, err
		}
		list = append(list, address)
	}
	return list, nil
}

func (u *PostgresAddress) GetById(ctx context.Context, addressId string, userId string) (*address.Address, error) {
	rows := u.db.QueryRowContext(ctx, "select id, district, street, number, zipcode from address where id = $1 and user_id = $2", addressId, userId)

	var address address.Address
	if err := rows.Scan(&address.Id, &address.District, &address.Street, &address.Number, &address.ZipCode); err != nil {
		return nil, err
	}
	return &address, nil
}

func (u *PostgresAddress) Save(ctx context.Context, userId string, address *address.Address) error {
	if err := u.db.QueryRowContext(ctx,
		`insert into address ("district", "street", "number", "zipcode", "user_id") values ($1, $2, $3, $4, $5)`,
		address.District, address.Street, address.Number, address.ZipCode, userId,
	).Scan(); err != nil {
		if err.Error() != "sql: no rows in result set" {
			return err
		}
	}
	return nil
}
