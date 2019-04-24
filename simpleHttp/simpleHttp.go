package simpleHttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//用于保存用户账号信息（注意，结构体里面的属性必须首字母大写）
type User struct {
	Id       uint `json:"-"` //其中`json:"-"`代表忽视此属性
	Username string
	Password string
}

//用于保存回馈信息
type Status struct {
	State  bool
	Detail string
}

var userArr = make([]User, 0)  //用于存储用户的信息
var userId uint = 1            //用户ID
var status = Status{false, ""} //反馈信息

//查看用户名是否已经存在
func Existed(user User) bool {
	for _, value := range userArr {
		if value.Username == user.Username { //判断用户名是否存在
			return true
		}
	}
	return false
}

//验证用户输入的是否正确
func Verify(user User) bool {
	for _, value := range userArr {
		//判断用户名与密码是否都相同
		if value.Username == user.Username && value.Password == user.Password {
			return true
		}
	}
	return false
}

//注册
func Register(userInfo []byte) {
	var user User
	json.Unmarshal(userInfo, &user) //将json转换成结构体
	if Existed(user) {              //判断是否已经注册过
		status = Status{false, "用户名已存在"} //将状态回馈信息写入
		return                           //一旦失败，程序就没必要再执行了
	}
	user.Id = userId
	userId += 1
	userArr = append(userArr, user) //将这个用户的信息保存到切片中
	status = Status{true, "注册成功"}
}

//登录
func LoginIn(userInfo []byte) {
	var user User
	json.Unmarshal(userInfo, &user)
	if !Existed(user) { //首先判断用户输入的用户名是否存在
		status = Status{false, "用户名不存在"}
		return
	}
	if !Verify(user) { //判断用户名与密码是否一一对应
		status = Status{false, "用户名或密码错误"}
		return
	}
	status = Status{true, "登录成功"}
}

//将回馈信息转换成json，byte
func Feedbook(finfo Status) []byte {
	s, _ := json.Marshal(finfo)
	return s
}

//注册的具体交互函数
func register(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" { //判断是不是POST请求
		s, _ := ioutil.ReadAll(req.Body) //读取数据
		Register(s)                      //注册
		res.Write(Feedbook(status))      //注册的结果反馈给用户
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

//登录的具体交互函数
func login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s, _ := ioutil.ReadAll(req.Body)
		LoginIn(s) //登录
		res.Write(Feedbook(status))
	} else {
		res.Write([]byte("{\"false\":\"只支持POST方式\"}"))
	}
}

func RunServer() {
	http.HandleFunc("/login", login)       //登录
	http.HandleFunc("/register", register) //注册
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("监听失败")
	}
}
