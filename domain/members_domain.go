package domain

import (
	"context"
	"errors"
	"fmt"
	"guguzaza-users/converters"
	"guguzaza-users/domain/entities"
	"guguzaza-users/domain/utils"
	"guguzaza-users/domain/validation"
	repo_ports "guguzaza-users/ports/repository"
	token_ports "guguzaza-users/ports/tokens"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type membersDomain struct {
	membersPort repo_ports.MembersRepositoryPort
	jwtUtil     token_ports.JwtUtilPort
}

func NewMembersDomain(membersPort repo_ports.MembersRepositoryPort, jwtUtil token_ports.JwtUtilPort) MembersDomain {
	return &membersDomain{
		membersPort: membersPort,
		jwtUtil:     jwtUtil,
	}
}

func (md *membersDomain) RegisterMember(c context.Context, memberData entities.MemberCreate) (id int, err error) {
	valid, errMsg := validation.ValidateMemberCreate(memberData)
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
	memberDataCp.Uuid = uuid.NewString()

	memberModel := converters.MemberBaseModelFromMemberCreate(memberDataCp, time.Now())

	return md.membersPort.RegisterMember(c, memberModel)
}

func (md *membersDomain) LoginMember(c context.Context, nickname, password string) (jwt string, err error) {
	memberBase, err := md.membersPort.GetMemberUserBaseByNickname(c, nickname)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(memberBase.Password), []byte(password)); err != nil {
		return "", errors.New("неверный пароль")
	}

	return md.jwtUtil.CreateJwt(c, memberBase.Uuid)
}

func (md *membersDomain) GetMemberByID(c context.Context, id int) (member entities.MemberPublic, err error) {
	memberModel, err := md.membersPort.GetMemberByID(c, id)
	if err != nil {
		return
	}

	return converters.MemberPublicFromModelToEntity(memberModel), nil
}

func (md *membersDomain) GetMembersPaginated(c context.Context, page, limit int) (members []entities.MemberPublic, err error) {
	if page < 1 {
		return []entities.MemberPublic{}, errors.New("значение 'page' не должно быть меньше 1")
	}

	if limit < 10 || limit > 50 {
		return []entities.MemberPublic{}, errors.New("значение 'limit' не должно быть в диапазоне от 10 до 50 (включительно)")
	}

	offset := (page - 1) * limit
	membersEntity, err := md.membersPort.GetMembersPaginated(c, offset, limit)
	if err != nil {
		return members, err
	}

	members = make([]entities.MemberPublic, 0, len(membersEntity))
	for _, member := range membersEntity {
		members = append(members, converters.MemberPublicFromModelToEntity(member))
	}

	return members, nil
}

func (md *membersDomain) GetTotalMembers(c context.Context) (total int64, err error) {
	return md.membersPort.GetTotalMembers(c)
}

func (md *membersDomain) UpdateMember(c context.Context, id int, updateData entities.MemberUpdate) (err error) {
	updates := converters.MemberUpdateEntityToUpdatesMap(updateData)

	return md.membersPort.UpdateMember(c, id, updates)
}

func (md *membersDomain) DeleteMember(c context.Context, id int) (err error) {
	return md.membersPort.DeleteMember(c, id)
}
