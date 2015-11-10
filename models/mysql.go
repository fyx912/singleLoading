package models

import (
	"database/sql"
	"fmt"
	"log"
	_ "singleLoading/utils/mysql"
)

//主要go外部文件引用函数,函数名称头字母要大写

//打开数据库
func OpenDatabase() (db *sql.DB, err error) {
	return sql.Open("mysql", "root:123456@tcp(192.168.1.135:3306)/testGo?charset=utf8")
}

//数据库关闭连接
func CloseDatebase() {
	db, err := OpenDatabase()
	CheckError(err)
	db.Close()
}

//检测errort

//主要go外部文件引用函数,函数名称头字母要大写func C
func CheckError(err error) {
	if err != nil {
		fmt.Println("err==", err)
		log.Println(err)
	}
}
