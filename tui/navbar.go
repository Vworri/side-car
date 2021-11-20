package tui

import "github.com/rivo/tview"

type NavBar struct {
	File *tview.DropDown
	Edit *tview.DropDown

}


func NewNavBar()(*NavBar){
	nav := new(NavBar)
	nav.File = tview.NewDropDown()
	nav.Edit = tview.NewDropDown()
	return nav
}
