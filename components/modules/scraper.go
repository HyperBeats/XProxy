package modules

import (
	"Proxy/components/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/zenthangplus/goccm"
)

func ScrapeUrl(Url string) {
	res, err := http.Get(Url)
	if utils.HandleError(err) {
		return
	}
	
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if utils.HandleError(err) {
		return
	}
	
	lines := strings.Split(string(content), "\n")

	for _, proxy := range lines {
		if proxy == "" {
			continue
		}

		utils.AppendFile("proxies.txt", proxy)
	}
}

func Scrape() {
	url_list, err := utils.ReadLines("url.csv")
	if utils.HandleError(err) {
		return
	}
	
	StartTime := time.Now()
	c := goccm.New(utils.Config.Options.ScrapeThreads)

	for _, url := range url_list {
		c.Wait()

		// type,url
		s := strings.Split(url, ",")

		go func(u string) {
			ScrapeUrl(u)
			c.Done()
		}(s[1])
	}

	c.WaitAllDone()
	utils.Log(fmt.Sprintf("Scraped %d urls in %fs", len(url_list), time.Since(StartTime).Seconds()))
}