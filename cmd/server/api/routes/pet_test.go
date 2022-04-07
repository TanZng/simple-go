package routes

import (
	"net/http"
	"simple-go/cmd/server/api"
	apiTest "simple-go/permissions/test_helpers"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func Test_petRoutes_RegisterRoutes(t *testing.T) {

	tests := []apiTest.HttpTestDescription{
		{
			Description:   "Returns error if required parameter Name is not set (name field)",
			Route:         "/pet",
			Method:        "POST",
			ExpectedCode:  http.StatusBadRequest,
			Authorization: "Bearer token",
			Body: []byte(`{
				"kind": "Perro"
			}`),
			ExpectedBody: []api.ErrorResponse{
				{
					FailedField: "Pet.Name",
					Tag:         "required",
					Value:       "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Description, func(t *testing.T) {
			app := fiber.New()
			apiInstance := api.New(app)

			u := petRoutes{}
			u.RegisterRoutes(&apiInstance)

			apiTest.AssertHttpRequest(t, tt, app)
		})
	}
}
