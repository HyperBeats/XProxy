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
	timeout := time.Duration(3 * time.Second)
	proxyURL, err := url.Parse("http://" + proxy)
	reqURL, err := url.Parse(req)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	res, err = client.Get(reqURL.String())
	return res, err
}

func CheckProxy(Proxy string) {
	_, err := ProxyReq("https://discord.com", Proxy)

	if err != nil {
		utils.Log(fmt.Sprintf("[INVALID] %s", Proxy))
		return
	}
	
	utils.Log(fmt.Sprintf("[VALID] %s", Proxy))
	utils.AppendFile("checked.txt", Proxy)
}

func main() {
	utils.PrintLogo()

	proxies, err := utils.ReadLines("proxies.txt")
	if err != nil {
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
