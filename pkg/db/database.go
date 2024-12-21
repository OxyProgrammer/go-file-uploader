package db

import (
	"go-file-uploader/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

func NewDB(dbFilename string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{
		CreateBatchSize: 1000,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Land{})
	if err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

func (d *DB) Close() error {
	sqlDB, err := d.conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *DB) CreateLand(land *models.Land) error {
	return d.conn.Create(land).Error
}

func (d *DB) CreateLands(lands []*models.Land) error {
	return d.conn.Create(lands).Error
}

func (d *DB) DeleteAllLands() error {
	return d.conn.Unscoped().Delete(&models.Land{}).Error
}
