package trackactivities

import "gorm.io/gorm"

type (
	TrackActivity struct {
		gorm.Model
		UserID    uint   `gorm:"not null"`
		SpotifyID string `gorm:"not null"`
		Isliked   *bool
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	TrackActivityRequest struct {
		SpotifyID string `json:"spotify_id"`
		Isliked   *bool  `json:"isliked"`
	}
)
