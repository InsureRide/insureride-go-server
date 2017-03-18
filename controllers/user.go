package controllers

import (
	"github.com/scmo/starthack/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}


// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, err := u.GetInt(":uid")
	if err != nil {
		u.Data["json"] = err.Error()
	}
	user, err := models.GetUser(uid)
	u.Data["json"] = user

	u.ServeJSON()
}

