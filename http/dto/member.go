package dto

import "time"

type MemberCreate struct {
	UserBase
	Sex   string `json:"sex"`
	About string `json:"about"`
}

type MemberBase struct {
	UserBase
	JoinTime time.Time `json:"joinTime"`
	Sex      string    `json:"sex"`
	About    string    `json:"about"`
}

type Member struct {
	ID int `json:"id"`
	MemberBase
}
