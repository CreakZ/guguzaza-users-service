package domain

import (
	"context"
	"errors"
	"guguzaza-users/adapters/repository/models"
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

type adminsDomain struct {
	adminsRepo       repo_ports.AdminsRepositoryPort
	jwtUtil          token_ports.JwtUtilPort
	inviteTokensUtil token_ports.InviteTokensUtilPort
}

func NewAdminsDomain(
	adminsRepo repo_ports.AdminsRepositoryPort,
	jwtUtil token_ports.JwtUtilPort,
	inviteTokensUtil token_ports.InviteTokensUtilPort,
) AdminsDomain {
	return adminsDomain{
		adminsRepo:       adminsRepo,
		jwtUtil:          jwtUtil,
		inviteTokensUtil: inviteTokensUtil,
	}
}

func (ad adminsDomain) CreateInviteToken(c context.Context, adminUuid string, positionID int) (token string, err error) {
	admin, err := ad.adminsRepo.GetAdminByUuid(c, adminUuid)
	if err != nil {
		return "", err
	}
	_ = admin // use this for logging

	return ad.inviteTokensUtil.CreateToken(c, positionID)
}

func (ad adminsDomain) RegisterAdmin(c context.Context, adminData entities.AdminCreate) (id int, err error) {
	positionID, err := ad.inviteTokensUtil.LookupPositionID(c, adminData.InviteToken)
	if err != nil {
		return 0, err
	}

	valid, errMsg := validation.ValidateAdminCreate(adminData)
	if !valid {
		return 0, errors.New(errMsg)
	}

	unique, err := ad.adminsRepo.CheckAdminNicknameUniqueness(c, adminData.Nickname)
	if err != nil {
		return 0, err
	}

	if !unique {
		return 0, errors.New("пользователь с таким никнеймом уже существует")
	}

	hashedPassword, err := utils.HashPassword(adminData.Password)
	if err != nil {
		return 0, err
	}

	adminDataModel := models.AdminRegister{
		UserBase: models.UserBase{
			Nickname: adminData.Nickname,
			Password: hashedPassword,
			Uuid:     uuid.NewString(),
		},
		PositionID: positionID,
		JoinDate:   time.Now(),
	}

	id, err = ad.adminsRepo.RegisterAdmin(c, adminDataModel)
	if err == nil {
		ad.inviteTokensUtil.DeleteToken(context.Background(), adminData.InviteToken)
	}

	return id, err
}

func (ad adminsDomain) LoginAdmin(c context.Context, nickname, password string) (jwt string, err error) {
	adminBase, err := ad.adminsRepo.GetAdminUserBaseByNickname(c, nickname)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminBase.Password), []byte(password))
	if err != nil {
		return "", errors.New("неверный пароль")
	}

	return ad.jwtUtil.CreateJwt(c, adminBase.Uuid)
}

func (ad adminsDomain) GetAdminByID(c context.Context, id int) (admin entities.AdminPublic, err error) {
	adminModel, err := ad.adminsRepo.GetAdminByID(c, id)
	if err != nil {
		return entities.AdminPublic{}, err
	}

	return converters.AdminPublicFromModelToEntity(adminModel), nil
}

func (ad adminsDomain) GetAdminByUuid(c context.Context, uuid string) (admin entities.AdminPublic, err error) {
	adminModel, err := ad.adminsRepo.GetAdminByUuid(c, uuid)
	if err != nil {
		return entities.AdminPublic{}, err
	}

	return converters.AdminPublicFromModelToEntity(adminModel), nil
}

func (ad adminsDomain) GetAdminsPaginated(c context.Context, page, limit int64) (admins entities.AdminsPaginated, err error) {
	if page < 1 {
		return entities.AdminsPaginated{}, errors.New("значение 'page' не должно быть меньше 1")
	}

	if limit < 10 || limit > 50 {
		return entities.AdminsPaginated{}, errors.New("значение 'limit' не должно быть в диапазоне от 10 до 50 (включительно)")
	}

	offset := (page - 1) * limit
	adminsPaginatedModel, err := ad.adminsRepo.GetAdminsPaginated(c, offset, limit)
	if err != nil {
		return entities.AdminsPaginated{}, err
	}

	totalPages := adminsPaginatedModel.TotalCount/adminsPaginatedModel.Limit + 1

	return converters.AdminsPaginatedFromModelToEntity(adminsPaginatedModel, totalPages), nil
}

func (ad adminsDomain) DeleteAdminByID(c context.Context, id int) (err error) {
	return ad.adminsRepo.DeleteAdminByID(c, id)
}

func (ad adminsDomain) DeleteAdminByUuid(c context.Context, uuid string) (err error) {
	return ad.adminsRepo.DeleteAdminByUuid(c, uuid)
}
