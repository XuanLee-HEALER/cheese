package entity

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	birthReg = `(\d+)月(\d+)日`
)

type RoleBirth struct {
	month int
	day   int
}

func (rb RoleBirth) GetRoleBirthDetail() (int, int) {
	return rb.month, rb.day
}

func (rb RoleBirth) String() string {
	tc := RoleBirth{}
	if rb == tc {
		return ""
	}
	return fmt.Sprintf("%d月%d日", rb.month, rb.day)
}

func NewRoleBirth(birthText string) RoleBirth {
	reg, err := regexp.Compile(birthReg)
	if err != nil {
		return RoleBirth{}
	}
	birth := reg.FindSubmatch([]byte(birthText))
	if len(birth) == 0 {
		return RoleBirth{}
	}
	tm, _ := strconv.Atoi(string(birth[1]))
	td, _ := strconv.Atoi(string(birth[2]))
	return RoleBirth{
		month: tm,
		day:   td,
	}
}

type Role struct {
	Name        string
	ElementType byte
	Birth       RoleBirth
	From        string
	Feature     string
	Weapon      byte
	Destiny     string
	Dub         string
}

func (r Role) String() string {
	return fmt.Sprintf("【姓名】：%s\n【元素】：%s\n【生日】：%s\n【故乡】：%s\n【特点】：%s\n【武器类型】：%s\n【命座】：%s\n【称号】：%s\n", r.Name, ToElementType(r.ElementType), r.Birth, r.From, r.Feature, ToWeapon(r.Weapon), r.Destiny, r.Dub)
}

func (r Role) ToInsert() string {
	return fmt.Sprintf(`
	insert into role(
		name, 
		elementType, 
		birth, 
		fromWhere, 
		feature, 
		weapon, 
		destiny, 
		dub) values (%s, %d, %s, %s, %s, %d, %s, %s)`, r.Name, r.ElementType, r.Birth, r.From, r.Feature, r.Weapon, r.Destiny, r.Dub)
}

// func (r Role) ToUpdate() string {

// }

// func (r Role) ToDelete() string {

// }
