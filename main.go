package main

import (
	_ "beego-app/routers"

	"beego-app/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/auth"
	"log"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Limiter initialization
	r, err := utils.InitLimiter()
	if err != nil {
		log.Fatal(err)
	}

	// Package utils/ratelimit provides handlers to rate limit
	beego.InsertFilter("/*", beego.BeforeRouter, func(c *context.Context) {
		utils.RateLimit(r, c)
	})

	// Package auth provides handlers to enable basic auth support
	authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
	beego.InsertFilter("*", beego.BeforeRouter, authPlugin)

	beego.Run()
}

func SecretAuth(username, password string) bool {
	// your own basic auth logic
	return username == "test" && password == "123456"
}
