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

func (Cursor *Cursor) getMaxLineLength(v *gocui.View) int {
	return len(v.BufferLines())
}

func (cursor *Cursor) MoveToFirst(g *gocui.Gui, v *gocui.View) error {
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	return nil
}
func (cursor *Cursor) lineBelow(g *gocui.Gui, v *gocui.View, d int) bool {
	yOffset, yCurrent, _ := cursor.FindPosition(g, v.Name())
	_, err := v.Line(yOffset + yCurrent + d)
	return err == nil && cursor.getMaxLineLength(v) >= yOffset+yCurrent+d && 0 <= yOffset+yCurrent+d
}

func (cursor *Cursor) Move(g *gocui.Gui, v *gocui.View, d int, callback func(int, int) error) bool {
	dir := 1
	if d < 0 {
		dir = -1
	}
	distance := int(math.Abs(float64(d)))
	for ; distance > 0; distance-- {
		if cursor.lineBelow(g, v, distance*dir) {
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
