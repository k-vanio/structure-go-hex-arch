package db

import (
	"database/sql"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driver, source string) (*Adapter, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (ad *Adapter) Close() {
	err := ad.db.Close()
	if err != nil {
		log.Println(err)
	}
}

func (ad *Adapter) AddHistory(answer int32, operation string) error {
	sqlString, args, err := sq.Insert("logs").Columns("value").Values(fmt.Sprintf("r: %v o: %v", answer, operation)).ToSql()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = ad.db.Exec(sqlString, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
