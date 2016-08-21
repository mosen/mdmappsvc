package vpp

import "time"

type SToken struct {
	Token string
	ExpDate *time.Time
	OrgName string
}
