package handler

import (
	"fmt"
	"heimdall/errors"
	"strconv"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

// Output Explain..
type Output struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

// Context Explain..
type Context struct {
	UserID    string
	Clearance int
}

type queryData map[string]string

type paramsData map[string]interface{}

// Data Explain..
type Data struct {
	Body   func(payload interface{}) error
	Query  queryData
	Params paramsData
}

type handler func(data Data, context Context) (interface{}, error)

// Params Explain..
type Params = []string

// --

func irisHandler(handler handler, params interface{}) iris.Handler {
	defaultError := iris.NewProblem().Detail("WHOOPS").Status(500)

	return func(ctx iris.Context) {
		// Assign Query
		queryPayload := ctx.URLParams()

		// Assign Params
		paramsPayload := map[string]interface{}{}
		if rawParam, ok := params.([]string); ok {
			for _, param := range rawParam {
				if p := ctx.Params().Get(param); p != "" {
					paramsPayload[param] = p
				}
			}
		}

		// Construct Data and Context
		data := Data{
			Body:   ctx.ReadJSON,
			Query:  queryPayload,
			Params: paramsPayload,
		}

		var context Context
		if user, ok := ctx.Values().Get("auth").(*jwt.Token); ok {
			claims := user.Claims.(jwt.MapClaims)
			context = Context{
				UserID:    claims["uid"].(string),
				Clearance: int(claims["acc"].(float64)),
			}
		}

		if data, err := handler(data, context); err == nil {
			ctx.JSON(data)
		} else if httpError, ok := err.(*errors.HTTPError); ok {
			ctx.Problem(iris.NewProblem().Detail(httpError.Error()).Status(httpError.Status).Title(httpError.Name))
		} else {
			ctx.Problem(defaultError)
		}
	}
}

func (p paramsData) GetInt(key string, def int) (int, error) {
	if value, ok := p[key].(int); ok {
		return value, nil
	} else if value, err := strconv.Atoi(p[key].(string)); err == nil {
		return value, nil
	}

	return 0, fmt.Errorf("error converting value")
}

func (p paramsData) GetString(key string) (string, error) {
	if value, ok := p[key].(string); ok {
		return value, nil
	}

	return "", fmt.Errorf("error converting value")
}
