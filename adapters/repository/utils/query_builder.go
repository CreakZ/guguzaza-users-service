package repository_utils

import (
	"github.com/Masterminds/squirrel"
)

func BuildUserUpdateQuery(table string, userID int, updates map[string]interface{}) (query string, args []interface{}, err error) {
	updateBuilder := squirrel.Update(table)

	for column, value := range updates {
		updateBuilder = updateBuilder.Set(column, value)
	}

	return updateBuilder.Where(squirrel.Eq{"id": userID}).PlaceholderFormat(squirrel.Dollar).ToSql()
}
