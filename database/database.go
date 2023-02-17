package database

import (
	"database/sql"
	"fmt"
)
import migrate "github.com/rubenv/sql-migrate"
import "github.com/gobuffalo/packr/v2"

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam
	fmt.Println("Applied", n, "migrations!")
}
