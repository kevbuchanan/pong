package pong

import (
	nc "github.com/rthornton128/goncurses"
	"log"
	"os"
	"os/signal"
)

type NCursesWindow struct {
	window *nc.Window
}

func (window NCursesWindow) ColorOn(color int) {
	window.window.ColorOn(int16(color))
}

func (window NCursesWindow) ColorOff(color int) {
	window.window.ColorOff(int16(color))
}

func (window NCursesWindow) AttributeOn(attr int) {
}

func (window NCursesWindow) AttributeOff(attr int) {
}

func (window NCursesWindow) Print(y int, x int, text string) {
	window.window.MovePrint(y, x, text)
}

func (window NCursesWindow) Move(y int, x int) {
	window.window.MoveWindow(y, x)
}

func (window NCursesWindow) MaxRowAndColumn() (int, int) {
	return window.window.MaxYX()
}

func (window NCursesWindow) CurrentRowAndColumn() (int, int) {
	return window.window.YX()
}

func (window NCursesWindow) Delete() {
	window.window.Delete()
}

type NCursesUi struct {
	screen *nc.Window
}

func (ui NCursesUi) Draw(window Window) {
	ncwindow := window.(NCursesWindow)
	ui.screen.Overlay(ncwindow.window)
}

func (ui NCursesUi) GetChar() rune {
	return rune(ui.screen.GetChar())
}

func (ui NCursesUi) Erase() {
	ui.screen.Erase()
}

func (ui NCursesUi) Refresh() {
	ui.screen.Refresh()
}

func (ui NCursesUi) NewWindow(height int, width int) Window {
	window, _ := nc.NewWindow(height, width, 0, 0)
	return NCursesWindow{
		window: window,
	}
}

func (ui NCursesUi) MaxRowAndColumn() (int, int) {
	log.Println(ui.screen.MaxYX())
	return ui.screen.MaxYX()
}

func handleInterrupt(signals chan os.Signal) {
	for _ = range signals {
		nc.End()
		os.Exit(1)
	}
}

func initNc() Ui {
	screen, err := nc.Init()
	if err != nil {
		os.Exit(1)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, os.Kill)
	go handleInterrupt(signals)
	nc.StartColor()
	nc.Cursor(0)
	nc.Echo(false)

	nc.InitPair(1, nc.C_WHITE, nc.C_BLACK)
	nc.InitPair(2, nc.C_YELLOW, nc.C_BLACK)
	nc.InitPair(3, nc.C_BLACK, nc.C_WHITE)

	return NCursesUi{
		screen: screen,
	}
}

func NewUi() Ui {
	return initNc()
}
