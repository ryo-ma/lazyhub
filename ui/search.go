package ui

import (
	"github.com/jroimartin/gocui"
)

type SearchPanel struct {
	ViewName     string
	viewPosition ViewPosition
}

func NewSearchPanel() (*SearchPanel, error) {
	searchPanel := SearchPanel{
		ViewName: "search",
		viewPosition: ViewPosition{
			x0: Position{0.1, 0},
			y0: Position{0.35, 0},
			x1: Position{0.9, 2},
			y1: Position{0.5, 2},
		},
	}
	return &searchPanel, nil
}

func (searchPanel *SearchPanel) DrawView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	x0, y0, x1, y1 := searchPanel.viewPosition.GetCoordinates(maxX, maxY)
	if v, err := g.SetView(searchPanel.ViewName, x0, y0, x1, y1); err != nil {
		v.SelFgColor = gocui.ColorBlack
		v.Editable = true
		v.Title = " Search Repository "
		_, err := g.SetCurrentView(searchPanel.ViewName)
		if err != nil {
			return err
		}
	}
	return nil
}
