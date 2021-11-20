package tui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type WorkoutDashboard struct {
	RoutinePannel *widgets.TabPane
}

func (w *WorkoutDashboard) InitRoutinePannel() {
	w.RoutinePannel = widgets.NewTabPane("pierwszy", "drugi", "trzeci", "żółw", "four", "five")
	w.RoutinePannel.SetRect(0, 1, 50, 4)
	w.RoutinePannel.Border = true
	ui.Render(w.RoutinePannel)
}
