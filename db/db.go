package db

import (
	"database/sql"
	"fmt"

)


func DBopener() {
	Db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/coinkeeper")
	if err != nil {
		panic(err)
	}
	defer Db.Close()
	fmt.Println("Есть контакт")
}
