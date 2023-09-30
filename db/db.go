package db

import (
	"database/sql"
	// "os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//create struct of type database

type Database struct {
	gorm.Model
	db *sql.DB
}

// host     = "localhost"
// port     = 5432
// user     = "postgres"
// password = "kiprop"
// dbname   = "mydb"postgresql://godbtest:31877101@go-chat.cydlpjqntwzu.us-west-2.rds.amazonaws.com:5432/go-chat?sslmode=disable
func NewDatabase() (*Database, error) {
	// Dbdriver := os.Getenv("DB_DRIVER")
	// DbHost := os.Getenv("DB_HOST")
	// DbUser := os.Getenv("POSTGRES_USER")
	// DbPassword := os.Getenv("POSTGRES_PASSWORD")
	// DbName := os.Getenv("POSTGRES_DB")
	// DbPort := os.Getenv("DB_PORT")
	db, err := sql.Open("postgres", "postgresql://kiprop:31877101@kiprop.cs7pago4pgpd.us-west-2.rds.amazonaws.com:5432/gochat?sslmode=disable")
	if err != nil {
		return nil, err
	}
	//return pointer database struct
	return &Database{db: db}, nil
}

//create func to close the db
func (d *Database) Close() {
	d.db.Close()
}

//get db fucn because the database connection is encapsulated
func (d *Database) GetDB() *sql.DB {
	return d.db
}
