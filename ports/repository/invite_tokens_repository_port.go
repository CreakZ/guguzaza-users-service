package ports

import (
	"context"
	"guguzaza-users/adapters/repository/models"
)

type InviteTokensRepositoryPort interface {
	CreateToken(c context.Context, tokenData models.InviteTokenCreate)

	CheckToken(c context.Context, token string) (exists bool, err error)

	DeleteToken(c context.Context, token string) (err error)
}
