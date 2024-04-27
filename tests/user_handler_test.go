package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SilverLuhtoja/TNVisual/internal/api"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	config := CreateTestConfig()
	server := httptest.NewServer(http.HandlerFunc(config.CreateUserHandler))

	var USERNAME string = "karu"
	var PASSWORD string = "ott"

	t.Run("Should succeed", func(t *testing.T) {
		// ARRANGE
		var params api.CreateUserRequest = api.CreateUserRequest{
			Username: USERNAME,
			Password: PASSWORD,
		}

		// ACT
		bodyReq, _ := json.Marshal(params)
		resp, err := http.Post(server.URL, "", bytes.NewBuffer(bodyReq))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		// ASSERT
		assert.Equal(t, 201, resp.StatusCode)
		assert.Equal(t, `"User created successfully"`, getBodyString(t, resp))

		ClearTable("users")
	})

	t.Run("Throws decoding error when bad request body", func(t *testing.T) {
		resp, err := http.Post(server.URL, "", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		assert.Equal(t, 400, resp.StatusCode)
		assert.Equal(t, `{"error":"createUserHandler - couldn't decode parameters"}`, getBodyString(t, resp))
	})

	t.Run("Throws duplicate error when username already present", func(t *testing.T) {
		// ARRANGE
		var params api.CreateUserRequest = api.CreateUserRequest{
			Username: USERNAME,
			Password: PASSWORD,
		}
		InsertData("users", []string{"username", "password"}, []string{USERNAME, PASSWORD})

		// ACT
		bodyReq, _ := json.Marshal(params)
		resp, err := http.Post(server.URL, "", bytes.NewBuffer(bodyReq))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		// ASSERT
		assert.Equal(t, 500, resp.StatusCode)
		assert.Contains(t, strings.Split(getBodyString(t, resp), " "), "duplicate")

		ClearTable("users")
	})

}

func getBodyString(t *testing.T, resp *http.Response) string {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return string(b)
}
