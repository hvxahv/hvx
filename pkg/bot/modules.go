package bot

import "time"

type AccountNotice struct {
	Name string
	Data time.Time
}

type ServicesRunNotice struct {
	Name string
	Port string
	Data time.Time
}