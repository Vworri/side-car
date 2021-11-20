package tui

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.con/vworri/side-car/database"
)

var taskCategories = []string{"fun", "health", "adulting"}

type TaskView struct {
	SidecarPage

	header       []string
	errorPannel  *tview.TextView
	taskData     *[]database.Task
	selectedTask int
	view         *tview.Table
	form         *tview.Form
}

func (t *TaskView) InitTaskView() {
	//set title
	t.Title = "Tasks"
	// set file and edit options for this view
	t.setEditOptions()
	// get data
	t.taskData = t.db.GetTasks()
	t.header = t.db.TaskFields()
	// create widgets
	t.form = tview.NewForm()
	t.createTaskTable()
	t.createErrorPanel()
	// set grid for the widgets
	t.page = tview.NewFlex().SetDirection(0)
	// initialize selected tasks
	t.selectedTask = 0
	// fill form
	t.taskForm()

	t.page.AddItem(t.view, 0, 1, true).
		AddItem(t.form, 0, 1, true).
		AddItem(t.errorPannel, 0, 1, false)
}

func (t *TaskView) createTaskTable() {
	t.view = tview.NewTable().SetSelectable(true, false).SetSelectedFunc(
		func(row, column int) {
			t.selectedTask = row - 1
			t.form.Clear(true)
			t.taskForm()

		}).
		SetWrapSelection(true, true)

	t.CreateHeader()
	t.FillData()

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
				t.view.SetCellSimple(row+1, col, fmt.Sprintf("%v", v.Due.Format("2006-01-02")))
				continue
			}
			t.view.SetCellSimple(row+1, col, fmt.Sprintf("%v", data.Field(col)))
		}

	}
}

func (t *TaskView) setEditOptions() {
	t.editOpions = tview.NewDropDown().SetLabel("Edit:")
	t.editOpions.AddOption(
		"new", func() {
			t.db.AddNewTask()
			t.refreshTasks()
		})

	t.editOpions.AddOption(
		"delete", func() {
			t.deleteTask()
		})
}

func (t TaskView) refreshTasks() {
	t.db.SyncTasks(t.taskData)
	t.selectedTask = len(*t.taskData) - 1
	t.view.Clear()
	t.CreateHeader()
	t.FillData()
	t.view.Select(t.selectedTask+1, 0)
	t.form.Clear(true)
	t.taskForm()
	t.errorPannel.Clear()

}

func (t TaskView) deleteTask() {
	task := (*t.taskData)[t.selectedTask]
	t.db.DeleteTask(&task)
	t.refreshTasks()
}

func (t *TaskView) taskForm() {

	task := (*t.taskData)[t.selectedTask]
	var duedate = task.Due.Format("2006-01-02")
	var cat_indx = 0
	for i, val := range taskCategories {
		if val == task.Category {
			cat_indx = i
			break
		}
	}
	t.form.
		AddInputField("Name", task.Name, 20, nil, func(text string) {
			task.Name = text
		}).
		AddInputField("Description", task.Description, 1000, nil, func(text string) {
			task.Description = text
		}).
		AddInputField("Due (2006-01-02)", duedate, 1000, nil, func(text string) {
			duedate = text

		}).
		AddCheckbox("complete", task.Complete, func(checked bool) {
			task.Complete = checked
		}).
		AddDropDown("Category", taskCategories, cat_indx, func(option string, optionIndex int) {
			task.Category = option
		}).
		AddButton("Save", func() {
			due, err := time.Parse("2006-01-02", duedate)
			if err != nil {
				t.showError("Cannot Parse Date")
				return
			} else {
				task.Due = due
				t.db.Save(task)
				t.refreshTasks()
			}

		}).
		AddButton("Delete", func() {
			t.deleteTask()
		})
	t.form.SetBorder(true).SetTitle(fmt.Sprintf("Showing Task #%d", task.Id)).SetTitleAlign(tview.AlignLeft)
}

func (app *App) SetupTaskPage() {
	app.taskPage = new(TaskView)
	app.taskPage.db = app.db
	app.taskPage.InitTaskView()
	app.navBar.Edit = app.taskPage.editOpions
	app.pages.AddPage(app.taskPage.Title,
		app.taskPage.page, false, true)
}
func (t *TaskView) showError(message string) {
	t.errorPannel.SetText(message)
}
func (t *TaskView) createErrorPanel() {
	t.errorPannel = tview.NewTextView().SetTextColor(
		tcell.ColorIndianRed).SetWrap(true)

}
