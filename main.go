package main

import (
	_ "github.com/scmo/insureride-go-server/routers"

	"github.com/astaxie/beego"
	"github.com/scmo/insureride-go-server/ethereum"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	// Setup Ethereum Connection
	ethereum.Init()
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}
