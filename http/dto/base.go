package dto

type UserBase struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Uuid     string `json:"uuid"`
}
