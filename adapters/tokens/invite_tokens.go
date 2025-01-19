package tokens

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	repo_ports "guguzaza-users/ports/repository"
	token_ports "guguzaza-users/ports/tokens"
)

type inviteTokensUtil struct {
	inviteTokensRepo repo_ports.InviteTokensRepositoryPort
}

func NewInviteTokensUtil(inviteTokensRepo repo_ports.InviteTokensRepositoryPort) token_ports.InviteTokensUtilPort {
	return inviteTokensUtil{
		inviteTokensRepo: inviteTokensRepo,
	}
}

func (itu inviteTokensUtil) CreateToken(c context.Context, positionID int) (token string, err error) {
	// token bytes is 15 bytes long slice because hex.EncodeToString returns
	// hexadecimal encoding of token bytes, which is 30 bytes long
	tokenBytes := make([]byte, 15)
	if _, err = rand.Read(tokenBytes); err != nil {
		return "", err
	}

	token = hex.EncodeToString(tokenBytes)
	err = itu.inviteTokensRepo.CreateToken(c, token, positionID)

	return
}

func (itu inviteTokensUtil) LookupPositionID(c context.Context, token string) (positionID int, err error) {
	return itu.inviteTokensRepo.LookupPositionID(c, token)
}

func (itu inviteTokensUtil) DeleteToken(c context.Context, token string) (err error) {
	return itu.inviteTokensRepo.DeleteToken(c, token)
}
