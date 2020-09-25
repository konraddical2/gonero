package wallet

import "github.com/konraddical2/gonero"

//#####################
// JSON RPC Structs
//#####################

// SetDaemonRequest is a struct for SetDaemon() requests
type SetDaemonRequest struct {
	// (Optional;Default: "") The URL of the daemon to connect to.
	Address string `json:"address,omitempty"`
	// (Optional;Default: false) If false, some RPC wallet methods will be disabled.
	Trusted bool `json:"trusted,omitempty"`
	// (Optional;Default: autodetect;Accepts: disabled, enabled, autodetect) Specifies whether the Daemon uses SSL encryption.
	SslSupport string `json:"ssl_support,omitempty"`
	// (Optional) The file path location of the SSL key.
	SslPrivateKeyPath string `json:"ssl_private_key_path,omitempty"`
	// (Optional) The file path location of the SSL certificate.
	SslCertificatePath string `json:"ssl_certificate_path,omitempty"`
	// (Optional) The file path location of the certificate authority file.
	SslCaFile string `json:"ssl_ca_file,omitempty"`
	// (Optional) The SHA1 fingerprints accepted by the SSL certificate.
	SslAllowedFingerprints []string `json:"ssl_allowed_fingerprints,omitempty"`
	// (Optional;Default: false) If false, the certificate must be signed by a trusted certificate authority.
	SslAllowAnyCert bool `json:"ssl_allow_any_cert,omitempty"`
}

// SetDaemonResponse is a struct for SetDaemon() responses
type SetDaemonResponse struct {
	// None
}

// GetBalanceRequest is a struct for GetBalance() requests
type GetBalanceRequest struct {
	// Return balance for this account.
	AccountIndex uint64 `json:"account_index"`
	// (Optional) Return balance detail for those subaddresses.
	AddressIndices []uint64 `json:"address_indices,omitempty"`
}

// GetBalanceResponse is a struct for GetBalance() responses
type GetBalanceResponse struct {
	// The total balance of the current monero-wallet-rpc in session.
	Balance gonero.AtomicXMR `json:"balance"`
	// Unlocked funds are those funds that are sufficiently deep enough in the Monero blockchain to be considered safe to spend.
	UnlockedBalance gonero.AtomicXMR `json:"unlocked_balance"`
	// True if importing multisig data is needed for returning a correct balance.
	MultisigImportNeeded bool `json:"multisig_import_needed"`
	// Balance information for each subaddress in an account.
	PerSubaddress []SubaddressBalance `json:"per_subaddress"`
}

// SubaddressBalance is a struct returned by GetBalance()
// that contains balance information for a subaddress in an account
type SubaddressBalance struct {
	//  Index of the subaddress in the account.
	AddressIndex uint64 `json:"address_index"`
	//  Address at this index. Base58 representation of the public keys.
	Address string `json:"address"`
	//  Balance for the subaddress (locked or unlocked).
	Balance gonero.AtomicXMR `json:"balance"`
	//  Unlocked balance for the subaddress.
	UnlockedBalance gonero.AtomicXMR `json:"unlocked_balance"`
	//  Label for the subaddress.
	Label string `json:"label"`
	//  Number of unspent outputs available for the subaddress.
	NumUnspentOutputs uint64 `json:"num_unspent_outputs"`
}

// GetAddressRequest is a struct for GetAddress() requests
type GetAddressRequest struct {
	// Return subaddresses for this account.
	AccountIndex uint64 `json:"account_index"`
	// (Optional) List of subaddresses to return from an account.
	AddressIndex []uint64 `json:"address_index,omitempty"`
}

// GetAddressResponse is a struct for GetAddress() responses
type GetAddressResponse struct {
	// The 95-character hex address string of the monero-wallet-rpc in session.
	Address string `json:"address"`
	// array of addresses informations
	Addresses []Address `json:"addresses"`
}

// Address is a struct for GetAddress()
type Address struct {
	// the 95 character hex (sub)address string.
	Address string `json:"address"`
	// label of the (sub)address
	Label string `json:"label"`
	// index of the subaddress
	AddressIndex uint64 `json:"address_index"`
	// states if the (sub)address has already received funds
	Used bool `json:"used"`
}

// GetAddressIndexRequest is a struct for GetAddressIndex() requests
type GetAddressIndexRequest struct {
	// (sub)address to look for.
	Address string `json:"address"`
}

// GetAddressIndexResponse is a struct for GetAddressIndex() responses
type GetAddressIndexResponse struct {
	// Subaddress account indices
	Index SubaddressIndex `json:"index"`
}

// SubaddressIndex is a struct that contains subaddress index information
type SubaddressIndex struct {
	// Major account index
	Major uint64 `json:"major"`
	// Minor account index
	Minor uint64 `json:"minor"`
}

// CreateAddressRequest is a struct for CreateAddress() requests
type CreateAddressRequest struct {
	// Create a new address for this account.
	AccountIndex uint64 `json:"account_index"`
	// (Optional) Label for the new address.
	Label string `json:"label,omitempty"`
}

// CreateAddressResponse is a struct for CreateAddress() responses
type CreateAddressResponse struct {
	// Newly created address. Base58 representation of the public keys.
	Address string `json:"address"`
	// Index of the new address under the input account.
	AddressIndex uint64 `json:"address_index"`
}

// LabelAddressRequest is a struct for LabelAddress() requests
type LabelAddressRequest struct {
	// JSON Object containing the major & minor address index
	Index SubaddressIndex `json:"index"`
	// Account index for the subaddress.
	Major uint64 `json:"major"`
	// Index of the subaddress in the account.
	Minor uint64 `json:"minor"`
	// Label for the address.
	Label string `json:"label"`
}

// LabelAddressResponse is a struct for LabelAddress() responses
type LabelAddressResponse struct {
	// None
}

// ValidateAddressRequest is a struct for ValidateAddress() requests
type ValidateAddressRequest struct {
	// The address to validate.
	Address string `json:"address"`
	// If true, consider addresses belonging to any of the three Monero
	// networks (mainnet, stagenet, and testnet) valid. Otherwise,
	// only consider an address valid if it belongs to the network on
	// which the rpc-wallet's current daemon is running (Defaults to false).
	AnyNetType bool `json:"any_net_type,omitempty"`
	// If true, consider OpenAlias-formatted addresses valid (Defaults to false).
	AllowOpenalias bool `json:"allow_openalias,omitempty"`
}

