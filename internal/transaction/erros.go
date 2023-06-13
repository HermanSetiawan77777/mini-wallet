package transaction

import "herman-technical-julo/internal/errors"

var ErrTransactionWalletDeactive = errors.NewDefaultValidationError("Wallet disabled", "wallet disabled")
