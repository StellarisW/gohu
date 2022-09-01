// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameNotificationSubject = "notification_subject"

// NotificationSubject mapped from table <notification_subject>
type NotificationSubject struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`             // 主键
	UserID      int64     `gorm:"column:user_id;not null" json:"user_id"`                        // 用户 ID
	MessageType int32     `gorm:"column:message_type;not null" json:"message_type"`              // 对象类型 (0-全部 1-关注我的 2-赞同与喜欢 3-评论与回复 4-邀请 5-提到我的)
	CreateTime  time.Time `gorm:"autoCreateTime;column:create_time;not null" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"autoUpdateTime;column:update_time;not null" json:"update_time"` // 修改时间
}

// TableName NotificationSubject's table name
func (*NotificationSubject) TableName() string {
	return TableNameNotificationSubject
}
