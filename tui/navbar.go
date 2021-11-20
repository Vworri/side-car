package tui

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

type NavBar struct {
	container *tview.Flex
	File      *tview.DropDown
	Edit      *tview.DropDown
	pages     *tview.Table
}

func (app *App) NewNavBar() {
	app.navBar = new(NavBar)
	app.navBar.Edit = tview.NewDropDown()
	app.navBar.pages = tview.NewTable().SetSelectable(false, true)
	app.navBar.container = tview.NewFlex().SetDirection(1)
	app.registerPages()

	app.navBar.container.AddItem(app.navBar.Edit, 0, 1, false).
		AddItem(app.navBar.pages, 0, 1, false)
}

func (app *App) registerPages() {
	count := app.pages.GetPageCount()
	fmt.Println(count)
	for i := 0; i < count; i++ {
		p, _ := app.pages.GetFrontPage()
		fmt.Println(p)
		app.navBar.pages.SetCell(0, i, tview.NewTableCell(p).SetTextColor(header_color).
			SetAlign(tview.AlignCenter))
		app.pages.SendToBack(p)

	}
	os.Exit(0)
}
