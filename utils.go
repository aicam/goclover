package goclover

import "encoding/json"

func handleFailwithStatusOk(body []byte) string {
	var respErrJson RespFailStruct
	err := json.Unmarshal(body, &respErrJson)

	if err != nil || respErrJson.Message == "" {
		return string(body)
	}
	return respErrJson.Message
}
