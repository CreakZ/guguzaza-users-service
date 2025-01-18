package domain

import (
	"context"
	"guguzaza-users/domain/entities"
)

type AdminsDomain interface {
}

type MembersDomain interface {
	RegisterMember(c context.Context, memberData entities.MemberCreate) (id int, err error)

	LoginMember(c context.Context, nickname, password string) (jwt string, err error)

	GetMemberByID(c context.Context, id int) (member entities.MemberPublic, err error)
	GetMembersPaginated(c context.Context, offset, limit int) (members []entities.MemberPublic, err error)
	GetTotalMembers(c context.Context) (total int64, err error)

	UpdateMember(c context.Context, id int, updateData entities.MemberUpdate) (err error)

	DeleteMember(c context.Context, id int) (err error)
}
