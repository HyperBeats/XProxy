package modules

import (
	"Proxy/components/utils"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func ProxyReq(req string, proxy string) (res *http.Response, err error) {
	ProxyUrl, err := url.Parse("http://" + proxy)
	if utils.HandleError(err) {
		return nil, err
	}

	ReqUrl, err := url.Parse(req)
	if utils.HandleError(err) {
		return nil, err
	}

	client := &http.Client{
		Timeout:   time.Duration(time.Duration(utils.Config.Filter.Timeout) * time.Second),
		Transport: &http.Transport{Proxy: http.ProxyURL(ProxyUrl)},
	}

	res, err = client.Get(ReqUrl.String())
	return res, err
}

func CheckProxy(Proxy string) {
	_, err := ProxyReq(utils.Config.Filter.Domain, Proxy)

	if err != nil {
		utils.Log(fmt.Sprintf("[DEAD]  %s", Proxy))
		return
	}
	
	utils.Valid++
	utils.Log(fmt.Sprintf("[ALIVE] %s", Proxy))
	utils.AppendFile("checked.txt", Proxy)
}