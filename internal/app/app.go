package app

import (
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/wallet"
)

type Application struct {
	Services *Services
}

type Services struct {
	AuthService   auth.AuthIService
	WalletService wallet.WalletIService
}
