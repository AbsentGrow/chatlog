package backup

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Contact represents a WeChat contact for backup
type Contact struct {
	UserName string `gorm:"primaryKey;column:user_name"`
	Alias    string `gorm:"column:alias"`
	Remark   string `gorm:"column:remark"`
	NickName string `gorm:"column:nick_name"`
	IsFriend int    `gorm:"column:is_friend"` // 0 or 1
}

// TableName overrides the table name
func (Contact) TableName() string {
	return "wechat_contact"
}

// ChatRoom represents a WeChat chat room for backup
type ChatRoom struct {
	Name     string         `gorm:"primaryKey;column:name"`
	Owner    string         `gorm:"column:owner"`
	Remark   string         `gorm:"column:remark"`
	NickName string         `gorm:"column:nick_name"`
	Users    JSONStringList `gorm:"column:users;type:text"` // Store as JSON string
}

// TableName overrides the table name
func (ChatRoom) TableName() string {
	return "wechat_chat_room"
}

// Message represents a WeChat message for backup
type Message struct {
	Seq        uint64      `gorm:"primaryKey;column:seq"` // Using Seq as primary key, though it might not be unique across sessions strictly speaking, but usually fine for backup per DB
	Time       time.Time   `gorm:"column:time"`
	Talker     string      `gorm:"column:talker;index"`
	TalkerName string      `gorm:"column:talker_name"`
	IsChatRoom int         `gorm:"column:is_chat_room"`
	Sender     string      `gorm:"column:sender"`
	IsSelf     int         `gorm:"column:is_self"`
	Type       int         `gorm:"column:type"`
	SubType    int         `gorm:"column:sub_type"`
	Content    string      `gorm:"column:content"`
	Contents   JSONContent `gorm:"column:contents;type:text"` // Store as JSON string
}

// TableName overrides the table name
func (Message) TableName() string {
	return "wechat_message"
}

// JSONStringList handles []string serialization to JSON
type JSONStringList []string

func (j *JSONStringList) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		str, ok := value.(string)
		if !ok {
			return errors.New(fmt.Sprint("Failed to unmarshal JSONStringList value:", value))
		}
		bytes = []byte(str)
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONStringList) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// JSONContent handles arbitrary JSON content serialization
type JSONContent map[string]interface{}

func (j *JSONContent) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		str, ok := value.(string)
		if !ok {
			return errors.New(fmt.Sprint("Failed to unmarshal JSONContent value:", value))
		}
		bytes = []byte(str)
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONContent) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
