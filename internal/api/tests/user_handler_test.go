package api_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SilverLuhtoja/TNVisual/internal/api"
	"github.com/SilverLuhtoja/TNVisual/internal/test_utils"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestCreateUserHandler(t *testing.T) {
	t.Run("Throws decoding error when request has no body", func(t *testing.T) {
		config := test_utils.CreateTestConfig()
		server := httptest.NewServer(http.HandlerFunc(config.CreateUserHandler))

		resp, err := http.Post(server.URL, "", bytes.NewBuffer([]byte("")))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		assert.Equal(t, resp.StatusCode, 500)
		assert.Equal(t, `{"error":"createUserHandler - couldn't decode parameters"}`, getBodyString(t, resp))
	})

	t.Run("Throws duplicate error when username already present", func(t *testing.T) {
		config := test_utils.CreateTestConfig()
		server := httptest.NewServer(http.HandlerFunc(config.CreateUserHandler))

		// set up expectations
		params := api.CreateUserRequest{
			Username: "s2",
			Password: "p2",
		}

		// ACT
		bodyReq, _ := json.Marshal(params)
		resp, err := http.Post(server.URL, "", bytes.NewBuffer(bodyReq))
		http.Post(server.URL, "", bytes.NewBuffer(bodyReq))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		// ASSERT
		assert.Equal(t, 500, resp.StatusCode)
		assert.Contains(t, strings.Split(getBodyString(t, resp), " "), "duplicate")

		test_utils.ClearTable("users")
	})
}

func getBodyString(t *testing.T, resp *http.Response) string {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return string(b)
}
