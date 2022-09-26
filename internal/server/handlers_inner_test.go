package server

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bejaneps/volume-assignment/internal/server/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandleCalculatePost(t *testing.T) {
	testCases := map[string]struct {
		setup              func() (*mocks.Service, *http.Request)
		expectedStatusCode int
		expectedBody       string
	}{
		"success": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				svc.On(
					"FindStartEndAirports",
					[][]string{
						{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"},
					},
				).Return(
					[]string{"SFO", "EWR"},
					nil,
				)

				req := httptest.NewRequest(
					http.MethodPost,
					`/calculate"`,
					strings.NewReader(`
						[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
					`),
				)

				return svc, req
			},
			expectedStatusCode: 200,
			expectedBody:       `["SFO","EWR"]`,
		},
		"fail-invalid-request-body": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				req := httptest.NewRequest(
					http.MethodPost,
					`/calculate"`,
					strings.NewReader(`invalid request body`),
				)

				return svc, req
			},
			expectedStatusCode: 400,
			expectedBody:       `Bad Request`,
		},
		"fail-service-error": {
			setup: func() (*mocks.Service, *http.Request) {
				svc := new(mocks.Service)

				svc.On(
					"FindStartEndAirports",
					[][]string{
						{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"},
					},
				).Return(
					nil,
					errors.New("random service error"),
				)

				req := httptest.NewRequest(
					http.MethodPost,
					`/calculate"`,
					strings.NewReader(`
						[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
					`),
				)

				return svc, req
			},
			expectedStatusCode: 500,
			expectedBody:       `Internal Server Error`,
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			svc, req := tt.setup()
			rr := httptest.NewRecorder()

			handler := handleCalculatePost(svc)
			handler(rr, req)

			body, err := io.ReadAll(rr.Result().Body)
			assert.NoError(t, err)

			body = bytes.TrimSuffix(body, []byte("\n"))
			assert.Equal(t, tt.expectedStatusCode, rr.Result().StatusCode)
			assert.Equal(t, tt.expectedBody, string(body))
		})
	}
}
