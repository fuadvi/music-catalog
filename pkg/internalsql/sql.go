package internalsql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(dataSource string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	return db, nil
}
