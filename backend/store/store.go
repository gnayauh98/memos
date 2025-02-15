package store

import (
	"database/sql"
	"log"

	"github.com/anqzi/memos/store/db"
)

type Store struct {
	db *sql.DB
}

func NewStore(config db.DBConfig) (Store, error) {
	_db, err := db.NewPostgresDB(config)
	if err != nil {
		return Store{}, err
	}
	err = _db.Ping()
	if err != nil {
		log.Println("数据库连接失败")
		return Store{}, err
	}
	log.Println("数据库连接成功")
	return Store{
		db: _db,
	}, nil
}
