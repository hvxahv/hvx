package models

import "github.com/jinzhu/gorm"

// Accounts 用户账户的 struct
type Accounts struct {
	gorm.Model
	Username	string`gorm:"username"`
	Password	string`gorm:"password"`
	Avatar		string`gorm:"avatar"`
	Bio			string`gorm:"bio"`
	Name 		string`gorm:"name"`
	EMail		string`gorm:"email"`
	Phone		int`gorm:"phone"`
	Telegram	string`gorm:"telegram"`
	Social      string`gorm:"social"`
	Private     int`gorm:"private"`
	PrivateKey  string`gorm:"private_key"`
	PublicKey	string`gorm:"public_key"`
}

// WebFinger 和 WebFingerLinks 组成标准 Activitypub 的 JSON-LD 协议
type WebFinger struct {
	Subject string `json:"subject"`
	Links   []WebFingerLinks `json:"links"`
}

// WebFingerLinks 供 WebFinger 使用
type WebFingerLinks struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}
