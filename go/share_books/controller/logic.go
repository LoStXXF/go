package controller

import (
	"regexp"
	"share_books/models"
	"strconv"
	"time"
	_ "fmt"
)

// 验证手机号码格式是否正确
func verification_mobile_number (phone string) bool {
	reg, err := regexp.Compile("^1[0-9]{10}$")
	if err != nil {
		__log.Println(err)
	}
	return reg.MatchString(phone);
}

// 根据用户验证信息，第一个是要验证的字段名称，第二个是其对应的值，
// 验证用户唯一性，唯一就代表在数据库中没有重复的，那么返回true
func verification_user_unique(key string, value interface{}) bool {
	var user models.Users_tb
	var count int
	__mysql.Where(key + " = ?", value).First(&user).Count(&count)
	if count != 0 {
		return false
	}
	return true
}

// 验证用户名格式是否正确
func verification_username (username string) bool {
	// 获取用户规则信息
	reg, err := regexp.Compile("^[a-zA-Z0-9_]{" + __userminlen + "," + __usermaxlen + "}$")
	if err != nil {
		__log.Println(err)
	}
	return reg.MatchString(username)
}

// 验证用户输入的用户名与密码是否正确 目前只判断输入的手机号是否正确
// 1 代表成功
// 2 代表验证码错误
// 3 代表验证码已经过期
func verification_user_code (phone string, code int) int {
	var count int 
	var code_contain models.Verification_codes_tb
	__mysql.Select("code, stime").Where("code = ?", code).First(&code_contain).Count(&count)
	if count != 0 {
		now_time := time.Now().Unix()
		code_time_s := code_contain.Stime.Unix()
		code_time, _ := strconv.ParseInt(__codetime, 10, 64)
		if (now_time - code_time_s) < code_time {
			return 1
		}
		return 3
	} else {
		return 2
	}
}

// 专门处理验证码的
func handle_phone_verification_code (phone string, code int, ret *RetJSON, seccess func(ret *RetJSON)) {
	ret_code := verification_user_code(phone, code)
	if ret_code == 1 {
		seccess(ret)
	} else if ret_code == 3 {
		*ret = RetJSON{0, "你输入的验证码已经过期，请重新发送"}
	} else {
		*ret = RetJSON{0, "你输入的验证码不正确"}
		return 
	}
	__mysql.Delete(&models.Verification_codes_tb{Phone:phone})
}