package ori

import (
	"fmt"
	"log"

	"github.com/go-mysql-org/go-mysql/client"
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

func (db Db) InsertRoleUrl(role, url string) (int, error) {
	var sql string = fmt.Sprintf(`insert into role_url values ("%s", "%s")`, role, url)
	r, err := db.conn.Execute(sql)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	return int(r.AffectedRows), nil
}

func (db Db) SelectRoleUrlByName(roleName string) (string, error) {
	var sql string = fmt.Sprintf(`select roleUrl from role_url where roleName="%s"`, roleName)
	r, err := db.conn.Execute(sql)
	if err != nil {
		return "", err
	}
	defer r.Close()

	res, err := r.GetString(0, 0)
	if err != nil {
		return "", err
	}
	return res, nil
}
