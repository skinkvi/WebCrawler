package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// mac os const connStr = "user=postgres dbname=WebCrawler sslmode=disable port=5433"
// linux "user=postgres dbname=WebCrawler sslmode=disable port=5432 password=1234"
const connStr = "user=postgres dbname=WebCrawler sslmode=disable port=5432 password=1234"

func init() {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}
