package main

import (
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
	"image"
	"time"
)

var count int
var wnd nucular.MasterWindow

type BackendConfig struct {
	HomeAssistantServer string
	HomeAssistantToken  string
}

func main() {
	wnd = nucular.NewMasterWindowSize(0, "L:A_N:application_ID:test", image.Pt(768, 1024-24), updatefn)
	wnd.SetStyle(style.FromTheme(style.WhiteTheme, 1.25))
	go clear(wnd)
	wnd.Main()
}

func updatefn(w *nucular.Window) {
	w.Row(50).Dynamic(2)
	if w.ButtonText("Exit") {
		wnd.Close()
	}
	if w.ButtonText(fmt.Sprintf("increment: %d", count)) {
		count++
	}
}

// hack to flash the buffer dark to clear ghosting
func clear(w nucular.MasterWindow) {
	originalTheme := *w.Style()
	time.Sleep(1000 * time.Millisecond)
	w.SetStyle(style.FromTheme(style.DarkTheme, 0))
	w.Changed()
	time.Sleep(200 * time.Millisecond)
	w.SetStyle(&originalTheme)
	w.Changed()
}
