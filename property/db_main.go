package property

import (
	"database/sql"
	"cmdb_backend/logger"

	_ "github.com/lib/pq"
)

func Db_main() *sql.DB {
	db, err := sql.Open("postgres", "user=cmdb_user password=000000 dbname=cmdb sslmode=disable")
	if err != nil {
		logger.Error.Fatal("can't open database: ", err)
	}
	return db
}
