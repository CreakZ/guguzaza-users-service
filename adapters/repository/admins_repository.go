package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"guguzaza-users/adapters/repository/models"
	ports "guguzaza-users/ports/repository"

	sq "github.com/Masterminds/squirrel"
)

type adminsRepository struct {
	db *sql.DB
}

func NewAdminsRepository(db *sql.DB) ports.AdminsRepositoryPort {
	return adminsRepository{
		db: db,
	}
}

func (ar adminsRepository) CheckAdminNicknameUniqueness(c context.Context, nickname string) (unique bool, err error) {
	query, args, err := sq.
		Select("1").
		From("admins").
		Where(sq.Eq{"nickname": nickname}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, err
	}

	row := ar.db.QueryRowContext(c, query, args...)

	var one int
	if err = row.Scan(&one); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}

	return false, nil
}

func (ar adminsRepository) RegisterAdmin(c context.Context, adminData models.AdminRegister) (id int, err error) {
	query, args, err := sq.
		Insert("admins").
		Columns("nickname", "password", "admin_uuid", "position_id", "join_date").
		Values(adminData.Nickname, adminData.Password, adminData.Uuid, adminData.PositionID, adminData.JoinDate).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	row := ar.db.QueryRowContext(c, query, args...)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ar adminsRepository) GetAdminByID(c context.Context, id int) (admin models.AdminPublic, err error) {
	query, args, err := sq.
		Select("a.id", "a.nickname", "a.admin_uuid", "a.join_date", "p.position").
		From("admins a").
		Join("positions p ON a.position_id = p.id").
		Where(sq.Eq{"a.id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return models.AdminPublic{}, err
	}

	row := ar.db.QueryRowContext(c, query, args...)
	err = row.Scan(&admin.ID, &admin.Nickname, &admin.Uuid, &admin.JoinDate, &admin.Position)
	if err != nil {
		return models.AdminPublic{}, err
	}

	return admin, nil
}

func (ar adminsRepository) GetAdminByUuid(c context.Context, uuid string) (admin models.AdminPublic, err error) {
	query, args, err := sq.
		Select("a.id", "a.nickname", "a.admin_uuid", "a.join_date", "p.position").
		From("admins a").
		Join("positions p ON a.position_id = p.id").
		Where(sq.Eq{"a.admin_uuid": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return models.AdminPublic{}, err
	}

	row := ar.db.QueryRowContext(c, query, args...)
	err = row.Scan(&admin.ID, &admin.Nickname, &admin.Uuid, &admin.JoinDate, &admin.Position)
	if err != nil {
		return models.AdminPublic{}, err
	}

	return admin, nil
}

func (ar adminsRepository) GetAdminUserBaseByNickname(c context.Context, nickname string) (adminBase models.UserBase, err error) {
	query, args, err := sq.
		Select("nickname", "password", "admin_uuid").
		From("admins").
		Where(sq.Eq{"nickname": nickname}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return models.UserBase{}, err
	}

	row := ar.db.QueryRowContext(c, query, args...)
	err = row.Scan(&adminBase.Nickname, &adminBase.Password, &adminBase.Uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("пользователь не найден")
		}
		return models.UserBase{}, err
	}

	return adminBase, nil
}

func (ar adminsRepository) GetAdminsPaginated(c context.Context, offset, limit int64) (admins models.AdminsPaginated, err error) {
	admins.Admins = make([]models.AdminPublic, 0, limit)
	admins.Limit = int(limit)

	query, args, err := sq.
		Select("a.id", "a.nickname", "p.position", "a.admin_uuid", "a.join_date", "COUNT(*) OVER() AS total_count").
		From("admins").
		Join("positions p ON a.position_id = p.id").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return models.AdminsPaginated{}, err
	}

	rows, err := ar.db.QueryContext(c, query, args...)
	if err != nil {
		return models.AdminsPaginated{}, err
	}
	defer rows.Close()

	for rows.Next() {
		admin := models.AdminPublic{}

		err = rows.Scan(&admin.ID, &admin.Nickname, &admin.Position, &admin.Uuid, &admin.JoinDate, &admins.TotalCount)
		if err != nil {
			return models.AdminsPaginated{}, err
		}

		admins.Admins = append(admins.Admins, admin)
	}

	return admins, nil
}

func (ar adminsRepository) DeleteAdminByID(c context.Context, id int) (err error) {
	query, args, err := sq.
		Delete("admins").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = ar.db.ExecContext(c, query, args...); err != nil {
		return err
	}

	return nil
}

func (ar adminsRepository) DeleteAdminByUuid(c context.Context, uuid string) (err error) {
	query, args, err := sq.
		Delete("admins").
		Where(sq.Eq{"admin_uuid": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = ar.db.ExecContext(c, query, args...); err != nil {
		return err
	}

	return nil
}
