package tui

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rivo/tview"
	"github.con/vworri/side-car/database"
)

type TaskView struct {
	Title        string
	header       []string
	taskData     *[]database.Task
	selectedTask int
	page         *tview.Flex
	view         *tview.Table
	actions      *tview.List
	fileOptions  *tview.DropDown
	editOpions   *tview.DropDown
	db           *database.Database
	form         *tview.Form
}

func (t *TaskView) InitTaskView() {
	t.Title = "Tasks"
	t.form = tview.NewForm()
	t.view = tview.NewTable().SetSelectable(true, false).SetSelectedFunc(
		func(row, column int) {
			t.selectedTask = row - 1
			t.form.Clear(true)
			t.TaskForm()

		})
	t.page = tview.NewFlex()
	t.taskData = t.db.GetTasks()
	t.selectedTask = len(*t.taskData) - 1
	t.header = t.db.TaskFields()
	t.TaskForm()
	t.CreateHeader()
	t.FillData()
	t.setEditOptions()
	t.page.AddItem(t.view, 0, 1, true).
		AddItem(t.form, 0, 1, true)
}

func (t *TaskView) CreateHeader() {
	for i, v := range t.header {
		t.view.SetCell(0, i, tview.NewTableCell(v).SetTextColor(header_color).
			SetAlign(tview.AlignCenter))
	}
}

func (t *TaskView) FillData() {
	for row, v := range *t.taskData {
		data := reflect.ValueOf(v)
		count := data.NumField()
		for col := 0; col < count; col++ {
			if strings.Contains(t.header[col], "due") {
				fmt.Printf("found you at %d\n", col)
				val, _ := v.Due.Value()
				t.view.SetCellSimple(row+1, col, fmt.Sprintf("%v", val))
				continue
			}
			t.view.SetCellSimple(row+1, col, fmt.Sprintf("%v", data.Field(col)))
		}

	}
}

func (t *TaskView) setEditOptions() {
	t.editOpions = tview.NewDropDown()
	t.editOpions.AddOption(
		"new", func() {
			t.db.AddNewTask()
			t.refreshTasks()
		})

	t.editOpions.AddOption(
		"delete", func() {
			t.deleteTask()
		})
	t.editOpions.AddOption(
		"edit", func() {
			t.editTask()
		})
}

func (t TaskView) refreshTasks() {
	t.taskData = t.db.GetTasks()
	t.selectedTask = len(*t.taskData) - 1
	t.view.Clear()
	t.CreateHeader()
	t.FillData()
	t.view.Select(t.selectedTask+1, 0)
	t.form.Clear(true)
	t.TaskForm()

}

func (t TaskView) editTask() {
}
func (t TaskView) deleteTask() {
}

func (t *TaskView) TaskForm() {
	task := (*t.taskData)[t.selectedTask]
	t.form.
		AddInputField("Name", task.Name, 20, nil, func(text string) {
			task.Name = text
		}).
		AddInputField("Description", task.Description, 1000, nil, func(text string) {
			task.Description = text
		}).
		AddCheckbox("complete", task.Complete, func(checked bool) {
			task.Complete = checked
		}).
		AddButton("Save", func() {
			t.db.Save(task)
			t.refreshTasks()
		}).
		AddButton("Delete", func() {
			t.db.DeleteTask(&task)
			t.refreshTasks()
		})
	t.form.SetBorder(true).SetTitle(fmt.Sprintf("Showing Task #%d", task.Id)).SetTitleAlign(tview.AlignLeft)
}

func (app *App) GetTasks() {
	app.taskPage = new(TaskView)
	app.taskPage.db = app.db
	app.taskPage.InitTaskView()
	app.navBar.Edit = app.taskPage.editOpions
	app.pages.AddPage(app.taskPage.Title,
		app.taskPage.page, false, true)
}
