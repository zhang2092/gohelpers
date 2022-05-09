package authorize

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type validAccessTokenResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

// ValidAccessTokenRequest 校验access_token请求
type ValidAccessTokenRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}

// ValidAccessToken 校验Token
func ValidAccessToken(url string, parameter ValidAccessTokenRequest, timeout int) bool {
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	args, err := json.Marshal(parameter)
	if err != nil {
		return false
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(args))
	if err != nil {
		return false
	}

	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return false
	}

	if response.StatusCode != 200 {
		return false
	}

	all, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	var resp validAccessTokenResponse
	err = json.Unmarshal(all, &resp)
	if err != nil {
		return false
	}

	if resp.Code == 200 && resp.Message == "ok" {
		return true
	}

	return false
}
