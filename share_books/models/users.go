package models

import (
	"time"
)

type Users_tb struct {
	U_id	int						// 用户id
	Username string					// 用户名
	Mobile	string					// 手机号
	Registration_time	time.Time	// 注册时间
	Code int 		`gorm:"-"`
}

