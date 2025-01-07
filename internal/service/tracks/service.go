package tracks

import (
	"context"
	"github.com/fuadvi/music-catalog/internal/models/trackactivities"
	"github.com/fuadvi/music-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type trackActivitiesRepository interface {
	Create(ctx context.Context, model trackactivities.TrackActivity) error
	Update(ctx context.Context, model trackactivities.TrackActivity) error
	Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error)
	GetBullSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error)
}

type Service struct {
	spotifyOutbound     spotifyOutbound
	trackActivitiesRepo trackActivitiesRepository
}

func NewService(spotifyOutbound spotifyOutbound, trackActivitiesRepo trackActivitiesRepository) *Service {
	return &Service{
		spotifyOutbound:     spotifyOutbound,
		trackActivitiesRepo: trackActivitiesRepo,
	}
}
