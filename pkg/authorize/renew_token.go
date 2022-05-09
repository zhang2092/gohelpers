package authorize

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type renewAccessTokenResponse struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	Data      RenewAccessTokenBody `json:"data"`
	RequestId string               `json:"request_id"`
}

// RenewAccessTokenBody 刷新access_token返回信息
type RenewAccessTokenBody struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

// RenewAccessTokenRequest 刷新access_token请求
type RenewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// RenewAccessToken 刷新Token
func RenewAccessToken(url string, parameter RenewAccessTokenRequest, timeout int) (*RenewAccessTokenBody, error) {
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	args, err := json.Marshal(parameter)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(args))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, err
	}

	all, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp renewAccessTokenResponse
	err = json.Unmarshal(all, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code == 200 && resp.Message == "ok" {
		return &resp.Data, nil
	}

	return nil, errors.New("刷新access_token失败")
}
