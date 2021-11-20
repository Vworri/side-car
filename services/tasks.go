package services

import "time"







var Status = [4]string{"backlog", "todo", "started", "done"}

type Task struct {
	name string
	description string
	status int
	due time.Time
}



func LoadTask(){

}

func CreateTask(name string, description string, status int, due time.Time){
	
}