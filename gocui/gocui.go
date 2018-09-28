package gocui

import (
	"fmt"
	"sync"

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

func (f *Frontend) Draw(selected int) (int, bool, error) {
	defer f.g.Close()
	f.SetIndex(selected)

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

	_, sY := main.Size()

	if err = main.SetOrigin(0, 0); err != nil {
		return err
	}

	if i >= sY {
		if err = main.SetOrigin(0, i-sY+1); err != nil {
			return err
		}
		if err := main.SetCursor(0, sY-1); err != nil {
			return err
		}
	} else if err := main.SetCursor(0, i); err != nil {
		return err
	}

	main.Clear()
	main.Title = " Servers "
	main.Highlight = true
	main.SelBgColor = gocui.ColorRed
	main.SelFgColor = gocui.ColorWhite

	for _, s := range f.conf.Servers {

		if s.Index < 10 {
			fmt.Fprintf(main, "  %d: %s \n", s.Index, s.Name)
		} else {
			fmt.Fprintf(main, " %d: %s \n", s.Index, s.Name)
		}

	}

	detail, err := g.SetView("detail", 0, maxY-9, maxX-31, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	detail.Clear()
	detail.Title = " Connection Details "
	fmt.Fprintln(detail, " ")
	fmt.Fprintf(detail, " Host     : %s\n", f.conf.Servers[i].IPAddress)
	fmt.Fprintf(detail, " Username : %s\n\n", f.conf.Servers[i].Username)
	fmt.Fprintf(detail, " Profile  : %s\n", f.conf.Servers[i].Profile)

	help, err := g.SetView("help", maxX-30, maxY-9, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	help.Clear()
	help.Title = " Keybindings "
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

	if err := g.SetKeybinding("", gocui.KeyF3, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(2)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF4, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(3)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF5, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(4)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF6, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(5)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF7, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(6)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF8, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(7)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF9, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(8)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(9)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF11, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(10)
		f.SetHasSelected(true)
		return gocui.ErrQuit
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyF12, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		f.SetIndex(11)
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

func (f *Frontend) MaxIndex() int {
	f.mux.Lock()
	defer f.mux.Unlock()
	return len(f.conf.Servers)
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
