package controllers

import (
	"github.com/scmo/starthack/models"

	"github.com/astaxie/beego"
	"github.com/scmo/starthack/ethereum"
	"strconv"
	"encoding/json"
)

// Operations about object
type CarController struct {
	beego.Controller
}



// @Title Get Car by UserId
// @Description find object by userId
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Car
// @Failure 403 :carId is empty
// @router /:userId [get]
func (c *CarController) GetCar() {
	userId := c.Ctx.Input.Param(":userId")
	if userId != "" {
		carObj := models.Car{}

		i, err := strconv.Atoi(userId)
		user, err := models.GetUser(i)
		if( err != nil){
			beego.Error(err)
		}
		carObj.ContractAddress = user.CarAddress
		beego.Debug("car address", carObj.ContractAddress)
		car, err := ethereum.GetCar(&carObj)
		_ = car
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = car
		}
	}
	c.ServeJSON()
}

// @Title Get drives by car
// @Description find object by carId
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Car
// @Failure 403 :carId is empty
// @router /:carId/drive [get]
func (c *CarController) GetDriveByCar() {
	carId := c.Ctx.Input.Param(":carId")

	//TODO
	c.Data["json"] = carId
	c.ServeJSON()
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Drive	true		"The object content"
// @Success 200 {string} models.Drive.Id
// @Failure 403 body is empty
// @router /:carId/drive [post]
func (c *CarController) Post() {

	carId := c.Ctx.Input.Param(":carId")

	carObj := models.Car{}
	carObj.ContractAddress = carId
	car, _ := ethereum.GetCar(&carObj)

	_ = car
	var drive models.Drive
	json.Unmarshal(c.Ctx.Input.RequestBody, &drive)

	drive, _ = ethereum.AddDrive(drive)

	ethereum.AddDriveToCar(car.ContractAddress, drive.ContractAddress)

	//driveId := models.AddDrive(drive)
	c.Data["json"] = drive
	c.ServeJSON()
}






