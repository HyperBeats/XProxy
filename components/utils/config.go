package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/BurntSushi/toml"
)

var (
	Config = ConfigStruct{}
	ActualIp string
	Valid = 0
)

type ConfigStruct struct {
	Filter struct {
		Timeout int    `toml:"timeout"`
	} `toml:"filter"`

	Options struct {
		Scrape  bool `toml:"scrape"`
		Threads int  `toml:"threads"`
		ScrapeThreads int `toml:"scrape_threads"`
		SaveTransparent bool `toml:"save_transparent"`
	} `toml:"options"`
}

func GetActualIp() string {
	res, err := http.Get("https://api.ipify.org")
	if HandleError(err) {
		return ""
	}
	
	defer res.Body.Close()

	if res.StatusCode == 403 || res.StatusCode == 404 {
		return ""
	}

	content, err := ioutil.ReadAll(res.Body)
	if HandleError(err) {
		return ""
	}

	return string(content)
}

func LoadConfig() {
	if _, err := toml.DecodeFile("script/config.toml", &Config); err != nil {
		panic(err)
	}
	ActualIp = GetActualIp()
}