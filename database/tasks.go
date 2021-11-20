package database

import (
	"time"

	"gorm.io/datatypes"
)

type Task struct {
	Id          int            `gorm:"primary_key"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	Complete    bool           `gorm:"column:complete"`
	Due         datatypes.Date `gorm:"column:due"`
	Category    string         `gorm:"column:category"`
}

func (db *Database) CreateTask(task *Task) {
	results := db.Create(task)
	if results.Error != nil {
		panic(results)
	}
}

func (db *Database) DeleteTask(task *Task) {
	db.Delete(task)
}

func (db *Database) UpdateTask(task *Task) {
	db.Save(task)
}

func (db *Database) GetTasks() *[]Task {
	tasks := new([]Task)
	db.Find(&tasks)
	if len(*tasks) < 1 {
		db.AddNewTask()
		db.Find(&tasks)
	}

	return tasks
}

func (db *Database) AddNewTask() {
	t := Task{
		Name:        "New Task",
		Description: "Everyone has something to do",
		Complete:    false,
		Due:         datatypes.Date(time.Now()),
		Category:    "fun",
	}
	db.CreateTask(&t)
}

func (db *Database) ToggleComplete(task *Task) {
	task.Complete = !task.Complete
	db.UpdateTask(task)
}

func (db *Database) TaskFields() []string {
	var task_feilds []string

	result, _ := db.Debug().Migrator().ColumnTypes(&Task{})
	for _, v := range result {
		task_feilds = append(task_feilds, v.Name())
	}
	return task_feilds
}
