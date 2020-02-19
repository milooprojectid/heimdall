package main

import (
	handler "heimdall/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
	app := iris.Default()

	handler.Register(app)

	return app
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading environment file")
	}

	app := newApp()
	port := os.Getenv("APP_PORT")
	app.Run(iris.Addr(":" + port))
}
