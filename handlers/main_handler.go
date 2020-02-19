package handler

func getProfileDetail(data Data, context Context) (interface{}, error) {
	return Output{
		Message: "profile data retrieved",
	}, nil
}

func getRootResponse(data Data, context Context) (interface{}, error) {
	return Output{
		Message: "Heimdall AI Model's Broker",
	}, nil
}
