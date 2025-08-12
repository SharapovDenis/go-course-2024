package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAd(t *testing.T) {
	client := getTestClient()

	// Сперва создаём 1 пользователя
	_, err := client.createUser("first", "world@example.com")
	assert.NoError(t, err)

	// Затем создаём 2 пользователя
	secondUsr, err := client.createUser("second", "world@example.com")
	assert.NoError(t, err)

	response, err := client.createAd(secondUsr.Data.ID, "hello", "world")
	assert.NoError(t, err)
	assert.Zero(t, response.Data.ID)
	assert.Equal(t, response.Data.Title, "hello")
	assert.Equal(t, response.Data.Text, "world")
	assert.Equal(t, response.Data.AuthorID, int64(1))
	assert.False(t, response.Data.Published)
}

func TestChangeAdStatus(t *testing.T) {
	client := getTestClient()

	// Сперва создаём пользователя
	usr, err := client.createUser("first", "world@example.com")
	assert.NoError(t, err)

	response, err := client.createAd(usr.Data.ID, "hello", "world")
	assert.NoError(t, err)

	response, err = client.changeAdStatus(usr.Data.ID, response.Data.ID, true)
	assert.NoError(t, err)
	assert.True(t, response.Data.Published)

	response, err = client.changeAdStatus(usr.Data.ID, response.Data.ID, false)
	assert.NoError(t, err)
	assert.False(t, response.Data.Published)

	response, err = client.changeAdStatus(usr.Data.ID, response.Data.ID, false)
	assert.NoError(t, err)
	assert.False(t, response.Data.Published)
}

func TestUpdateAd(t *testing.T) {
	client := getTestClient()

	// Сперва создаём пользователя
	usr, err := client.createUser("first", "world@example.com")
	assert.NoError(t, err)

	response, err := client.createAd(usr.Data.ID, "hello", "world")
	assert.NoError(t, err)

	response, err = client.updateAd(usr.Data.ID, response.Data.ID, "привет", "мир")
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Title, "привет")
	assert.Equal(t, response.Data.Text, "мир")
}

func TestListAds(t *testing.T) {
	client := getTestClient()

	// Сперва создаём пользователя
	usr, err := client.createUser("first", "world@example.com")
	assert.NoError(t, err)

	response, err := client.createAd(usr.Data.ID, "hello", "world")
	assert.NoError(t, err)

	publishedAd, err := client.changeAdStatus(usr.Data.ID, response.Data.ID, true)
	assert.NoError(t, err)

	_, err = client.createAd(usr.Data.ID, "best cat", "not for sale")
	assert.NoError(t, err)

	// Hint: По умолчанию выводятся только published объявления
	ads, err := client.listAds()
	assert.NoError(t, err)
	assert.Len(t, ads.Data, 1)
	assert.Equal(t, ads.Data[0].ID, publishedAd.Data.ID)
	assert.Equal(t, ads.Data[0].Title, publishedAd.Data.Title)
	assert.Equal(t, ads.Data[0].Text, publishedAd.Data.Text)
	assert.Equal(t, ads.Data[0].AuthorID, publishedAd.Data.AuthorID)
	assert.True(t, ads.Data[0].Published)
}

func TestCreateUser(t *testing.T) {
	client := getTestClient()

	response, err := client.createUser("Alice", "alice@example.com")
	assert.NoError(t, err)
	assert.Zero(t, response.Data.ID)
	assert.Equal(t, response.Data.Name, "Alice")
	assert.Equal(t, response.Data.Email, "alice@example.com")
}

func TestUpdateUser(t *testing.T) {
	client := getTestClient()

	// Сначала создаем пользователя
	createResp, err := client.createUser("Bob", "bob@example.com")
	assert.NoError(t, err)

	// Обновляем пользователя
	updateResp, err := client.updateUser(createResp.Data.ID, "Robert", "robert@example.com")
	assert.NoError(t, err)
	assert.Equal(t, updateResp.Data.ID, createResp.Data.ID)
	assert.Equal(t, updateResp.Data.Name, "Robert")
	assert.Equal(t, updateResp.Data.Email, "robert@example.com")
}

func TestListUsers(t *testing.T) {
	client := getTestClient()

	// Создаем нескольких пользователей
	_, err := client.createUser("User1", "user1@example.com")
	assert.NoError(t, err)
	_, err = client.createUser("User2", "user2@example.com")
	assert.NoError(t, err)

	usersResp, err := client.listUsers()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(usersResp.Data), 2)
}
