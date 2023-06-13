package wallet

import "herman-technical-julo/internal/errors"

var ErrWalletIdNil = errors.NewDefaultValidationError("Wallet id cannot be empty", "Wallet id cannot be empty")
var ErrCustomerXidNil = errors.NewDefaultValidationError("Customer xid cannot be empty", "customer xid cannot be empty")
