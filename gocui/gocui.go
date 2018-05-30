package gocui

import (
	"fmt"
	"sync"
	"time"

	"github.com/iwittkau/ssh-select"

	"github.com/jroimartin/gocui"
)

type Frontend struct {
	g           *gocui.Gui
	index       int
	hasSelected bool
	conf        *sshselect.Configuration
	views       []string
	mux         *sync.Mutex
}

func New(conf *sshselect.Configuration) (*Frontend, error) {

	g, err := gocui.NewGui(gocui.OutputNormal)
	f := &Frontend{
		g,
		0,
		false,
		conf,
		[]string{"main"},
		&sync.Mutex{},
	}
	g.SetManagerFunc(f.layout)

	if err := f.initKeybindings(f.g); err != nil {
		return nil, err
	}

	return f, err
}

func (f *Frontend) Draw() (int, bool, error) {
	defer f.g.Close()

	if err := f.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return f.Index(), f.HasSelected(), err
	}

	return f.Index(), f.HasSelected(), nil
}

func (f *Frontend) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	i := f.Index()

	main, err := g.SetView("main", 0, 0, maxX-1, maxY-10)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	main.Clear()
	main.Highlight = true
	main.SelBgColor = gocui.ColorRed
	main.SelFgColor = gocui.ColorWhite

	for _, s := range f.conf.Servers {
		str := fmt.Sprintf("[%d] %s", s.Index, s.Name)
		fmt.Fprintln(main, str)
	}

	if err := main.SetCursor(0, i); err != nil {
		return err
	}

	detail, err := g.SetView("detail", 0, maxY-9, maxX-31, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	detail.Clear()
	fmt.Fprintln(detail, " ")
	fmt.Fprintf(detail, " Selected : %d\n", i+1)
	fmt.Fprintf(detail, " IP       : %s\n", f.conf.Servers[i].IpAddress)
	fmt.Fprintf(detail, " Username : %s\n", f.conf.Servers[i].Username)
	fmt.Fprintf(detail, " %s", time.Now().String())

	help, err := g.SetView("help", maxX-30, maxY-9, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	help.Clear()
	fmt.Fprintln(help, " ")
	fmt.Fprintln(help, "    ↑ ↓: Select")
	fmt.Fprintln(help, "     ↵ : Connect")
	fmt.Fprintln(help, "     ^C: Exit")
	fmt.Fprintln(help, " F1-F12: Direct selection")

	if _, err := g.SetCurrentView("main"); err != nil {
		return err
	}

	return nil
}

func (f *Frontend) initKeybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, f.selected); err != nil {
		return err
	}

	if err := g.SetKeybinding("main", gocui.KeyArrowUp, gocui.ModNone, f.scrollUp); err != nil {
		g.Update(f.layout)
		return err
	}

	if err := g.SetKeybinding("main", gocui.KeyArrowDown, gocui.ModNone, f.scrollDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF1, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(0)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(1)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (f *Frontend) selected(g *gocui.Gui, v *gocui.View) error {
	f.SetHasSelected(true)
	return gocui.ErrQuit
}

func (f *Frontend) scrollUp(g *gocui.Gui, v *gocui.View) error {
	i := f.Index()
	if i == 0 {
		f.SetIndex(len(f.conf.Servers) - 1)
	} else {
		f.SetIndex(i - 1)
	}

	return nil

}

func (f *Frontend) scrollDown(g *gocui.Gui, v *gocui.View) error {
	i := f.Index()
	if i >= len(f.conf.Servers)-1 {
		f.SetIndex(0)
	} else {
		f.SetIndex(i + 1)
	}

	return nil
}

func (f *Frontend) SetIndex(i int) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.index = i
}

func (f *Frontend) Index() int {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.index
}

func (f *Frontend) SetHasSelected(t bool) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.hasSelected = t
}

func (f *Frontend) HasSelected() bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.hasSelected
}
