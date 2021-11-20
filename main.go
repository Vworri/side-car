// +build ignore

package main

import (
	"github.con/vworri/side-car/tui"
)

func main() {

	app := tui.CreateApplicaion()

	if err := app.EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
