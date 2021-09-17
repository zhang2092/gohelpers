package util

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(20 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error status code", resp.StatusCode)
		return nil, fmt.Errorf("error status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return io.ReadAll(utf8Reader)
}

func FetchHost(reqUrl, proxyIP string) ([]byte, error) {
	<-rateLimiter

	client := &http.Client{}

	//是否使用代理IP
	if proxyIP != "" {
		proxy, err := url.Parse(proxyIP)
		if err != nil {
			log.Fatalf("1:%v\n", err)
		}
		netTransport := &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		}
		client = &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		}
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Host", "www.stats.gov.cn")
	//req.Header.Set("Cookie", "_trs_uv=kma79tru_6_1yw6; SF_cookie_1=37059734; wzws_cid=ed0db2d09a630ccd20292459e4bfc8091d460d5990f415e3d928761ee970d90f7ba6254f9c4b8e9b79ad456094d4a1381305d0b065ff9cc539ee1775ec262f946af61ddbed371e2a6dae2bc2041a30fc")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.54")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error status code", resp.StatusCode)
		return nil, fmt.Errorf("error status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return io.ReadAll(utf8Reader)
}

// HttpPost 模拟请求方法
func HttpPost(postUrl string, headers map[string]string, jsonMap map[string]interface{}, proxyIP string) ([]byte, error) {
	client := &http.Client{}
	//转换成postBody
	bytesData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	postBody := bytes.NewReader(bytesData)

	//是否使用代理IP
	if proxyIP != "" {
		//proxy := func(_ *http.Request) (*url.URL, error) {
		//	return url.Parse(proxyIP)
		//}
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		proxyUrl, err := url.Parse(proxyIP)
		if err == nil { // 使用传入代理
			transport.Proxy = http.ProxyURL(proxyUrl)
		}
		//&http.Transport{Proxy: proxy}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}

	// get请求
	req, err := http.NewRequest("GET", postUrl, postBody)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error status code", resp.StatusCode)
		return nil, fmt.Errorf("error status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return io.ReadAll(utf8Reader)
}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil && err != io.EOF {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

// PostRequest 请求
func PostRequest(url, parameter string) ([]byte, error) {
	client := &http.Client{}
	byteParameter := bytes.NewBuffer([]byte(parameter))
	request, _ := http.NewRequest("POST", url, byteParameter)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		return nil, errors.New("网络请求失败")
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("读取网络内容失败")
	}
	return all, nil
}

// PostRequestString 请求
func PostRequestString(url, parameter string) ([]byte, error) {
	client := &http.Client{}
	byteParameter := bytes.NewBuffer([]byte(parameter))
	request, _ := http.NewRequest("POST", url, byteParameter)
	request.Header.Set("Content-type", "application/x-www-form-urlencoded")
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		return nil, errors.New("网络请求失败")
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("读取网络内容失败")
	}
	return all, nil
}
