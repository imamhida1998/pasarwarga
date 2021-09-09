package postgresql

import (
	"database/sql"
	"pasarwarga/database/postgresql/config"

	_ "github.com/lib/pq"
)

type Connection struct {
	Driver string
}

func (connection *Connection) connectionString(driver string, user string, password string,
	dbName string, sslMode string) string {

	connection.Driver = driver

	connString := "user=" + user
	connString += " password=" + password
	connString += " dbname=" + dbName
	connString += " sslmode=" + sslMode

	return connString

}

func (connection Connection) OpenConnection(driver string, user string, password string,
	dbName string, sslMode string) *sql.DB {

	connString := connection.connectionString(driver, user, password, dbName, sslMode)
	db, err := sql.Open(driver, connString)
	if err != nil {
		panic(err)
	}

	return db
}

func OpenConnection(phase string) *sql.DB {

	connection := Connection{}
	dbConf := config.DB_CONFIGS[phase]

	db := connection.OpenConnection(dbConf.Driver,
		dbConf.User,
		dbConf.Password,
		dbConf.DbName,
		dbConf.SslMode,
	)

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
