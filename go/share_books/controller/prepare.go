package controller

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
	"os"
	"github.com/larspensjo/config"
	"runtime"
	"flag"
	"errors"
)

var __mysql *gorm.DB
var __log *log.Logger
var __userminlen string
var __usermaxlen string
var __codetime string
var config_file = flag.String("configfile", "config/share_books.ini", "General configuration file")
func init() {
	var err error
	// 日志文件，方便写入日志
	handle_log("share_books.log")

	// 连接数据库，先读取配置文件，将数据库连接参数拿到手
	con_info, _ := read_config_item("mysql")
	con_str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
				con_info["USERNAME"], 
				con_info["PASSWORD"],
				con_info["HOSTNAME"], 
				con_info["POST"], 
				con_info["DATABASE"])
	__mysql, err = gorm.Open("mysql", con_str)
	if err != nil {
		__log.Println(err)	
	}
	__mysql.SingularTable(true)
	user_conf_info, _ := read_config_item("user")
	__userminlen = user_conf_info["MINLEN"]
	__usermaxlen = user_conf_info["MAXLEN"]
	__codetime = user_conf_info["CODETIME"]
}

func handle_log(log_file_name string) {
	log_file, err := os.OpenFile("logs/" + log_file_name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0)
    if err != nil{
        log.Fatalln("读取日志文件失败",err)
    }
    __log=log.New(log_file,"",log.Ldate|log.Ltime)
}


// 读取配置文件中的指定项的内容
func read_config_item (load_item string) (map[string]string, error) {
	var load_ret = make(map[string]string)
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	cfg, err := config.ReadDefault(*config_file)
	if err != nil {
		return nil, err
	}
	if cfg.HasSection(load_item) {
		section, err := cfg.SectionOptions(load_item)
		if err == nil {
			for _, value := range section {
				options, err := cfg.String(load_item, value)
				if err == nil {
					load_ret[value] = options
				} else {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New(load_item + "项不存在，请重新传参")
	}
	return load_ret, nil
}