package db

import (
	"Infotecs/internal/entity"
	"Infotecs/pkg/logging"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	once sync.Once
	numberOfWallets = 10
)

func InitDB(logger *logging.Logger) (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("FATAL: Ошибка загрузки .env файла: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")
	
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("FATAL: Ошибка подключения к БД: %v", err)
	}

	if err := db.AutoMigrate(&entity.Transaction{}, &entity.Wallet{}); err != nil {
		logger.Fatalf("FATAL: Ошибка миграции: %v", err)
	}

	once.Do(func() {
        generateWallets(logger)
    })

	return db, nil
}


//Функция для генерации 10 кошельков
func generateWallets(logger *logging.Logger) {
	var count int64
	err := db.Model(&entity.Wallet{}).Count(&count).Error
	if err != nil {
		logger.Print("Could not generate 10 wallets for db")
		return
	}
	if count == 0 {
		var wallet entity.Wallet
		for i := 0; i < numberOfWallets; i++{
			wallet = entity.Wallet{
				ID: uuid.NewString(),
				Amount: 100,
			}
			logger.Info(wallet)
			err := db.Create(&wallet).Error
			if err != nil {
				return
			}
		}
	} else {
		logger.Info("Wallets were created earlier")
	}
}