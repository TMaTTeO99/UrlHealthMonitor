package external

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/TMaTTeO99/UrlHealthMonitor/config"
)

type ReqFactoryInterface interface {
	BuildTotalVirusPostReq(urlRequest string) (*http.Request, error)
	BuildTotalVirusGetReq(urlRequest string) (*http.Request, error)
}

type ReqFactoryImpl struct {
	Config *config.ConfigData
}

func (c *ReqFactoryImpl) BuildTotalVirusPostReq(urlRequest string) (*http.Request, error) {

	form := url.Values{}
	form.Set("url", urlRequest)

	req, err := http.NewRequest("POST", c.Config.VIRUSTOTAL_BASE_URL, strings.NewReader(form.Encode()))

	if err != nil {
		return nil, err
	}
	req.Header.Add("x-apikey", c.Config.API_KEY)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	return req, nil

}

func (c *ReqFactoryImpl) BuildTotalVirusGetReq(id string) (*http.Request, error) {

	fullUrl := fmt.Sprintf("%s/%s", c.Config.ANALIZE_URL_BASE_URL, id)
	fmt.Println(fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", c.Config.API_KEY)

	return req, nil

}
