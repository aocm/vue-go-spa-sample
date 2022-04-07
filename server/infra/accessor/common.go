package accessor

import (
	"fmt"
	"log"
	"time"

	"github.com/go-gorp/gorp"
)

// AccessDBとセットで利用
type IDBAccessor interface {
	initDb() *gorp.DbMap
	connectDb() *gorp.DbMap
	atTransactional(*gorp.DbMap) (*gorp.Transaction, error)
}

func AccessDB(acsr IDBAccessor) *gorp.DbMap {
	return acsr.initDb()
}
func ConnectDB(acsr IDBAccessor) *gorp.DbMap {
	return acsr.connectDb()
}

func StartTransaction(acsr IDBAccessor, dbmap *gorp.DbMap) (*gorp.Transaction, error) {
	return acsr.atTransactional(dbmap)
}

// mysqlとsql serverの共通構造
type DBBase struct {
	dbType   string
	user     string
	pass     string
	protocol string
	dbName   string
}

// 共通エラーハンドリング
func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
		fmt.Println(err)
	}
}

// TableBase テーブル共通構造部分
type TableBase struct {
	Created time.Time `json:"created" db:"created,notnull"`
	Updated time.Time `json:"updated" db:"updated,notnull"`
}

// PreInsert テーブル共通構造部分の共通処理（Insert）
func (tb *TableBase) PreInsert(s gorp.SqlExecutor) error {
	tb.Created = time.Now()
	tb.Updated = tb.Created
	return nil
}

// PreUpdate テーブル共通構造部分の共通処理（Update）
func (tb *TableBase) PreUpdate(s gorp.SqlExecutor) error {
	tb.Updated = time.Now()
	return nil
}

// Hoge ほげというドメインモデル。テーブル定義も兼ねている。
type Hoge struct {
	ID   int64  `json:"id" db:"id,primarykey,autoincrement"`
	Name string `json:"name" db:"name,notnull,size:200"`
	TableBase
}

// Message メッセージテーブル
type Message struct {
	ID   int64  `json:"id" db:"id,primarykey,autoincrement"`
	Text string `json:"text" db:"text,notnull,size:200"`
	TableBase
}
