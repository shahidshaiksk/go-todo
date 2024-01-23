package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type DB interface {
	Insert(tableName string, args map[string]string) error
	DeleteWithCondition(tableName, condition string) error
	SelectWithCondition(tableName, condition string) (*sql.Rows, error)
	Update(tableName, condition string, args map[string]string) error
	Select(tableName string) (*sql.Rows, error)
}

type Db struct {
	db *sql.DB
}

func NewPostgresDb() (*Db, error) {
	db := &Db{}
	var err error
	db.db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/task?sslmode=disable")
	return db, err
}

func (db *Db) Insert(tableName string, args map[string]string) error {
	keys := make([]string, 0, len(args))
	mapValues := make([]string, 0, len(args))
	for k, v := range args {
		keys = append(keys, k)
		mapValues = append(mapValues, v)
	}
	columns := strings.Join(keys, ", ")
	values := strings.Join(mapValues, ", ")

	query := fmt.Sprintf("INSERT INTO %v(%v) VALUES (%v)", tableName, columns, values)

	_, err := db.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) DeleteWithCondition(tableName, condition string) error {
	query := fmt.Sprintf("DELETE FROM %v WHERE %v", tableName, condition)

	_, err := db.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) SelectWithCondition(tableName, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE %v", tableName, condition)

	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *Db) Select(tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %v", tableName)

	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *Db) Update(tableName, condition string, args map[string]string) error {
	var setters string
	for key, value := range args {
		setters += key + "=" + value + ","
	}
	setters = setters[:len(setters)-1]
	query := fmt.Sprintf("UPDATE %v SET %v WHERE %v", tableName, setters, condition)

	_, err := db.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
