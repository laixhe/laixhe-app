package models

import (
	"time"

	"gorm.io/gorm"
)

// MessageData 用户的聊天数据表
type MessageData struct {
	ID          int64          `gorm:"column:id;type:uint;not null;primary_key;autoIncrement"`
	MessageID   string         `gorm:"column:message_id;type:string;size:40;not null;unique;default:'';comment:消息id (以: m_ 开头)"`
	DialogID    string         `gorm:"column:dialog_id;type:string;size:40;not null;index;default:'';comment:会话id"`
	ChatType    int32          `gorm:"column:chat_type;type:uint;not null;default:0;comment:聊天类型:0=私聊 1=群聊 ..."`
	MessageType int32          `gorm:"column:message_type;type:uint;not null;default:0;comment:消息类型:0=文本1=图片2=文件3=语音4=视频 ..."`
	UserID      string         `gorm:"column:user_id;type:string;size:40;not null;default:'';comment:用户id"`
	ToID        string         `gorm:"column:to_id;type:string;size:40;not null;default:'';comment:接收者(用户ID/群ID)"`
	Pts         int64          `gorm:"column:pts;type:uint;not null;default:0;comment:会话的序列id(唯一值)"`
	TextContent string         `gorm:"column:text_content;type:string;size:4096;not null;default:'';comment:消息的内容"`
	MessageCopy []byte         `gorm:"column:message_copy;type:bytes;size:4096;not null;default:'';comment:(pb格式)消息其它内容"`
	LatestTime  int64          `gorm:"column:latest_time;type:uint;not null;default:0;comment:最近发生的时间(时间戳)"`
	IsRead      bool           `gorm:"column:is_read;type:bool;not null;default:0;comment:是否已读"`
	IsWithdraw  bool           `gorm:"column:is_withdraw;type:bool;not null;default:0;comment:是否撤回"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

// TableName -
func (u *MessageData) TableName() string {
	return "message_data"
}
