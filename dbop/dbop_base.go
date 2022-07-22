package dbop

import (
	"errors"
	"log"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
)

type ToStructVal func(fields []mysql.FieldValue) any

type Db struct {
	conn *client.Conn
}

var DbInst Db

const (
	dbAddr = "localhost:3306"
	dbUser = "root"
	dbPswd = "123456aa"
	dbName = "cheese"
)

func init() {
	conn, err := client.Connect(dbAddr, dbUser, dbPswd, dbName)
	if err != nil {
		log.Fatal("connect db failed.", err)
	}
	DbInst = Db{
		conn: conn,
	}
}

func (db Db) InsertOne(sql string) (int, error) {
	r, err := db.conn.Execute(sql)
	if err != nil {
		return 0, err
	}
	if r.AffectedRows > 0 {
		return int(r.AffectedRows), nil
	}
	return 0, errors.New("insert failed")
}

func (db Db) InsertMany(sql string) (int, error) {
	db.conn.Begin()
	r, err := db.conn.Execute(sql)
	if err != nil {
		return 0, err
	}
	db.conn.Commit()
	if r.AffectedRows > 0 {
		return int(r.AffectedRows), nil
	}
	return 0, errors.New("insert failed")
}

func (db Db) updateOne(sql string) (int, error) {
	return 0, nil
}

func (db Db) updateMany(sql string) (int, error) {
	return 0, nil
}

func (db Db) Delete(sql string) (int, error) {
	r, err := db.conn.Execute(sql)
	if err != nil {
		return 0, err
	}
	return int(r.AffectedRows), nil
}

func (db Db) Query(sql string, trans ToStructVal) ([]any, error) {
	r, err := db.conn.Execute(sql)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := []any{}
	for _, row := range r.Values {
		res = append(res, trans(row))
	}
	return res, nil
}
