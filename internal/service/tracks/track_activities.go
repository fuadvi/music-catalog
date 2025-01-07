package tracks

import (
	"context"
	"errors"
	"fmt"
	"github.com/fuadvi/music-catalog/internal/models/trackactivities"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *Service) UpsertTrackActivities(ctx context.Context, userID uint, request trackactivities.TrackActivityRequest) error {
	activity, err := s.trackActivitiesRepo.Get(ctx, userID, request.SpotifyID)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Msg("failed to retrieve track activity")
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || activity == nil {
		activity := trackactivities.TrackActivity{
			UserID:    userID,
			SpotifyID: request.SpotifyID,
			Isliked:   request.Isliked,
			CreatedBy: fmt.Sprintf("%d", userID),
			UpdatedBy: fmt.Sprintf("%d", userID),
		}
		err = s.trackActivitiesRepo.Create(ctx, activity)
		if err != nil {
			log.Error().Err(err).Msg("error create record to database")
			return err
		}
		return nil
	}

	activity.Isliked = request.Isliked
	activity.UpdatedBy = fmt.Sprintf("%d", userID)
	err = s.trackActivitiesRepo.Update(ctx, *activity)
	if err != nil {
		log.Error().Err(err).Msg("error update record to database")
		return err
	}
	return nil
}
