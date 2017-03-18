package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:CarController"] = append(beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:CarController"],
		beego.ControllerComments{
			Method: "GetCar",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:CarController"] = append(beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:CarController"],
		beego.ControllerComments{
			Method: "GetDriveByCar",
			Router: `/:carId/drive`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:CarController"] = append(beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:CarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:carId/drive`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/scmo/insureride-go-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
