package spotify

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fuadvi/music-catalog/internal/configs"
	"github.com/fuadvi/music-catalog/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func Test_Outbound_Search(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHttpClient := httpclient.NewMockHTTPClient(mockCtrl)
	next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7"
	type args struct {
		query  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *SpotifySearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			want: &SpotifySearchResponse{
				Tracks: SpotifyTracks{
					Href:     "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=id-ID,id;q%3D0.9,en-US;q%3D0.8,en;q%3D0.7",
					Limit:    10,
					Next:     &next,
					Offset:   0,
					Previous: nil,
					Total:    897,
					Items: []SpotifyTrackObjects{
						{
							Album: SpotifyAlbumObjects{
								AlbumType:   "album",
								TotalTracks: 22,
								Images: []SpotifyAlbumImage{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
									},
								},
								Name: "Bohemian Rhapsody (The Original Soundtrack)",
							},
							Artists: []SpotifyArtistObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
							ID:       "3z8h0TU7ReDPLIbEnYhWZb",
							Name:     "Bohemian Rhapsody",
						},
						{
							Album: SpotifyAlbumObjects{
								AlbumType:   "compilation",
								TotalTracks: 17,
								Images: []SpotifyAlbumImage{
									{
										Url: "https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263",
									},
									{
										Url: "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263",
									},
								},
								Name: "Greatest Hits (Remastered)",
							},
							Artists: []SpotifyArtistObject{
								{
									Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
									Name: "Queen",
								},
							},
							Explicit: false,
							Href:     "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
							ID:       "2OBofMJx94NryV2SK8p8Zf",
							Name:     "Bohemian Rhapsody - Remastered 2011",
						},
					},
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := "https://api.spotify.com/v1/search"
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)
				req.Header.Add("Authorization", "Bearer accessToken")

				mockHttpClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(searchResponse)),
				}, nil)
			},
		},
		{
			name: "error",
			args: args{
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				params := url.Values{}
				params.Set("q", args.query)
				params.Set("type", "track")
				params.Set("limit", strconv.Itoa(args.limit))
				params.Set("offset", strconv.Itoa(args.offset))

				basePath := "https://api.spotify.com/v1/search"
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				assert.NoError(t, err)
				req.Header.Add("Authorization", "Bearer accessToken")

				mockHttpClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: 500,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`internal server error`)),
				}, err)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			o := &Outbound{
				cfg:         &configs.Config{},
				client:      mockHttpClient,
				AccessToken: "accessToken",
				TokenType:   "Bearer",
				ExpiredAt:   time.Now().Add(1 * time.Hour),
			}
			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
