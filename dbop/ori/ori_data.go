package ori

// func (db Db) InsertRoleUrl(role, url string) (int, error) {
// 	var sql string = fmt.Sprintf(`insert into role_url values ("%s", "%s")`, role, url)
// 	r, err := db.conn.Execute(sql)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer r.Close()

// 	return int(r.AffectedRows), nil
// }

// func (db Db) SelectRoleUrlByName(roleName string) (string, error) {
// 	var sql string = fmt.Sprintf(`select roleUrl from role_url where roleName="%s"`, roleName)
// 	r, err := db.conn.Execute(sql)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer r.Close()

// 	res, err := r.GetString(0, 0)
// 	if err != nil {
// 		return "", err
// 	}
// 	return res, nil
// }

// func (db Db) SelectAllRoleUrl() ([]RoleUrl, error) {
// 	var res = []RoleUrl{}
// 	var sql string = "select roleName, roleUrl from role_url"
// 	var r mysql.Result
// 	err := db.conn.ExecuteSelectStreaming(sql, &r, func(row []mysql.FieldValue) error {
// 		var tRoleName, tRoleurl string
// 		for idx, field := range row {
// 			if string(r.Fields[idx].Name) == "roleName" {
// 				tRoleName = string(field.AsString())
// 			}
// 			if string(r.Fields[idx].Name) == "roleUrl" {
// 				tRoleurl = string(field.AsString())
// 			}
// 		}
// 		res = append(res, RoleUrl{
// 			RoleName: tRoleName,
// 			RoleUrl:  tRoleurl,
// 		})
// 		return nil
// 	}, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
