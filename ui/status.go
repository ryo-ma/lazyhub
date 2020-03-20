package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type StatusPanel struct {
	ViewName     string
	viewPosition ViewPosition
}

func NewStatusPanel() (*StatusPanel, error) {
	statusPanel := StatusPanel{
		ViewName: "status",
		viewPosition: ViewPosition{
			x0: Position{0.0, 0},
			y0: Position{0.89, 0},
			x1: Position{1.0, 2},
			y1: Position{1.0, 2},
		},
	}
	return &statusPanel, nil
}

func (statusPanel *StatusPanel) DrawView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	x0, y0, x1, y1 := statusPanel.viewPosition.GetCoordinates(maxX, maxY)
	if v, err := g.SetView(statusPanel.ViewName, x0, y0, x1, y1); err != nil {
		v.SelFgColor = gocui.ColorBlack
		v.Title = " Status"
	}
	return nil
}

func (statusPanel *StatusPanel) DrawText(g *gocui.Gui, message string) error {
	v, err := g.View(statusPanel.ViewName)
	if err != nil {
		return err
	}
	v.Clear()
	fmt.Fprintln(v, message)

	return nil
}
