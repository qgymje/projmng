package entity

import "time"

type ContentItem struct {
	Id          int         `json:"id"`
	Key         string      `json:"key"`
	Val         interface{} `json:"val"`
	Description string      `json:"description"`
	Creator     *User       `json:"creator"`
	CreatedAt   time.Time   `json:"created_at"`
	DeletedAt   time.Time   `json:"deleted_at"`
}

type IContentItem interface {
	Create(item *ContentItem) error
	Creates(items []*ContentItem) error
}