// ValidateAddressResponse is a struct for ValidateAddress() responses
type ValidateAddressResponse struct {
	// True if the input address is a valid Monero address.
	Valid bool `json:"valid"`
	// True if the given address is an integrated address.
	Integrated bool `json:"integrated"`
	// True if the given address is a subaddress
	Subaddress bool `json:"subaddress"`
	// Specifies which of the three Monero networks (mainnet, stagenet, and testnet) the address belongs to.
	Nettype string `json:"nettype"`
	// True if the address is OpenAlias-formatted.
	OpenaliasAddress bool `json:"openalias_address"`
}

// GetAccountsRequest is a struct for GetAccounts() requests
type GetAccountsRequest struct {
	// (Optional) Tag for filtering accounts.
	Tag string `json:"tag,omitempty"`
}

// GetAccountsResponse is a struct for GetAccounts() responses
type GetAccountsResponse struct {
	// array of subaddress account information
	SubaddressAccounts []Subaddress `json:"subaddress_accounts,omitempty"`
	// Index of the account.
	AccountIndex uint64 `json:"account_index"`
	// Balance of the account (locked or unlocked).
	Balance gonero.AtomicXMR `json:"balance"`
	// Base64 representation of the first subaddress in the account.
	BaseAddress string `json:"base_address"`
	// (Optional) Label of the account.
	Label string `json:"label,omitempty"`
	// (Optional) Tag for filtering accounts.
	Tag string `json:"tag,omitempty"`
	// Unlocked balance for the account.
	UnlockedBalance gonero.AtomicXMR `json:"unlocked_balance"`
	// Total balance of the selected accounts (locked or unlocked).
	TotalBalance gonero.AtomicXMR `json:"total_balance"`
	// Total unlocked balance of the selected accounts.
	TotalUnlockedBalance gonero.AtomicXMR `json:"total_unlocked_balance"`
}

// Subaddress is a struct for GetAccounts()
type Subaddress struct {
	//  Index of the account.
	AccountIndex uint64 `json:"account_index"`
	//  Balance of the account (locked or unlocked).
	Balance gonero.AtomicXMR `json:"balance"`
	//  Base64 representation of the first subaddress in the account.
	BaseAddress string `json:"base_address"`
	//  (Optional) Label of the account.
	Label string `json:"label,omitempty"`
	//  (Optional) Tag for filtering accounts.
	Tag string `json:"tag,omitempty"`
	//  Unlocked balance for the account.
	UnlockedBalance gonero.AtomicXMR `json:"unlocked_balance"`
}

// CreateAccountRequest is a struct for CreateAccount() requests
type CreateAccountRequest struct {
	// (Optional) Label for the account.
	Label string `json:"label,omitempty"`
}

// CreateAccountResponse is a struct for CreateAccount() responses
type CreateAccountResponse struct {
	// Index of the new account.
	AccountIndex uint64 `json:"account_index"`
	// Address for this account. Base58 representation of the public keys.
	Address string `json:"address"`
}

// LabelAccountRequest is a struct for LabelAccount() requests
type LabelAccountRequest struct {
	// Apply label to account at this index.
	AccountIndex uint64 `json:"account_index"`
	// Label for the account.
	Label string `json:"label"`
}

// LabelAccountResponse is a struct for LabelAccount() responses
type LabelAccountResponse struct {
	// None
}

// GetAccountTagsRequest is a struct for GetAccountTags() requests
type GetAccountTagsRequest struct {
	// None
}

// GetAccountTagsResponse is a struct for GetAccountTags() responses
type GetAccountTagsResponse struct {
	// array of account tag information
	AccountTags []AccountTag `json:"account_tags"`
}

// AccountTag is a struct for GetAccountTags()
type AccountTag struct {
	//  Filter tag.
	Tag string `json:"tag"`
	//  Label for the tag.
	Label string `json:"label"`
	// List of tagged account indices.
	Accounts []uint64 `json:"accounts"`
}

// TagAccountsRequest is a struct for TagAccounts() requests
type TagAccountsRequest struct {
	// Tag for the accounts.
	Tag string `json:"tag"`
	// Tag this list of accounts.
	Accounts []uint64 `json:"accounts"`
}

// TagAccountsResponse is a struct for TagAccounts() responses
type TagAccountsResponse struct {
	// None
}

// UntagAccountsRequest is a struct for UntagAccounts() requests
type UntagAccountsRequest struct {
	// Remove tag from this list of accounts.
	Accounts []uint64 `json:"accounts"`
}

// UntagAccountsResponse is a struct for UntagAccounts() responses
type UntagAccountsResponse struct {
	// None
}

// SetAccountTagDescriptionRequest is a struct for SetAccountTagDescription() requests
type SetAccountTagDescriptionRequest struct {
	// Set a description for this tag.
	Tag string `json:"tag"`
	// Description for the tag.
	Description string `json:"description"`
}

// SetAccountTagDescriptionResponse is a struct for SetAccountTagDescription() responses
type SetAccountTagDescriptionResponse struct {
	// None
}

// GetHeightRequest is a struct for GetHeight() requests
type GetHeightRequest struct {
	// None
}

// GetHeightResponse is a struct for GetHeight() responses
type GetHeightResponse struct {
	// The current monero-wallet-rpc's blockchain height.
	// If the wallet has been offline for a long time,
	// it may need to catch up with the daemon.
	Height uint64 `json:"height"`
}

// TransferRequest is a struct for Transfer() requests
type TransferRequest struct {
	// array of destinations to receive XMR
	Destinations []Destination `json:"destinations"`
	// Amount to send to each destination, in atomic units.
	Amount gonero.AtomicXMR `json:"amount"`
	// Destination public address.
	Address string `json:"address"`
	// (Optional) Transfer from this account index. (Defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`
	// (Optional) Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
	// Set a priority for the transaction. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority PriorityType `json:"priority"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Number of outputs to mix in the transaction (this output + N decoys from the blockchain).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// (Optional) Return the transaction key after sending.
	GetTxKey bool `json:"get_tx_key,omitempty"`
	// (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// Return the transaction as hex string after sending (Defaults to false)
	GetTxHex bool `json:"get_tx_hex"`
	// Return the metadata needed to relay the transaction. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata"`
}

