package structures

type GetAccessTokenResp struct {
	AccessToken string `json:"access_token"`
}

type RespFailStruct struct {
	Message string `json:"message"`
}
