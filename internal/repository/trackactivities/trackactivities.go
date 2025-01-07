package trackactivities

import (
	"context"
	"github.com/fuadvi/music-catalog/internal/models/trackactivities"
)

func (r *Repository) Create(ctx context.Context, model trackactivities.TrackActivity) error {
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *Repository) Update(ctx context.Context, model trackactivities.TrackActivity) error {
	return r.db.WithContext(ctx).Save(&model).Error
}

func (r *Repository) Get(ctx context.Context, userID uint, spotifyID string) (*trackactivities.TrackActivity, error) {

	activity := trackactivities.TrackActivity{}
	res := r.db.WithContext(ctx).Where("user_id = ?", userID).Where("spotify_id = ?", spotifyID).First(&activity)

	if res.Error != nil {
		return nil, res.Error
	}

	return &activity, nil
}

func (r *Repository) GetBullSpotifyIDs(ctx context.Context, userID uint, spotifyIDs []string) (map[string]trackactivities.TrackActivity, error) {

	var activities []trackactivities.TrackActivity
	res := r.db.WithContext(ctx).Where("user_id = ?", userID).Where("spotify_id IN ?", spotifyIDs).First(&activities)

	if res.Error != nil {
		return nil, res.Error
	}

	result := make(map[string]trackactivities.TrackActivity)
	for _, activity := range activities {
		result[activity.SpotifyID] = activity
	}

	return result, nil
}
