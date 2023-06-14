package http

import (
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/httpserver/response"
	"herman-technical-julo/internal/transaction"
	"net/http"
)

type GetViewTransactionResponse struct {
	transaction.ViewTransactionWallet `json:"transaction"`
}

func HandleViewWallet(transactionService transaction.TransactionWalletIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentSession, err := auth.GetSessionFromContext(r.Context())
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}

		targetWallet, err := transactionService.ViewMyTransactionWallet(r.Context(), currentSession.WalletId)
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		response.WithData(w, http.StatusOK, targetWallet, "Success")
	}
}
