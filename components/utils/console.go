package utils

import (
	"fmt"
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
	fmt.Println(Content)
}
