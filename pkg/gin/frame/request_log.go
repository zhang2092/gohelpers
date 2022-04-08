package frame

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyydhl-zhang/common-module/pkg/logger"
)

type requestLog struct {
	RequestID       string      `json:"request_id"`
	RequestTime     string      `json:"request_time"`
	RequestMethod   string      `json:"request_method"`
	RequestUri      string      `json:"request_uri"`
	RequestProto    string      `json:"request_proto"`
	RequestUa       string      `json:"request_ua"`
	RequestReferer  string      `json:"request_referer"`
	RequestPostData string      `json:"request_post_data"`
	RequestClientIp string      `json:"request_client_ip"`
	ResponseCode    int         `json:"response_code"`
	ResponseMsg     string      `json:"response_msg"`
	ResponseData    interface{} `json:"response_data"`
	TimeUsed        string      `json:"time_used"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 获取请求参数
		requestBody := getRequestBody(c)

		// 处理请求
		c.Next()

		var requestID string
		var responseCode int
		var responseMsg string
		var responseData interface{}
		responseBody := bodyLogWriter.body.String()
		if len(responseBody) > 0 {
			response := response{}
			err := json.Unmarshal([]byte(responseBody), &response)
			if err == nil {
				requestID = response.RequestId
				responseCode = response.Code
				responseMsg = response.Message
				responseData = response.Data
			}
		}

		// 日志格式
		accessLog := &requestLog{
			RequestID:       requestID,
			RequestTime:     startTime.Format("2006-01-02 15:04:05.999999999"),
			RequestMethod:   c.Request.Method,
			RequestUri:      c.Request.RequestURI,
			RequestProto:    c.Request.Proto,
			RequestUa:       c.Request.UserAgent(),
			RequestReferer:  c.Request.Referer(),
			RequestPostData: requestBody,
			RequestClientIp: c.ClientIP(),
			ResponseCode:    responseCode,
			ResponseMsg:     responseMsg,
			ResponseData:    responseData,
			TimeUsed:        fmt.Sprintf("%s", time.Since(startTime)), // 记录请求所用时间
		}

		l, err := json.Marshal(accessLog)
		if err == nil {
			logger.Logger.Infof("%s", l)
		} else {
			logger.Logger.Infof("%v", accessLog)
		}
	}
}

// getRequestBody 获取请求参数
func getRequestBody(c *gin.Context) string {
	switch c.Request.Method {
	case http.MethodPost:
		fallthrough
	case http.MethodPut:
		fallthrough
	case http.MethodPatch:
		// ioutil
		//var bodyBytes []byte
		//bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		//if err != nil {
		//	return ""
		//}
		//
		//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//return string(bodyBytes)

		// buffer
		var buffer [512]byte
		result := bytes.NewBuffer(nil)
		for {
			n, err := c.Request.Body.Read(buffer[0:])
			result.Write(buffer[0:n])
			if err != nil && err == io.EOF {
				break
			} else if err != nil {
				return ""
			}
		}

		c.Request.Body = ioutil.NopCloser(result)
		return formatStr(result.String())
	}

	return ""
}

func formatStr(content string) string {
	if len(content) == 0 {
		return ""
	}

	content = strings.Replace(content, " ", "", -1)
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "\r", "", -1)
	return content
}
