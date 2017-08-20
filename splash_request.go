package splash

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

type SplashClient struct {
	SplashHost string
	SplashPort string
}

// check the splash service availability
func checkSplashUrl(splash_url string) bool {
	resp, err := http.Get(splash_url)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return false
	}
	JSON, err := simplejson.NewJson(body)
	if err != nil {
		return false
	}

	err_code, _ := JSON.Get("error").Int()
	return err_code == 400
}

// format the request url with option
func (this *SplashClient) formatRequestUrl(url string, option *Option) string {
	request_url := fmt.Sprintf("http://%s:%s/render.json?url=%s&html=1",
		this.SplashHost, this.SplashPort, url)
	if option.Png {
		request_url += "&png=1"
	}
	if option.Timeout != 0 {
		request_url += fmt.Sprintf("&timeout=%d", option.Timeout)
	}
	if option.Wait >= 0.1 {
		request_url += fmt.Sprintf("&wait=%.1f", option.Timeout)
	}
	return request_url
}

func NewSplashClient(splash_host string, splash_port string) (*SplashClient, error) {
	splash_url := fmt.Sprintf("http://%s:%s", splash_host, splash_port)
	if !checkSplashUrl(splash_url + "/render.json") {
		return nil, errors.New("Splash service is unavailable at " + splash_url)
	}

	client := SplashClient{
		SplashHost: splash_host,
		SplashPort: splash_port,
	}
	return &client, nil
}

func (this *SplashClient) Get(url string, option *Option) (*Response, error) {
	request_url := this.formatRequestUrl(url, option)
	resp, err := http.Get(request_url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	JSON, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}
	response := Response{}
	response.Title, _ = JSON.Get("title").String()
	response.Url, _ = JSON.Get("url").String()
	response.RequestedUrl, _ = JSON.Get("requestedUrl").String()
	response.Html, _ = JSON.Get("html").String()
	response.Base64Png, _ = JSON.Get("png").String()

	return &response, nil
}
