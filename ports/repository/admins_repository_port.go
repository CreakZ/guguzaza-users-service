package ports

import (
	"context"
	"guguzaza-users/adapters/repository/models"
)

type AdminsRepositoryPort interface {
	CheckAdminNicknameUniqueness(c context.Context, nickname string) (unique bool, err error)
	RegisterAdmin(c context.Context, adminData models.AdminRegister) (id int, err error)

	GetAdminByID(c context.Context, id int) (admin models.AdminPublic, err error)
	GetAdminByUuid(c context.Context, uuid string) (admin models.AdminPublic, err error)
	GetAdminUserBaseByNickname(c context.Context, nickname string) (adminBase models.UserBase, err error)
	GetAdminsPaginated(c context.Context, offset, limit int64) (admins models.AdminsPaginated, err error)

	DeleteAdminByID(c context.Context, id int) (err error)
	DeleteAdminByUuid(c context.Context, uuid string) (err error)
}
