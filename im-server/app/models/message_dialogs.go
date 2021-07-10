package models

import (
	"time"

	"gorm.io/gorm"
)

// MessageDialogs 用户的会话表
type MessageDialogs struct {
	ID          int64          `gorm:"column:id;type:uint;not null;primary_key;autoIncrement"`
	DialogID    string         `gorm:"column:dialog_id;type:string;size:40;not null;unique;default:'';comment:会话id (id规则: d_{user_id}_{char_type}_{to_id} )"`
	UserID      string         `gorm:"column:user_id;type:string;size:40;not null;index:user_to_chat;default:'';comment:用户id"`
	ToID        string         `gorm:"column:to_id;type:string;size:40;not null;index:user_to_chat;default:'';comment:接收者(用户ID/群ID)"`
	ChatType    int32          `gorm:"column:chat_type;type:uint;not null;default:0;comment:聊天类型:0=私聊 1=群聊 ..."`
	MessageID   string         `gorm:"column:message_id;type:string;size:40;not null;index:user_to_chat;default:'';comment:消息ID(唯一值)"`
	Pts         int64          `gorm:"column:pts;type:uint;not null;default:0;comment:会话的序列id(唯一值)"`
	UnreadCount int32          `gorm:"column:unread_count;type:uint;not null;default:0;comment:未读数"`
	LatestTime  int64          `gorm:"column:latest_time;type:uint;not null;default:0;comment:最近发生的时间(时间戳)"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName -
func (u *MessageDialogs) TableName() string {
	return "message_dialogs"
}
