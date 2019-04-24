package models

import (
	"time"
)

type trades_tb struct {
	t_id int
	b_id int
	bu_id int
	su_id int 
	trade_time time.Time
	final_time time.Time
}