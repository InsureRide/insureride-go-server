package models

import (
	"errors"
	"github.com/astaxie/beego"
)

var (
	UserList map[int]*User
)

func init() {
	UserList = make(map[int]*User)
	u := User{1, "astaxie", "11111", Profile{"male", 25, "Gøteborg", "volvofan@gmail.com"}, beego.AppConfig.String("demoCarAddress")}
	UserList[u.Id] = &u
}

type User struct {
	Id       int
	Username string
	Password string
	Profile  Profile
	CarAddress string
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}



func GetUser(uId int) (u *User, err error) {
	if u, ok := UserList[uId]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[int]*User {
	return UserList
}


