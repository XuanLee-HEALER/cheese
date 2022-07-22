package dbop_test

import (
	"cheese/dbop"
	"cheese/entity"
	"testing"
)

func TestQuery(t *testing.T) {
	r, err := dbop.DbInst.Query(entity.RoleUrlQueryAll, entity.TransToRoleUrl)
	if err != nil {
		t.Error(err)
	}
	if len(r) != 53 {
		t.Error("query error")
	}
}

func TestInsertOne(t *testing.T) {
	tt := entity.RoleUrl{
		RoleName: "test",
		RoleUrl:  "test.html",
	}
	sql, err := entity.RoleUrlToInsert(tt)
	if err != nil {
		t.Error(err)
	}
	r, err := dbop.DbInst.InsertOne(sql)
	if err != nil {
		t.Error(err)
	}
	if r != 1 {
		t.Error("insert error")
	}
}

func TestInsertMany(t *testing.T) {
	rus := []entity.RoleUrl{
		{
			RoleName: "test1",
			RoleUrl:  "test.html",
		},
		{
			RoleName: "test2",
			RoleUrl:  "test.html",
		},
		{
			RoleName: "test3",
			RoleUrl:  "test.html",
		},
	}
	sql, err := entity.RoleUrlToInsert(rus...)
	t.Log("sql:", sql)
	if err != nil {
		t.Error(err)
	}
	r, err := dbop.DbInst.InsertMany(sql)
	if err != nil {
		t.Error(err)
	}
	if r != len(rus) {
		t.Error("insert error")
	}
}

func TestDelete(t *testing.T) {
	rus := []entity.RoleUrl{
		{
			RoleName: "test",
		},
		{
			RoleName: "test1",
		},
		{
			RoleName: "test2",
		},
		{
			RoleName: "test3",
			RoleUrl:  "test.html",
		},
	}
	sql, err := entity.RoleUrlToDelete(rus...)
	t.Log("sql:", sql)
	if err != nil {
		t.Error(err)
	}
	r, err := dbop.DbInst.Delete(sql)
	if err != nil {
		t.Error(err)
	}
	if r != len(rus) {
		t.Error("insert error")
	}
}
