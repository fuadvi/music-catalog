package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/fuadvi/music-catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)
	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name: "failed to bind json",
			mockFn: func() {
			},
			expectedStatusCode: 400,
		},
		{
			name: "failed to create user",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}).Return(errors.New("username and email exist "))
			},
			expectedStatusCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				Service: mockSvc,
			}
			h.RegisterRoute()

			w := httptest.NewRecorder()

			endpoint := "/memberships/sign-up"
			var body *bytes.Buffer
			if tt.name == "failed to bind json" {
				body = bytes.NewBuffer([]byte("invalid json"))
			} else {
				model := memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				}
				val, err := json.Marshal(model)
				assert.NoError(t, err)
				body = bytes.NewBuffer(val)
			}
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.expectedStatusCode)
		})
	}
}