// TransferResponse is a struct for Transfer() responses
type TransferResponse struct {
	// Amount transferred for the transaction.
	Amount gonero.AtomicXMR `json:"amount"`
	// Integer value of the fee charged for the txn.
	Fee uint64 `json:"fee"`
	// Set of multisig transactions in the process of being signed (empty for non multisig).
	MultisigTxset string `json:"multisig_txset"`
	// Raw transaction represented as hex string, if get_tx_hex is true.
	TxBlob string `json:"tx_blob"`
	// String for the publically searchable transaction hash.
	TxHash string `json:"tx_hash"`
	// String for the transaction key if get_tx_key is true, otherwise, blank string.
	TxKey string `json:"tx_key"`
	// Set of transaction metadata needed to relay this transfer later, if get_tx_metadata is true.
	TxMetadata string `json:"tx_metadata"`
	// signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// Destination is a struct for Transfer() and  TransferSplit()
type Destination struct {
	//  Amount to send to each destination, in atomic units.
	Amount gonero.AtomicXMR `json:"amount"`
	//  Destination public address.
	Address string `json:"address"`
}

// TransferSplitRequest is a struct for TransferSplit() requests
type TransferSplitRequest struct {
	// amount;unsigned int;Amount to send to each destination, in atomic units. address - string;Destination public address.
	Destinations []Destination `json:"destinations"`
	// Amount to send to each destination, in atomic units.
	Amount gonero.AtomicXMR `json:"amount"`
	// Destination public address.
	Address string `json:"address"`
	// (Optional) Transfer from this account index. (Defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`
	// (Optional) Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// Set a priority for the transactions. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority PriorityType `json:"priority"`
	// (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// Return the transactions as hex string after sending
	GetTxHex bool `json:"get_tx_hex"`
	// True to use the new transaction construction algorithm, defaults to false.
	NewAlgorithm bool `json:"new_algorithm"`
	// Return list of transaction metadata needed to relay the transfer later.
	GetTxMetadata bool `json:"get_tx_metadata"`
}

// TransferSplitResponse is a struct for TransferSplit() responses
type TransferSplitResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []int64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []int64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// multisig).
	MultisigTxset string `json:"multisig_txset"`
	// signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// SignTransferRequest is a struct for SignTransfer() requests
type SignTransferRequest struct {
	// string. Set of unsigned tx returned by "transfer" or "transfer_split" methods.
	UnsignedTxset string `json:"unsigned_txset"`
	// (Optional) If true, return the raw transaction data. (Defaults to false)
	ExportRaw bool `json:"export_raw,omitempty"`
}

// SignTransferResponse is a struct for SignTransfer() responses
type SignTransferResponse struct {
	// string. Set of signed tx to be used for submitting transfer.
	SignedTxset string `json:"signed_txset"`
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The tx raw data of every transaction.
	TxRawList []string `json:"tx_raw_list"`
}

// SubmitTransferRequest is a struct for SubmitTransfer() requests
type SubmitTransferRequest struct {
	// Set of signed tx returned by "sign_transfer"
	TxDataHex string `json:"tx_data_hex"`
}

// SubmitTransferResponse is a struct for SubmitTransfer() responses
type SubmitTransferResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
}

