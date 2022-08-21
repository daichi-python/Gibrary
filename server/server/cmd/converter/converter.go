package converter

import (
	"database/sql"
	"reflect"
	"server/cmd/database"
	"strings"

	"github.com/sirupsen/logrus"
)

func ToInsertQueryParams(s JsonStruct) database.QueryParams {
	params := database.QueryParams{}

	refType := reflect.TypeOf(s)
	refValue := reflect.ValueOf(s)

	for i := 0; i < refType.NumField(); i++ {
		TypeField := refType.Field(i)
		ValueField := refValue.Field(i)

		field := TypeField.Name
		field = strings.ToLower(field)
		value := ValueField.String()

		if value != "" {
			params[field] = value
		}
	}

	return params
}

func ToUpdateQueryParams(s JsonStruct) (id string, params database.QueryParams) {
	id = s.GetID()
	params = database.QueryParams{}

	refType := reflect.TypeOf(s)
	refValue := reflect.ValueOf(s)

	for i := 0; i < refType.NumField(); i++ {
		TypeField := refType.Field(i)
		ValueField := refValue.Field(i)

		field := TypeField.Name
		field = strings.ToLower(field)
		value := ValueField.String()

		if value != "" && field != "id" {
			params[field] = value
		}
	}

	return id, params
}

func ToSelectQueryParams(s JsonStruct) database.QueryParams {
	params := database.QueryParams{}

	refType := reflect.TypeOf(s)
	refValue := reflect.ValueOf(s)

	for i := 0; i < refType.NumField(); i++ {
		TypeField := refType.Field(i)
		ValueField := refValue.Field(i)

		field := TypeField.Name
		field = strings.ToLower(field)
		value := ValueField.String()

		if value != "" {
			params[field] = value
		}
	}

	return params
}

//TODO
//VALUEがNULLの場合でも上手く動作するように編集する。
func UserRowToStruct(row *sql.Row) (*User, error) {
	var user User
	var id, name, email, gender, birthday, is_active string

	err := row.Scan(&id, &name, &email, &gender, &birthday, &is_active)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Row":   "User",
			"Error": err,
		}).Errorln("Failed to scan record.")
		return nil, err
	}

	user.ID = id
	user.Name = name
	user.Email = email
	user.Gender = gender
	user.Birthday = birthday
	user.Is_active = is_active

	return &user, nil
}
