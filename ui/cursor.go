package ui

import (
	"github.com/jroimartin/gocui"
	"math"
)

type Cursor struct{}

func (cursor *Cursor) FindPosition(g *gocui.Gui, viewName string) (int, int, error) {
	v, err := g.View(viewName)
	if err != nil {
		return 0, 0, err
	}
	_, yOffset := v.Origin()
	_, yCurrent := v.Cursor()
	return yOffset, yCurrent, nil
}
func (cursor *Cursor) lineBelow(v *gocui.View, d int) bool {
	_, y := v.Cursor()
	_, err := v.Line(y + d)
	return err == nil
}
func (cursor *Cursor) MoveToFirst(g *gocui.Gui, v *gocui.View) error {
	yOffset, yCurrent, err := cursor.FindPosition(g, v.Name())
	if err != nil {
		panic(err)
	}
	cursor.Move(g, v, -(yOffset + yCurrent), nil)
	return nil
}

func (cursor *Cursor) Move(g *gocui.Gui, v *gocui.View, d int, callback func(int, int) error) bool {
	dir := 1
	if d < 0 {
		dir = -1
	}
	distance := int(math.Abs(float64(d)))
	for ; distance > 0; distance-- {
		if cursor.lineBelow(v, distance*dir) {
			v.MoveCursor(0, distance*dir, false)
			yOffset, yCurrent, _ := cursor.FindPosition(g, v.Name())
			if callback != nil {
				callback(yOffset, yCurrent)
			}

			return true
		}
	}
	return false
}
