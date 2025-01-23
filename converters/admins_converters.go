package converters

import (
	"guguzaza-users/adapters/repository/models"
	"guguzaza-users/domain/entities"
	"guguzaza-users/http/dto"
)

func AdminCreateFromDtoToEntity(admin dto.AdminCreate) entities.AdminCreate {
	return entities.AdminCreate{
		Nickname:    admin.Nickname,
		Password:    admin.Password,
		InviteToken: admin.InviteToken,
	}
}

func AdminPublicFromEntityToDto(admin entities.AdminPublic) dto.AdminPublic {
	return dto.AdminPublic{
		ID:       admin.ID,
		Nickname: admin.Nickname,
		Uuid:     admin.Uuid,
		Position: admin.Position,
		JoinTime: dto.JoinTime{
			Time: admin.JoinDate,
		},
	}
}

func AdminPublicFromModelToEntity(admin models.AdminPublic) entities.AdminPublic {
	return entities.AdminPublic{
		ID:       admin.ID,
		Nickname: admin.Nickname,
		Uuid:     admin.Uuid,
		Position: admin.Position,
		JoinDate: admin.JoinDate,
	}
}

func AdminsPaginatedFromModelToEntity(admins models.AdminsPaginated, totalPages int) entities.AdminsPaginated {
	items := make([]entities.AdminPublic, 0, len(admins.Admins))

	for _, admin := range admins.Admins {
		items = append(items, AdminPublicFromModelToEntity(admin))
	}

	return entities.AdminsPaginated{
		Limit:      admins.Limit,
		TotalPages: totalPages,
		Admins:     items,
	}
}

func AdminsPaginatedFromEntityToDto(admins entities.AdminsPaginated) dto.AdminsPaginated {
	items := make([]dto.AdminPublic, 0, len(admins.Admins))

	for _, admin := range admins.Admins {
		items = append(items, AdminPublicFromEntityToDto(admin))
	}

	return dto.AdminsPaginated{
		Limit:      admins.Limit,
		TotalPages: admins.TotalPages,
		Admins:     items,
	}
}
