package wallet

// PriorityType is a transaction priority
type PriorityType uint

// Accepted Values are: 0-3 for:
// default, unimportant, normal, and elevated priority.
const (
	PriorityDefault     PriorityType = 0
	PriorityUnimportant PriorityType = 1
	PriorityNormal      PriorityType = 2
	PriorityElevated    PriorityType = 3
)

// TransferType is a string used in IncomingTransferRequest
// that contains the possible types:
// "all": all the transfers;
// "available": only transfers which are not yet spent;
// "unavailable": only transfers which are already spent.
type TransferType string

const (
	// TransferAll - all the transfers
	TransferAll TransferType = "all"
	// TransferAvailable - only transfers which are not yet spent
	TransferAvailable TransferType = "available"
	// TransferUnavailable - only transfers which are already spent
	TransferUnavailable TransferType = "unavailable"
)

// QueryKeyType is the parameter to send with client.QueryKey()
type QueryKeyType string

const (
	// QueryMnemonicKey is the mnemonic seed
	QueryMnemonicKey QueryKeyType = "mnemonic"
	// QueryViewKey is the private view key
	QueryViewKey QueryKeyType = "view_key"
	// QuerySpendKey is the private spend key
	QuerySpendKey QueryKeyType = "spend_key"
)
