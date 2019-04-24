package models

import (
	"time"
)

type books_tb struct {
	b_id	int						//书籍自增id
	book_name	int					//书名
	g_id	int						//书所属类型
	status	int						//状态（）
	author	string					//作者
	publisher	string				//出版商
	su_id	int						//书籍所有者
	au_id	int						//书籍所在这个人的手中
	abstract	string				//简介
	label	string					//标签
	printing_time time.Time			//出版时间（印刷时间）
}