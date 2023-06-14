package walletlog

import "herman-technical-julo/internal/errors"

var ErrWalletIdNil = errors.NewDefaultValidationError("Wallet id cannot empty", "wallet id cannot empty")
var ErrStatusZero = errors.NewDefaultValidationError("Status id cannot 0", "status id cannot 0")
