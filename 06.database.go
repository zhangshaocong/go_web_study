package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "root:123456@/hello?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("warcello1","研发部","2021-01-11")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	//更新数据
	stmt1, err := db.Prepare("update userinfo set username = ? where uid=?")
	checkErr(err)
	res, err = stmt1.Exec("hello", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid,&username, &departname,&created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)
	}
	//删除数据
	del, err := db.Prepare("delete from userinfo where uid = ?")
	checkErr(err)
	res1, err := del.Exec(2)
	checkErr(err)
	affect1, err := res1.RowsAffected()
	checkErr(err)
	fmt.Println(affect1)
}
func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}