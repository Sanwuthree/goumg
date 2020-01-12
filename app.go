package main

import (
	"fmt"
	"github.com/Sanwuthree/gocui"
	"log"
)

var (
	preButton *gocui.View
)

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panic(err)
	}
	g.Cursor = true
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SetManagerFunc(layout)
	if err := keyBinding(g); err != nil {
		log.Panic(err)
	}
	if err := g.MainLoop(); err != nil {
		log.Panic(err)
	}
}
func layout(g *gocui.Gui) error {
	maxX, _ := g.Size()
	if v, err := g.SetView("bt1", 1, 1, 9, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//v.Highlight=true
		//g.BgColor=gocui.ColorRed
		//g.FgColor=gocui.ColorGreen
		//v.SelBgColor=gocui.ColorBlack
		//v.SelFgColor=gocui.ColorGreen
		fmt.Fprint(v, "Button1")
		//g.SetViewOnTop("bt1")
	}
	if v, err := g.SetView("bt2", 10, 1, 18, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//g.BgColor=gocui.ColorRed
		//g.FgColor=gocui.ColorGreen
		//v.SelBgColor=gocui.ColorBlack
		//v.SelFgColor=gocui.ColorGreen
		fmt.Fprint(v, "Button2")
	}
	if v, err := g.SetView("alert1", maxX/2-17, 3, maxX/2+17, 20); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorGreen
		v.Title = "提示"
		fmt.Fprint(v, "Button2")
	}
	_, y0, x1, _, _ := g.ViewPosition("alert1")
	if v2, err := g.SetView("aClose", x1+2, y0, x1+5, y0+2); err != nil {
		fmt.Fprint(v2, "xsdds")
		v2.Highlight = true
		v2.SelBgColor = gocui.ColorRed
		v2.SelFgColor = gocui.ColorWhite
	}
	return nil
}

func keyBinding(g *gocui.Gui) error {
	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, onClick)
	g.SetKeybinding("aClose", gocui.MouseLeft, gocui.ModNone, onCloseAlert)
	return nil
}

func onCloseAlert(gui *gocui.Gui, view *gocui.View) error {
	//v, err := gui.View("alert1")
	//if err!=nil{
	//	log.Panic(err)
	//	return err
	//}
	if err := gui.DeleteView("alert1"); err != nil {
		panic(err)
		return nil
	}
	return nil
}

func onClick(gui *gocui.Gui, view *gocui.View) error {
	if preButton != nil {
		preButton.Highlight = false
	}
	view.Highlight = true
	preButton = view
	gui.SetCurrentView(view.Name())

	return nil
}

func quit(g *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
