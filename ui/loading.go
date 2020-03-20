package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type LoadingPanel struct {
	ViewName     string
	viewPosition ViewPosition
}

func NewLoadingPanel() (*LoadingPanel, error) {
	loadingPanel := LoadingPanel{
		ViewName: "loading",
		viewPosition: ViewPosition{
			x0: Position{0.1, 0},
			y0: Position{0.35, 0},
			x1: Position{0.9, 2},
			y1: Position{0.5, 2},
		},
	}
	return &loadingPanel, nil
}

func (loadingPanel *LoadingPanel) ShowLoading(g *gocui.Gui, callback func()) error {
	g.Update(func(g *gocui.Gui) error {
		loadingPanel.DrawView(g)
		g.Update(func(g *gocui.Gui) error {
			callback()
			g.DeleteView(loadingPanel.ViewName)
			return nil
		})
		return nil

	})
	return nil
}

func (loadingPanel *LoadingPanel) DrawView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	x0, y0, x1, y1 := loadingPanel.viewPosition.GetCoordinates(maxX, maxY)
	if v, err := g.SetView(loadingPanel.ViewName, x0, y0, x1, y1); err != nil {
		v.SelBgColor = gocui.ColorMagenta
		g.SetCurrentView(loadingPanel.ViewName)
		_, err := g.SetCurrentView(loadingPanel.ViewName)
		if err != nil {
			return err
		}
		fmt.Fprintf(v, "LOADING....")
	}
	return nil
}
