package main

import (
	"log"

	"encoding/base64"
	"github.com/jroimartin/gocui"
	"github.com/ryo-ma/lazyhub/lib"
	"github.com/ryo-ma/lazyhub/ui"
)

var client *lib.Client
var repositoryPanel *ui.RepositoryPanel
var textPanel *ui.TextPanel
var statusPanel *ui.StatusPanel
var searchPanel *ui.SearchPanel
var cursor *ui.Cursor

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	_ = err
	defer g.Close()
	g.SetManagerFunc(layout)

	client, _ = lib.NewClient()
	result, _ := client.GetTrendingRepository("", "")
	repositoryPanel, _ = ui.NewRepositoryPanel()
	repositoryPanel.Result = result
	textPanel, _ = ui.NewTextPanel()
	statusPanel, _ = ui.NewStatusPanel()
	searchPanel, _ = ui.NewSearchPanel()
	cursor = &ui.Cursor{}

	repositoryPanel.DrawView(g)
	textPanel.DrawView(g)
	statusPanel.DrawView(g)
	textPanel.DrawText(g, &repositoryPanel.Result.Items[0])
	g.SetCurrentView(repositoryPanel.ViewName)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(repositoryPanel.ViewName, 'r', gocui.ModNone, drawReadme); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'k', gocui.ModNone, cursorMovement(-1)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'j', gocui.ModNone, cursorMovement(1)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlU, gocui.ModNone, cursorMovement(-5)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlD, gocui.ModNone, cursorMovement(5)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlF, gocui.ModNone, drawSearchEditor); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding(searchPanel.ViewName, gocui.KeyEnter, gocui.ModNone, searchRepositoryByTopic); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, cursorMovement(-1)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, cursorMovement(1)); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}

func layout(g *gocui.Gui) error {
	repositoryPanel.DrawView(g)
	statusPanel.DrawView(g)
	textPanel.DrawView(g)
	return nil
}

func cursorMovement(d int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		cursor.Move(g, v, d, func(yOffset int, yCurrent int) error {
			if g.CurrentView().Name() == repositoryPanel.ViewName {
				if yOffset+yCurrent >= len(repositoryPanel.Result.Items) {
					return nil
				}
				textPanel.DrawText(g, &repositoryPanel.Result.Items[yOffset+yCurrent])
			}
			return nil
		})
		return nil
	}
}

func drawSearchEditor(g *gocui.Gui, _ *gocui.View) error {
	err := searchPanel.DrawView(g)
	if err != nil {
		return err
	}
	return nil
}

func drawReadme(g *gocui.Gui, _ *gocui.View) error {
	yOffset, yCurrent, _ := cursor.FindPosition(g, repositoryPanel.ViewName)
	currentItem := repositoryPanel.Result.Items[yCurrent+yOffset]
	readme, _ := client.GetReadme(currentItem)
	b, _ := base64.StdEncoding.DecodeString(readme.Content)
	textPanel.DrawReadme(g, &currentItem, string(b))
	return nil
}

func searchRepositoryByTopic(g *gocui.Gui, v *gocui.View) error {
	topic, err := v.Line(0)
	if topic == "" {
		g.DeleteView(searchPanel.ViewName)
		return nil
	}
	if err != nil {
		return err
	}
	repositoryPanel.Result, _ = client.SearchRepository(topic)
	g.DeleteView(searchPanel.ViewName)
	vr, err := g.View(repositoryPanel.ViewName)
	if err != nil {
		return err
	}
	vr.Clear()
	vr.Title = " Search [" + topic + "]"
	repositoryPanel.Result.Draw(vr)
	g.SetCurrentView(repositoryPanel.ViewName)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
