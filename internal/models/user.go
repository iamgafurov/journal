package models

import "time"

type User struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name"`
	Login       string    `json:"login"`
	Token       string    `json:"token"`
	Service     string    `json:"service"`
	UchprocId   int64     `json:"uchprocId"`
	UchprocCode int64     `json:"uchprocCode"`
	RegDate     time.Time `json:"regDate"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ExpireAt    time.Time `json:"expire_at"`
	Status      string    `json:"status"`
}
