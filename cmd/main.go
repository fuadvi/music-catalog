package main

import (
	"github.com/fuadvi/music-catalog/internal/configs"
	"github.com/fuadvi/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
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

	r := gin.Default()
	r.Run(cfg.Service.Port)
}
