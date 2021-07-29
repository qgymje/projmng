package entity

import "time"

type TaskNotification struct {
	Id     int
	Task   *Task
	SendTo *User

	NotificationTemplate *NotificationTemplate
	TemplateValue        map[string]interface{}

	CreatedAt  time.Time
	NotifiedAt time.Time
}

type ITaskNotification interface {
	Create(tn *TaskNotification) error
}
