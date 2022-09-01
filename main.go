package main

import (
	"Proxy/components/utils"
	"github.com/zenthangplus/goccm"
)

func CheckProxy(Proxy string) {

}

func main() {
	utils.PrintLogo()

	proxies, err := utils.ReadLines("proxies.txt")
	if err != nil {
		return
	}

	c := goccm.New(30)
	for _, proxy := range proxies {
		c.Wait()

		go func(proxy string) {
			CheckProxy(proxy)
			c.Done()
		}(proxy)
	}

	c.WaitAllDone()
}
