package entity

import "time"

type Task struct {
	Id       int      `json:"id"`
	Project  *Project `json:"project"`
	Creator  *User    `json:"creator"`
	Executor *User    `json:"executor"`

	WillNofity      bool                   `json:"will_nofity"`
	ContentTemplate *ContentTemplate       `json:"content_template"` // Task -> ContentTemplate -> TemplateItem
	TemplateValue   map[string]interface{} `json:"template_value"`

	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	FinishedAt time.Time `json:"finished_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type ITask interface {
	Create(task *Task) error
	Update(task *Task) error
	Delete(task *Task) error
	ListByUser(uid int) ([]*Task, error)
}
