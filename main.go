package main

import (
	_ "github.com/scmo/insureride-go-server/routers"

	"github.com/astaxie/beego"
	"github.com/scmo/insureride-go-server/ethereum"
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
	beego.Run()
}
