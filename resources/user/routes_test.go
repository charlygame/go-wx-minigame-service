package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/charlygame/CatGameService/test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	rg := r.Group("/v1")
	AddUserRoutes(rg)
	return r
}

func createUser(t *testing.T, r *gin.Engine) (int, gin.H) {
	body := `{
		"username": "test",
		"session_key": "test",
		"open_id": "test",
		"score": 0
	}`

	req, _ := http.NewRequest("POST", "/v1/user/", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	return w.Code, got
}

func TestCreateUserShouldReturn200(t *testing.T) {
	test.Init("../../test.env")
	defer test.Clear()

	r := setupRouter()

	code, result := createUser(t, r)

	assert.Equal(t, http.StatusOK, code)
	assert.NotEmpty(t, result["id"])
}

func TestGetUserShouldReturn200(t *testing.T) {
	test.Init("../../test.env")
	defer test.Clear()

	r := setupRouter()
	_, result := createUser(t, r)

	req, _ := http.NewRequest("GET", "/v1/user/"+result["id"].(string), nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "test", got["username"])
}

func TestUpdateUserShouldReturn200(t *testing.T) {
	test.Init("../../test.env")
	defer test.Clear()

	r := setupRouter()
	_, result := createUser(t, r)

	body := `{
		"score": 100
	}`

	req, _ := http.NewRequest("PUT", "/v1/user/"+result["id"].(string), strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", got)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 100.0, got["score"])
}

func TestWXLoginShouldReturn200(t *testing.T) {
	test.Init("../../test.env")
	defer test.Clear()

	r := setupRouter()

	body := `{
		"code": "test"
	}`

	req, _ := http.NewRequest("POST", "/v1/user/wx_login", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", got)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, got["session_key"])
	assert.NotEmpty(t, got["openid"])
	assert.NotEmpty(t, got["unionid"])

}
