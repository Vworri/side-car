package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.con/vworri/side-car/database"
)

const header_color = tcell.ColorYellow

type App struct {
	*tview.Application
	navBar   *NavBar
	grid     *tview.Grid
	pages    *tview.Pages
	taskPage *TaskView
	db       *database.Database
}

func CreateApplicaion() App {
	app := App{
		Application: tview.NewApplication(),
		db:          database.NewDatabase(),
		pages:       tview.NewPages(),
		grid:        tview.NewGrid(),
		navBar:      NewNavBar(),
	}
	app.GetTasks()
	app.setGrid()
	app.pages.SetChangedFunc(func() {
		app.Draw()
	})
	app.SetRoot(app.grid, true)
	return app

}

func (app *App) setGrid() {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	app.grid.
		SetRows(3, 0, 3).
		// SetColumns(30, 0, 30).
		SetBorders(true)
	_, first_page := app.pages.GetFrontPage()
	app.grid.AddItem(first_page, 1, 0, 1, 3, 0, 0, false).
		AddItem(app.navBar.Edit, 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

}
