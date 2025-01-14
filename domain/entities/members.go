package entities

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

type MemberCreate struct {
	UserBase
	Sex, About string
}

type MemberPublic struct {
	ID         int
	Nickname   string
	JoinDate   time.Time
	Sex, About string
}

type MemberUpdate struct {
	Sex, About *string
}
