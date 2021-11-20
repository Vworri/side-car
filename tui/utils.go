package tui

import (
	"github.com/rivo/tview"
	"github.con/vworri/side-car/database"
)

type SidecarPage struct {
	Title      string
	editOpions *tview.DropDown
	page       *tview.Flex
	db         *database.Database
}
