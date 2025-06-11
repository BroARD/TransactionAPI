package db

import (
	"TransactionAPI/internal/models"
	"fmt"
	"log"
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

func InitDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("FATAL: Ошибка загрузки .env файла: %v", err)
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
		log.Fatalf("FATAL: Ошибка подключения к БД: %v", err)
	}

	if err := db.AutoMigrate(&models.Transaction{}, &models.Wallet{}); err != nil {
		log.Fatalf("FATAL: Ошибка миграции: %v", err)
	}

	once.Do(generateWallets)

	return db, nil
}


//Функция для генерации 10 кошельков
func generateWallets() {
	var count int64
	err := db.Model(&models.Wallet{}).Count(&count).Error
	if err != nil {
		log.Print("Could not generate 10 wallets for db")
		return
	}
	if count == 0 {
		var wallet models.Wallet
		for i := 0; i < numberOfWallets; i++{
			wallet = models.Wallet{
				ID: uuid.NewString(),
				Amount: 100,
			}
			log.Println(wallet)
			err := db.Create(&wallet).Error
			if err != nil {
				return
			}
		}
	} else {
		log.Println("Wallets were created earlier")
	}
}