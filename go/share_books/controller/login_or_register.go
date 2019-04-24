package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"share_books/models"
	"io/ioutil"
	"time"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"fmt"
	"math/rand"
)



// user用户登录
func api_login(c *gin.Context) {
	var user_login models.Verification_codes_tb
	user_info, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(user_info, &user_login)
	var ret RetJSON
	// 处理验证码
	handle_phone_verification_code(user_login.Phone, user_login.Code, &ret, func(ret *RetJSON) {
		var user models.Users_tb
		__mysql.Where("mobile = ?", user_login.Phone).First(&user)
		json_user, _ := json.Marshal(user)
		*ret = RetJSON{1, string(json_user)}
	})
	c.JSON(http.StatusOK, ret)
}

// user用户注册
func api_register(c *gin.Context) {
	var user models.Users_tb
	user_info, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(user_info, &user)
	
	// 把当前的时间保存起来
	user.Registration_time = time.Now()
	
	var ret RetJSON
	handle_phone_verification_code(user.Mobile, user.Code, &ret, func(ret *RetJSON) {})
	if !verification_mobile_number(user.Mobile) {							// 验证手机号码
		ret = RetJSON{2, "你输入的手机格式有误，请重新输入"}
	} else if !verification_user_unique("mobile", user.Mobile) {			// 验证手机号是否重复
		ret = RetJSON{3, "你输入的手机号已经被注册过了"}
	} else if !verification_username(user.Username) {						// 验证用户名格式是否正确
		ret = RetJSON{4, "你输入的用户名格式不正确"}
	} else if !verification_user_unique("username", user.Username) {		// 验证用户名是否已经注册过
		ret = RetJSON{5, "你输入的用户名已经注册过，请重新填写"}
	} else {
		// 处理验证码
		handle_phone_verification_code(user.Mobile, user.Code, &ret, func(ret *RetJSON) {
			if err := __mysql.Create(user).Error; err != nil {				// 将这条数据插入到数据库中
				__log.Println(err)
				*ret = RetJSON{0, "注册失败"}
			} else {
				json_user, _ := json.Marshal(user)
				*ret = RetJSON{1, string(json_user)}
			}
		})
	}
	c.JSON(http.StatusOK, ret)
}

// user用户获取个人信息
func api_update_personal_info(c *gin.Context) {
	var user models.Users_tb
	id := c.Param("id");
	__mysql.Where("u_id = ?", id).First(&user)
	data, _ := json.Marshal(user)
	var ret RetJSON
	ret = RetJSON{1, string(data)}
	c.JSON(http.StatusOK, ret)
}

// user用户个人信息认证
func api_real_name_authentication(c *gin.Context) {
	
}


// user用户获取手机验证码
func api_get_verification_code (c *gin.Context) {
	phone := c.Param("phone")
	clnt := ypclnt.New("9128797cb55f4865e0a614df63495df6")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phone
	// 设置随机种子
	rand.Seed(int64(time.Now().UnixNano()))
	// 获取随机数并保存起来，可以发送给注册的人或者登陆的人
	vc := rand.Int()%100000
	param[ypclnt.TEXT] = fmt.Sprintf("【吴敬测试】您的验证码是%d", vc)
	code := models.Verification_codes_tb{
		Phone: phone,
		Code: vc,
		Stime: time.Now(),
	}
	var ret RetJSON
	if err := __mysql.Create(&code).Error; err != nil {
		__log.Println(err)
		ret = RetJSON{0, "发送验证码失败"}
	} else {
		r := clnt.Sms().SingleSend(param)
		if r.Code == 0 {
			ret = RetJSON{1, "验证码发送成功，请注意查收"}
		} else {
			ret = RetJSON{0, r.Msg}
			__log.Println(r)
		}
	}
	c.JSON(http.StatusOK,ret)
}