package vpp

import "time"

type ServiceToken struct {
	Token   string
	ExpDate *time.Time
	OrgName string
}
