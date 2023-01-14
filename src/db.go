package main

import "time"

type SettingEntity struct {
	ID    int64
	Type  string
	Value string
	End   time.Time
}
