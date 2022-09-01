package main

import (
	"Proxy/components/utils"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/zenthangplus/goccm"
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
	if utils.HandleError(err) {
		return nil, err
	}

	return res, nil
}

func CheckProxy(Proxy string) {
	_, err := ProxyReq(utils.Config.Filter.Domain, Proxy)

	if err != nil {
		utils.Log(fmt.Sprintf("[DEAD]  %s", Proxy))
		return
	}
	
	utils.Log(fmt.Sprintf("[ALIVE] %s", Proxy))
	utils.AppendFile("checked.txt", Proxy)
}

func main() {
	utils.PrintLogo()
	utils.LoadConfig()

	proxies, err := utils.ReadLines("proxies.txt")
	if utils.HandleError(err) {
		return
	}

	utils.Log(fmt.Sprintf("Loaded %d proxies", len(proxies)))

	StartTime := time.Now()
	c := goccm.New(800)
	for _, proxy := range proxies {
		c.Wait()

		go func(proxy string) {
			CheckProxy(proxy)
			c.Done()
		}(proxy)
	}

	c.WaitAllDone()
	utils.Log(fmt.Sprintf("Checked %d proxies in %fs", len(proxies), time.Since(StartTime).Seconds()))
}
