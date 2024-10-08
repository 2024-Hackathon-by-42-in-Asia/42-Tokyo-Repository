package accessdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"carbon_driver_app/internal/loadconfig"
	"log"
)

// User model
type User struct {
	UID      string `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Icon     string
	// DeliveryScores is a one-to-many relationship with DeliveryScore
	DeliveryScores []DeliveryScore `gorm:"foreignKey:UserID;references:UID"`
}

// DeliveryScore model
type DeliveryScore struct {
	UID       string `gorm:"primaryKey"`
	UserID    string `gorm:"not null"` // UserID is the foreign key
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserID;references:UID"`
	CreatedAt string
	UpdatedAt string
	Score     int
}

func InitDB(config *loadconfig.Config) (*gorm.DB, error) {
	dsn := "host=" + config.DBHost + " user=" + config.DBUser + " password=" + config.DBPassword + " dbname=" + config.DBName + " port=" + config.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &DeliveryScore{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	return err
}
