package apicall

import (
	"io"
	"log"
	"net/http"
)

func MakeGetAccessTokenReq(clientId, code, clientSecret, baseUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"oauth/token", nil)
	if err != nil {
		log.Print(err)
		return []byte(""), err
	}

	q := req.URL.Query()
	q.Add("client_id", clientId)
	q.Add("client_secret", clientSecret)
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}
