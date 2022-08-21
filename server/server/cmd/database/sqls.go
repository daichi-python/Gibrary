package database

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

type QueryParams map[string]string

type ColumnParams []string

func (DBc *DBController) InsertUser(params QueryParams) (sql.Result, error) {
	query := InsertQueryHelper("users", params)
	logrus.WithFields(logrus.Fields{
		"table": "users",
		"query": query,
	}).Debugln("[--InsertUser--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"name": params["name"],
		}).Errorln("Failed to Insert user record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) UpdateUser(id string, params QueryParams) (sql.Result, error) {
	query := UpdateQueryHelper("users", id, params)

	logrus.WithFields(logrus.Fields{
		"table": "users",
		"id":    id,
		"query": query,
	}).Debugln("[--UpdateUser--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "users",
			"id":    id,
		}).Errorln("Failed to Update user record.")
		return nil, err
	}
	return result, nil
}

// The only values that params accepts are ID or EMAIL.
func (DBc *DBController) SelectUser(params QueryParams) *sql.Row {
	columns := ColumnParams{
		"id", "name", "email", "gender", "birthday", "is_active",
	}
	query := SelectQueryHelper("users", params, columns)

	row := DBc.db.QueryRow(query)
	return row
}

func (DBc *DBController) InsertGroupy(params QueryParams) (sql.Result, error) {
	query := InsertQueryHelper("groupies", params)
	logrus.WithFields(logrus.Fields{
		"table": "groupies",
		"query": query,
	}).Debugln("[--InsertGroupy--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "groupies",
			"name":  params["name"],
		}).Errorln("Failed to create groupies record.")
		return nil, err
	}
	return result, nil
}

func (DBc *DBController) UpdateGroupy(id string, params QueryParams) (sql.Result, error) {
	query := UpdateQueryHelper("groupies", id, params)
	logrus.WithFields(logrus.Fields{
		"table": "groupies",
		"query": query,
	}).Debugln("[--UpdateGroupy--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "groupies",
			"name":  params["name"],
		}).Errorln("Failed to update groupies record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) DeleteGroupy(id string) (sql.Result, error) {
	query := DeleteQueryHelper("groupies", id)
	logrus.WithFields(logrus.Fields{
		"table": "groupies",
		"query": query,
	}).Debugln("[--DeleteGroupy--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "groupies",
			"id":    id,
		}).Errorln("Failed to delete groupies record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) InsertUserInGroupy(params QueryParams) (sql.Result, error) {
	query := InsertQueryHelper("user_in_groupy", params)
	logrus.WithFields(logrus.Fields{
		"table": "user_in_groupy",
		"query": query,
	}).Debugln("[--InsertUserInGroupy--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "groupies",
		}).Errorln("Failed to insert groupies record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) DeleteUserInGroupy(id string) (sql.Result, error) {
	query := DeleteQueryHelper("user_in_groupy", id)
	logrus.WithFields(logrus.Fields{
		"table": "user_in_groupy",
		"query": query,
	}).Debugln("[--DeleteUserInGroupy--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "groupies",
			"id":    id,
		}).Errorln("Failed to delete groupies record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) InsertBoardItem(params QueryParams) (sql.Result, error) {
	query := InsertQueryHelper("board_items", params)
	logrus.WithFields(logrus.Fields{
		"table": "board_items",
		"query": query,
	}).Debugln("[--InsertBoardItem--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "board_items",
			"title": params["title"],
		}).Errorln("Failed to insert board_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) UpdateBoardItem(id string, params QueryParams) (sql.Result, error) {
	query := UpdateQueryHelper("board_items", id, params)
	logrus.WithFields(logrus.Fields{
		"table": "board_items",
		"id":    id,
		"query": query,
	}).Debugln("[--UpdateBoardItem--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "board_items",
			"title": params["title"],
		}).Errorln("Failed to update board_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) DeleteBoardItem(id string) (sql.Result, error) {
	query := DeleteQueryHelper("board_items", id)
	logrus.WithFields(logrus.Fields{
		"table": "board_items",
		"query": query,
	}).Debugln("[--DeleteBoardItem--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "board_items",
			"id":    id,
		}).Errorln("Failed to delete board_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) InsertGroupyBoardItem(params QueryParams) (sql.Result, error) {
	query := InsertQueryHelper("groupy_board_items", params)
	logrus.WithFields(logrus.Fields{
		"table": "groupy_board_items",
		"query": query,
	}).Debugln("[--InsertGroupyBoardItem--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "groupy_board_items",
		}).Errorln("Failed to insert groupy_board_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) DeleteGroupyBoardItem(id string) (sql.Result, error) {
	query := DeleteQueryHelper("groupy_board_items", id)
	logrus.WithFields(logrus.Fields{
		"table": "groupy_board_items",
		"query": query,
	}).Debugln("[--DeleteGroupyBoardItem--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "groupy_board_items",
			"id":    id,
		}).Errorln("Failed to delete groupy_board_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) InsertUserLikeBoardItem(params QueryParams) (sql.Result, error) {
	query := InsertQueryHelper("user_like_board_items", params)
	logrus.WithFields(logrus.Fields{
		"table": "user_like_board_items",
		"query": query,
	}).Debugln("[--InsertUserLikeBoardItem--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "user_like_board_items",
		}).Errorln("Failed to insert user_like_board_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) DeleteUserLikeBoardItem(id string) (sql.Result, error) {
	query := DeleteQueryHelper("user_like_board_items", id)
	logrus.WithFields(logrus.Fields{
		"table": "user_like_board_items",
		"query": query,
	}).Debugln("[--DeleteUserLikeBoardItem--] prepared query")

	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"table": "user_like_board_items",
		}).Errorln("Failed to delete user_like_board_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) InsertHomeItem(params QueryParams) (sql.Result, error) {
	query := InsertQueryHelper("home_items", params)
	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Table": "home_items",
			"Error": err,
		}).Errorln("Failed to insert home_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) UpdateHomeItem(id string, params QueryParams) (sql.Result, error) {
	query := UpdateQueryHelper("home_items", id, params)
	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Table": "home_items",
			"Error": err,
		}).Errorln("Failed to update home_items record.")
		return nil, err
	}

	return result, nil
}

