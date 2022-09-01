package utils

import "github.com/BurntSushi/toml"

var (
	Config = ConfigStruct{}
)

type ConfigStruct struct {
	Filter struct {
		Domain  string `toml:"domain"`
		Timeout int    `toml:"timeout"`
	} `toml:"filter"`

	Options struct {
		Scrape  bool `toml:"scrape"`
		Threads int  `toml:"threads"`
	} `toml:"options"`
}

func LoadConfig() {
	if _, err := toml.DecodeFile("script/config.toml", &Config); err != nil {
		panic(err)
	}
}