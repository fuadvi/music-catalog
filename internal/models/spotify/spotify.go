package spotify

type SearchResponse struct {
	Limit  int                   `json:"limit"`
	Offset int                   `json:"offset"`
	Total  int                   `json:"total"`
	Items  []SpotifyTrackObjects `json:"items"`
}

type SpotifyTrackObjects struct {
	AlbumType        string   `json:"album_type"`
	AlbumTotalTracks int      `json:"album_total_tracks"`
	AlbumImages      []string `json:"album_images"`
	AlbumName        string   `json:"album_name"`

	ArtistsName []string `json:"artist_name"`
	Explicit    bool     `json:"explicit"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	IsLiked     *bool    `json:"isLiked"`
}
