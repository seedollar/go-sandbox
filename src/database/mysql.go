package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "gouser:gouser@/gotest?charset=utf8")
	checkErr(err)

	defer db.Close()

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username = ?, departname = ?, created = ?")
	checkErr(err)

	// Execute statement
	res, err := stmt.Exec("seedollar", "Development", "2016-02-29")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	stmt,err = db.Prepare("UPDATE userinfo SET username = ? WHERE uid = ?")
	checkErr(err)

	res, err = stmt.Exec("seedollarNew", id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affected)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var deptName string
		var created string

		err := rows.Scan(&uid, &username, &deptName, &created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(deptName)
		fmt.Println(created)
	}

	// Delete row
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid = ?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affected, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affected)


}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
