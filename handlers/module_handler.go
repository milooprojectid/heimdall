package handler

import (
	e "heimdall/errors"
	s "heimdall/service"
)

func getSentimentAnalysis(data Data, context Context) (interface{}, error) {
	var body map[string]string
	data.Body(&body)

	if text, ok := body["text"]; !ok || len(text) == 0 {
		return nil, &e.HTTPError{"Bad Request", 400, "text field required"}
	}

	service, err := s.GetService("morbius")
	if err != nil {
		return nil, &e.HTTPError{"Internal Server Error", 500, "Fail getting service, " + err.Error()}
	}

	result, err := service.Call("sentiment", body)
	if err != nil {
		return nil, &e.HTTPError{"Internal Server Error", 500, "Fail calling service, " + err.Error()}
	}

	return result, nil
}
