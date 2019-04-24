package models

import (
	"time"
)

type Verification_codes_tb struct {
	V_id int
	Phone string
	Code int
	Stime time.Time
}