package models

import "time"

type AdminRegister struct {
	UserBase
	InviteToken string
}

type Admin struct {
	ID int
	UserBase
}

type AdminPublic struct {
	ID                 int
	Nickname, Position string
	JoinDate           time.Time
}
