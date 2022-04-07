package test_helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func AssertHttpRequest(t assert.TestingT, tt HttpTestDescription, app *fiber.App) {
	req := httptest.NewRequest(tt.Method, tt.Route, bytes.NewReader(tt.Body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if tt.Authorization != "" {
		req.Header.Set("Authorization", tt.Authorization)
	}

	resp, _ := app.Test(req, 1)
	defer resp.Body.Close()

	assert.Equalf(t, tt.ExpectedCode, resp.StatusCode, tt.Description)

	if tt.ExpectedBody != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		expectBody, _ := json.Marshal(tt.ExpectedBody)
		assert.Equalf(t, string(expectBody), string(body), tt.Description)
	}
}
