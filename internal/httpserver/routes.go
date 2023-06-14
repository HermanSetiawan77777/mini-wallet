package httpserver

import (
	"herman-technical-julo/internal/app"
	authhttp "herman-technical-julo/internal/auth/http"
	"herman-technical-julo/internal/httpserver/middleware/authenticator"
	muxauth "herman-technical-julo/internal/httpserver/middleware/authenticator/mux"
	"herman-technical-julo/internal/httpserver/response"
	transactionhttp "herman-technical-julo/internal/transaction/http"
	wallethttp "herman-technical-julo/internal/wallet/http"
	"net/http"

	"github.com/gorilla/mux"
)

func buildRoutes(appContainer *app.Application) http.Handler {
	root := mux.NewRouter()

	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.WithData(w, http.StatusOK, "Server is running", "success")
	}).Methods(http.MethodGet)

	root.HandleFunc("/api/v1/init", authhttp.HandleGetToken(appContainer.Services.AuthService)).Methods(http.MethodPost)

	authRouter := root.NewRoute().Subrouter()
	authRouter.Use(muxauth.AuthenticateRequest(authenticator.NewJULOWebAuthenticator(appContainer.Services.AuthService)))
	authRouter.HandleFunc("/test/authenticate", func(w http.ResponseWriter, r *http.Request) {
		response.WithData(w, http.StatusOK, "authenticated", "Authenticated")
	}).Methods(http.MethodGet)

	authRouter.HandleFunc("/api/v1/wallet", wallethttp.HandleEnableWallet(appContainer.Services.WalletService)).Methods(http.MethodPost)
	authRouter.HandleFunc("/api/v1/wallet", wallethttp.HandleDisabledWallet(appContainer.Services.WalletService)).Methods(http.MethodPatch)
	authRouter.HandleFunc("/api/v1/wallet", wallethttp.HandleViewWallet(appContainer.Services.WalletService)).Methods(http.MethodGet)

	authRouter.HandleFunc("/api/v1/wallet/transactions", transactionhttp.HandleViewTransaction(appContainer.Services.TransactionService)).Methods(http.MethodGet)
	authRouter.HandleFunc("/api/v1/wallet/deposits", transactionhttp.HandleDepositedTransaction(appContainer.Services.TransactionService)).Methods(http.MethodPost)
	authRouter.HandleFunc("/api/v1/wallet/withdrawals", transactionhttp.HandleWithdrawalTransaction(appContainer.Services.TransactionService)).Methods(http.MethodPost)

	return root
}
