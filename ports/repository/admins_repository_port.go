package ports

import (
	"context"
	"guguzaza-users/adapters/repository/models"
)

type AdminsRepositoryPort interface {
	RegisterAdminUser(c context.Context, adminData models.AdminRegister) (id int, err error)

	GetAdminUserById(c context.Context, id int) (admin models.Admin, err error)
	GetAdminUserPasswordByNickname(c context.Context, nickname string) (password string, err error)
	GetAdminUsersPaginated(c context.Context, offset, limit int64) []models.Admin

	DeleteAdminUser(c context.Context)
}
