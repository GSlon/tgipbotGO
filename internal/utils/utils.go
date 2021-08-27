package utils

import (
	"net"
	"net/http"
    "io/ioutil"
	"encoding/json"
	"fmt"
	"time"
)

type Info struct {
    City   string `json:"city"`
    Region   string `json:"region"`
    Country string    `json:"country_name"`
}

func ValidateIPv4(ip string) bool {
	if net.ParseIP(ip) == nil {
        return false
    } 

	return true
}

// https://ipapi.co/
func GetIpInfo(ip string) (string, error) {
	client := &http.Client{
        Timeout: time.Second * 10,
    }

	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "golang application")

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var info Info 
	json.Unmarshal(body, &info)

	res := fmt.Sprintf("city: %s, region: %s, country: %s", 
						info.City, info.Region, info.Country)
	return res, nil
}