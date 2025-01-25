package dto

type Credentials struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type UserBase struct {
	Credentials
	Uuid string `json:"uuid"`
}
