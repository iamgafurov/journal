package dto

import "time"

type TokenizePayload struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expireAt"`
}

type TokenHash struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
	Time  int64  `json:"time"`
}
