package model

import "time"

type TokenBlacklist struct {
	ID         int
	Token      string
	ExpiredAt  time.Time
}