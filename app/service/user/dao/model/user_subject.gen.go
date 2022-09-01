// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserSubject = "user_subject"

// UserSubject mapped from table <user_subject>
type UserSubject struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`             // 主键
	Username   string    `gorm:"column:username;not null" json:"username"`                      // 用户名 (登陆用)
	Password   string    `gorm:"column:password;not null" json:"password"`                      // 密码
	Nickname   string    `gorm:"column:nickname;not null" json:"nickname"`                      // 昵称
	Email      string    `gorm:"column:email" json:"email"`                                     // 邮箱
	Phone      string    `gorm:"column:phone" json:"phone"`                                     // 手机号
	LastIP     string    `gorm:"column:last_ip;not null" json:"last_ip"`                        // 最近登录 ip 地址
	Vip        int32     `gorm:"column:vip;not null" json:"vip"`                                // vip 等级
	Follower   int32     `gorm:"column:follower;not null" json:"follower"`                      // 被关注数
	State      int32     `gorm:"column:state;not null" json:"state"`                            // 状态 (0-正常 1-冻结 2-封禁)
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time;not null" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time;not null" json:"update_time"` // 修改时间
}

// TableName UserSubject's table name
func (*UserSubject) TableName() string {
	return TableNameUserSubject
}
