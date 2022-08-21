package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"server/cmd/converter"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	url           string = "http://localhost:4000"
	user_id       string
	groupy_id     string
	board_item_id string
)

func TestCreateUserHandler(t *testing.T) {
	user := converter.User{
		Name:  generateRandomString(25),
		Email: fmt.Sprintf("%s@gmail.com", generateRandomString(40)),
	}

	data, err := json.Marshal(user)
	assert.Nil(t, err)

	res, err := http.Post(url+"/user/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_user converter.User
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &res_user)

	assert.Equal(t, user.Name, res_user.Name)
	assert.NotEqual(t, "", res_user.ID)

	user_id = res_user.ID
}

func TestUpdateUserHandler(t *testing.T) {
	user := converter.User{
		ID:       user_id,
		Birthday: "2000-01-01",
		Gender:   "men",
	}

	data, err := json.Marshal(user)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url+"/user/update", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err := client.Do(req)
	assert.Nil(t, err)

	var res_user converter.User
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_user)
	assert.Nil(t, err)

	assert.Equal(t, user.Birthday, res_user.Birthday)
}

func TestCreateGroupyHandler(t *testing.T) {
	groupy := converter.Groupy{
		Name:      generateRandomString(10),
		Introduce: generateRandomString(100),
	}

	data, err := json.Marshal(groupy)
	assert.Nil(t, err)

	res, err := http.Post(url+"/groupy/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_groupy converter.Groupy
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_groupy)
	assert.Nil(t, err)

	assert.Equal(t, groupy.Name, res_groupy.Name)

	groupy_id = fmt.Sprint(res_groupy.ID)
}

func TestUpdateGroupyHandler(t *testing.T) {
	groupy := converter.Groupy{
		ID:        groupy_id,
		Is_opened: "FALSE",
	}

	data, err := json.Marshal(groupy)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url+"/groupy/update", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err := client.Do(req)
	assert.Nil(t, err)

	var res_groupy converter.Groupy
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_groupy)
	assert.Nil(t, err)

	assert.Equal(t, groupy.Is_opened, res_groupy.Is_opened)
}

func TestDeleteGroupyHandler(t *testing.T) {
	addGroupy := converter.Groupy{
		Name:      generateRandomString(10),
		Introduce: generateRandomString(50),
	}

	data, err := json.Marshal(addGroupy)
	assert.Nil(t, err)

	res, err := http.Post(url+"/groupy/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_groupy1 converter.Groupy
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_groupy1)
	assert.Nil(t, err)

	groupy := converter.Groupy{
		ID: res_groupy1.ID,
	}

	data, err = json.Marshal(groupy)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url+"/groupy/delete", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err = client.Do(req)
	assert.Nil(t, err)

	var res_groupy converter.Groupy
	b, _ = io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_groupy)
	assert.Nil(t, err)

	assert.Equal(t, groupy.ID, res_groupy.ID)
}

func TestUserInGroupyHandler(t *testing.T) {
	uig := converter.UserInGroupy{
		User:   user_id,
		Groupy: groupy_id,
	}

	data, err := json.Marshal(uig)
	assert.Nil(t, err)

	res, err := http.Post(url+"/uig/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_uig converter.UserInGroupy
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_uig)
	assert.Nil(t, err)

	assert.Equal(t, uig.User, res_uig.User)

	duig := converter.UserInGroupy{
		ID: res_uig.ID,
	}

	data, err = json.Marshal(duig)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url+"/uig/delete", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err = client.Do(req)
	assert.Nil(t, err)

	var dres_uig converter.UserInGroupy
	b, _ = io.ReadAll(res.Body)
	err = json.Unmarshal(b, &dres_uig)
	assert.Nil(t, err)

	assert.Equal(t, res_uig.ID, dres_uig.ID)
}

func TestCreateBoardItemHandler(t *testing.T) {
	board_item := converter.BoardItem{
		Category: "1",
		Type:     "1",
		Title:    generateRandomString(6),
		Detail:   generateRandomString(30),
	}

	data, err := json.Marshal(board_item)
	assert.Nil(t, err)

	res, err := http.Post(url+"/boarditem/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_board_item converter.BoardItem
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_board_item)
	assert.Nil(t, err)

	assert.Equal(t, board_item.Type, res_board_item.Type)
	assert.Equal(t, board_item.Title, res_board_item.Title)

	board_item_id = res_board_item.ID
}

