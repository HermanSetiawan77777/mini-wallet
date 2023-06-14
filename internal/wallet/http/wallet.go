package http

import (
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/httpserver/response"
	"herman-technical-julo/internal/wallet"
	"net/http"
)

type GetEnabledWalletResponse struct {
	Wallet *wallet.EnabledWalletDetail `json:"wallet"`
}

type GetDisableWalletResponse struct {
	Wallet *wallet.DisableWalletDetail `json:"wallet"`
}

func HandleEnableWallet(walletService wallet.WalletIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentSession, err := auth.GetSessionFromContext(r.Context())
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}

		targetWallet, err := walletService.EnableWallet(r.Context(), &wallet.EnableDisableWalletParam{
			WalletId: currentSession.WalletId,
		})
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		response.WithData(w, http.StatusOK, &GetEnabledWalletResponse{
			Wallet: targetWallet,
		}, "Success")
	}
}

func HandleViewWallet(walletService wallet.WalletIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentSession, err := auth.GetSessionFromContext(r.Context())
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}

		targetWallet, err := walletService.GetByLinkedWallet(r.Context(), currentSession.WalletId)
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		response.WithData(w, http.StatusOK, &GetEnabledWalletResponse{
			Wallet: targetWallet,
		}, "Success")
	}
}

func HandleDisabledWallet(walletService wallet.WalletIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentSession, err := auth.GetSessionFromContext(r.Context())
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}

		targetWallet, err := walletService.DisableWallet(r.Context(), &wallet.EnableDisableWalletParam{
			WalletId: currentSession.WalletId,
		})
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}

		payload := &wallet.DisableWalletDetail{
			WalletId:    targetWallet.WalletId,
			CustomerXid: targetWallet.CustomerXid,
			Status:      targetWallet.Status,
			DateLog:     targetWallet.DateLog,
			Balance:     targetWallet.Balance,
		}
		response.WithData(w, http.StatusOK, &GetDisableWalletResponse{
			Wallet: payload,
		}, "Success")
	}
}
