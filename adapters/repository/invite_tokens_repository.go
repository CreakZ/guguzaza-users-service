package repository

import (
	"context"
	"database/sql"
	"errors"
	ports "guguzaza-users/ports/repository"

	sq "github.com/Masterminds/squirrel"
)

type inviteTokensRepository struct {
	db *sql.DB
}

func NewInviteTokensRepository(db *sql.DB) ports.InviteTokensRepositoryPort {
	return inviteTokensRepository{
		db: db,
	}
}

func (itr inviteTokensRepository) CreateToken(c context.Context, token string, positionID int) (err error) {
	query, args, err := sq.
		Insert("invite_tokens").
		Columns("token", "position_id").
		Values(token, positionID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = itr.db.ExecContext(c, query, args...)

	return err
}

func (itr inviteTokensRepository) LookupPositionID(c context.Context, token string) (positionID int, err error) {
	query, args, err := sq.
		Select("position_id").
		From("invite_tokens").
		Where(sq.Eq{"token": token}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	row := itr.db.QueryRowContext(c, query, args...)
	if err = row.Scan(&positionID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("токен не найден")
		}
	}

	return
}

func (itr inviteTokensRepository) DeleteToken(c context.Context, token string) (err error) {
	query, args, err := sq.
		Delete("invite_tokens").
		Where(sq.Eq{"token": token}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = itr.db.ExecContext(c, query, args...)
	switch err {
	case sql.ErrNoRows:
		return errors.New("токен не найден")
	default:
		return
	}
}
