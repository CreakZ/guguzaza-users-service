package domain

import (
	"context"
	"guguzaza-users/domain/entities"
)

type AdminsDomain interface {
	CreateInviteToken(c context.Context, adminUuid string, positionID int) (token string, err error)

	RegisterAdmin(c context.Context, adminData entities.AdminCreate) (id int, err error)

	LoginAdmin(c context.Context, nickname, password string) (jwt string, err error)

	GetAdminByID(c context.Context, id int) (admin entities.AdminPublic, err error)
	GetAdminByUuid(c context.Context, uuid string) (admin entities.AdminPublic, err error)
	GetAdminsPaginated(c context.Context, offset, limit int64) (admins entities.AdminsPaginated, err error)

	DeleteAdminByID(c context.Context, id int) (err error)
	DeleteAdminByUuid(c context.Context, uuid string) (err error)
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
