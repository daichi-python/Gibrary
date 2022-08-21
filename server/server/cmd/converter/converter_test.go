package converter

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"server/cmd/database"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var controller *database.DBController

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

	controller = database.NewDBController(dbEngine)
}

func TestMain(m *testing.M) {
	setup()
	log.Println("SUCCESS: START TEST")
	code := m.Run()
	os.Exit(code)
}

func TestToInsertQueryParams(t *testing.T) {
	user := User{}
	user.Email = fmt.Sprintf("%s@gmail.com", generateRandomString(21))
	user.Name = generateRandomString(10)

	params := ToInsertQueryParams(user)
	assert.Equal(t, user.Email, params["email"])
	assert.Equal(t, user.Name, params["name"])
}

func TestUpdateQueryParams(t *testing.T) {
	user := User{}
	user.Name = generateRandomString(12)
	user.Birthday = "2000-06-30"
	user.ID = "1"

	id, params := ToUpdateQueryParams(user)
	assert.Equal(t, user.ID, id)
	assert.Equal(t, user.Birthday, params["birthday"])
}

func TestToSelectQueryParams(t *testing.T) {
	user := User{}
	user.ID = "1"

	params := ToSelectQueryParams(user)
	assert.Equal(t, params["id"], "1")
}

func TestUserRowToSelect(t *testing.T) {
	params := database.QueryParams{
		"id": "1",
	}
	row := controller.SelectUser(params)
	user, err := UserRowToStruct(row)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, "1")
	assert.Contains(t, user.Email, "@gmail.com")
}

func generateRandomString(num int) string {
	stringMap := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := ""

	for i := 0; i < num; i++ {
		randomString += string(stringMap[rand.Intn(len(stringMap))])
	}

	return randomString
}
