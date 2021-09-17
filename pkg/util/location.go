package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LocationResponse struct {
	ShowApiResCode  int                  `json:"showapi_res_code"`
	ShowApiResError string               `json:"showapi_res_error"`
	ShowApiResBody  LocationBodyResponse `json:"showapi_res_body"`
}

type LocationBodyResponse struct {

	// 0为成功，其他失败。失败时不扣点数
	RetCode int `json:"ret_code"`

	// 1移动    2电信    3联通
	Type int `json:"type"`

	// 号段 1890871
	Num int `json:"num"`

	// 网络 服务商
	Name string `json:"name"`

	// 此地区身份证号开头几位 530000
	ProvCode string `json:"prov_code"`

	// 邮编 650000
	PostCode string `json:"post_code"`

	// 省份 云南
	Prov string `json:"prov"`

	// 城市 昆明
	City string `json:"city"`

	// 区号 0871
	AreaCode string `json:"area_code"`
}

func GetLocationWithPhone(phone string) (LocationResponse, error) {
	url := "http://showphone.market.alicloudapi.com/6-1?num=" + phone
	appCode := "ee6e2d724f3a42f4881c36ec90b9ef4c"
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", "APPCODE "+appCode)

	var out LocationResponse
	get, err := client.Do(req)
	if err != nil {
		return out, fmt.Errorf("接口请求失败：%v", err)
	}
	defer get.Body.Close()

	res, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return out, fmt.Errorf("body读取失败：%v", err)
	}

	log.Printf("1: %#v", res)

	err = json.Unmarshal(res, &out)
	if err != nil {
		return out, fmt.Errorf("json反序列化失败：%v", err)
	}

	return out, nil
}

type RegionResponse struct {
	ShowApiResCode  int                `json:"showapi_res_code"`
	ShowApiResError string             `json:"showapi_res_error"`
	ShowApiResBody  RegionBodyResponse `json:"showapi_res_body"`
}

type RegionBodyResponse struct {

	// 0为成功，其他失败。失败时不扣点数
	RetCode int `json:"ret_code"`

	// 县
	county string `json:"county"`

	// 网络 服务商
	ISP string `json:"isp"`

	// IP
	IP string `json:"ip"`

	// 国家
	Country string `json:"country"`

	// 省份 云南
	Region string `json:"region"`

	// 城市 昆明
	City string `json:"city"`
}

func GetRegionWithIP(ip string) (RegionResponse, error) {
	url := "http://saip.market.alicloudapi.com/ip?ip=" + ip
	appCode := "ee6e2d724f3a42f4881c36ec90b9ef4c"
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", "APPCODE "+appCode)

	var out RegionResponse
	get, err := client.Do(req)
	if err != nil {
		return out, fmt.Errorf("接口请求失败：%v", err)
	}
	defer get.Body.Close()

	res, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return out, fmt.Errorf("body读取失败：%v", err)
	}

	log.Printf("2: %#v", res)

	err = json.Unmarshal(res, &out)
	if err != nil {
		return out, fmt.Errorf("json反序列化失败：%v", err)
	}

	return out, nil
}
