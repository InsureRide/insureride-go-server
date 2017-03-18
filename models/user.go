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
	u := User{1, "tonystark", Profile{ "Stark", "Tony", 9.8, "male", 25, "Gøteborg", "tony@starkindustires.com", "+41 79 700 88 00"}, beego.AppConfig.String("demoCarAddress")}
	UserList[u.Id] = &u
}

type User struct {
	Id       int
	Username string
	Profile  Profile
	CarAddress string
}

type Profile struct {
	Lastname	string
	Prename string
	Rating float32
	Gender  string
	Age     int
	Address string
	Email   string
	Phone string
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


