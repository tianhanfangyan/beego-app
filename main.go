package main

import (
	_ "beego-app/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/auth"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Package auth provides handlers to enable basic auth support
	authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
	beego.InsertFilter("*", beego.BeforeRouter, authPlugin)

	beego.Run()
}

func SecretAuth(username, password string) bool {
	// your own basic auth logic
	return username == "test" && password == "123456"
}