func (DBc *DBController) DeleteHomeItem(id string) (sql.Result, error) {
	query := DeleteQueryHelper("home_items", id)
	result, err := DBc.db.Exec(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Table": "home_items",
			"Error": err,
		}).Errorln("Failed to delete home_items record.")
		return nil, err
	}

	return result, nil
}

func InsertQueryHelper(table string, params QueryParams) string {
	reg := regexp.MustCompile("^is_/*")

	columns := ""
	values := ""
	for key, value := range params {
		if reg.MatchString(key) {
			columns += fmt.Sprintf("%s, ", key)
			values += fmt.Sprintf("%s, ", value)
		} else {
			columns += fmt.Sprintf("%s, ", key)
			values += fmt.Sprintf("'%s', ", value)
		}
	}
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", table, columns[:len(columns)-2], values[:len(values)-2])
	return query
}

func UpdateQueryHelper(table, id string, params QueryParams) string {
	reg := regexp.MustCompile("^is_.*")

	valueSets := ""
	for key, value := range params {
		if reg.MatchString(key) {
			valueSets += fmt.Sprintf("%s = %s, ", key, value)
		} else {
			valueSets += fmt.Sprintf("%s = '%s', ", key, value)
		}
	}
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %s", table, valueSets[:len(valueSets)-2], id)
	return query
}

func DeleteQueryHelper(table, id string) string {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %s", table, id)
	return query
}

// Receive a single param
func SelectQueryHelper(table string, params QueryParams, cp ColumnParams) string {
	columns := ""
	for _, value := range cp {
		columns += fmt.Sprintf("%s, ", value)
	}

	values := ""
	for key, value := range params {
		if _, err := strconv.Atoi(value); err == nil {
			values += fmt.Sprintf("WHERE %s = %s", key, value)
		} else {
			pm := "%" + value + "%"
			values += fmt.Sprintf("WHERE %s LIKE %s", key, pm)
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s %s;", columns[:len(columns)-2], table, values)

	return query
}

// Helper function to score each column and extract records with a default score or higher
func SelectQueryWitchCaseHelper(table string, params QueryParams, meter int) string {
	caseQuery := ""
	total := ""

	for key, value := range params {
		cq, column := CaseQueryHelper(key, value)

		if caseQuery == "" {
			caseQuery += cq
			total += column
		} else {
			caseQuery += fmt.Sprintf(", %s", cq)
			total += fmt.Sprintf(", %s", column)
		}
	}

	query := fmt.Sprintf("SELECT *, %s AS total FROM (SELECT *, %s FROM %s) AS temp_table HAVING total >= %v ORDER BY total DESC;", total, caseQuery, table, meter)

	return query
}

// Returns the name of the table to be used temporarily with the query of case
func CaseQueryHelper(column string, value string) (string, string) {
	query := ""

	// If value is a string, use pattern matching
	if _, err := strconv.Atoi(value); err == nil {
		query = fmt.Sprintf("CASE WHEN %s = %v THEN 1 ELSE 0 END AS %s_match", column, value, column)
	} else {
		// Because values cannot be enclosed in % in fmt.Sprintf
		pm := "'%" + value + "%'"
		query = fmt.Sprintf("CASE WHEN %s LIKE %s THEN 1 ELSE 0 END AS %s_match", column, pm, column)
	}

	return query, fmt.Sprintf("%s_match", column)
}
