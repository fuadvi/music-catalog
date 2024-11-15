package main

import (
	"github.com/fuadvi/music-catalog/internal/configs"
	membershipHANDLER "github.com/fuadvi/music-catalog/internal/handler/memberships"
	"github.com/fuadvi/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/fuadvi/music-catalog/internal/repository/memberships"
	membershipSVC "github.com/fuadvi/music-catalog/internal/service/memberships"
	"github.com/fuadvi/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"log"
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

	membershipRepo := membershipsRepo.NewRepository(db)
	membershipSvc := membershipSVC.NewService(cfg, membershipRepo)
	membershipHandler := membershipHANDLER.NewHandler(r, membershipSvc)

	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
