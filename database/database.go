package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/monstrasitix/gopnik/env"
)

func ConnectionString() string {
    return fmt.Sprintf("%s:%s@/%s",
        env.DB_USER(),
        env.DB_PASSWORD(),
        env.DB_NAME())
}

func Open() (*sql.DB, error) {
    return sql.Open("mysql", ConnectionString())
}

func HandleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
