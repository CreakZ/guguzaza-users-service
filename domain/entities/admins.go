package entities

import "time"

type AdminCreate struct {
	Nickname, Password, InviteToken string
}

type AdminBase struct {
	Nickname, Password, Uuid string
}

type AdminPublic struct {
	ID                       int
	Nickname, Uuid, Position string
	JoinDate                 time.Time
}

type AdminsPaginated struct {
	Limit, TotalPages int
	Admins            []AdminPublic
}
