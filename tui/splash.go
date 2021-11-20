package tui

import (
	"github.com/rivo/tview"
)

type grid struct {
	*tview.Flex
}

func (app *App) HomePage() {

	splashGrid := createFlex()
	if err := app.SetRoot(splashGrid, true).Run(); err != nil {
		panic(err)
	}

}

func createFlex() grid {
	return grid{Flex: tview.NewFlex()}
}
