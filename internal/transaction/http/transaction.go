package http

import (
	"herman-technical-julo/internal/auth"
	"herman-technical-julo/internal/errors"
	"herman-technical-julo/internal/httpserver/request"
	"herman-technical-julo/internal/httpserver/response"
	"herman-technical-julo/internal/transaction"
	"net/http"
)

type GetViewTransactionResponse struct {
	transaction.ViewTransactionWallet `json:"transaction"`
}

type DepositedReponse struct {
	Transaction *transaction.ViewTransactionDepositWallet `json:"deposit"`
}
type WithdrawalReponse struct {
	Transaction *transaction.ViewTransactionWithdrawalWallet `json:"withdrawal"`
}

type TransactionRequest struct {
	Amount      int    `json:"amount"`
	ReferenceId string `json:"reference_id"`
}

func HandleViewTransaction(transactionService transaction.TransactionWalletIService) http.HandlerFunc {
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

func HandleDepositedTransaction(serviceTransaction transaction.TransactionWalletIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := &TransactionRequest{}
		err := request.DecodeBody(r, &payload)
		if err != nil {
			if err == errors.ErrEmptyPayload {
				response.WithError(w, err, "fail")
				return
			}
			response.WithError(w, errors.ErrUnprocessablePayload, "fail")
			return
		}
		currentSession, err := auth.GetSessionFromContext(r.Context())
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		data, err := serviceTransaction.DepositTransaction(r.Context(), &transaction.CreateTransactionParam{
			WalletId:      currentSession.WalletId,
			Amount:        payload.Amount,
			ReferenceId:   payload.ReferenceId,
			TransactionBy: currentSession.CustomerXid,
		})
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		response.WithData(w, http.StatusOK, &DepositedReponse{
			Transaction: data,
		}, "success")
	}
}

func HandleWithdrawalTransaction(serviceTransaction transaction.TransactionWalletIService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := &TransactionRequest{}
		err := request.DecodeBody(r, &payload)
		if err != nil {
			if err == errors.ErrEmptyPayload {
				response.WithError(w, err, "fail")
				return
			}
			response.WithError(w, errors.ErrUnprocessablePayload, "fail")
			return
		}
		currentSession, err := auth.GetSessionFromContext(r.Context())
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		data, err := serviceTransaction.WithdrawalTransaction(r.Context(), &transaction.CreateTransactionParam{
			WalletId:      currentSession.WalletId,
			Amount:        payload.Amount,
			ReferenceId:   payload.ReferenceId,
			TransactionBy: currentSession.CustomerXid,
		})
		if err != nil {
			response.WithError(w, err, "fail")
			return
		}
		response.WithData(w, http.StatusOK, &WithdrawalReponse{
			Transaction: data,
		}, "success")
	}
}
