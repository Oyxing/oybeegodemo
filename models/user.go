package models

import (
	"fmt"
	"strconv"
)

// 添加用户
func AddUser(name string, age int64) (int64, error) {
	return db.Insert(&User{Name: name, Age: age})
}

// 添加mont
func AddMont(id string, msg string, code int64) (int64, error) {
	return db.Insert(&Mode{Uuid: id, Msg: msg, Code: code})
}

// 更新用户
func UpdateUser(user *User) (int64, error) {
	fmt.Println(user)
	return db.Id(1).Update(user)
}
func GetUsermsg() ([]User, error) {
	user := make([]User, 0)
	db.Find(&user)
	for k, v := range user {
		mode := GetMode(strconv.Itoa(v.Id))
		user[k].Mode = mode
	}

	return user, nil
}
func GetUser() ([]User, error) {
	user := make([]User, 0)
	db.Find(&user)
	return user, nil
}

func GetMode(uuid string) []Mode {
	mode := make([]Mode, 0)
	db.Where("uuid=?", uuid).Find(&mode)
	return mode
}

func Queryuser() []User {
	user := []User{}
	db.Find(&user)
	return user
}
