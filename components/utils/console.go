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
	██╗  ██╗██████╗ ██████╗  ██████╗ ██╗  ██╗██╗   ██╗
	╚██╗██╔╝██╔══██╗██╔══██╗██╔═══██╗╚██╗██╔╝╚██╗ ██╔╝
	 ╚███╔╝ ██████╔╝██████╔╝██║   ██║ ╚███╔╝  ╚████╔╝ 
	 ██╔██╗ ██╔═══╝ ██╔══██╗██║   ██║ ██╔██╗   ╚██╔╝  
	██╔╝ ██╗██║     ██║  ██║╚██████╔╝██╔╝ ██╗   ██║   
	╚═╝  ╚═╝╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   
													
	`)
}

func Log(Content string) {
	date := strings.ReplaceAll(time.Now().Format("15:04:05"), ":", "<fg=353a3b>:</>")
	content := fmt.Sprintf("[%s] %s.", date, Content)

	content = strings.ReplaceAll(content, "DEAD", "<fg=f5291b>DEAD</>")
	content = strings.ReplaceAll(content, "ALIVE", "<fg=61eb42>ALIVE</>")

	color.Println(content)
}
