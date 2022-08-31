package converter

import (
	"database/sql"
	"reflect"
	"strings"

	"server/cmd/database"

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

func UserRowsToStruct(rows *sql.Rows) ([]User, error) {
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Gender, &user.Birthday, &user.Is_active)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Row":   "User",
				"Error": err,
			}).Errorln("Failed to scan user record.")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GroupyRowsToStruct(rows *sql.Rows) ([]Groupy, error) {
	var groupies []Groupy

	for rows.Next() {
		var groupy Groupy
		err := rows.Scan(&groupy.ID, &groupy.Name, &groupy.Introduce, &groupy.Max_people, &groupy.Group_key, &groupy.Is_opened)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Row":   "Groupy",
				"Error": err,
			}).Errorln("Failed to scan groupy record.")
			return nil, err
		}
		groupies = append(groupies, groupy)
	}
	return groupies, nil
}

func UserInGroupyRowsToStruct(rows *sql.Rows) ([]UserInGroupy, error) {
	var uigs []UserInGroupy

	for rows.Next() {
		var uig UserInGroupy
		err := rows.Scan(&uig.ID, &uig.User, &uig.Groupy)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Row":   "UserInGroupy",
				"Error": err,
			}).Errorln("Failed to scan UserInGroupy record.")
		}
		uigs = append(uigs, uig)
	}
	return uigs, nil
}
