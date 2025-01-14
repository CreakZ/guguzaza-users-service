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

type membersRepository struct {
	db *sql.DB
}

func NewMembersRepository(db *sql.DB) ports.MembersRepositoryPort {
	return &membersRepository{
		db: db,
	}
}

func (mr *membersRepository) CheckMemberNicknameUniqueness(c context.Context, nickname string) (unique bool, err error) {
	query, args, err := sq.
		Select("1").
		From("members").
		Where(sq.Eq{"nickname": nickname}).
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, err
	}

	row := mr.db.QueryRowContext(c, query, args...)

	var one int
	if err = row.Scan(&one); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}

	return false, nil
}

func (mr *membersRepository) RegisterMember(c context.Context, memberData models.MemberBase) (id int, err error) {
	tx, err := mr.db.BeginTx(c, nil)
	if err != nil {
		return 0, err
	}

	query, args, err := sq.
		Insert("members").
		Columns("nickname", "password", "member_uuid", "join_date", "sex", "about").
		Values(memberData.Nickname, memberData.Password, memberData.Uuid, memberData.JoinDate, memberData.Sex, memberData.About).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	row := tx.QueryRowContext(c, query, args...)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("ошибка завершения транзакции: %w", err)
	}

	return id, nil
}

func (mr *membersRepository) GetMemberUserBaseByNickname(c context.Context, nickname string) (memberBase models.UserBase, err error) {
	query, args, err := sq.
		Select("nickname", "password", "member_uuid").
		From("members").
		Where(sq.Eq{"nickname": nickname}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	row := mr.db.QueryRowContext(c, query, args...)
	if err = row.Scan(&memberBase.Nickname, &memberBase.Password, &memberBase.Uuid); err != nil {
		return models.UserBase{}, err
	}

	return memberBase, nil
}

// change models.Member to models.MemberPublic
func (mr *membersRepository) GetMemberByID(c context.Context, id int) (member models.Member, err error) {
	query, args, err := sq.
		Select("id", "nickname", "member_uuid", "join_date", "sex", "about"). // password is skipped intentionally
		From("members").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return member, err // write precise error message
	}

	row := mr.db.QueryRowContext(c, query, args...)
	err = row.Scan(&member.ID, &member.Nickname, &member.Uuid, &member.JoinDate, &member.Sex, &member.About)
	if err != nil {
		return models.Member{}, err
	}

	return member, nil
}

func (mr *membersRepository) GetMemberIDByUuid(c context.Context, uuid string) (id int, err error) {
	query, args, err := sq.
		Select("id").
		From("members").
		Where(sq.Eq{"member_uuid": uuid}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	row := mr.db.QueryRowContext(c, query, args...)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (mr *membersRepository) GetMembersPaginated(c context.Context, offset, limit int) (members []models.MemberPublic, err error) {
	query, args, err := sq.
		Select("id", "nickname", "member_uuid", "join_date", "sex", "about").
		From("members").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return members, err // write precise error message
	}

	rows, err := mr.db.QueryContext(c, query, args...)
	if err != nil {
		return members, err
	}

	member := models.MemberPublic{}
	for rows.Next() {
		err = rows.Scan(&member.ID, &member.Nickname, &member.Uuid, &member.JoinDate, &member.Sex, &member.About)
		if err != nil {
			return []models.MemberPublic{}, err
		}
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		return []models.MemberPublic{}, err
	}

	return members, nil
}

func (mr *membersRepository) GetTotalMembers(c context.Context) (total int, err error) {
	row := mr.db.QueryRowContext(c, "SELECT COUNT(*) FROM members")
	if err = row.Scan(&total); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		} else {
			return 0, err
		}
	}

	return total, nil
}

// updates must be non-nil and should contain at least 1 key-value pair
func (mr *membersRepository) UpdateMember(c context.Context, id int, updates map[string]interface{}) (err error) {
	tx, err := mr.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	updBuilder := sq.Update("members")

	for column, value := range updates {
		updBuilder = updBuilder.Set(column, value)
	}

	query, args, err := updBuilder.
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := mr.db.ExecContext(c, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		return fmt.Errorf("ошибка обновления данных пользователя: указанный пользователь не найден")
	}

	return
}

func (mr *membersRepository) DeleteMember(c context.Context, id int) (err error) {
	tx, err := mr.db.BeginTx(c, nil)
	if err != nil {
		return
	}

	query, args, err := sq.
		Delete("members").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	res, err := tx.ExecContext(c, query, args...)
	if err != nil {
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		return fmt.Errorf("ошибка удаления пользователя: указанный пользователь не найден")
	}

	return nil
}
