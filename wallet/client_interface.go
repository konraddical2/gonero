package wallet

// Client is a monero-wallet-rpc client
type Client interface {
	// JSON RPC Methods
	// Connect the RPC server to a Monero daemon.
	SetDaemon(*SetDaemonRequest) (*SetDaemonResponse, error)
	// Return the wallet's balance.
	GetBalance(*GetBalanceRequest) (*GetBalanceResponse, error)
	// Return the wallet's addresses for an account. Optionally filter for specific set of subaddresses.
	GetAddress(*GetAddressRequest) (*GetAddressResponse, error)
	// Get account and address indexes from a specific (sub)address
	GetAddressIndex(*GetAddressIndexRequest) (*GetAddressIndexResponse, error)
	// Create a new address for an account. Optionally, label the new address.
	CreateAddress(*CreateAddressRequest) (*CreateAddressResponse, error)
	// Label an address.
	LabelAddress(*LabelAddressRequest) (*LabelAddressResponse, error)
	// Analyzes a string to determine whether it is a valid monero wallet address and returns the result and the address specifications.
	ValidateAddress(*ValidateAddressRequest) (*ValidateAddressResponse, error)
	// Get all accounts for a wallet. Optionally filter accounts by tag.
	GetAccounts(*GetAccountsRequest) (*GetAccountsResponse, error)
	// Create a new account with an optional label.
	CreateAccount(*CreateAccountRequest) (*CreateAccountResponse, error)
	// Label an account.
	LabelAccount(*LabelAccountRequest) (*LabelAccountResponse, error)
	// Get a list of user-defined account tags.
	GetAccountTags(*GetAccountTagsRequest) (*GetAccountTagsResponse, error)
	// Apply a filtering tag to a list of accounts.
	TagAccounts(*TagAccountsRequest) (*TagAccountsResponse, error)
	// Remove filtering tag from a list of accounts.
	UntagAccounts(*UntagAccountsRequest) (*UntagAccountsResponse, error)
	// Set description for an account tag.
	SetAccountTagDescription(*SetAccountTagDescriptionRequest) (*SetAccountTagDescriptionResponse, error)
	// Returns the wallet's current block height.
	GetHeight(*GetHeightRequest) (*GetHeightResponse, error)
	// Send monero to a number of recipients.
	Transfer(*TransferRequest) (*TransferResponse, error)
	// Same as transfer, but can split into more than one tx if necessary.
	TransferSplit(*TransferSplitRequest) (*TransferSplitResponse, error)
	// Sign a transaction created on a read-only wallet (in cold-signing process)
	SignTransfer(*SignTransferRequest) (*SignTransferResponse, error)
	// Submit a previously signed transaction on a read-only wallet (in cold-signing process).
	SubmitTransfer(*SubmitTransferRequest) (*SubmitTransferResponse, error)
	// Send all dust outputs back to the wallet's, to make them easier to spend (and mix).
	SweepDust(*SweepDustRequest) (*SweepDustResponse, error)
	// Send all unlocked balance to an address.
	SweepAll(*SweepAllRequest) (*SweepAllResponse, error)
	// Send all of a specific unlocked output to an address.
	SweepSingle(*SweepSingleRequest) (*SweepSingleResponse, error)
	// Relay a transaction previously created with "do_not_relay":true.
	RelayTx(*RelayTxRequest) (*RelayTxResponse, error)
	// Save the wallet file.
	Store(*StoreRequest) (*StoreResponse, error)
	// Get a list of incoming payments using a given payment id.
	GetPayments(*GetPaymentsRequest) (*GetPaymentsResponse, error)
	// Get a list of incoming payments using a given payment id, or a list of payments ids, from a given height. This method is the preferred method over get_payments because it has the same functionality but is more extendable. Either is fine for looking up transactions by a single payment ID.
	GetBulkPayments(*GetBulkPaymentsRequest) (*GetBulkPaymentsResponse, error)
	// Return a list of incoming transfers to the wallet.
	IncomingTransfers(*IncomingTransfersRequest) (*IncomingTransfersResponse, error)
	// Return the spend or view private key.
	QueryKey(*QueryKeyRequest) (*QueryKeyResponse, error)
	// Make an integrated address from the wallet address and a payment id.
	MakeIntegratedAddress(*MakeIntegratedAddressRequest) (*MakeIntegratedAddressResponse, error)
	// Retrieve the standard address and payment id corresponding to an integrated address.
	SplitIntegratedAddress(*SplitIntegratedAddressRequest) (*SplitIntegratedAddressResponse, error)
	// Stops the wallet, storing the current state.
	StopWallet(*StopWalletRequest) (*StopWalletResponse, error)
	// Rescan the blockchain from scratch, losing any information which can not be recovered from the blockchain itself. This includes destination addresses, tx secret keys, tx notes, etc.
	RescanBlockchain(*RescanBlockchainRequest) (*RescanBlockchainResponse, error)
	// Set arbitrary string notes for transactions.
	SetTxNotes(*SetTxNotesRequest) (*SetTxNotesResponse, error)
	// Get string notes for transactions.
	GetTxNotes(*GetTxNotesRequest) (*GetTxNotesResponse, error)
	// Set arbitrary attribute.
	SetAttribute(*SetAttributeRequest) (*SetAttributeResponse, error)
	// Get attribute value by name.
	GetAttribute(*GetAttributeRequest) (*GetAttributeResponse, error)
	// Get transaction secret key from transaction id.
	GetTxKey(*GetTxKeyRequest) (*GetTxKeyResponse, error)
	// Check a transaction in the blockchain with its secret key.
	CheckTxKey(*CheckTxKeyRequest) (*CheckTxKeyResponse, error)
	// Get transaction signature to prove it.
	GetTxProof(*GetTxProofRequest) (*GetTxProofResponse, error)
	// Prove a transaction by checking its signature.
	CheckTxProof(*CheckTxProofRequest) (*CheckTxProofResponse, error)
	// Generate a signature to prove a spend. Unlike proving a transaction, it does not requires the destination public address.
	GetSpendProof(*GetSpendProofRequest) (*GetSpendProofResponse, error)
	// Prove a spend using a signature. Unlike proving a transaction, it does not requires the destination public address.
	CheckSpendProof(*CheckSpendProofRequest) (*CheckSpendProofResponse, error)
	// Generate a signature to prove of an available amount in a wallet.
	GetReserveProof(*GetReserveProofRequest) (*GetReserveProofResponse, error)
	// Proves a wallet has a disposable reserve using a signature.
	CheckReserveProof(*CheckReserveProofRequest) (*CheckReserveProofResponse, error)
	// Returns a list of transfers.
	GetTransfers(*GetTransfersRequest) (*GetTransfersResponse, error)
	// Show information about a transfer to/from this address.
	GetTransferByTxid(*GetTransferByTxidRequest) (*GetTransferByTxidResponse, error)
	// Returns details for each transaction in an unsigned or multisig transaction set. Transaction sets are obtained as return values from one of the following RPC methods:
	DescribeTransfer(*DescribeTransferRequest) (*DescribeTransferResponse, error)
	// Sign a string.
	Sign(*SignRequest) (*SignResponse, error)
	// Verify a signature on a string.
	Verify(*VerifyRequest) (*VerifyResponse, error)
	// Export all outputs in hex format.
	ExportOutputs(*ExportOutputsRequest) (*ExportOutputsResponse, error)
	// Import outputs in hex format.
	ImportOutputs(*ImportOutputsRequest) (*ImportOutputsResponse, error)
	// Export a signed set of key images.
	ExportKeyImages(*ExportKeyImagesRequest) (*ExportKeyImagesResponse, error)
	// Import signed key images list and verify their spent status.
	ImportKeyImages(*ImportKeyImagesRequest) (*ImportKeyImagesResponse, error)
	// Create a payment URI using the official URI spec.
	MakeURI(*MakeURIRequest) (*MakeURIResponse, error)
	// Parse a payment URI to get payment information.
	ParseURI(*ParseURIRequest) (*ParseURIResponse, error)
	// Retrieves entries from the address book.
	GetAddressBook(*GetAddressBookRequest) (*GetAddressBookResponse, error)
	// Add an entry to the address book.
	AddAddressBook(*AddAddressBookRequest) (*AddAddressBookResponse, error)
	// Edit an existing address book entry.
	EditAddressBook(*EditAddressBookRequest) (*EditAddressBookResponse, error)
	// Delete an entry from the address book.
	DeleteAddressBook(*DeleteAddressBookRequest) (*DeleteAddressBookResponse, error)
	// Refresh a wallet after openning.
	Refresh(*RefreshRequest) (*RefreshResponse, error)
	// Set whether and how often to automatically refresh the current wallet.
	AutoRefresh(*AutoRefreshRequest) (*AutoRefreshResponse, error)
	// Rescan the blockchain for spent outputs.
	RescanSpent(*RescanSpentRequest) (*RescanSpentResponse, error)
	// Start mining in the Monero daemon.
	StartMining(*StartMiningRequest) (*StartMiningResponse, error)
	// Stop mining in the Monero daemon.
	StopMining(*StopMiningRequest) (*StopMiningResponse, error)
	// Get a list of available languages for your wallet's seed.
	GetLanguages(*GetLanguagesRequest) (*GetLanguagesResponse, error)
	// Create a new wallet. You need to have set the argument "âwallet-dir" when launching monero-wallet-rpc to make this work.
	CreateWallet(*CreateWalletRequest) (*CreateWalletResponse, error)
	// Restores a wallet from a given wallet address, view key, and optional spend key.
	GenerateFromKeys(*GenerateFromKeysRequest) (*GenerateFromKeysResponse, error)
	// Open a wallet. You need to have set the argument "âwallet-dir" when launching monero-wallet-rpc to make this work.
	OpenWallet(*OpenWalletRequest) (*OpenWalletResponse, error)
	// Create and open a wallet on the RPC server from an existing mnemonic phrase and close the currently open wallet.
	RestoreDeterministicWallet(*RestoreDeterministicWalletRequest) (*RestoreDeterministicWalletResponse, error)
	// Close the currently opened wallet, after trying to save it.
	CloseWallet(*CloseWalletRequest) (*CloseWalletResponse, error)
	// Change a wallet password.
	ChangeWalletPassword(*ChangeWalletPasswordRequest) (*ChangeWalletPasswordResponse, error)
	// Check if a wallet is a multisig one.
	IsMultisig(*IsMultisigRequest) (*IsMultisigResponse, error)
	// Prepare a wallet for multisig by generating a multisig string to share with peers.
	PrepareMultisig(*PrepareMultisigRequest) (*PrepareMultisigResponse, error)
	// Make a wallet multisig by importing peers multisig string.
	MakeMultisig(*MakeMultisigRequest) (*MakeMultisigResponse, error)
	// Export multisig info for other participants.
	ExportMultisigInfo(*ExportMultisigInfoRequest) (*ExportMultisigInfoResponse, error)
	// Import multisig info from other participants.
	ImportMultisigInfo(*ImportMultisigInfoRequest) (*ImportMultisigInfoResponse, error)
	// Turn this wallet into a multisig wallet, extra step for N-1/N wallets.
	FinalizeMultisig(*FinalizeMultisigRequest) (*FinalizeMultisigResponse, error)
	// Sign a transaction in multisig.
	SignMultisig(*SignMultisigRequest) (*SignMultisigResponse, error)
	// Submit a signed multisig transaction.
	SubmitMultisig(*SubmitMultisigRequest) (*SubmitMultisigResponse, error)
	// Get RPC version Major & Minor integer-format, where Major is the first 16 bits and Minor the last 16 bits.
	GetVersion(*GetVersionRequest) (*GetVersionResponse, error)

	// Other RPC Methods

}
