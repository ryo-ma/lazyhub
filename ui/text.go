package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/ryo-ma/lazyhub/lib"
)

type TextPanel struct {
	ViewName     string
	viewPosition ViewPosition
}

func NewTextPanel() (*TextPanel, error) {
	textPanel := TextPanel{
		ViewName: "text",
		viewPosition: ViewPosition{
			x0: Position{0.3, 0},
			y0: Position{0.0, 0},
			x1: Position{1.0, 2},
			y1: Position{0.9, 2},
		},
	}
	return &textPanel, nil
}

func (textPanel *TextPanel) DrawView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	x0, y0, x1, y1 := textPanel.viewPosition.GetCoordinates(maxX, maxY)
	if v, err := g.SetView(textPanel.ViewName, x0, y0, x1, y1); err != nil {
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.Title = " text "
	}
	return nil
}

func (textPanel *TextPanel) DrawText(g *gocui.Gui, item *lib.Item) error {
	v, err := g.View(textPanel.ViewName)
	if err != nil {
		return err
	}
	v.Clear()
	v.Title = " " + item.GetRepositoryName() + " "
	fmt.Fprintln(v, item.String())

	return nil
}

func (textPanel *TextPanel) DrawReadme(g *gocui.Gui, item *lib.Item, readme string) error {
	v, err := g.View(textPanel.ViewName)
	if err != nil {
		return err
	}
	v.Clear()
	v.Title = " " + item.GetRepositoryName() + " "
	fmt.Fprintln(v, readme)
	g.SetCurrentView(textPanel.ViewName)
	v.Highlight = true
	return nil
}
