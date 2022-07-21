package dbop

import (
	"errors"
	"log"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
)

type (
	Insertable interface {
		ToInsert() string
	}
	Updatable interface {
		ToUpdate() string
	}
	Deletable interface {
		ToDelete() string
	}
	Queriable interface {
		ToQuery(id int, args ...interface{}) string
	}
)

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

func (db Db) InsertOne(o Insertable) (int, error) {
	r, err := db.conn.Execute(o.ToInsert())
	if err != nil {
		return 0, err
	}
	if r.AffectedRows > 0 {
		return int(r.AffectedRows), nil
	}
	return 0, errors.New("insert failed")
}

func (db Db) insertMany(sql string) (int, error) {
	return 0, nil
}

func (db Db) updateOne(sql string) (int, error) {
	return 0, nil
}

func (db Db) updateMany(sql string) (int, error) {
	return 0, nil
}

func (db Db) delete(sql string) (int, error) {
	return 0, nil
}

func (db Db) queryOne(sql string) (int, error) {
	return 0, nil
}

func (db Db) QueryMany(o Queriable, id int, args ...interface{}) ([]Queriable, error) {
	var result mysql.Result
	var res []Queriable = make([]Queriable, 48)
	err := db.conn.ExecuteSelectStreaming(o.ToQuery(id, args...), &result, func(row []mysql.FieldValue) error {
		var t Queriable
		err := toStructVal(row, &t)
		if err != nil {
			return err
		}
		res = append(res, t)
		return nil
	}, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func toStructVal(row []mysql.FieldValue, pStruct *Queriable) error {
	return nil
}
