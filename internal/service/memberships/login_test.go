package memberships

import (
	"fmt"
	"github.com/fuadvi/music-catalog/internal/configs"
	"github.com/fuadvi/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func TestService_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).
					Return(&memberships.User{
						Model: gorm.Model{
							ID: 1,
						},
						Email:    "test@gmail.com",
						Username: "testusername",
						Password: "$2a$10$qOTw9.MDmIfi60PZThHCMec5g4XWI0CskDKR4bbsIOAEhBs3cQvmG",
					}, nil)
			},
		},
		{
			name: "error when login",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).
					Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &Service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJwt: "abc",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.NotEmpty(t, got)
			}

			if tt.wantErr {
				fmt.Printf("tess masuk error rt= %+v", got)
				assert.Empty(t, got)
			}
		})
	}
}
