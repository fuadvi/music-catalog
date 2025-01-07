package main

import (
	"github.com/fuadvi/music-catalog/internal/configs"
	membershipHANDLER "github.com/fuadvi/music-catalog/internal/handler/memberships"
	tracksHandler "github.com/fuadvi/music-catalog/internal/handler/tracks"
	"github.com/fuadvi/music-catalog/internal/models/trackactivities"
	trackactivitiesRepo "github.com/fuadvi/music-catalog/internal/repository/trackactivities"

	"github.com/fuadvi/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/fuadvi/music-catalog/internal/repository/memberships"
	"github.com/fuadvi/music-catalog/internal/repository/spotify"
	membershipSVC "github.com/fuadvi/music-catalog/internal/service/memberships"
	"github.com/fuadvi/music-catalog/internal/service/tracks"
	"github.com/fuadvi/music-catalog/pkg/httpclient"
	"github.com/fuadvi/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config")
	}
	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi databases")
	}
	db.AutoMigrate(&memberships.User{})
	db.AutoMigrate(&trackactivities.TrackActivity{})

	httpClient := httpclient.NewClient(&http.Client{})

	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpClient)

	membershipRepo := membershipsRepo.NewRepository(db)
	trackactivityRepo := trackactivitiesRepo.NewRepository(db)

	membershipSvc := membershipSVC.NewService(cfg, membershipRepo)
	spotifySvc := tracks.NewService(spotifyOutbound, trackactivityRepo)

	membershipHandler := membershipHANDLER.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	trackHandler := tracksHandler.NewHandler(r, spotifySvc)
	trackHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
