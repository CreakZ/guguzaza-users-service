package dto

type AdminCreate struct {
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	InviteToken string `json:"inviteToken"`
}

type AdminPublic struct {
	ID       int      `json:"id"`
	Nickname string   `json:"nickname"`
	Uuid     string   `json:"uuid"`
	Position string   `json:"position"`
	JoinTime JoinTime `json:"joinTime"`
}

type AdminsPaginated struct {
	Limit      int           `json:"limit"`
	TotalPages int           `json:"totalPages"`
	Admins     []AdminPublic `json:"admins"`
}
