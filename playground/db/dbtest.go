package main

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/client"
)

func main() {
	conn, _ := client.Connect("localhost:3306", "root", "123456aa", "cheese")
	defer conn.Close()

	conn.Ping()

	r, _ := conn.Execute(`insert into role_url(roleName, roleUrl) values ("test", "/obs/test")`)
	fmt.Println(r.InsertId, " ", r.AffectedRows)

	r, _ = conn.Execute(`select roleName, roleUrl from role_url`)
	defer r.Close()

	for _, rows := range r.Values {
		for _, row := range rows {
			fmt.Printf("%s,", string(row.AsString()))
		}
		fmt.Println()
	}
}
