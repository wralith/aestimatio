package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/wralith/aestimatio/server/api-gateway/mock"
	"github.com/wralith/aestimatio/server/api-gateway/pkg/vld"
)

var h *AuthHandler
var e *echo.Echo

func TestMain(m *testing.M) {
	mockSvc := mock.NewMockAuthClient()
	h = NewAuthHandler(mockSvc)
	e = echo.New()
	e.Validator = &vld.Validator{Validator: validator.New()}

	m.Run()
}

func TestAuthHandler_Login(t *testing.T) {
	tests := []struct {
		name string
		body string
		code int
	}{
		{
			name: "Happy",
			body: `{"email":"test@test.com","password":"1234567"}`,
			code: http.StatusOK,
		},
		{
			name: "Invalid Values",
			body: `{"email":"test@test.com","password":"123"}`,
			code: http.StatusBadRequest,
		},
		{
			name: "Invalid Body",
			body: `{"username":"test""password":"1234567"}`,
			code: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/tasks/list", strings.NewReader(test.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.Login(c)

			require.NoError(t, err)
			require.Equal(t, test.code, rec.Code)

		})
	}
}

func TestAuthHandler_Register(t *testing.T) {
	tests := []struct {
		name string
		body string
		code int
	}{
		{
			name: "Happy",
			body: `{"email":"test@test.com","username":"test","password":"1234567"}`,
			code: http.StatusCreated,
		},
		{
			name: "Invalid Values",
			body: `{"email":"test@test.com","username":"test","password":"123"}`,
			code: http.StatusBadRequest,
		},
		{
			name: "Invalid Body",
			body: `{"email":"test@test.com","username":"test""password":"1234567"}`,
			code: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(test.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.Register(c)

			require.NoError(t, err)
			require.Equal(t, test.code, rec.Code)

		})
	}
}

func TestAuthHandler_Validate(t *testing.T) {
	tests := []struct {
		name   string
		header string
		code   int
	}{
		{
			name:   "Happy",
			header: "Bearer dummjwt", // Test passes if it can parse whether it is valid jwt or not
			code:   http.StatusOK,
		},
		{
			name:   "Invalid Header Type",
			header: "NotBearer xyz",
			code:   http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/restricted", nil)
			req.Header.Set("Authorization", test.header)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.Validate(c)

			require.NoError(t, err)
			require.Equal(t, test.code, rec.Code)
		})
	}
}
