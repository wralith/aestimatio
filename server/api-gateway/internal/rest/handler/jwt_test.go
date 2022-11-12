package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/wralith/aestimatio/server/api-gateway/internal/rest/request"
)

func TestAuthHandler_Authenticate(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		code    int
		isValid bool
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
		{
			name:   "No Bearer",
			header: "",
			code:   http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e.GET("/", func(c echo.Context) error {
				return c.JSON(http.StatusOK, "ok")
			})

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", test.header)
			rec := httptest.NewRecorder()

			e.Use(h.Authenticate)
			e.ServeHTTP(rec, req)

			require.Equal(t, test.code, rec.Code)
		})
	}
}

func TestAuthHandler_CallAuthClient(t *testing.T) {
	// Stub does not test if jwt valid or not, for test purpose it returns value passed by context

	req := request.ValidateRequest{
		JWT: "dummy valid",
	}

	ctx := context.WithValue(context.Background(), "isValid", true)
	got := h.callAuthClient(ctx, req)
	require.True(t, got)

	ctx = context.WithValue(context.Background(), "isValid", false)
	got = h.callAuthClient(ctx, req)
	require.False(t, got)
}
