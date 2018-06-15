package main

import (
	_ "app-ms-profile/routers"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	runtime.GOMAXPROCS(4)

	// Setup logs.
	setupLogs()

	// Setup CORS.
	setupCORS()

	// Setup static routes.
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	// Run and serve.
	logs.Info("The app.rest is set up correctly.")
	logs.Info("Listen and serve at %s", beego.AppConfig.String("httpport"))
	beego.Run()
}

/** Setup application logger. */
func setupLogs() {
	// Async logs to improve performance.
	logs.Async()

	// Show method and line number configuration.
	// Set to false to improve performance.
	logs.EnableFuncCallDepth(true)

	// Setup logger.
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"app-ms-profile.log", 
	"daily":true, "maxdays":7, "separate":["error", "info", "debug"]}`)
}

/** Setup CORS. */
func setupCORS() {
	// Allow CORS.
	// Allowed methods.
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
