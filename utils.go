package goclover

import (
	"encoding/json"
	"github.com/aicam/goclover/structures"
)

func handleFailwithStatusOk(body []byte) string {
	var respErrJson structures.RespFailStruct
	err := json.Unmarshal(body, &respErrJson)

	if err != nil || respErrJson.Message == "" {
		return string(body)
	}
	return respErrJson.Message
}
