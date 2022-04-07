package accessor

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlAccessor struct {
	DBBase
}

func (a MysqlAccessor) atTransactional(dbmap *gorp.DbMap) (*gorp.Transaction, error) {
	return dbmap.Begin()
}

func (a MysqlAccessor) initDb() *gorp.DbMap {
	// connect mysql
	a.dbType = "mysql"
	a.user = "root"
	a.pass = "root"
	a.protocol = "tcp(db:3306)"
	a.dbName = "vgs"

	CONNECT := a.user + ":" + a.pass + "@" + a.protocol + "/"
	CONNECTDB := CONNECT + a.dbName + "?parseTime=true"

	db, err := sql.Open(a.dbType, CONNECT)
	checkError(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// create db if not exists
	_, err = dbmap.Exec("CREATE DATABASE IF NOT EXISTS vgs DEFAULT CHARACTER SET utf8;")
	checkError(err, "create db failed")

	// connect db
	db, err = sql.Open("mysql", CONNECTDB)
	checkError(err, "sql.Open failed")
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add table
	dbmap.AddTableWithName(Hoge{}, "hoge").SetKeys(true, "ID")
	dbmap.AddTableWithName(Message{}, "message").SetKeys(true, "ID")

	// create the table
	err = dbmap.CreateTablesIfNotExists()
	checkError(err, "Create tables failed")

	return dbmap
}
func (a MysqlAccessor) connectDb() *gorp.DbMap {
	// connect mysql
	a.dbType = "mysql"
	a.user = "root"
	a.pass = "root"
	a.protocol = "tcp(db:3306)"
	a.dbName = "vgs"

	CONNECT := a.user + ":" + a.pass + "@" + a.protocol + "/"
	CONNECTDB := CONNECT + a.dbName + "?parseTime=true"

	// connect db
	db, err := sql.Open("mysql", CONNECTDB)
	checkError(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Hoge{}, "hoge").SetKeys(true, "ID")
	dbmap.AddTableWithName(Message{}, "message").SetKeys(true, "ID")
	return dbmap
}
