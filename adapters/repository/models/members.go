package models

import "time"

type MemberBase struct {
	UserBase
	JoinDate   time.Time
	Sex, About string
}

type Member struct {
	ID int
	MemberBase
}

type MemberPublic struct {
	ID             int
	Nickname, Uuid string
	JoinDate       time.Time
	Sex, About     string
}
