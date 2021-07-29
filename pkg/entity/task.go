package entity

import "time"

type Task struct {
	Id       int
	Project  *Project
	Creator  *User
	Executor *User

	WillNofity      bool
	ContentTemplate *ContentTemplate
	TemplateValue   map[string]interface{}

	StartedAt  time.Time
	EndedAt    time.Time
	FinishedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type ITask interface {
	Create(task *Task) error
	Update(task *Task) error
	Delete(task *Task) error
	ListByUser(uid int) ([]*Task, error)
}
