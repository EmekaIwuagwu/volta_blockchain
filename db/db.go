package db

import (
    "log"
    "os"
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "github.com/shopspring/decimal"
)

var DB *gorm.DB

// User represents a user in the blockchain
type User struct {
    Address   string          `gorm:"primary_key"`
    Passkey   string
    UUID      string
    Balance   decimal.Decimal `gorm:"type:decimal(36,18);"`
    CreatedAt time.Time
}

// Transaction represents a transaction in the blockchain
type Transaction struct {
    ID                uint `gorm:"primary_key"`
    AddressFrom       string
    AddressTo         string
    Amount            decimal.Decimal `gorm:"type:decimal(36,18);"`
    DateOfTransaction time.Time
    CreatedAt         time.Time
}

// Block represents a block in the blockchain
type Block struct {
    BlockID      string `gorm:"primary_key"`
    Timestamp    time.Time
    Hash         string
    PreviousHash string
    Address      string
}

// Init initializes the database connection and migrates the schema
func Init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " sslmode=disable password=" + dbPassword
    DB, err = gorm.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&User{}, &Transaction{}, &Block{})
}

