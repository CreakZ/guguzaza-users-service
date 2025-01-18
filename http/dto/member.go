package dto

type MemberCreate struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	About    string `json:"about"`
}

type MemberBase struct {
	UserBase
	JoinTime JoinTime `json:"joinTime"`
	Sex      string   `json:"sex"`
	About    string   `json:"about"`
}

type MemberPublic struct {
	ID       int      `json:"id"`
	Nickname string   `json:"nickname"`
	Uuid     string   `json:"uuid"`
	JoinTime JoinTime `json:"joinTime"`
	Sex      string   `json:"sex"`
	About    string   `json:"about"`
}

type Member struct {
	ID int `json:"id"`
	MemberBase
}
