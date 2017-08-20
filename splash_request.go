package splash

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
)

type SplashRequest struct {
	SplashHost string
	SplashPort string
}

func NewSplashRequest(splash_host string, splash_port string) (*SplashRequest, error) {
	request := SplashRequest{
		SplashHost: splash_host,
		SplashPort: splash_port,
	}
	return &request, nil
}

func (this *SplashRequest) formatRequestUrl(url string, option *Option, js_render bool) string {
	return ""
}

func (this *SplashRequest) Get(url string, option *Option) (*Response, error) {
	request_url := fmt.Sprintf("http://localhost:8050/render.json?url=%s&png=1&html=1", url)
	resp, err := http.Get(request_url)
	if err != nil {
		return nil, err
	}
	byte_data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	JSON, err := simplejson.NewJson(byte_data)
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

func (this *SplashRequest) GetWithJS() {
	// request_url =
}
