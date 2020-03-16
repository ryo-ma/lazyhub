package main

import (
	"fmt"
	"log"
	"math"

	"./lib"
	"github.com/jroimartin/gocui"
)

const (
	repositoryView = "repository"
	textView       = "text"
	pathView       = "path"
	helpView       = "help"
	searchView     = "search"
)

type position struct {
	prc    float32
	margin int
}

func (p position) getCoordinate(max int) int {
	return int(p.prc*float32(max)) - p.margin
}

type viewPosition struct {
	x0, y0, x1, y1 position
}

func (vp viewPosition) getCoordinates(maxX int, maxY int) (int, int, int, int) {
	x0 := vp.x0.getCoordinate(maxX)
	y0 := vp.y0.getCoordinate(maxY)
	x1 := vp.x1.getCoordinate(maxX)
	y1 := vp.y1.getCoordinate(maxY)
	return x0, y0, x1, y1
}

var viewPositions = map[string]viewPosition{
	repositoryView: {
		position{0.0, 0},
		position{0.0, 0},
		position{0.3, 2},
		position{0.9, 2},
	},
	textView: {
		position{0.3, 0},
		position{0.0, 0},
		position{1.0, 2},
		position{0.9, 2},
	},
	pathView: {
		position{0.0, 0},
		position{0.89, 0},
		position{1.0, 2},
		position{1.0, 2},
	},
	searchView: {
		position{0.1, 0},
		position{0.35, 0},
		position{0.9, 2},
		position{0.5, 2},
	},
}

var client *lib.Client
var result *lib.Result
var isSearch bool

func main() {
	client, _ = lib.NewClient()
	result, _ = client.GetTrendingRepository("", "")
	g, err := gocui.NewGui(gocui.OutputNormal)
	_ = err
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(repositoryView, 'k', gocui.ModNone, cursorMovement(-1)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(repositoryView, 'j', gocui.ModNone, cursorMovement(1)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(repositoryView, gocui.KeyCtrlU, gocui.ModNone, cursorMovement(-5)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(repositoryView, gocui.KeyCtrlD, gocui.ModNone, cursorMovement(5)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlF, gocui.ModNone, drawSearchEditor); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(searchView, gocui.KeyEnter, gocui.ModNone, searchRepositoryByTopic); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(repositoryView, gocui.KeyArrowUp, gocui.ModNone, cursorMovement(-1)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(repositoryView, gocui.KeyArrowDown, gocui.ModNone, cursorMovement(1)); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	views := []string{repositoryView, textView, pathView}
	for _, view := range views {
		x0, y0, x1, y1 := viewPositions[view].getCoordinates(maxX, maxY)
		if v, err := g.SetView(view, x0, y0, x1, y1); err != nil {
			v.SelFgColor = gocui.ColorBlack
			v.SelBgColor = gocui.ColorGreen
			v.Title = " " + view + " "
			if view == repositoryView {
				v.Highlight = true
				result.Draw(v)
				v.Title = " Today Trending "
			}
			if view == textView {
				drawText(g)
			}
		}
	}
	if !isSearch {
		_, err := g.SetCurrentView(repositoryView)
		if err != nil {
			log.Fatal("failed to set current view: ", err)
		}
	}
	return nil
}

func lineBelow(v *gocui.View, d int) bool {
	_, y := v.Cursor()
	line, err := v.Line(y + d)
	return err == nil && line != ""
}

func cursorMovement(d int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		dir := 1
		if d < 0 {
			dir = -1
		}
		distance := int(math.Abs(float64(d)))
		for ; distance > 0; distance-- {
			if lineBelow(v, distance*dir) {
				v.MoveCursor(0, distance*dir, false)
				drawPosition(g)
				drawText(g)

				return nil
			}
		}
		return nil
	}
}

func drawSearchEditor(g *gocui.Gui, _ *gocui.View) error {
	maxX, maxY := g.Size()
	x0, y0, x1, y1 := viewPositions[searchView].getCoordinates(maxX, maxY)
	if v, err := g.SetView(searchView, x0, y0, x1, y1); err != nil {
		v.SelFgColor = gocui.ColorBlack
		v.Title = " " + searchView + " "
		v.Editable = true
	}
	_, err := g.SetCurrentView(searchView)
	if err != nil {
		return err
	}
	isSearch = true
	return nil
}

func drawPosition(g *gocui.Gui) error {
	v, err := g.View(pathView)
	if err != nil {
		return err
	}
	v.Clear()
	yOffset, yCurrent, _ := findCursorPosition(g)
	fmt.Fprintf(v, "yOffset: %d ", yOffset)
	fmt.Fprintf(v, "yCurrent: %d ", yCurrent)
	return nil
}

func findCursorPosition(g *gocui.Gui) (int, int, error) {
	v, err := g.View(repositoryView)
	if err != nil {
		return 0, 0, err
	}
	_, yOffset := v.Origin()
	_, yCurrent := v.Cursor()
	return yOffset, yCurrent, nil
}

func drawText(g *gocui.Gui) error {
	v, err := g.View(textView)
	if err != nil {
		return err
	}
	v.Clear()
	yOffset, yCurrent, _ := findCursorPosition(g)
	v.Title = " " + result.Items[yCurrent+yOffset].GetRepositoryName() + " "
	fmt.Fprintln(v, result.Items[yCurrent+yOffset].String())

	return nil
}

func searchRepositoryByTopic(g *gocui.Gui, v *gocui.View) error {
	topic, err := v.Line(0)
	if topic == "" {
		isSearch = false
		g.DeleteView(searchView)
		return nil
	}
	if err != nil {
		return err
	}
	result, _ = client.SearchRepository(topic)
	g.DeleteView(searchView)
	vr, err := g.View(repositoryView)
	if err != nil {
		return err
	}
	vr.Clear()
	vr.Title = " Search [" + topic + "]"
	result.Draw(vr)
	isSearch = false
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
