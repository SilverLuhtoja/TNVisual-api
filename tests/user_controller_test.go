package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SilverLuhtoja/TNVisual/internal/api/user"
	"github.com/SilverLuhtoja/TNVisual/internal/api/user/resources"
	"github.com/SilverLuhtoja/TNVisual/tests/test_utils"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestUserControllerIntegration(t *testing.T) {
	userRepo := user.NewUserRepostitory(GetDatabaseQueries())
	userInteractor := user.NewUserInteractor(userRepo)
	controller := *user.NewUserController(*userInteractor)
	server := httptest.NewServer(http.HandlerFunc(controller.Create))

	var USERNAME string = "karu"
	var PASSWORD string = "ott"

	t.Run("Should succeed", func(t *testing.T) {
		// ARRANGE
		requestParams := getCreateUserRequest(USERNAME, PASSWORD)

		// ACT
		bodyReq, _ := json.Marshal(requestParams)
		resp, err := http.Post(server.URL, "", bytes.NewBuffer(bodyReq))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		// ASSERT
		assert.Equal(t, 201, resp.StatusCode)
		assert.Contains(t, test_utils.GetBodyString(t, resp), `User created successfully`)

		ClearTable("users")
	})

	t.Run("Throws decoding error when body is empty", func(t *testing.T) {
		resp, err := http.Post(server.URL, "", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		assert.Equal(t, 400, resp.StatusCode)
		assert.Contains(t, test_utils.GetBodyString(t, resp), `decoding error`)
	})

	t.Run("Throws duplicate error when username already present", func(t *testing.T) {
		// ARRANGE
		requestParams := getCreateUserRequest(USERNAME, PASSWORD)
		InsertData("users", []string{"username", "password"}, []string{USERNAME, PASSWORD})

		// ACT
		bodyReq, _ := json.Marshal(requestParams)
		resp, err := http.Post(server.URL, "", bytes.NewBuffer(bodyReq))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		// ASSERT
		assert.Equal(t, 500, resp.StatusCode)
		assert.Contains(t, test_utils.GetBodyString(t, resp), `duplicate`)

		ClearTable("users")
	})

}

func getCreateUserRequest(USERNAME, PASSWORD string) resources.CreateUserRequest {
	return resources.CreateUserRequest{
		Username: USERNAME,
		Password: PASSWORD,
	}
}
