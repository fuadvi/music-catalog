package tracks

import (
	"encoding/json"
	"github.com/fuadvi/music-catalog/internal/models/spotify"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := NewMockservice(mockCtrl)
	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       spotify.SearchResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			expectedBody: spotify.SearchResponse{
				Limit:  10,
				Offset: 0,
				Total:  897,
				Items: []spotify.SpotifyTrackObjects{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 22,
						AlbumImages:      []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "3z8h0TU7ReDPLIbEnYhWZb",
						Name:             "Bohemian Rhapsody",
					},
					{
						AlbumType:        "compilation",
						AlbumTotalTracks: 17,
						AlbumImages:      []string{"https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263"},
						AlbumName:        "Greatest Hits (Remastered)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "2OBofMJx94NryV2SK8p8Zf",
						Name:             "Bohemian Rhapsody - Remastered 2011",
					},
				},
			},
			wantErr: false,
			mockFn: func() {
				mockService.EXPECT().Search(gomock.Any(), "bohemian rhapsody", 10, 1).Return(&spotify.SearchResponse{
					Limit:  10,
					Offset: 0,
					Total:  897,
					Items: []spotify.SpotifyTrackObjects{
						{
							AlbumType:        "album",
							AlbumTotalTracks: 22,
							AlbumImages:      []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
							AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
							ArtistsName:      []string{"Queen"},
							Explicit:         false,
							ID:               "3z8h0TU7ReDPLIbEnYhWZb",
							Name:             "Bohemian Rhapsody",
						},
						{
							AlbumType:        "compilation",
							AlbumTotalTracks: 17,
							AlbumImages:      []string{"https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263"},
							AlbumName:        "Greatest Hits (Remastered)",
							ArtistsName:      []string{"Queen"},
							Explicit:         false,
							ID:               "2OBofMJx94NryV2SK8p8Zf",
							Name:             "Bohemian Rhapsody - Remastered 2011",
						},
					},
				}, nil)
			},
		},
		{
			name:               "success",
			expectedStatusCode: 400,
			expectedBody:       spotify.SearchResponse{},
			wantErr:            true,
			mockFn: func() {
				mockService.EXPECT().Search(gomock.Any(), "bohemian rhapsody", 10, 1).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				Service: mockService,
			}
			h.RegisterRoute()

			w := httptest.NewRecorder()

			endpoint := "/tracks/search?query=bohemian+rhapsody&pageSize=10&pageIndex=1"

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.expectedStatusCode)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.SearchResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
				assert.ObjectsAreEqual(tt.expectedBody, response)
				//assert.Equal(t, tt.expectedBody.AccessToken, response.AccessToken)
			}
		})
	}
}
