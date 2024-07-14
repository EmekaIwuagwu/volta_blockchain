package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"volta_blockchain/db"
	"volta_blockchain/proto"
	"volta_blockchain/utils"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Server struct{}

func (s *Server) CreateAddress(ctx context.Context, req *proto.CreateAddressRequest) (*proto.CreateAddressResponse, error) {
	address := "VT" + utils.GenerateHex(10)
	passkey := utils.GenerateHash(fmt.Sprintf("%d", rand.Intn(10)))
	userUUID := uuid.New().String()
	user := db.User{
		Address:   address,
		Passkey:   passkey,
		UUID:      userUUID,
		Balance:   decimal.NewFromFloat(0), // Initialize balance with 0
		CreatedAt: time.Now(),
	}
	db.DB.Create(&user)

	return &proto.CreateAddressResponse{
		Message: "New address generated",
		Address: address,
		Balance: user.Balance.String(), // Convert decimal.Decimal to string
		Uuid:    userUUID,
		Passkey: passkey,
	}, nil
}

func (s *Server) SendTokens(ctx context.Context, req *proto.SendTokensRequest) (*proto.SendTokensResponse, error) {
	var fromUser, toUser db.User
	db.DB.Where("address = ?", req.AddressFrom).First(&fromUser)
	db.DB.Where("address = ?", req.AddressTo).First(&toUser)

	if fromUser.Passkey != req.Passkey {
		return nil, fmt.Errorf("invalid passkey")
	}

	fromBalance := fromUser.Balance // Use decimal.Decimal directly
	amount, _ := decimal.NewFromString(req.Amount)

	if fromBalance.LessThan(amount) {
		return nil, fmt.Errorf("insufficient balance")
	}

	toBalance := toUser.Balance // Use decimal.Decimal directly
	fromUser.Balance = fromBalance.Sub(amount) // Perform decimal.Decimal operations directly
	toUser.Balance = toBalance.Add(amount) // Perform decimal.Decimal operations directly

	db.DB.Save(&fromUser)
	db.DB.Save(&toUser)

	transaction := db.Transaction{
		AddressFrom:       req.AddressFrom,
		AddressTo:         req.AddressTo,
		Amount:            amount, // Store decimal.Decimal directly
		DateOfTransaction: time.Now(),
		CreatedAt:         time.Now(),
	}
	db.DB.Create(&transaction)

	return &proto.SendTokensResponse{
		Message:    "Transaction sent successfully",
		AddressFrom: req.AddressFrom,
		AddressTo:   req.AddressTo,
		Amount:      amount.String(), // Convert decimal.Decimal to string
	}, nil
}

func (s *Server) CheckBalance(ctx context.Context, req *proto.CheckBalanceRequest) (*proto.CheckBalanceResponse, error) {
	var user db.User
	db.DB.Where("address = ?", req.Address).First(&user)

	return &proto.CheckBalanceResponse{
		Message: "Balance fetched successfully",
		Address: user.Address,
		Balance: user.Balance.String(), // Convert decimal.Decimal to string
	}, nil
}

func (s *Server) CheckTransactions(ctx context.Context, req *proto.CheckTransactionsRequest) (*proto.CheckTransactionsResponse, error) {
	var transactions []db.Transaction
	db.DB.Where("address_from = ? OR address_to = ?", req.Address, req.Address).Find(&transactions)

	var txnList []*proto.Transaction
	for _, txn := range transactions {
		txnList = append(txnList, &proto.Transaction{
			AddressFrom:       txn.AddressFrom,
			AddressTo:         txn.AddressTo,
			Amount:            txn.Amount.String(), // Convert decimal.Decimal to string
			DateOfTransaction: txn.DateOfTransaction.Format(time.RFC3339), // Convert time.Time to string
			CreatedAt:         txn.CreatedAt.Format(time.RFC3339), // Convert time.Time to string
		})
	}

	return &proto.CheckTransactionsResponse{
		Message:      "Transactions fetched successfully",
		Transactions: txnList,
	}, nil
}
