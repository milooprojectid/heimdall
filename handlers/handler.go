package handler

import (
	"github.com/kataras/iris/v12"
)

// Register explain..
func Register(app *iris.Application) {
	app.Get("/", irisHandler(getRootResponse, nil))
	app.Get("/profile", irisHandler(getProfileDetail, nil))

	app.Post("/module/sentiment-analysis", irisHandler(getSentimentAnalysis, nil))
}
