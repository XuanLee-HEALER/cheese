package entity

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-mysql-org/go-mysql/mysql"
)

const (
	RoleUrlQueryAll = `select roleName, roleUrl from role_url`
	RoleUrlCount    = `select count(*) from role_url`
)

type RoleUrl struct {
	RoleName string
	RoleUrl  string
}

func (ru RoleUrl) String() string {
	return fmt.Sprintf("RoleUrl: name={%s} url={%s}", ru.RoleName, ru.RoleUrl)
}

// func (ru RoleUrl) ToQueryWith(args ...any) string {
// 	return fmt.Sprintf(`insert into Role`, ru.RoleName, ru.RoleUrl)
// }

func RoleUrlToInsert(rus ...RoleUrl) (string, error) {
	ll := len(rus)
	switch {
	case ll <= 0:
		return "", errors.New("objects array length less than or equal zero")
	case ll == 1:
		return fmt.Sprintf(`insert into role_url values ("%s", "%s")`, rus[0].RoleName, rus[0].RoleUrl), nil
	case ll > 1:
		subs := make([]string, 0)
		for _, ru := range rus {
			subs = append(subs, fmt.Sprintf(`("%s", "%s")`, ru.RoleName, ru.RoleUrl))
		}
		return `insert into role_url values ` + strings.Join(subs, ","), nil
	}
	return "", errors.New("roleurl to insert sql unknown error")
}

func TransToRoleUrl(fields []mysql.FieldValue) any {
	var tName, tUrl string
	for idx, val := range fields {
		if idx == 0 {
			tName = string(val.AsString())
		}
		if idx == 1 {
			tUrl = string(val.AsString())
		}
	}
	return RoleUrl{
		RoleName: tName,
		RoleUrl:  tUrl,
	}
}
