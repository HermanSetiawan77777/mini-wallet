package http

import (
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/httpserver/response"
	"herman-technical-julo/internal/wallet"
	"net/http"
)

type GetEnabledResponse struct {
	Wallet *wallet.WalletDetail `json:"wallet"`
}

func HandleEnableWallet(walletService wallet.WalletIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentSession, err := auth.GetSessionFromContext(r.Context())
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}

		targetWallet, err := walletService.EnableWallet(r.Context(), &wallet.UpdateWalletParam{
			WalletId: currentSession.WalletId,
		})
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		response.WithData(w, http.StatusOK, &GetEnabledResponse{
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
		response.WithData(w, http.StatusOK, &GetEnabledResponse{
			Wallet: targetWallet,
		}, "Success")
	}
}
