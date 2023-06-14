package transaction

import "herman-technical-julo/internal/errors"

var ErrTransactionWalletDeactive = errors.NewDefaultValidationError("Wallet disabled", "wallet disabled")
var ErrTransactionIdNil = errors.NewDefaultValidationError("Transaction id cannot empty", "transaction id cannot empty")
var ErrWalletIdNil = errors.NewDefaultValidationError("Wallet id cannot empty", "wallet id cannot empty")
var ErrTransactionTypeNil = errors.NewDefaultValidationError("Transaction type id cannot empty", "transaction type id cannot empty")
var ErrTransactionByNil = errors.NewDefaultValidationError("Transaction by cannot empty", "transaction by cannot empty")
var ErrAmountZero = errors.NewDefaultValidationError("Amount cannot 0", "amount cannot 0")
var ErrReferenceIdNil = errors.NewDefaultValidationError("Reference id cannot empty", "reference id cannot empty")
var ErrWalletDeactive = errors.NewDefaultValidationError("Wallet disabled", "wallet disabled")
