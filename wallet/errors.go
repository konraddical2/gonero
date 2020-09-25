package wallet

import (
	"fmt"

	"github.com/gorilla/rpc/v2/json2"
)

// ErrorCode is a monero-wallet-rpc error code
type ErrorCode int

const (
	// ErrUnkownError - WALLET_RPC_ERROR_CODE_UNKNOWN_ERROR
	ErrUnkownError ErrorCode = -1
	// ErrWrongAddress - WALLET_RPC_ERROR_CODE_WRONG_ADDRESS
	ErrWrongAddress ErrorCode = -2
	// ErrDaemonIsBusy - WALLET_RPC_ERROR_CODE_DAEMON_IS_BUSY
	ErrDaemonIsBusy ErrorCode = -3
	// ErrGenericTransferError - WALLET_RPC_ERROR_CODE_GENERIC_TRANSFER_ERROR
	ErrGenericTransferError ErrorCode = -4
	// ErrWrongPaymentID - WALLET_RPC_ERROR_CODE_WRONG_PAYMENT_ID
	ErrWrongPaymentID ErrorCode = -5
	// ErrTransferType - WALLET_RPC_ERROR_CODE_TRANSFER_TYPE
	ErrTransferType ErrorCode = -6
	// ErrDenied - WALLET_RPC_ERROR_CODE_DENIED
	ErrDenied ErrorCode = -7
	// ErrWrongTxid - WALLET_RPC_ERROR_CODE_WRONG_TXID
	ErrWrongTxid ErrorCode = -8
	// ErrWrongSignature - WALLET_RPC_ERROR_CODE_WRONG_SIGNATURE
	ErrWrongSignature ErrorCode = -9
	// ErrWrongKeyImage - WALLET_RPC_ERROR_CODE_WRONG_KEY_IMAGE
	ErrWrongKeyImage ErrorCode = -10
	// ErrWrongURI - WALLET_RPC_ERROR_CODE_WRONG_URI
	ErrWrongURI ErrorCode = -11
	// ErrWrongIndex - WALLET_RPC_ERROR_CODE_WRONG_INDEX
	ErrWrongIndex ErrorCode = -12
	// ErrNotOpen - WALLET_RPC_ERROR_CODE_NOT_OPEN
	ErrNotOpen ErrorCode = -13
)

// RPCError is the error structured returned by the monero-wallet-rpc
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
