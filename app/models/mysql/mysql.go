package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Info struct {
	Driver   string
	Hostname string
	Database string
	Username string
	Password string
	Port     string
}

// Connect return the connection to mysqlConnect
// with struct that has the option for the connection
func (i *Info) Connect() (*sql.DB, error) {
	connStr := i.Username + ":" + i.Password + "@tcp(" + i.Hostname + ")/" + i.Database + "?parseTime=true"
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	return db, err
}
