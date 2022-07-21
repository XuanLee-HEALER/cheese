package entity

import "fmt"

type RoleUrl struct {
	RoleName string
	RoleUrl  string
}

func (ru RoleUrl) String() string {
	return fmt.Sprintf("RoleUrl: name={%s} url={%s}", ru.RoleName, ru.RoleUrl)
}

func (ru RoleUrl) ToQuery(id int, args ...interface{}) string {
	switch id {
	case 1:
		return ru.queryAll()
	}
	return ""
}

func (ru RoleUrl) queryAll() string {
	return `select roleName, roleUrl from role_url`
}