// SweepDustRequest is a struct for SweepDust() requests
type SweepDustRequest struct {
	// (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// (Optional) Return the transactions as hex string after sending. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`
	// (Optional) Return list of transaction metadata needed to relay the transfer later. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

// SweepDustResponse is a struct for SweepDust() responses
type SweepDustResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []int64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []int64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// multisig).
	MultisigTxset string `json:"multisig_txset"`
	// signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// SweepAllRequest is a struct for SweepAll() requests
type SweepAllRequest struct {
	// Destination public address.
	Address string `json:"address"`
	// Sweep transactions from this account.
	AccountIndex uint64 `json:"account_index"`
	// (Optional) Sweep from this set of subaddresses in the account.
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
	// (Optional) Priority for sending the sweep transfer, partially determines fee.
	Priority PriorityType `json:"priority,omitempty"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// (Optional) Include outputs below this amount.
	BelowAmount gonero.AtomicXMR `json:"below_amount,omitempty"`
	// (Optional) If true, do not relay this sweep transfer. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// (Optional) return the transactions as hex encoded string. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`
	// (Optional) return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

// SweepAllResponse is a struct for SweepAll() responses
type SweepAllResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []int64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []int64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// multisig).
	MultisigTxset string `json:"multisig_txset"`
	// signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// SweepSingleRequest is a struct for SweepSingle() requests
type SweepSingleRequest struct {
	// Destination public address.
	Address string `json:"address"`
	// Sweep transactions from this account.
	AccountIndex uint64 `json:"account_index"`
	// (Optional) Sweep from this set of subaddresses in the account.
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
	// (Optional) Priority for sending the sweep transfer, partially determines fee.
	Priority PriorityType `json:"priority,omitempty"`
	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`
	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// Key image of specific output to sweep.
	KeyImage string `json:"key_image"`
	// (Optional) Include outputs below this amount.
	BelowAmount gonero.AtomicXMR `json:"below_amount,omitempty"`
	// (Optional) If true, do not relay this sweep transfer. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// (Optional) return the transactions as hex encoded string. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`
	// (Optional) return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

// SweepSingleResponse is a struct for SweepSingle() responses
type SweepSingleResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
	// The amount transferred for every transaction.
	AmountList []int64 `json:"amount_list"`
	// The amount of fees paid for every transaction.
	FeeList []int64 `json:"fee_list"`
	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`
	// multisig).
	MultisigTxset string `json:"multisig_txset"`
	// signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// RelayTxRequest is a struct for RelayTx() requests
type RelayTxRequest struct {
	// transaction metadata returned from a transfer method with get_tx_metadata set to true.
	Hex string `json:"hex"`
}

// RelayTxResponse is a struct for RelayTx() responses
type RelayTxResponse struct {
	// String for the publically searchable transaction hash.
	TxHash string `json:"tx_hash"`
}

// StoreRequest is a struct for Store() requests
type StoreRequest struct {
	// None
}

// StoreResponse is a struct for Store() responses
type StoreResponse struct {
	// None
}

// GetPaymentsRequest is a struct for GetPayments() requests
type GetPaymentsRequest struct {
	// Payment ID used to find the payments (16 characters hex).
	PaymentID string `json:"payment_id"`
}

// GetPaymentsResponse is a struct for GetPayments() responses
type GetPaymentsResponse struct {
	// list of Payments
	Payments []Payment `json:"payments"`
}

// Payment is a struct that contains information for a payment
type Payment struct {
	// Payment ID matching the input parameter.
	PaymentID string `json:"payment_id"`
	// Transaction hash used as the transaction ID.
	TxHash string `json:"tx_hash"`
	// Amount for this payment.
	Amount gonero.AtomicXMR `json:"amount"`
	// Height of the block that first confirmed this payment.
	BlockHeight uint64 `json:"block_height"`
	// Time (in block height) until this payment is safe to spend.
	UnlockTime uint64 `json:"unlock_time"`
	// subaddress index
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`
	// Account index for the subaddress.
	Major uint64 `json:"major"`
	// Index of the subaddress in the account.
	Minor uint64 `json:"minor"`
	// Address receiving the payment; Base58 representation of the public keys.
	Address string `json:"address"`
}

// GetBulkPaymentsRequest is a struct for GetBulkPayments() requests
type GetBulkPaymentsRequest struct {
	// Payment IDs used to find the payments (16 characters hex).
	PaymentIDs []string `json:"payment_ids"`
	// The block height at which to start looking for payments.
	MinBlockHeight uint64 `json:"min_block_height"`
}

// GetBulkPaymentsResponse is a struct for GetBulkPayments() responses
type GetBulkPaymentsResponse struct {
	// list of payments.
	Payments []Payment `json:"payments"`
}

// IncomingTransfersRequest is a struct for IncomingTransfers() requests
type IncomingTransfersRequest struct {
	// "all": all the transfers,
	// "available": only transfers which are not yet spent,
	// OR "unavailable": only transfers which are already spent.
	TransferType TransferType `json:"transfer_type"`
	// (Optional) Return transfers for this account. (defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`
	// (Optional) Return transfers sent to these subaddresses.
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
	// (Optional) Enable verbose output, return key image if true.
	Verbose bool `json:"verbose,omitempty"`
}

// IncomingTransfersResponse is a struct for IncomingTransfers() responses
type IncomingTransfersResponse struct {
	// list of transfers
	Transfers []IncomingTransfer `json:"transfers"`
}

// IncomingTransfer is a struct that contains information for a transfer
type IncomingTransfer struct {
	// Amount of this transfer.
	Amount gonero.AtomicXMR `json:"amount"`
	// Mostly internal use, can be ignored by most users.
	GlobalIndex uint64 `json:"global_index"`
	// Key image for the incoming transfer's unspent output (empty unless verbose is true).
	KeyImage string `json:"key_image"`
	// Indicates if this transfer has been spent.
	Spent bool `json:"spent"`
	// Subaddress index for incoming transfer.
	SubaddrIndex uint64 `json:"subaddr_index"`
	// Several incoming transfers may share the same hash if they were in the same transaction.
	TxHash string `json:"tx_hash"`
	// Size of transaction in bytes.
	TxSize uint64 `json:"tx_size"`
}

// QueryKeyRequest is a struct for QueryKey() requests
type QueryKeyRequest struct {
	// Which key to retrieve: "mnemonic" - the mnemonic seed (older wallets do not have one) OR "view_key" - the view key
	KeyType QueryKeyType `json:"key_type"`
}

// QueryKeyResponse is a struct for QueryKey() responses
type QueryKeyResponse struct {
	// The view key will be hex encoded, while the mnemonic will be a string of words.
	Key string `json:"key"`
}

// MakeIntegratedAddressRequest is a struct for MakeIntegratedAddress() requests
type MakeIntegratedAddressRequest struct {
	// (Optional, defaults to primary address) Destination public address.
	StandardAddress string `json:"standard_address,omitempty"`
	// (Optional, defaults to a random ID) 16 characters hex encoded.
	PaymentID string `json:"payment_id,omitempty"`
}

// MakeIntegratedAddressResponse is a struct for MakeIntegratedAddress() responses
type MakeIntegratedAddressResponse struct {
	// string
	IntegratedAddress string `json:"integrated_address"`
	// hex encoded
	PaymentID string `json:"payment_id"`
}

// SplitIntegratedAddressRequest is a struct for SplitIntegratedAddress() requests
type SplitIntegratedAddressRequest struct {
	// string
	IntegratedAddress string `json:"integrated_address"`
}

// SplitIntegratedAddressResponse is a struct for SplitIntegratedAddress() responses
type SplitIntegratedAddressResponse struct {
	// States if the address is a subaddress
	IsSubaddress bool `json:"is_subaddress"`
	// hex encoded
	Payment string `json:"payment"`
	// string
	StandardAddress string `json:"standard_address"`
}

// StopWalletRequest is a struct for StopWallet() requests
type StopWalletRequest struct {
	// None
}

// StopWalletResponse is a struct for StopWallet() responses
type StopWalletResponse struct {
	// None
}

// RescanBlockchainRequest is a struct for RescanBlockchain() requests
type RescanBlockchainRequest struct {
	// None
}

// RescanBlockchainResponse is a struct for RescanBlockchain() responses
type RescanBlockchainResponse struct {
	// None
}

// SetTxNotesRequest is a struct for SetTxNotes() requests
type SetTxNotesRequest struct {
	// transaction ids
	Txids []string `json:"txids"`
	// notes for the transactions
	Notes []string `json:"notes"`
}

// SetTxNotesResponse is a struct for SetTxNotes() responses
type SetTxNotesResponse struct {
	// None
}

// GetTxNotesRequest is a struct for GetTxNotes() requests
type GetTxNotesRequest struct {
	// transaction ids
	Txids []string `json:"txids"`
}

// GetTxNotesResponse is a struct for GetTxNotes() responses
type GetTxNotesResponse struct {
	// notes for the transactions
	Notes []string `json:"notes"`
}

// SetAttributeRequest is a struct for SetAttribute() requests
type SetAttributeRequest struct {
	// attribute name
	Key string `json:"key"`
	// attribute value
	Value string `json:"value"`
}

// SetAttributeResponse is a struct for SetAttribute() responses
type SetAttributeResponse struct {
	// None
}

// GetAttributeRequest is a struct for GetAttribute() requests
type GetAttributeRequest struct {
	// attribute name
	Key string `json:"key"`
}

// GetAttributeResponse is a struct for GetAttribute() responses
type GetAttributeResponse struct {
	// attribute value
	Value string `json:"value"`
}

// GetTxKeyRequest is a struct for GetTxKey() requests
type GetTxKeyRequest struct {
	// transaction id.
	Txid string `json:"txid"`
}

// GetTxKeyResponse is a struct for GetTxKey() responses
type GetTxKeyResponse struct {
	// transaction secret key.
	TxKey string `json:"tx_key"`
}

// CheckTxKeyRequest is a struct for CheckTxKey() requests
type CheckTxKeyRequest struct {
	// transaction id.
	Txid string `json:"txid"`
	// transaction secret key.
	TxKey string `json:"tx_key"`
	// destination public address of the transaction.
	Address string `json:"address"`
}

// CheckTxKeyResponse is a struct for CheckTxKey() responses
type CheckTxKeyResponse struct {
	// Number of block mined after the one with the transaction.
	Confirmations uint64 `json:"confirmations"`
	// States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`
	// Amount of the transaction.
	Received uint64 `json:"received"`
}

// GetTxProofRequest is a struct for GetTxProof() requests
type GetTxProofRequest struct {
	// transaction id.
	Txid string `json:"txid"`
	// destination public address of the transaction.
	Address string `json:"address"`
	// (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message,omitempty"`
}

// GetTxProofResponse is a struct for GetTxProof() responses
type GetTxProofResponse struct {
	// transaction signature.
	Signature string `json:"signature"`
}

// CheckTxProofRequest is a struct for CheckTxProof() requests
type CheckTxProofRequest struct {
	// transaction id.
	Txid string `json:"txid"`
	// destination public address of the transaction.
	Address string `json:"address"`
	// (Optional) Should be the same message used in get_tx_proof.
	Message string `json:"message,omitempty"`
	// transaction signature to confirm.
	Signature string `json:"signature"`
}

// CheckTxProofResponse is a struct for CheckTxProof() responses
type CheckTxProofResponse struct {
	// Number of block mined after the one with the transaction.
	Confirmations uint64 `json:"confirmations"`
	// States if the inputs proves the transaction.
	Good bool `json:"good"`
	// States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`
	// Amount of the transaction.
	Received uint64 `json:"received"`
}

// GetSpendProofRequest is a struct for GetSpendProof() requests
type GetSpendProofRequest struct {
	// transaction id.
	Txid string `json:"txid"`
	// (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message,omitempty"`
}

// GetSpendProofResponse is a struct for GetSpendProof() responses
type GetSpendProofResponse struct {
	// spend signature.
	Signature string `json:"signature"`
}

// CheckSpendProofRequest is a struct for CheckSpendProof() requests
type CheckSpendProofRequest struct {
	// transaction id.
	Txid string `json:"txid"`
	// (Optional) Should be the same message used in get_spend_proof.
	Message string `json:"message,omitempty"`
	// spend signature to confirm.
	Signature string `json:"signature"`
}

// CheckSpendProofResponse is a struct for CheckSpendProof() responses
type CheckSpendProofResponse struct {
	// States if the inputs proves the spend.
	Good bool `json:"good"`
}

// GetReserveProofRequest is a struct for GetReserveProof() requests
type GetReserveProofRequest struct {
	// Proves all wallet balance to be disposable.
	All bool `json:"all"`
	// Specify the account from witch to prove reserve. (ignored if all is set to true)
	AccountIndex uint64 `json:"account_index"`
	// Amount (in atomic units) to prove the account has for reserve. (ignored if all is set to true)
	Amount gonero.AtomicXMR `json:"amount"`
	// (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message,omitempty"`
}

// GetReserveProofResponse is a struct for GetReserveProof() responses
type GetReserveProofResponse struct {
	// reserve signature.
	Signature string `json:"signature"`
}

// CheckReserveProofRequest is a struct for CheckReserveProof() requests
type CheckReserveProofRequest struct {
	// Public address of the wallet.
	Address string `json:"address"`
	// (Optional) Should be the same message used in get_reserve_proof.
	Message string `json:"message,omitempty"`
	// reserve signature to confirm.
	Signature string `json:"signature"`
}

// CheckReserveProofResponse is a struct for CheckReserveProof() responses
type CheckReserveProofResponse struct {
	// States if the inputs proves the reserve.
	Good bool `json:"good"`
}

// GetTransfersRequest is a struct for GetTransfers() requests
type GetTransfersRequest struct {
	// (defaults to false) Include incoming transfers.
	In bool `json:"in"`
	// (defaults to false) Include outgoing transfers.
	Out bool `json:"out"`
	// (defaults to false) Include pending transfers.
	Pending bool `json:"pending"`
	// (defaults to false) Include failed transfers.
	Failed bool `json:"failed"`
	// (defaults to false) Include transfers from the daemon's transaction pool.
	Pool bool `json:"pool"`
	// (Optional) Filter transfers by block height.
	FilterByHeight bool `json:"filter_by_height,omitempty"`
	// (Optional) Minimum block height to scan for transfers, if filtering by height is enabled.
	MinHeight uint64 `json:"min_height,omitempty"`
	// (Optional) Maximum block height to scan for transfers, if filtering by height is enabled (defaults to max block height).
	MaxHeight uint64 `json:"max_height,omitempty"`
	// (Optional) Index of the account to query for transfers. (defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`
	// (Optional) List of subaddress indices to query for transfers. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
}

// GetTransfersResponse is a struct for GetTransfers() responses
type GetTransfersResponse struct {
	// array of incoming transfers
	In []Transfer `json:"in"`
	// array of outgoing transfers
	Out []Transfer `json:"out array of transfers (see above)."`
	// array of pending transfers
	Pending []Transfer `json:"pending array of transfers (see above)."`
	// array of failed transfers
	Failed []Transfer `json:"failed array of transfers (see above)."`
	// array of transfers in the mempool
	Pool []Transfer `json:"pool array of transfers (see above)."`
}

// Transfer is a struct for GetTransfers()
type Transfer struct {
	//  Public address of the transfer.
	Address string `json:"address"`
	//  Amount transferred.
	Amount gonero.AtomicXMR `json:"amount"`
	// Number of block mined since the block containing this
	// transaction (or block height at which the transaction
	// should be added to a block if not yet confirmed).
	Confirmations uint64 `json:"confirmations"`
	//  True if the key image(s) for the transfer have been seen before.
	DoubleSpendSeen bool `json:"double_spend_seen"`
	//  Transaction fee for this transfer.
	Fee uint64 `json:"fee"`
	//  Height of the first block that confirmed this transfer (0 if not mined yet).
	Height uint64 `json:"height"`
	//  Note about this transfer.
	Note string `json:"note"`
	//  Payment ID for this transfer.
	PaymentID string `json:"payment_id"`
	// JSON object containing the major & minor subaddress index
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`
	//  Account index for the subaddress.
	Major uint64 `json:"major"`
	//  Index of the subaddress under the account.
	Minor uint64 `json:"minor"`
	//  Estimation of the confirmations needed for the transaction to be included in a block.
	SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
	//  POSIX timestamp for when this transfer was first confirmed in a block (or timestamp submission if not mined yet).
	Timestamp uint64 `json:"timestamp"`
	//  Transaction ID for this transfer.
	Txid string `json:"txid"`
	//  Transfer type: "in"
	Type string `json:"type"`
	//  Number of blocks until transfer is safely spendable.
	UnlockTime uint64 `json:"unlock_time"`
}

// GetTransferByTxidRequest is a struct for GetTransferByTxid() requests
type GetTransferByTxidRequest struct {
	// Transaction ID used to find the transfer.
	Txid string `json:"txid"`
	// (Optional) Index of the account to query for the transfer.
	AccountIndex uint64 `json:"account_index,omitempty"`
}

// GetTransferByTxidResponse is a struct for GetTransferByTxid() responses
type GetTransferByTxidResponse struct {
	// JSON object containing payment information
	Transfer []TransferDetail `json:"transfer"`
}

// TransferDetail is
type TransferDetail struct {
	// Address that transferred the funds. Base58 representation of the public keys.
	Address string `json:"address"`
	// Amount of this transfer.
	Amount gonero.AtomicXMR `json:"amount"`
	// Number of block mined since the block containing this transaction
	// (or block height at which the transaction should be added to
	// a block if not yet confirmed).
	Confirmations uint64 `json:"confirmations"`
	// array of JSON objects containing transfer destinations
	Destinations []Destination `json:"destinations"`
	// True if the key image(s) for the transfer have been seen before.
	DoubleSpendSeen bool `json:"double_spend_seen"`
	// Transaction fee for this transfer.
	Fee uint64 `json:"fee"`
	// Height of the first block that confirmed this transfer.
	Height uint64 `json:"height"`
	// Note about this transfer.
	Note string `json:"note"`
	// Payment ID for this transfer.
	PaymentID string `json:"payment_id"`
	// JSON object containing the major & minor subaddress index
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`
	// Account index for the subaddress.
	Major uint64 `json:"major"`
	// Index of the subaddress under the account.
	Minor uint64 `json:"minor"`
	// Estimation of the confirmations needed for the transaction to be included in a block.
	SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`
	// POSIX timestamp for the block that confirmed this transfer (or timestamp submission if not mined yet).
	Timestamp uint64 `json:"timestamp"`
	// Transaction ID of this transfer (same as input TXID).
	Txid string `json:"txid"`
	// Type of transfer, one of the following: "in", "out", "pending", "failed", "pool"
	Type string `json:"type"`
	// Number of blocks until transfer is safely spendable.
	UnlockTime uint64 `json:"unlock_time"`
}

// DescribeTransferRequest is a struct for DescribeTransfer() requests
type DescribeTransferRequest struct {
	// (Optional) A hexadecimal string representing a set of unsigned transactions (empty for multisig transactions;non-multisig signed transactions are not supported).
	UnsignedTxset string `json:"unsigned_txset,omitempty"`
	// (Optional) A hexadecimal string representing the set of signing keys used in a multisig transaction (empty for unsigned transactions;non-multisig signed transactions are not supported).
	MultisigTxset string `json:"multisig_txset,omitempty"`
}

// DescribeTransferResponse is a struct for DescribeTransfer() responses
type DescribeTransferResponse struct {
	// (Optional) A hexadecimal string representing a set of unsigned transactions (empty for multisig transactions;non-multisig signed transactions are not supported).
	UnsignedTxset string `json:"unsigned_txset,omitempty"`
	// (Optional) A hexadecimal string representing the set of signing keys used in a multisig transaction (empty for unsigned transactions;non-multisig signed transactions are not supported).
	MultisigTxset string `json:"multisig_txset,omitempty"`
}

// SignRequest is a struct for Sign() requests
type SignRequest struct {
	// Anything you need to sign.
	Data string `json:"data"`
}

// SignResponse is a struct for Sign() responses
type SignResponse struct {
	// Signature generated against the "data" and the account public address.
	Signature string `json:"signature"`
}

// VerifyRequest is a struct for Verify() requests
type VerifyRequest struct {
	// What should have been signed.
	Data string `json:"data"`
	// Public address of the wallet used to sign the data.
	Address string `json:"address"`
	// signature generated by sign method.
	Signature string `json:"signature"`
}

// VerifyResponse is a struct for Verify() responses
type VerifyResponse struct {
	// boolean
	Good bool `json:"good"`
}

// ExportOutputsRequest is a struct for ExportOutputs() requests
type ExportOutputsRequest struct {
	// None
}

// ExportOutputsResponse is a struct for ExportOutputs() responses
type ExportOutputsResponse struct {
	// wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

// ImportOutputsRequest is a struct for ImportOutputs() requests
type ImportOutputsRequest struct {
	// wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

// ImportOutputsResponse is a struct for ImportOutputs() responses
type ImportOutputsResponse struct {
	// number of outputs imported.
	NumImported uint64 `json:"num_imported"`
}

// ExportKeyImagesRequest is a struct for ExportKeyImages() requests
type ExportKeyImagesRequest struct {
	// None
}

// ExportKeyImagesResponse is a struct for ExportKeyImages() responses
type ExportKeyImagesResponse struct {
	// array of signed key images
	SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
	// string
	KeyImage string `json:"key_image"`
	// string
	Signature string `json:"signature"`
}

// SignedKeyImage is a struct for ExportKeyImages()
type SignedKeyImage struct {
	// string
	KeyImage string `json:"key_image"`
	// string
	Signature string `json:"signature"`
}

// ImportKeyImagesRequest is a struct for ImportKeyImages() requests
type ImportKeyImagesRequest struct {
	// key_image;string;signature - string
	SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
	// string
	KeyImage string `json:"key_image"`
	// string
	Signature string `json:"signature"`
}

// ImportKeyImagesResponse is a struct for ImportKeyImages() responses
type ImportKeyImagesResponse struct {
	// unsigned int
	Height uint64 `json:"height"`
	// Amount (in atomic units) spent from those key images.
	Spent uint64 `json:"spent"`
	// Amount (in atomic units) still available from those key images.
	Unspent uint64 `json:"unspent"`
}

// MakeURIRequest is a struct for MakeURI() requests
type MakeURIRequest struct {
	// Wallet address
	Address string `json:"address"`
	// (optional) the integer amount to receive, in atomic units
	Amount gonero.AtomicXMR `json:"amount,omitempty"`
	// (optional) 16 or 64 character hexadecimal payment id
	PaymentID string `json:"payment_id,omitempty"`
	// (optional) name of the payment recipient
	RecipientName string `json:"recipient_name,omitempty"`
	// (optional) Description of the reason for the tx
	TxDescription string `json:"tx_description,omitempty"`
}

// MakeURIResponse is a struct for MakeURI() responses
type MakeURIResponse struct {
	// This contains all the payment input information as a properly formatted payment URI
	URI string `json:"uri"`
}

// ParseURIRequest is a struct for ParseURI() requests
type ParseURIRequest struct {
	// This contains all the payment input information as a properly formatted payment URI
	URI string `json:"uri"`
}

// ParseURIResponse is a struct for ParseURI() responses
type ParseURIResponse struct {
	// JSON object containing payment information
	URI URI `json:"uri"`
}

// URI is a struct
type URI struct {
	// Wallet address
	Address string `json:"address"`
	// Integer amount to receive, in atomic units (0 if not provided)
	Amount gonero.AtomicXMR `json:"amount"`
	// 16 or 64 character hexadecimal payment id (empty if not provided)
	PaymentID string `json:"payment_id"`
	// Name of the payment recipient (empty if not provided)
	RecipientName string `json:"recipient_name"`
	// Description of the reason for the tx (empty if not provided)
	TxDescription string `json:"tx_description"`
}

// GetAddressBookRequest is a struct for GetAddressBook() requests
type GetAddressBookRequest struct {
	// indices of the requested address book entries
	Entries []uint64 `json:"entries"`
}

// GetAddressBookResponse is a struct for GetAddressBook() responses
type GetAddressBookResponse struct {
	// array of entries
	Entries []AddressBookEntry `json:"entries"`
}

// AddressBookEntry is a struct for an entry in the address book
type AddressBookEntry struct {
	//  Public address of the entry
	Address string `json:"address"`
	//  Description of this address entry
	Description string `json:"description"`
	//
	Index uint64 `json:"index"`
	//
	PaymentID string `json:"payment_id"`
}

// AddAddressBookRequest is a struct for AddAddressBook() requests
type AddAddressBookRequest struct {
	// string
	Address string `json:"address"`
	// (optional), defaults to "0000000000000000000000000000000000000000000000000000000000000000"
	PaymentID string `json:"payment_id,omitempty"`
	// (optional), defaults to ""
	Description string `json:"description,omitempty"`
}

// AddAddressBookResponse is a struct for AddAddressBook() responses
type AddAddressBookResponse struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// EditAddressBookRequest is a struct for EditAddressBook() requests
type EditAddressBookRequest struct {
	// Index of the address book entry to edit.
	Index uint64 `json:"index"`
	// If true, set the address for this entry to the value of "address".
	SetAddress bool `json:"set_address"`
	// (Optional) The 95-character public address to set.
	Address string `json:"address,omitempty"`
	// If true, set the description for this entry to the value of "description".
	SetDescription bool `json:"set_description"`
	// (Optional) Human-readable description for this entry.
	Description string `json:"description,omitempty"`
	// If true, set the payment ID for this entry to the value of "payment_id".
	SetPaymentID bool `json:"set_payment_id"`
	// (Optional) Payment ID for this address.
	PaymentID string `json:"payment_id,omitempty"`
}

// EditAddressBookResponse is a struct for EditAddressBook() responses
type EditAddressBookResponse struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// DeleteAddressBookRequest is a struct for DeleteAddressBook() requests
type DeleteAddressBookRequest struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// DeleteAddressBookResponse is a struct for DeleteAddressBook() responses
type DeleteAddressBookResponse struct {
	// None
}

// RefreshRequest is a struct for Refresh() requests
type RefreshRequest struct {
	// (Optional) The block height from which to start refreshing.
	StartHeight uint64 `json:"start_height,omitempty"`
}

// RefreshResponse is a struct for Refresh() responses
type RefreshResponse struct {
	// Number of new blocks scanned.
	BlocksFetched uint64 `json:"blocks_fetched"`
	// States if transactions to the wallet have been found in the blocks.
	ReceivedMoney bool `json:"received_money"`
}

// AutoRefreshRequest is a struct for AutoRefresh() requests
type AutoRefreshRequest struct {
	// (Optional) Enable or disable automatic refreshing (default = true).
	Enable bool `json:"enable,omitempty"`
	// (Optional) The period of the wallet refresh cycle (i.e. time between refreshes) in seconds.
	Period uint64 `json:"period,omitempty"`
}

// AutoRefreshResponse is a struct for AutoRefresh() responses
type AutoRefreshResponse struct {
	// Number of threads created for mining.
	ThreadsCount uint64 `json:"threads_count"`
	// Allow to start the miner in smart mining mode.
	DoBackgroundMining bool `json:"do_background_mining"`
	// Ignore battery status (for smart mining only)
	IgnoreBattery bool `json:"ignore_battery"`
}

// RescanSpentRequest is a struct for RescanSpent() requests
type RescanSpentRequest struct {
	// None
}

// RescanSpentResponse is a struct for RescanSpent() responses
type RescanSpentResponse struct {
	// None
}

// StartMiningRequest is a struct for StartMining() requests
type StartMiningRequest struct {
	// Number of threads created for mining.
	ThreadsCount uint64 `json:"threads_count"`
	// Allow to start the miner in smart mining mode.
	DoBackgroundMining bool `json:"do_background_mining"`
	// Ignore battery status (for smart mining only)
	IgnoreBattery bool `json:"ignore_battery"`
}

// StartMiningResponse is a struct for StartMining() responses
type StartMiningResponse struct {
	// None
}

// StopMiningRequest is a struct for StopMining() requests
type StopMiningRequest struct {
	// None
}

// StopMiningResponse is a struct for StopMining() responses
type StopMiningResponse struct {
	// None
}

// GetLanguagesRequest is a struct for GetLanguages() requests
type GetLanguagesRequest struct {
	// None
}

// GetLanguagesResponse is a struct for GetLanguages() responses
type GetLanguagesResponse struct {
	// List of available languages
	Languages []string `json:"languages"`
}

// CreateWalletRequest is a struct for CreateWallet() requests
type CreateWalletRequest struct {
	// Wallet file name.
	Filename string `json:"filename"`
	// (Optional) password to protect the wallet.
	Password string `json:"password,omitempty"`
	// Language for your wallets' seed.
	Language string `json:"language"`
}

// CreateWalletResponse is a struct for CreateWallet() responses
type CreateWalletResponse struct {
	// None
}

// GenerateFromKeysRequest is a struct for GenerateFromKeys() requests
type GenerateFromKeysRequest struct {
	// (Optional;defaults to 0) The block height to restore the wallet from.
	RestoreHeight int64 `json:"restore_height,omitempty"`
	// The wallet's file name on the RPC server.
	Filename string `json:"filename"`
	// The wallet's primary address.
	Address string `json:"address"`
	// (Optional;omit to create a view-only wallet) The wallet's private spend key.
	Spendkey string `json:"spendkey,omitempty"`
	// The wallet's private view key.
	Viewkey string `json:"viewkey"`
	// The wallet's password.
	Password string `json:"password"`
	// (Defaults to true) If true, save the current wallet before generating the new wallet.
	AutosaveCurrent bool `json:"autosave_current"`
}

// GenerateFromKeysResponse is a struct for GenerateFromKeys() responses
type GenerateFromKeysResponse struct {
	// The wallet's address.
	Address string `json:"address"`
	// Verification message indicating that the wallet was generated successfully and whether or not it is a view-only wallet.
	Info string `json:"info"`
}

// OpenWalletRequest is a struct for OpenWallet() requests
type OpenWalletRequest struct {
	// wallet name stored in âwallet-dir.
	Filename string `json:"filename"`
	// (Optional) only needed if the wallet has a password defined.
	Password string `json:"password,omitempty"`
}

// OpenWalletResponse is a struct for OpenWallet() responses
type OpenWalletResponse struct {
	// None
}

// RestoreDeterministicWalletRequest is a struct for RestoreDeterministicWallet() requests
type RestoreDeterministicWalletRequest struct {
	// Name of the wallet.
	Name string `json:"name"`
	// Password of the wallet.
	Password string `json:"password"`
	// Mnemonic phrase of the wallet to restore.
	Seed string `json:"seed"`
	// (Optional) Block height to restore the wallet from (default = 0).
	RestoreHeight int64 `json:"restore_height,omitempty"`
	// (Optional) Language of the mnemonic phrase in case the old language is invalid.
	Language string `json:"language,omitempty"`
	// (Optional) Offset used to derive a new seed from the given mnemonic to recover a secret wallet from the mnemonic phrase.
	SeedOffset string `json:"seed_offset,omitempty"`
	// Whether to save the currently open RPC wallet before closing it (Defaults to true).
	AutosaveCurrent bool `json:"autosave_current"`
}

// RestoreDeterministicWalletResponse is a struct for RestoreDeterministicWallet() responses
type RestoreDeterministicWalletResponse struct {
	// 95-character hexadecimal address of the restored wallet as a string.
	Address string `json:"address"`
	// Message describing the success or failure of the attempt to restore the wallet.
	Info string `json:"info"`
	// Mnemonic phrase of the restored wallet, which is updated if the wallet was restored from a deprecated-style mnemonic phrase.
	Seed string `json:"seed"`
	// Indicates if the restored wallet was created from a deprecated-style mnemonic phrase.
	WasDeprecated bool `json:"was_deprecated"`
}

// CloseWalletRequest is a struct for CloseWallet() requests
type CloseWalletRequest struct {
	// None
}

// CloseWalletResponse is a struct for CloseWallet() responses
type CloseWalletResponse struct {
	// None
}

// ChangeWalletPasswordRequest is a struct for ChangeWalletPassword() requests
type ChangeWalletPasswordRequest struct {
	// (Optional) Current wallet password, if defined.
	OldPassword string `json:"old_password,omitempty"`
	// (Optional) New wallet password, if not blank.
	NewPassword string `json:"new_password,omitempty"`
}

// ChangeWalletPasswordResponse is a struct for ChangeWalletPassword() responses
type ChangeWalletPasswordResponse struct {
	// None
}

// IsMultisigRequest is a struct for IsMultisig() requests
type IsMultisigRequest struct {
	// None
}

// IsMultisigResponse is a struct for IsMultisig() responses
type IsMultisigResponse struct {
	// States if the wallet is multisig
	Multisig bool `json:"multisig"`
	// boolean
	Ready bool `json:"ready"`
	// Amount of signature needed to sign a transfer.
	Threshold uint64 `json:"threshold"`
	// Total amount of signature in the multisig wallet.
	Total uint64 `json:"total"`
}

// PrepareMultisigRequest is a struct for PrepareMultisig() requests
type PrepareMultisigRequest struct {
	// None
}

// PrepareMultisigResponse is a struct for PrepareMultisig() responses
type PrepareMultisigResponse struct {
	// Multisig string to share with peers to create the multisig wallet.
	MultisigInfo string `json:"multisig_info"`
}

// MakeMultisigRequest is a struct for MakeMultisig() requests
type MakeMultisigRequest struct {
	// List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`
	// Amount of signatures needed to sign a transfer. Must be less or equal than the amount of signature in multisig_info.
	Threshold uint64 `json:"threshold"`
	// Wallet password
	Password string `json:"password"`
}

// MakeMultisigResponse is a struct for MakeMultisig() responses
type MakeMultisigResponse struct {
	// multisig wallet address.
	Address string `json:"address"`
	// Multisig string to share with peers to create the multisig wallet (extra step for N-1/N wallets).
	MultisigInfo string `json:"multisig_info"`
}

// ExportMultisigInfoRequest is a struct for ExportMultisigInfo() requests
type ExportMultisigInfoRequest struct {
	// None
}

// ExportMultisigInfoResponse is a struct for ExportMultisigInfo() responses
type ExportMultisigInfoResponse struct {
	// Multisig info in hex format for other participants.
	Info string `json:"info"`
}

// ImportMultisigInfoRequest is a struct for ImportMultisigInfo() requests
type ImportMultisigInfoRequest struct {
	// List of multisig info in hex format from other participants.
	Info []string `json:"info"`
}

// ImportMultisigInfoResponse is a struct for ImportMultisigInfo() responses
type ImportMultisigInfoResponse struct {
	// Number of outputs signed with those multisig info.
	NOutputs uint64 `json:"n_outputs"`
}

// FinalizeMultisigRequest is a struct for FinalizeMultisig() requests
type FinalizeMultisigRequest struct {
	// List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`
	// Wallet password
	Password string `json:"password"`
}

// FinalizeMultisigResponse is a struct for FinalizeMultisig() responses
type FinalizeMultisigResponse struct {
	// multisig wallet address.
	Address string `json:"address"`
}

// SignMultisigRequest is a struct for SignMultisig() requests
type SignMultisigRequest struct {
	// Multisig transaction in hex format, as returned by transfer under multisig_txset.
	TxDataHex string `json:"tx_data_hex"`
}

// SignMultisigResponse is a struct for SignMultisig() responses
type SignMultisigResponse struct {
	// Multisig transaction in hex format.
	TxDataHex string `json:"tx_data_hex"`
	// List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// SubmitMultisigRequest is a struct for SubmitMultisig() requests
type SubmitMultisigRequest struct {
	// Multisig transaction in hex format, as returned by sign_multisig under tx_data_hex.
	TxDataHex string `json:"tx_data_hex"`
}

// SubmitMultisigResponse is a struct for SubmitMultisig() responses
type SubmitMultisigResponse struct {
	// List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// GetVersionRequest is a struct for GetVersion() requests
type GetVersionRequest struct {
	// None
}

// GetVersionResponse is a struct for GetVersion() responses
type GetVersionResponse struct {
	// RPC version, formatted with Major * 2^16 + Minor (Major encoded over the first 16 bits, and Minor over the last 16 bits).
	Version uint64 `json:"version"`
}

//#####################
// Other RPC Structs
//#####################
