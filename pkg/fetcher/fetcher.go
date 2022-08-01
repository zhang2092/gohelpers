package fetcher

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// Get get请求
func Get(url string, timeout int) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("[get] new request err: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[get] client do err: %v", err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response err: %v", err)
	}

	return res, nil
}

// GetJson application/json get 请求
func GetJson(url string, parameter []byte, timeout int) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	byteParameter := bytes.NewBuffer(parameter)
	req, err := http.NewRequest("GET", url, byteParameter)
	if err != nil {
		return nil, fmt.Errorf("[getjson] new request err: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[getjson] client do err: %v", err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response err: %v", err)
	}

	return res, nil
}

// PostJson application/json post 请求
func PostJson(url, parameter string, timeout int) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	byteParameter := bytes.NewBuffer([]byte(parameter))
	request, err := http.NewRequest("POST", url, byteParameter)
	if err != nil {
		return nil, fmt.Errorf("[postjson] new request err: %v", err)
	}

	request.Header.Set("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("[postjson] client do err: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("request not equal 200")
	}

	all, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response err: %v", err)
	}

	return all, nil
}

// PostString application/x-www-form-urlencoded 请求
func PostString(url, parameter string, timeout int) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	byteParameter := bytes.NewBuffer([]byte(parameter))
	request, err := http.NewRequest("POST", url, byteParameter)
	if err != nil {
		return nil, fmt.Errorf("[poststring] new request err: %v", err)
	}

	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("[poststring] client do err: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("request not equal 200")
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response err: %v", err)
	}

	return all, nil
}

func determiningEncoding(r *bufio.Reader) encoding.Encoding {
	b, err := r.Peek(1024)
	if err != nil && err != io.EOF {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(b, "")
	return e
}
