package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/ryo-ma/lazyhub/lib"
)

type RepositoryPanel struct {
	ViewName     string
	viewPosition ViewPosition
	Result       *lib.Result
}

func NewRepositoryPanel() (*RepositoryPanel, error) {
	repositoryPanel := RepositoryPanel{
		ViewName: "repository",
		viewPosition: ViewPosition{
			x0: Position{0.0, 0},
			y0: Position{0.0, 0},
			x1: Position{0.3, 2},
			y1: Position{0.9, 2},
		},
	}
	return &repositoryPanel, nil
}

func (repositoryPanel *RepositoryPanel) DrawView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	x0, y0, x1, y1 := repositoryPanel.viewPosition.GetCoordinates(maxX, maxY)
	if v, err := g.SetView(repositoryPanel.ViewName, x0, y0, x1, y1); err != nil {
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.Highlight = true
		repositoryPanel.Result.Draw(v)
		v.Title = " Today Trending "
	}
	return nil
}
