package models

import "time"

type Detail_users_tb struct {
	Du_id int						// 表唯一id
	U_id int						// 用户id
	Zh	string						// 中文名
	Ch	string						// 英文名
	Birthday time.Time				// 生日
	Email	string					// 邮箱
	Age	int							// 年龄
	Sex	int							// 性别
	Identity	string				// 身份证
	Real_name	string				// 真实姓名
	Personal_portrait	string		// 个人头像
	Personalized_signature string	// 个性签名
}