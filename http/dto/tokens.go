package dto

type InviteTokenCreate struct {
	PositionID int `json:"positionId"`
}

type InviteToken struct {
	Token string `json:"inviteToken"`
}
