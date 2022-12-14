// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserCollection = "user_collection"

// UserCollection mapped from table <user_collection>
type UserCollection struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`                                // 主键
	UserID      int64     `gorm:"column:user_id;not null" json:"user_id"`                        // 用户 id
	CollectType int32     `gorm:"column:collect_type;not null" json:"collect_type"`              // 收藏类型 (1-喜欢 2-赞同 3-收藏 4-关注)
	ObjType     int32     `gorm:"column:obj_type;not null" json:"obj_type"`                      // 对象类型 喜欢-(1-回答 2-文章) 赞同-(1-回答 2-文章) 收藏-(0-默认对象 1-回答 2-文章) 关注-(1-用户 2-话题 3-专栏 4-问题 5-收藏夹)
	ObjID       int64     `gorm:"column:obj_id;not null" json:"obj_id"`                          // 对象 id
	CreateTime  time.Time `gorm:"autoCreateTime;column:create_time;not null" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"autoUpdateTime;column:update_time;not null" json:"update_time"` // 修改时间
}

// TableName UserCollection's table name
func (*UserCollection) TableName() string {
	return TableNameUserCollection
}
