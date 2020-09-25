package daemon

import (
	"fmt"

	"github.com/gorilla/rpc/v2/json2"
)

// ErrorCode is a monerod RPC error code
type ErrorCode int

const (
	// Copied from https://github.com/monero-project/monero/blob/master/src/rpc/core_rpc_server_error_codes.h

	// ErrWrongParam - CORE_RPC_ERROR_CODE_WRONG_PARAM
	ErrWrongParam ErrorCode = -1
	// ErrTooBigHeight - CORE_RPC_ERROR_CODE_TOO_BIG_HEIGHT
	ErrTooBigHeight ErrorCode = -2
	// ErrTooBigReserveSize - CORE_RPC_ERROR_CODE_TOO_BIG_RESERVE_SIZE
	ErrTooBigReserveSize ErrorCode = -3
	// ErrWrongWalletAddress - CORE_RPC_ERROR_CODE_WRONG_WALLET_ADDRESS
	ErrWrongWalletAddress ErrorCode = -4
	// ErrInternalError - CORE_RPC_ERROR_CODE_INTERNAL_ERROR
	ErrInternalError ErrorCode = -5
	// ErrWrongBlockblob - CORE_RPC_ERROR_CODE_WRONG_BLOCKBLOB
	ErrWrongBlockblob ErrorCode = -6
	// ErrBlockNotAccepted - CORE_RPC_ERROR_CODE_BLOCK_NOT_ACCEPTED
	ErrBlockNotAccepted ErrorCode = -7
	// ErrCoreBust - CORE_RPC_ERROR_CODE_CORE_BUSY
	ErrCoreBust ErrorCode = -1
	// ErrWrongBlockblobSize - CORE_RPC_ERROR_CODE_WRONG_BLOCKBLOB_SIZE
	ErrWrongBlockblobSize ErrorCode = -10
	// ErrUnsupportedRPC - CORE_RPC_ERROR_CODE_UNSUPPORTED_RPC
	ErrUnsupportedRPC ErrorCode = -11
	// ErrMiningToSubaddress - CORE_RPC_ERROR_CODE_MINING_TO_SUBADDRESS
	ErrMiningToSubaddress ErrorCode = -12
	// ErrRegtestRequired - CORE_RPC_ERROR_CODE_REGTEST_REQUIRED
	ErrRegtestRequired ErrorCode = -13
	// ErrPaymentRequired - CORE_RPC_ERROR_CODE_PAYMENT_REQUIRED
	ErrPaymentRequired ErrorCode = -14
	// ErrInvalidClient - CORE_RPC_ERROR_CODE_INVALID_CLIENT
	ErrInvalidClient ErrorCode = -15
	// ErrPaymentTooLow - CORE_RPC_ERROR_CODE_PAYMENT_TOO_LOW
	ErrPaymentTooLow ErrorCode = -16
	// ErrDuplicatePayment CORE_RPC_ERROR_CODE_DUPLICATE_PAYMENT
	ErrDuplicatePayment ErrorCode = -17
	// ErrStalePayment - CORE_RPC_ERROR_CODE_STALE_PAYMENT
	ErrStalePayment ErrorCode = -18
	// ErrCodeRestricted - CORE_RPC_ERROR_CODE_RESTRICTED
	ErrCodeRestricted ErrorCode = -19
)

// RPCError is the error structured returned by the monerod RPC
type RPCError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (re *RPCError) Error() string {
	return fmt.Sprintf("%v: %v", re.Code, re.Message)
}

// GetRPCError checks if an error interface is an rpc error.
func GetRPCError(err error) error {
	if err == nil {
		return err
	}
	gerr, ok := err.(*json2.Error)
	if !ok {
		return err
	}
	return &RPCError{
		Code:    ErrorCode(gerr.Code),
		Message: gerr.Message,
	}
}
