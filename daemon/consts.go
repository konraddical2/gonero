package daemon

// Daemon default ports
const (
	MainnetPort  uint = 18081
	TestnetPort  uint = 28081
	StagenetPort uint = 38081
)

// RPCStatus is status code returned from the monerod RPC
type RPCStatus string

const (
	// RPCStatusOk CORE_RPC_STATUS_OK   "OK"
	RPCStatusOk RPCStatus = "OK"
	// RPCStatusBusy CORE_RPC_STATUS_BUSY   "BUSY"
	RPCStatusBusy RPCStatus = "BUSY"
	// RPCStatusNotMining CORE_RPC_STATUS_NOT_MINING "NOT MINING"
	RPCStatusNotMining RPCStatus = "NOT_MINING"
	// RPCStatusPaymentRequired CORE_RPC_STATUS_PAYMENT_REQUIRED "PAYMENT REQUIRED"
	RPCStatusPaymentRequired RPCStatus = "PAYMENT REQUIRED"
)
