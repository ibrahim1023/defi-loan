package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/loan module sentinel errors
var (
	ErrInvalidSigner  = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 2, "wrong loan state")
	ErrDeadline       = sdkerrors.Register(ModuleName, 3, "deadline")
)
