package wallet

import "herman-technical-julo/internal/errors"

var ErrWalletIdNil = errors.NewDefaultValidationError("Wallet id cannot be empty", "Wallet id cannot be empty")
var ErrCustomerXidNil = errors.NewDefaultValidationError("Customer xid cannot be empty", "customer xid cannot be empty")
var ErrStatusZero = errors.NewDefaultValidationError("Status cannot 0", "status cannot 0")
var ErrWalletIdNotExist = errors.NewDefaultValidationError("Wallet id not exist", "Wallet id not exist")
var ErrWalletAlreadyEnabled = errors.NewDefaultValidationError("Failed, Wallet already enabled", "failed, wallet already enabled")
var ErrWalletDeactive = errors.NewDefaultValidationError("Wallet disabled", "wallet disabled")
var ErrAmountZero = errors.NewDefaultValidationError("Transaction amount cannot 0", "transaction amount cannot 0")
var ErrWalletAlreadyDisabled = errors.NewDefaultValidationError("Failed, Wallet already disabled", "failed, wallet already disabled")
