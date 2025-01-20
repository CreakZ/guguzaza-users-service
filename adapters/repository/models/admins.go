package models

import "time"

type AdminRegister struct {
	UserBase
	PositionID int
	JoinDate   time.Time
}

type Admin struct {
	ID int
	UserBase
}

type AdminPublic struct {
	ID                       int
	Nickname, Position, Uuid string
	JoinDate                 time.Time
}

type AdminsPaginated struct {
	Limit, TotalCount int
	Admins            []AdminPublic
}
