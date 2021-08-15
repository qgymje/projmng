package entity

import "time"

// 通知模板：
// 项目创建
// 任务指派
// 任务回复
// 任务完成
// 项目完成

type NotificationTemplate struct {
	Id        int
	Name      string
	Creator   *User
	Template  []*NotificationItem
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type NotificationItem struct {
	Key string
	Val string
}

type INotificationTemplate interface {
}
