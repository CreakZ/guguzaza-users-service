package domain

import (
	"context"
	"errors"
	"fmt"
	"guguzaza-users/converters"
	"guguzaza-users/domain/entities"
	"guguzaza-users/domain/utils"
	"guguzaza-users/domain/validation"
	ports "guguzaza-users/ports/repository"
	"time"
)

type membersDomain struct {
	membersPort ports.MembersRepositoryPort
}

func NewMembersDomain(membersPort ports.MembersRepositoryPort) MembersDomain {
	return &membersDomain{
		membersPort: membersPort,
	}
}

func (md *membersDomain) RegisterMember(c context.Context, memberData entities.MemberCreate) (id int, err error) {
	valid, errMsg := validation.CheckNicknameValidity(memberData.Nickname)
	if !valid {
		return 0, errors.New(errMsg)
	}

	valid, errMsg = validation.CheckPasswordValidity(memberData.Password)
	if !valid {
		return 0, errors.New(errMsg)
	}

	unique, err := md.membersPort.CheckMemberNicknameUniqueness(c, memberData.Nickname)
	if err != nil {
		return 0, err
	}

	if !unique {
		return 0, fmt.Errorf("пользователь с указанным логином уже существует")
	}

	hashedPassword, err := utils.HashPassword(memberData.Password)
	if err != nil {
		return 0, err
	}

	memberDataCp := memberData
	memberDataCp.Password = hashedPassword

	memberModel := converters.MemberBaseModelFromMemberCreate(memberDataCp, time.Now())

	return md.membersPort.RegisterMember(c, memberModel)
}

// TODO
func (md *membersDomain) AuthMember(c context.Context, memberData entities.UserBase) (jwt string, err error) {
	return
}

func (md *membersDomain) GetMemberByID(c context.Context, id int) (member entities.Member, err error) {
	memberModel, err := md.membersPort.GetMemberByID(c, id)
	if err != nil {
		return
	}

	return converters.MemberFromEntityToModel(memberModel), nil
}

func (md *membersDomain) GetMembersPaginated(c context.Context, offset, limit int) (members []entities.MemberPublic, err error) {
	membersEntity, err := md.membersPort.GetMembersPaginated(c, offset, limit)
	if err != nil {
		return members, err
	}

	members = make([]entities.MemberPublic, len(membersEntity))
	for _, member := range membersEntity {
		members = append(members, converters.MemberPublicFromModelToEntity(member))
	}

	return members, nil
}

// mocked right now
func (md *membersDomain) GetTotalMembers(c context.Context) (total int64, err error) {
	return 1337, nil
}

func (md *membersDomain) UpdateMember(c context.Context, id int, updateData entities.MemberUpdate) (err error) {
	updates := converters.MemberUpdateEntityToUpdatesMap(updateData)

	return md.membersPort.UpdateMember(c, id, updates)
}

func (md *membersDomain) DeleteMember(c context.Context, id int) (err error) {
	return md.membersPort.DeleteMember(c, id)
}
