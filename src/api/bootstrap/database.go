package bootstrap

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgresql driver
)

var (
	// SQL wrapper
	SQL *sqlx.DB
	//Databases Info
	databases Info
)

// Type is the type of database from a Type* constant
type Type string

const (
	// TypeMySQL is MySQL
	TypeMySQL Type = "MySQL"
	// TypePostgresql is Postgresql
	TypePostgresql Type = "Postgresql"
)

// Info contains the database configurations
type Info struct {
	// Database type
	Type Type
	// MySQL info if used
	MySQL MySQLInfo
	// Postgresl info if used
	Postgresql PostgresqlInfo
}

// MySQLInfo is the details for the database connection
type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// PostgresqlInfo is struct for connect to postgresql
type PostgresqlInfo struct {
	Username string
	Password string
	Name     string
}

//DSN returns the Data Source Name
func DSN(ci MySQLInfo) string {
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		fmt.Sprintf("%d", ci.Port) +
		")/" +
		ci.Name + ci.Parameter
}

// DbInfo function
func DbInfo(ci PostgresqlInfo) string {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", ci.Username, ci.Password, ci.Name)
	return dbinfo
}

//Connect to database
func Connect(d Info) {
	log.Println("call Connect func")
	var err error

	// Store the config
	databases = d
	log.Println(databases)
	switch d.Type {
	case TypeMySQL:
		// Connect to MySQL
		if SQL, err = sqlx.Connect("mysql", DSN(d.MySQL)); err != nil {
			log.Println("SQL Driver Error", err)
		}

		// Check if is alive
		if err = SQL.Ping(); err != nil {
			log.Println("Database Error", err)
		}
	case TypePostgresql:
		// Connect to PostgreSQL
		//Try to connect postgresql, but 1st call DbInfo
		db, err := sql.Open("postgres", DbInfo(d.Postgresql))
		//checkErr(err)
		if err != nil {
			log.Println("SQL Driver Error", err)
		}
		log.Println("SQL Driver Postgre Success")
		defer db.Close()
	default:
		log.Println("No registered database in config")
	}
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return databases
}
