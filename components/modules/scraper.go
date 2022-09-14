package modules

import (
	"Proxy/components/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/zenthangplus/goccm"
)

var (
	proxyRegex = regexp.MustCompile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}):([0-9]{1,5})")
)

func RemoveUrl(Url string, ProxyType string) {
	utils.Log(fmt.Sprintf("Dead link: %s", Url))
	utils.RemoveLine("url.csv", fmt.Sprintf("%s,%s", ProxyType, Url))
}

func ScrapeUrl(Url string, ProxyType string) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(utils.Config.Options.ScrapeTimeout),
	}

	res, err := client.Get(Url)
	if utils.HandleError(err) {
		if utils.Config.Options.RemoveUrlOnError {
			RemoveUrl(Url, ProxyType)
		}
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 403 || res.StatusCode == 404 || res.StatusCode == 401 {
		RemoveUrl(Url, ProxyType)
		return
	}

	content, err := ioutil.ReadAll(res.Body)
	if utils.HandleError(err) {
		return
	}

	for _, proxy := range proxyRegex.FindAllString(string(content), -1) {
		utils.AppendFile("proxies.txt", fmt.Sprintf("%s://%s", ProxyType, proxy))
	}
}

func Scrape() {
	url_list, err := utils.ReadLines("url.csv")
	if utils.HandleError(err) {
		return
	}

	StartTime, c, crawled := time.Now(), goccm.New(utils.Config.Options.ScrapeThreads), 0
	
	for i, url := range url_list {
		c.Wait()

		// * type,url
		s := strings.Split(url, ",")

		go func(u string, t string, n int) {
			ScrapeUrl(u, t)
			crawled++
			
			utils.Log(fmt.Sprintf("Scraped page #%d (%d/%d)", n, crawled, len(url_list)))
			c.Done()
		}(s[1], s[0], i)
	}

	c.WaitAllDone()
	utils.Log(fmt.Sprintf("Scraped %d urls in %fs", len(url_list), time.Since(StartTime).Seconds()))
}
