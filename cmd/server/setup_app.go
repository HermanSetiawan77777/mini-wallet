package main

import (
	"herman-technical-julo/config"
	"herman-technical-julo/internal/app"
	"herman-technical-julo/internal/auth"
	jwtservice "herman-technical-julo/internal/token/jwt"
	"herman-technical-julo/internal/transaction"
	transactionsql "herman-technical-julo/internal/transaction/sql"
	"herman-technical-julo/internal/wallet"
	walletsql "herman-technical-julo/internal/wallet/sql"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func setupAppContainer(dbs *databases) *app.Application {
	services := setupServices(dbs)

	return &app.Application{
		Services: services,
	}
}

func setupServices(dbs *databases) *app.Services {
	walletService := setupWalletService(dbs.julodb)
	authService := setupAuthService(dbs.julodb, walletService)
	transactionService := setupTransactionService(dbs.julodb)
	return &app.Services{
		WalletService:      walletService,
		AuthService:        authService,
		TransactionService: transactionService,
	}
}

func setupWalletService(julodb *gorm.DB) wallet.WalletIService {
	repo := walletsql.NewWalletSQLRepository(julodb)

	walletService := wallet.NewWalletService(repo)
	return walletService
}

func setupAuthService(julodb *gorm.DB, walletService wallet.WalletIService) auth.AuthIService {
	tokener := jwtservice.NewJWTService[*auth.Session](jwt.SigningMethodHS256, config.JwtSecret())
	authService := auth.NewAuthService(walletService, tokener)
	return authService
}

func setupTransactionService(julodb *gorm.DB) transaction.TransactionWalletIService {
	repo := transactionsql.NewTransactionWalletSQLRepository(julodb)

	transactionService := transaction.NewTransactionWalletService(repo, setupWalletService(julodb))
	return transactionService
}
