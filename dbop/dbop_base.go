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
	return db.transaction(sql, db.InsertOne)
}

func (db Db) Update(sql string) (int, error) {
	return db.transaction(sql, db.execute)
}

func (db Db) Delete(sql string) (int, error) {
	return db.transaction(sql, db.execute)
}

func (db Db) execute(sql string) (int, error) {
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

func (db Db) transaction(sql string, f func(string) (int, error)) (rn int, rErr error) {
	defer func() {
		if e := recover(); e != nil {
			err := db.conn.Rollback()
			if err != nil {
				rn = 0
				rErr = err
			}
			rn = 0
			rErr = e.(error)
		}
	}()
	err := db.conn.Begin()
	if err != nil {
		panic(err)
	}
	rn, rErr = f(sql)
	if rErr != nil {
		panic((err))
	}
	err = db.conn.Commit()
	if err != nil {
		panic(err)
	}
	return rn, rErr
}
