package ports

import "context"

type InviteTokensUtilPort interface {
	CreateToken(c context.Context, positionID int) (token string, err error)
	LookupPositionID(c context.Context, token string) (positionID int, err error)
	DeleteToken(c context.Context, token string) (err error)
}
