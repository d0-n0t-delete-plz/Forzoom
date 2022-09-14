package main

import (
	// "Forzoom/request"
	// "Forzoom/print"
	// "fmt"
	"Forzoom/render"
	"github.com/eiannone/keyboard"
	// "os/exec"
)

var (
	Button = []string{"Update", "Zoom", "Focus"}
	menus = render.Menus{Index: 1, Button: Button, Page: 1}
)

func main(){
	// conf := request.ReadConfig()
	// str, less := print.PrintCli(conf.Mode, conf.Focus)

	// fmt.Println(str)
	// fmt.Println(less)

	// fmt.Scanln()

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	menus.Print()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Key == keyboard.KeyEsc {
			break
		} else if event.Key == keyboard.KeyArrowUp {
			menus.Up()
		} else if event.Key == keyboard.KeyArrowDown {
			menus.Down()
		} else if event.Key == keyboard.KeyEnter {
			menus.Enter()
		}
	}
}
