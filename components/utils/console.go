package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/inancgumus/screen"
)

func PrintLogo() {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Println(`
	__  ______                      
	\ \/ /  _ \ _ __ _____  ___   _ 
	 \  /| |_) | '__/ _ \ \/ / | | |
	 /  \|  __/| | | (_) >  <| |_| |
	/_/\_\_|   |_|  \___/_/\_\\__, |
                                   |___/ 
   `)
}

func Log(Content string) {
	date := strings.ReplaceAll(time.Now().Format("15:04:05"), ":", "<fg=353a3b>:</>")
	content := fmt.Sprintf("[%s] [%d] %s.", date, Valid, Content)

	content = strings.ReplaceAll(content, "DEAD", "<fg=f5291b>DEAD</>")
	content = strings.ReplaceAll(content, "ALIVE", "<fg=61eb42>ALIVE</>")

	for _, element := range []string{"(", ")", "[", "]", "#"} {
		content = strings.ReplaceAll(content, element, fmt.Sprintf("<fg=3d3d3d>%s</>", element))
	}

	color.Println(content)
}

func HandleError(Err error) bool {
	if Err != nil {
		if Config.Dev.Debug {
			fmt.Println(Err)
		}
		return true
	}

	return false
}