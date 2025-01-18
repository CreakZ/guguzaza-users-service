package converters

import (
	"guguzaza-users/adapters/repository/models"
	"guguzaza-users/domain/entities"
	"guguzaza-users/http/dto"
	"time"
)

func MemberBaseFromDtoToEntity() {

}

func MemberBaseFromEntityToModel(member entities.MemberBase) models.MemberBase {
	return models.MemberBase{
		UserBase: UserBaseFromEntityToModel(member.UserBase),
		JoinDate: member.JoinDate,
		Sex:      member.Sex,
		About:    member.About,
	}
}

func MemberBaseFromModelToEntity(member models.MemberBase) entities.MemberBase {
	return entities.MemberBase{
		UserBase: UserBaseFromModelToEntity(member.UserBase),
		JoinDate: member.JoinDate,
		Sex:      member.Sex,
		About:    member.About,
	}
}

func MemberBaseFromEntityToDto() {

}

func MemberFromEntityToModel(member models.Member) entities.Member {
	return entities.Member{
		ID:         member.ID,
		MemberBase: MemberBaseFromModelToEntity(member.MemberBase),
	}
}

func MemberFromModelToEntity(member entities.Member) models.Member {
	return models.Member{
		ID:         member.ID,
		MemberBase: MemberBaseFromEntityToModel(member.MemberBase),
	}
}

func MemberPublicFromDtoToEntity() {

}

func MemberPublicFromEntityToModel(member entities.MemberPublic) models.MemberPublic {
	return models.MemberPublic{
		ID:       member.ID,
		Nickname: member.Nickname,
		Uuid:     member.Uuid,
		JoinDate: member.JoinDate,
		Sex:      member.Sex,
		About:    member.About,
	}
}

func MemberPublicFromModelToEntity(member models.MemberPublic) entities.MemberPublic {
	return entities.MemberPublic{
		ID:       member.ID,
		Nickname: member.Nickname,
		Uuid:     member.Uuid,
		JoinDate: member.JoinDate,
		Sex:      member.Sex,
		About:    member.About,
	}
}

func MemberPublicFromEntityToDto(member entities.MemberPublic) dto.MemberPublic {
	return dto.MemberPublic{
		ID:       member.ID,
		Nickname: member.Nickname,
		Uuid:     member.Uuid,
		JoinTime: dto.JoinTime{Time: member.JoinDate},
		Sex:      member.Sex,
		About:    member.About,
	}
}

func MemberUpdateEntityToUpdatesMap(member entities.MemberUpdate) map[string]interface{} {
	updates := make(map[string]interface{}, 0)

	if member.Sex != nil {
		updates["sex"] = *member.Sex
	}

	if member.About != nil {
		updates["about"] = *member.About
	}

	return updates
}

func MemberCreateFromDtoToEntity(member dto.MemberCreate) entities.MemberCreate {
	return entities.MemberCreate{
		UserBase: entities.UserBase{
			Nickname: member.Nickname,
			Password: member.Password,
		},
		Sex:   member.Sex,
		About: member.About,
	}
}

func MemberBaseModelFromMemberCreate(member entities.MemberCreate, joinDate time.Time) models.MemberBase {
	var sexModel string
	switch member.Sex {
	case "f":
		sexModel = "Женщинский"
	default:
		sexModel = "Мужчинский"
	}
	return models.MemberBase{
		UserBase: UserBaseFromEntityToModel(member.UserBase),
		JoinDate: joinDate,
		Sex:      sexModel,
		About:    member.About,
	}
}
