package models

import (
	// "database/sql"
	"fmt"
)

type User struct {
	Username string
	Password string
}

//说明:flag 为函数返回值;username,password为接收参数
func ValidationLogin(usernames string, passwords string) (falg bool) {
	db, err := OpenDatabase() //连接数据函数
	CheckError(err)
	rows, err := db.Query(" SELECT account,password FROM User where 1=1 and account=? and password=? ", usernames, passwords)
	// rows, err := db.Query(" SELECT account,password FROM User ")
	CheckError(err)
	defer rows.Close()
	columns, _ := rows.Columns()
	fmt.Println("columns=", columns)
	var user User
	for rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		CheckError(err)
		user.Username = username
		user.Password = password
		fmt.Printf("%s is %v \n", username, password)
	}
	if usernames == user.Username && passwords == user.Password {
		falg = true
		fmt.Println("数据验证成功!", true)
	} else {
		falg = false
		fmt.Println("数据验证失败!", false)
	}
	return falg
}

func ValidationLoginTest() {
	db, err := OpenDatabase()
	CheckError(err)
	usernames := "system"
	passwords := "123456"
	rows, err := db.Query(" SELECT account,password FROM User where 1=1 and account=? and password=? ", usernames, passwords)
	// rows, err := db.Query(" SELECT account,password FROM User ")
	CheckError(err)
	defer rows.Close()
	columns, _ := rows.Columns()
	fmt.Println("columns=", columns)
	var user User
	for rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		CheckError(err)
		user.Username = username
		user.Password = password
		fmt.Printf("%s is %v \n", username, password)
	}
	if usernames == user.Username && passwords == user.Password {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

//返回map类型,如: map[account:system password:123456]
func MapSelect() {
	db, err := OpenDatabase()
	CheckError(err)
	rows, err := db.Query(" SELECT account,password FROM User ")
	CheckError(err)
	defer rows.Close()
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//说明:flag 为函数返回值;username为接收参数
func ValidationUsername(usernames string) (falg bool) {
	db, err := OpenDatabase() //连接数据函数
	CheckError(err)
	rows, err := db.Query(" SELECT account,password FROM User where 1=1 and account=? ", usernames)
	// rows, err := db.Query(" SELECT account,password FROM User ")
	CheckError(err)
	defer rows.Close()
	columns, _ := rows.Columns()
	fmt.Println("columns=", columns)
	var user User
	for rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		CheckError(err)
		user.Username = username
		fmt.Println("%s ", username)
	}
	if usernames == user.Username {
		falg = true
		fmt.Println("数据验证成功!", true)
	} else {
		falg = false
		fmt.Println("数据验证失败!", false)
	}
	return falg
}
