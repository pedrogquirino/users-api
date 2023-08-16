package repository

import (
	"database/sql"
	"fmt"
	"log"

	postgresAddress "users-api/internal/adapters/repository/address"
	postgresUser "users-api/internal/adapters/repository/users"
	"users-api/internal/core/domain"
	"users-api/internal/core/ports"

	_ "github.com/lib/pq"
)

const (
	driverDB = "postgres"
)

func Start(options *domain.Options) {
	var usersRepository ports.UsersRepository
	var addressRepository ports.AddressRepository

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		options.Config.SqlDB.HOST, options.Config.SqlDB.User, options.Config.SqlDB.Password, options.Config.SqlDB.DB, options.Config.SqlDB.SSLMode)

	db, err := sql.Open(driverDB, psqlInfo)
	if err != nil {
		log.Fatal("Error conecting to database", err)
	}

	usersRepository = postgresUser.NewPostgresUser(db)
	addressRepository = postgresAddress.NewPostgresAddress(db)
	options.UsersRepository = usersRepository
	options.AddressRepository = addressRepository
}