func TestUpdateBoardItemHandler(t *testing.T) {
	board_item := converter.BoardItem{
		ID:        board_item_id,
		Detail:    generateRandomString(20),
		Is_opened: "FALSE",
	}

	data, err := json.Marshal(board_item)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url+"/boarditem/update", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err := client.Do(req)
	assert.Nil(t, err)

	var res_board_item converter.BoardItem
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_board_item)
	assert.Nil(t, err)

	assert.Equal(t, board_item.Detail, res_board_item.Detail)
	assert.Equal(t, board_item.Is_opened, res_board_item.Is_opened)
}

func TestDeleteBoardItemHandler(t *testing.T) {
	board_item := converter.BoardItem{
		Type:     "1",
		Category: "1",
		Title:    generateRandomString(8),
		Detail:   generateRandomString(30),
	}

	data, err := json.Marshal(board_item)
	assert.Nil(t, err)

	res, err := http.Post(url+"/boarditem/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_board_item converter.BoardItem
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_board_item)
	assert.Nil(t, err)

	dboard_item := converter.BoardItem{
		ID: res_board_item.ID,
	}

	data, err = json.Marshal(dboard_item)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url+"/boarditem/delete", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err = client.Do(req)
	assert.Nil(t, err)

	var dres_board_item converter.BoardItem
	b, _ = io.ReadAll(res.Body)
	err = json.Unmarshal(b, &dres_board_item)
	assert.Nil(t, err)

	assert.Equal(t, res_board_item.ID, dres_board_item.ID)
}

func TestGroupyBoardItemHandler(t *testing.T) {
	groupy_board_item := converter.GroupyBoardItem{
		Groupy:    groupy_id,
		BoardItem: board_item_id,
	}

	data, err := json.Marshal(groupy_board_item)
	assert.Nil(t, err)

	res, err := http.Post(url+"/gbi/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_groupy_board_item converter.GroupyBoardItem
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_groupy_board_item)
	assert.Nil(t, err)

	assert.Equal(t, groupy_board_item.Groupy, res_groupy_board_item.Groupy)
	assert.Equal(t, groupy_board_item.BoardItem, res_groupy_board_item.BoardItem)

	groupy_board_item = converter.GroupyBoardItem{
		ID: res_groupy_board_item.ID,
	}

	data, err = json.Marshal(groupy_board_item)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url+"/gbi/delete", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err = client.Do(req)
	assert.Nil(t, err)

	var dres_groupy_board_item converter.GroupyBoardItem
	b, _ = io.ReadAll(res.Body)
	err = json.Unmarshal(b, &dres_groupy_board_item)
	assert.Nil(t, err)

	assert.Equal(t, dres_groupy_board_item.ID, res_groupy_board_item.ID)
}

func TestUserLikeBoardItemHandler(t *testing.T) {
	ulbi := converter.UserLikeBoardItem{
		User:      user_id,
		BoardItem: board_item_id,
	}

	data, err := json.Marshal(ulbi)
	assert.Nil(t, err)

	res, err := http.Post(url+"/ulbi/create", "application/json", bytes.NewBuffer(data))
	assert.Nil(t, err)

	var res_ulbi converter.UserLikeBoardItem
	b, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(b, &res_ulbi)
	assert.Nil(t, err)

	assert.Equal(t, ulbi.User, res_ulbi.User)
	assert.Equal(t, ulbi.BoardItem, res_ulbi.BoardItem)

	dulbi := converter.UserLikeBoardItem{
		ID: res_ulbi.ID,
	}

	data, err = json.Marshal(dulbi)
	assert.Nil(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url+"/ulbi/delete", bytes.NewBuffer(data))
	assert.Nil(t, err)

	res, err = client.Do(req)
	assert.Nil(t, err)

	var dres_ulbi converter.UserLikeBoardItem
	b, _ = io.ReadAll(res.Body)
	err = json.Unmarshal(b, &dres_ulbi)
	assert.Nil(t, err)

	assert.Equal(t, res_ulbi.ID, dres_ulbi.ID)
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
