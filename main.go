package main

import (
	"Proxy/components/modules"
	"Proxy/components/utils"
	"fmt"
	"time"

	"github.com/zenthangplus/goccm"
)

func main() {
	utils.PrintLogo()
	utils.LoadConfig()

	if utils.Config.Options.Scrape {
		modules.Scrape()
	}
	
	proxies, err := utils.ReadLines("proxies.txt")
	if utils.HandleError(err) {
		return
	}

	proxies = utils.RemoveDuplicateStr(proxies)
	utils.Log(fmt.Sprintf("Loaded %d proxies", len(proxies)))

	StartTime := time.Now()
	c := goccm.New(utils.Config.Options.Threads)

	for _, proxy := range proxies {
		c.Wait()

		go func(proxy string) {
			modules.CheckProxy(proxy)
			c.Done()
		}(proxy)
	}
	
	c.WaitAllDone()
	utils.Log(fmt.Sprintf("Checked %d proxies in %fs", len(proxies), time.Since(StartTime).Seconds()))
}
