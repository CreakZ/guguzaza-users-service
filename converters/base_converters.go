package converters

import (
	"guguzaza-users/adapters/repository/models"
	"guguzaza-users/domain/entities"
	"guguzaza-users/http/dto"
)

func UserBaseFromDtoToEntity(user dto.UserBase) entities.UserBase {
	return entities.UserBase{
		Nickname: user.Nickname,
		Password: user.Password,
		Uuid:     user.Uuid,
	}
}

func UserBaseFromEntityToModel(user entities.UserBase) models.UserBase {
	return models.UserBase{
		Nickname: user.Nickname,
		Password: user.Password,
		Uuid:     user.Uuid,
	}
}

func UserBaseFromModelToEntity(user models.UserBase) entities.UserBase {
	return entities.UserBase{
		Nickname: user.Nickname,
		Password: user.Password,
		Uuid:     user.Uuid,
	}
}
