package converters

import (
	"guguzaza-users/adapters/repository/models"
	"guguzaza-users/domain/entities"
)

func AdminPublicFromModelToEntity(admin models.AdminPublic) entities.AdminPublic {
	return entities.AdminPublic{
		ID:       admin.ID,
		Nickname: admin.Nickname,
		Uuid:     admin.Uuid,
		Position: admin.Position,
		JoinDate: admin.JoinDate,
	}
}

func AdminsPaginatedFromModelToEntity(admins models.AdminsPaginated) entities.AdminsPaginated {
	adminsData := make([]entities.AdminPublic, 0, len(admins.Admins))

	for _, admin := range admins.Admins {
		adminsData = append(adminsData, AdminPublicFromModelToEntity(admin))
	}

	return entities.AdminsPaginated{
		Limit:      admins.Limit,
		TotalCount: admins.TotalCount,
		Admins:     adminsData,
	}
}
