package main

import (
    "log"
    "time"
    "volta_blockchain/db"
    "volta_blockchain/server"
    "volta_blockchain/utils"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

func main() {
    // Initialize the database
    db.Init()

    // Create admin user
    address := "VTA" + utils.GenerateHex(10)
    passkey := utils.GenerateHash("passkey")
    adminUUID := uuid.New().String()
    initialSupply := decimal.NewFromInt(10000000000000000000)
    admin := db.User{
        Address:   address,
        Passkey:   passkey,
        UUID:      adminUUID,
        Balance:   initialSupply.String(),
        CreatedAt: time.Now().Format(time.RFC3339),
    }
    db.DB.Create(&admin)

    // Create genesis block
    genesisBlock := db.Block{
        BlockID:      "0",
        Timestamp:    time.Now().Format(time.RFC3339),
        Hash:         "genesis_hash",
        PreviousHash: "0",
        Address:      address,
    }
    db.DB.Create(&genesisBlock)

    log.Println("Admin and Genesis Block created successfully")

    // Run the gRPC server
    server.RunServer()
}
