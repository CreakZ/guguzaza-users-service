package ports

import (
	"context"
)

type InviteTokensRepositoryPort interface {
	CreateToken(c context.Context, token string, positionID int) (err error)
	LookupPositionID(c context.Context, token string) (positionID int, err error)
	DeleteToken(c context.Context, token string) (err error)
}
