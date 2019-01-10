package models

import "time"

type User struct {
	Id         int
	Name       string
	Username   string
	Password   string
	Email      string
	LoginCount int
	LastTime   time.Time
	LastIp     string
	State      int8
	Created    time.Time
	Updated    time.Time
}

func (m *User) TableName() string {
	return TableName("user")
}

type RegForm struct {
	Username string `form:"email"`
	Pass     string `form:"pass"`
	Repass   string `form:"repass"`
	Name     string `form:"username"`
	Vercode  string `form:"vercode"`
}

// type LogForm struct {
// 	Username string
// 	Password string
// 	Email    string
// }
