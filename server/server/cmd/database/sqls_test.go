package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var controller *DBController

var (
	user_id       int64
	groupy_id     int64
	board_item_id int64
	home_item_id  int64
)

func setup() {
	log.SetPrefix("[--TESTING LOG--]")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error: Failed to read env file.")
		return
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	dbEngine, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(gibrary_test_db)/gibrary?charset=utf8mb4&parseTime=true", user, pass))
	if err != nil {
		log.Println("Error: Failed to connect MySQL server.")
	}

	controller = NewDBController(dbEngine)
}

func TestMain(m *testing.M) {
	setup()
	log.Println("SUCCESS: START TEST")
	code := m.Run()
	controller.db.Exec("DELETE FROM users;")
	os.Exit(code)
}

func TestInsertUser(t *testing.T) {
	email := generateRandomString(20)
	params := QueryParams{
		"email": fmt.Sprintf("%s@gmail.com", email),
		"name":  "富永大地",
	}

	result, err := controller.InsertUser(params)
	assert.Nil(t, err)

	user_id, _ = result.LastInsertId()
}

func TestUpdateUser(t *testing.T) {
	params := QueryParams{
		"name":      "富永台地",
		"is_active": "FALSE",
	}

	_, err := controller.UpdateUser(fmt.Sprint(user_id), params)
	assert.Nil(t, err)
}

func TestInsertAndDeleteGroupy(t *testing.T) {
	params := QueryParams{
		"name":      "富永家",
		"introduce": generateRandomString(100),
	}

	result, err := controller.InsertGroupy(params)
	assert.Nil(t, err)

	groupy_id, _ = result.LastInsertId()

	params2 := QueryParams{
		"name":      generateRandomString(10),
		"introduce": generateRandomString(100),
	}

	result, err = controller.InsertGroupy(params2)
	assert.Nil(t, err)

	num, _ := result.LastInsertId()
	_, err = controller.DeleteGroupy(fmt.Sprint(num))
	assert.Nil(t, err)
}

func TestUpdateGroupy(t *testing.T) {
	params := QueryParams{
		"name":       "富永屋",
		"max_people": "70",
	}

	_, err := controller.UpdateGroupy(fmt.Sprint(groupy_id), params)
	assert.Nil(t, err)
}

func TestInsertAndDeleteUserInGroupy(t *testing.T) {
	params := QueryParams{
		"User":   fmt.Sprint(user_id),
		"Groupy": fmt.Sprint(groupy_id),
	}

	result, err := controller.InsertUserInGroupy(params)
	assert.Nil(t, err)

	num, _ := result.LastInsertId()
	_, err = controller.DeleteUserInGroupy(fmt.Sprint(num))
	assert.Nil(t, err)
}

func TestInsertAndDeleteBoardItem(t *testing.T) {
	params := QueryParams{
		"Type":     "1",
		"Category": "2",
		"title":    generateRandomString(10),
	}

	result, err := controller.InsertBoardItem(params)
	assert.Nil(t, err)

	board_item_id, _ = result.LastInsertId()

	params2 := QueryParams{
		"Type":     "2",
		"Category": "1",
		"title":    generateRandomString(10),
	}

	result, err = controller.InsertBoardItem(params2)
	assert.Nil(t, err)

	num, _ := result.LastInsertId()
	_, err = controller.DeleteBoardItem(fmt.Sprint(num))
	assert.Nil(t, err)
}

func TestUpdateBoardItem(t *testing.T) {
	params := QueryParams{
		"title":      "updated title",
		"detail":     generateRandomString(100),
		"applicants": "20",
	}

	_, err := controller.UpdateBoardItem(fmt.Sprint(board_item_id), params)
	assert.Nil(t, err)
}

func TestInsertAndDeleteGroupyBoardItem(t *testing.T) {
	params := QueryParams{
		"Groupy":    fmt.Sprint(groupy_id),
		"BoardItem": fmt.Sprint(board_item_id),
	}

	result, err := controller.InsertGroupyBoardItem(params)
	assert.Nil(t, err)

	num, _ := result.LastInsertId()
	_, err = controller.DeleteGroupyBoardItem(fmt.Sprint(num))
	assert.Nil(t, err)
}

func TestInsertAndDeleteUserLikeBoardItem(t *testing.T) {
	params := QueryParams{
		"User":      fmt.Sprint(user_id),
		"BoardItem": fmt.Sprint(board_item_id),
	}

	result, err := controller.InsertUserLikeBoardItem(params)
	assert.Nil(t, err)

	num, _ := result.LastInsertId()
	_, err = controller.DeleteUserLikeBoardItem(fmt.Sprint(num))
	assert.Nil(t, err)
}

func TestInsertAndDeleteHomeItem(t *testing.T) {
	params := QueryParams{
		"Detail": generateRandomString(30),
	}

	result, err := controller.InsertHomeItem(params)
	assert.Nil(t, err)

	home_item_id, _ = result.LastInsertId()

	result, err = controller.InsertHomeItem(params)
	assert.Nil(t, err)

	num, _ := result.LastInsertId()
	_, err = controller.DeleteHomeItem(fmt.Sprint(num))
	assert.Nil(t, err)
}

func TestUpdateHomeItem(t *testing.T) {
	params := QueryParams{
		"Detail": generateRandomString(20),
	}

	_, err := controller.UpdateHomeItem(fmt.Sprint(home_item_id), params)
	assert.Nil(t, err)
}

func generateRandomString(num int) string {
	rand.Seed(time.Now().UnixNano())

	stringMap := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := ""

	for i := 0; i < num; i++ {
		randomString += string(stringMap[rand.Intn(len(stringMap))])
	}

	return randomString
}
