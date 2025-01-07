package tracks

import (
	"context"
	"github.com/fuadvi/music-catalog/internal/models/spotify"
	"github.com/fuadvi/music-catalog/internal/models/trackactivities"
	sporifyRepo "github.com/fuadvi/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *Service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageSize - 1) * pageIndex

	trackDetail, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error searching tracks to spotify")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetail.Tracks.Items))

	for idx, track := range trackDetail.Tracks.Items {
		trackIDs[idx] = track.ID
	}

	trackActivity, err := s.trackActivitiesRepo.GetBullSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get tracks activity from database")
		return nil, err
	}

	return modelTOResponse(trackDetail, trackActivity), nil
}

func modelTOResponse(data *sporifyRepo.SpotifySearchResponse, mapTrackActivity map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObjects, 0)
	for _, item := range data.Tracks.Items {

		artisName := make([]string, len(item.Artists))
		for idx, artis := range item.Artists {
			artisName[idx] = artis.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.Url
		}

		items = append(items, spotify.SpotifyTrackObjects{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImages:      imageUrls,
			AlbumName:        item.Album.Name,
			ArtistsName:      artisName,
			Explicit:         item.Explicit,
			ID:               item.ID,
			Name:             item.Name,
			IsLiked:          mapTrackActivity[item.ID].Isliked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Total:  data.Tracks.Total,
		Items:  items,
	}
}
