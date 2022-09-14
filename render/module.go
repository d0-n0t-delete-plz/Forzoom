package render

import (
	"fmt"
	"os"
    "os/exec"
	"runtime"
	"golang.org/x/sys/windows"
	"time"
	"Forzoom/print"
)

var (
	date = string([]rune(time.Now().Format(time.RFC3339))[0:10])
)

type Menus struct{
	Index int
	Button []string
	Page int
	Cache string
	Lesson []print.Less
}

func New(Button []string) Menus{
	return Menus{Index: 1, Button: Button, Page: 1}
}

func (m *Menus) Print(mod, focus string){
	Clear()
	m.Cache, m.Lesson = print.PrintCli(mod, focus, date)
}

func (m *Menus) UpdateCache(){
	for i:=range m.Button{
		if i+1 == m.Index{
			fmt.Print("> " + m.Button[i] + "  \n")
		} else {
			fmt.Print(m.Button[i] + "  \n")
		}
	}
}

func (m *Menus) Up(mod, focus string){
	Clear()
	m.Index-=1
	if m.Index<1{m.Index+=1}
	m.Print(mod, focus)
}

func (m *Menus) Down(mod, focus string){
	Clear()
	m.Index+=1
	if m.Index>len(m.Button){m.Index-=1}
	m.Print(mod, focus)
}

func (m *Menus) Enter(mod, focus string){
	Clear()
	m.Print(mod, focus)
	fmt.Printf("Choose: %v", m.Button[m.Index-1])
}

func Clear(){
	switch runtime.GOOS{
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		j, _ := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
		windows.SetConsoleCursorPosition(j, windows.Coord{X: 0, Y: 0})
	}
}
