package entity

import (
	"time"
)

type ContentTemplate struct {
	Id        int            `json:"id"`
	Task      *Task          `json:"task"`
	Content   []*ContentItem `json:"content"` // 保证了kv的插入顺序,Val可以支持嵌套ContentItem
	Creator   *User          `json:"creator"`
	CreatedAt time.Time      `json:"created_at"`
}

type IContentTemplate interface {
	Create(ct *ContentTemplate) error
}
