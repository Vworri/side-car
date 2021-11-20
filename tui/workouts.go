package tui

import "github.com/rivo/tview"

type WorkoutView struct {
	SidecarPage
}

func (w *WorkoutView) InitWorkoutView() {
	w.Title = "Workouts"
	w.page = tview.NewFlex().SetDirection(0)
	w.editOpions = tview.NewDropDown()

}

func (app *App) SetupWorkoutPage() {
	app.workoutPage = new(WorkoutView)
	app.workoutPage.db = app.db
	app.workoutPage.InitWorkoutView()
	app.pages.AddPage(app.workoutPage.Title,
		app.workoutPage.page, false, true)
}
