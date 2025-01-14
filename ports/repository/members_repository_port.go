package ports

import (
	"context"
	"guguzaza-users/adapters/repository/models"
)

type MembersRepositoryPort interface {
	CheckMemberNicknameUniqueness(c context.Context, nickname string) (unique bool, err error)
	RegisterMember(c context.Context, memberData models.MemberBase) (id int, err error)

	GetMemberById(c context.Context, id int) (member models.Member, err error)
	GetMemberPasswordByNickname(c context.Context, nickname string) (password string, err error)
	GetMembersPaginated(c context.Context, offset, limit int) (members []models.MemberPublic, err error)
	GetTotalMembers(c context.Context) (total int, err error)

	UpdateMember(c context.Context, id int, updates map[string]interface{}) (err error)

	DeleteMember(c context.Context, id int) (err error)
}
