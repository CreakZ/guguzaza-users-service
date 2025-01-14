package ports

import "context"

type JwtUtilPort interface {
	CreateJwt(c context.Context, userUuid string) (jwt string, err error) // userUuid is a 'sub' claim
	ParseJwtClaims(c context.Context, jwt string) (userUuid string, err error)
}
