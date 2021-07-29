package entity

import "time"

type TaskComment struct {
	Id        int
	Task      *Task
	Creator   *User
	Content   string
	CreatedAt time.Time
	DeletedAt time.Time
}
