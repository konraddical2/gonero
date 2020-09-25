package wallet

//#####################
// General RPC Methods
//#####################

//#####################
// JSON RPC Methods
//#####################

// SetDaemon Connect the RPC server to a Monero daemon.
func (c *client) SetDaemon(req *SetDaemonRequest) (resp *SetDaemonResponse, err error) {
	err = c.do("set_daemon", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBalance Return the wallet's balance.
func (c *client) GetBalance(req *GetBalanceRequest) (resp *GetBalanceResponse, err error) {
	err = c.do("get_balance", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAddress Return the wallet's addresses for an account. Optionally filter for specific set of subaddresses.
func (c *client) GetAddress(req *GetAddressRequest) (resp *GetAddressResponse, err error) {
	err = c.do("get_address", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAddressIndex Get account and address indexes from a specific (sub)address
func (c *client) GetAddressIndex(req *GetAddressIndexRequest) (resp *GetAddressIndexResponse, err error) {
	err = c.do("get_address_index", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CreateAddress Create a new address for an account. Optionally, label the new address.
func (c *client) CreateAddress(req *CreateAddressRequest) (resp *CreateAddressResponse, err error) {
	err = c.do("create_address", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// LabelAddress Label an address.
func (c *client) LabelAddress(req *LabelAddressRequest) (resp *LabelAddressResponse, err error) {
	err = c.do("label_address", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ValidateAddress Analyzes a string to determine whether it is a valid monero wallet address and returns the result and the address specifications.
func (c *client) ValidateAddress(req *ValidateAddressRequest) (resp *ValidateAddressResponse, err error) {
	err = c.do("validate_address", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAccounts Get all accounts for a wallet. Optionally filter accounts by tag.
func (c *client) GetAccounts(req *GetAccountsRequest) (resp *GetAccountsResponse, err error) {
	err = c.do("get_accounts", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CreateAccount Create a new account with an optional label.
func (c *client) CreateAccount(req *CreateAccountRequest) (resp *CreateAccountResponse, err error) {
	err = c.do("create_account", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// LabelAccount Label an account.
func (c *client) LabelAccount(req *LabelAccountRequest) (resp *LabelAccountResponse, err error) {
	err = c.do("label_account", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAccountTags Get a list of user-defined account tags.
func (c *client) GetAccountTags(req *GetAccountTagsRequest) (resp *GetAccountTagsResponse, err error) {
	err = c.do("get_account_tags", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// TagAccounts Apply a filtering tag to a list of accounts.
func (c *client) TagAccounts(req *TagAccountsRequest) (resp *TagAccountsResponse, err error) {
	err = c.do("tag_accounts", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// UntagAccounts Remove filtering tag from a list of accounts.
func (c *client) UntagAccounts(req *UntagAccountsRequest) (resp *UntagAccountsResponse, err error) {
	err = c.do("untag_accounts", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetAccountTagDescription Set description for an account tag.
func (c *client) SetAccountTagDescription(req *SetAccountTagDescriptionRequest) (resp *SetAccountTagDescriptionResponse, err error) {
	err = c.do("set_account_tag_description", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetHeight Returns the wallet's current block height.
func (c *client) GetHeight(req *GetHeightRequest) (resp *GetHeightResponse, err error) {
	err = c.do("get_height", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// Transfer Send monero to a number of recipients.
func (c *client) Transfer(req *TransferRequest) (resp *TransferResponse, err error) {
	err = c.do("transfer", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// TransferSplit Same as transfer, but can split into more than one tx if necessary.
func (c *client) TransferSplit(req *TransferSplitRequest) (resp *TransferSplitResponse, err error) {
	err = c.do("transfer_split", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SignTransfer Sign a transaction created on a read-only wallet (in cold-signing process)
func (c *client) SignTransfer(req *SignTransferRequest) (resp *SignTransferResponse, err error) {
	err = c.do("sign_transfer", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SubmitTransfer Submit a previously signed transaction on a read-only wallet (in cold-signing process).
func (c *client) SubmitTransfer(req *SubmitTransferRequest) (resp *SubmitTransferResponse, err error) {
	err = c.do("submit_transfer", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SweepDust Send all dust outputs back to the wallet's, to make them easier to spend (and mix).
func (c *client) SweepDust(req *SweepDustRequest) (resp *SweepDustResponse, err error) {
	err = c.do("sweep_dust", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SweepAll Send all unlocked balance to an address.
func (c *client) SweepAll(req *SweepAllRequest) (resp *SweepAllResponse, err error) {
	err = c.do("sweep_all", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SweepSingle Send all of a specific unlocked output to an address.
func (c *client) SweepSingle(req *SweepSingleRequest) (resp *SweepSingleResponse, err error) {
	err = c.do("sweep_single", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// RelayTx Relay a transaction previously created with "do_not_relay":true.
func (c *client) RelayTx(req *RelayTxRequest) (resp *RelayTxResponse, err error) {
	err = c.do("relay_tx", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// Store Save the wallet file.
func (c *client) Store(req *StoreRequest) (resp *StoreResponse, err error) {
	err = c.do("store", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetPayments Get a list of incoming payments using a given payment id.
func (c *client) GetPayments(req *GetPaymentsRequest) (resp *GetPaymentsResponse, err error) {
	err = c.do("get_payments", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBulkPayments Get a list of incoming payments using a given payment id, or a list of payments ids, from a given height. This method is the preferred method over get_payments because it has the same functionality but is more extendable. Either is fine for looking up transactions by a single payment ID.
func (c *client) GetBulkPayments(req *GetBulkPaymentsRequest) (resp *GetBulkPaymentsResponse, err error) {
	err = c.do("get_bulk_payments", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// IncomingTransfers Return a list of incoming transfers to the wallet.
func (c *client) IncomingTransfers(req *IncomingTransfersRequest) (resp *IncomingTransfersResponse, err error) {
	err = c.do("incoming_transfers", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// QueryKey Return the spend or view private key.
func (c *client) QueryKey(req *QueryKeyRequest) (resp *QueryKeyResponse, err error) {
	err = c.do("query_key", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// MakeIntegratedAddress Make an integrated address from the wallet address and a payment id.
func (c *client) MakeIntegratedAddress(req *MakeIntegratedAddressRequest) (resp *MakeIntegratedAddressResponse, err error) {
	err = c.do("make_integrated_address", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SplitIntegratedAddress Retrieve the standard address and payment id corresponding to an integrated address.
func (c *client) SplitIntegratedAddress(req *SplitIntegratedAddressRequest) (resp *SplitIntegratedAddressResponse, err error) {
	err = c.do("split_integrated_address", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// StopWallet Stops the wallet, storing the current state.
func (c *client) StopWallet(req *StopWalletRequest) (resp *StopWalletResponse, err error) {
	err = c.do("stop_wallet", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// RescanBlockchain Rescan the blockchain from scratch, losing any information which can not be recovered from the blockchain itself. This includes destination addresses, tx secret keys, tx notes, etc.
func (c *client) RescanBlockchain(req *RescanBlockchainRequest) (resp *RescanBlockchainResponse, err error) {
	err = c.do("rescan_blockchain", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetTxNotes Set arbitrary string notes for transactions.
func (c *client) SetTxNotes(req *SetTxNotesRequest) (resp *SetTxNotesResponse, err error) {
	err = c.do("set_tx_notes", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTxNotes Get string notes for transactions.
func (c *client) GetTxNotes(req *GetTxNotesRequest) (resp *GetTxNotesResponse, err error) {
	err = c.do("get_tx_notes", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetAttribute Set arbitrary attribute.
func (c *client) SetAttribute(req *SetAttributeRequest) (resp *SetAttributeResponse, err error) {
	err = c.do("set_attribute", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAttribute Get attribute value by name.
func (c *client) GetAttribute(req *GetAttributeRequest) (resp *GetAttributeResponse, err error) {
	err = c.do("get_attribute", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTxKey Get transaction secret key from transaction id.
func (c *client) GetTxKey(req *GetTxKeyRequest) (resp *GetTxKeyResponse, err error) {
	err = c.do("get_tx_key", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CheckTxKey Check a transaction in the blockchain with its secret key.
func (c *client) CheckTxKey(req *CheckTxKeyRequest) (resp *CheckTxKeyResponse, err error) {
	err = c.do("check_tx_key", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTxProof Get transaction signature to prove it.
func (c *client) GetTxProof(req *GetTxProofRequest) (resp *GetTxProofResponse, err error) {
	err = c.do("get_tx_proof", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CheckTxProof Prove a transaction by checking its signature.
func (c *client) CheckTxProof(req *CheckTxProofRequest) (resp *CheckTxProofResponse, err error) {
	err = c.do("check_tx_proof", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetSpendProof Generate a signature to prove a spend. Unlike proving a transaction, it does not requires the destination public address.
func (c *client) GetSpendProof(req *GetSpendProofRequest) (resp *GetSpendProofResponse, err error) {
	err = c.do("get_spend_proof", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CheckSpendProof Prove a spend using a signature. Unlike proving a transaction, it does not requires the destination public address.
func (c *client) CheckSpendProof(req *CheckSpendProofRequest) (resp *CheckSpendProofResponse, err error) {
	err = c.do("check_spend_proof", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetReserveProof Generate a signature to prove of an available amount in a wallet.
func (c *client) GetReserveProof(req *GetReserveProofRequest) (resp *GetReserveProofResponse, err error) {
	err = c.do("get_reserve_proof", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CheckReserveProof Proves a wallet has a disposable reserve using a signature.
func (c *client) CheckReserveProof(req *CheckReserveProofRequest) (resp *CheckReserveProofResponse, err error) {
	err = c.do("check_reserve_proof", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTransfers Returns a list of transfers.
func (c *client) GetTransfers(req *GetTransfersRequest) (resp *GetTransfersResponse, err error) {
	err = c.do("get_transfers", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTransferByTxid Show information about a transfer to/from this address.
func (c *client) GetTransferByTxid(req *GetTransferByTxidRequest) (resp *GetTransferByTxidResponse, err error) {
	err = c.do("get_transfer_by_txid", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// DescribeTransfer Returns details for each transaction in an unsigned or multisig transaction set. Transaction sets are obtained as return values from one of the following RPC methods:
func (c *client) DescribeTransfer(req *DescribeTransferRequest) (resp *DescribeTransferResponse, err error) {
	err = c.do("describe_transfer", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// Sign Sign a string.
func (c *client) Sign(req *SignRequest) (resp *SignResponse, err error) {
	err = c.do("sign", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// Verify Verify a signature on a string.
func (c *client) Verify(req *VerifyRequest) (resp *VerifyResponse, err error) {
	err = c.do("verify", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ExportOutputs Export all outputs in hex format.
func (c *client) ExportOutputs(req *ExportOutputsRequest) (resp *ExportOutputsResponse, err error) {
	err = c.do("export_outputs", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ImportOutputs Import outputs in hex format.
func (c *client) ImportOutputs(req *ImportOutputsRequest) (resp *ImportOutputsResponse, err error) {
	err = c.do("import_outputs", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ExportKeyImages Export a signed set of key images.
func (c *client) ExportKeyImages(req *ExportKeyImagesRequest) (resp *ExportKeyImagesResponse, err error) {
	err = c.do("export_key_images", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ImportKeyImages Import signed key images list and verify their spent status.
func (c *client) ImportKeyImages(req *ImportKeyImagesRequest) (resp *ImportKeyImagesResponse, err error) {
	err = c.do("import_key_images", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// MakeURI Create a payment URI using the official URI spec.
func (c *client) MakeURI(req *MakeURIRequest) (resp *MakeURIResponse, err error) {
	err = c.do("make_uri", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ParseURI Parse a payment URI to get payment information.
func (c *client) ParseURI(req *ParseURIRequest) (resp *ParseURIResponse, err error) {
	err = c.do("parse_uri", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAddressBook Retrieves entries from the address book.
func (c *client) GetAddressBook(req *GetAddressBookRequest) (resp *GetAddressBookResponse, err error) {
	err = c.do("get_address_book", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// AddAddressBook Add an entry to the address book.
func (c *client) AddAddressBook(req *AddAddressBookRequest) (resp *AddAddressBookResponse, err error) {
	err = c.do("add_address_book", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// EditAddressBook Edit an existing address book entry.
func (c *client) EditAddressBook(req *EditAddressBookRequest) (resp *EditAddressBookResponse, err error) {
	err = c.do("edit_address_book", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// DeleteAddressBook Delete an entry from the address book.
func (c *client) DeleteAddressBook(req *DeleteAddressBookRequest) (resp *DeleteAddressBookResponse, err error) {
	err = c.do("delete_address_book", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// Refresh Refresh a wallet after openning.
func (c *client) Refresh(req *RefreshRequest) (resp *RefreshResponse, err error) {
	err = c.do("refresh", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// AutoRefresh Set whether and how often to automatically refresh the current wallet.
func (c *client) AutoRefresh(req *AutoRefreshRequest) (resp *AutoRefreshResponse, err error) {
	err = c.do("auto_refresh", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// RescanSpent Rescan the blockchain for spent outputs.
func (c *client) RescanSpent(req *RescanSpentRequest) (resp *RescanSpentResponse, err error) {
	err = c.do("rescan_spent", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// StartMining Start mining in the Monero daemon.
func (c *client) StartMining(req *StartMiningRequest) (resp *StartMiningResponse, err error) {
	err = c.do("start_mining", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// StopMining Stop mining in the Monero daemon.
func (c *client) StopMining(req *StopMiningRequest) (resp *StopMiningResponse, err error) {
	err = c.do("stop_mining", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetLanguages Get a list of available languages for your wallet's seed.
func (c *client) GetLanguages(req *GetLanguagesRequest) (resp *GetLanguagesResponse, err error) {
	err = c.do("get_languages", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CreateWallet Create a new wallet. You need to have set the argument "âwallet-dir" when launching monero-wallet-rpc to make this work.
func (c *client) CreateWallet(req *CreateWalletRequest) (resp *CreateWalletResponse, err error) {
	err = c.do("create_wallet", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GenerateFromKeys Restores a wallet from a given wallet address, view key, and optional spend key.
func (c *client) GenerateFromKeys(req *GenerateFromKeysRequest) (resp *GenerateFromKeysResponse, err error) {
	err = c.do("generate_from_keys", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// OpenWallet Open a wallet. You need to have set the argument "âwallet-dir" when launching monero-wallet-rpc to make this work.
func (c *client) OpenWallet(req *OpenWalletRequest) (resp *OpenWalletResponse, err error) {
	err = c.do("open_wallet", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// RestoreDeterministicWallet Create and open a wallet on the RPC server from an existing mnemonic phrase and close the currently open wallet.
func (c *client) RestoreDeterministicWallet(req *RestoreDeterministicWalletRequest) (resp *RestoreDeterministicWalletResponse, err error) {
	err = c.do("restore_deterministic_wallet", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// CloseWallet Close the currently opened wallet, after trying to save it.
func (c *client) CloseWallet(req *CloseWalletRequest) (resp *CloseWalletResponse, err error) {
	err = c.do("close_wallet", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ChangeWalletPassword Change a wallet password.
func (c *client) ChangeWalletPassword(req *ChangeWalletPasswordRequest) (resp *ChangeWalletPasswordResponse, err error) {
	err = c.do("change_wallet_password", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// IsMultisig Check if a wallet is a multisig one.
func (c *client) IsMultisig(req *IsMultisigRequest) (resp *IsMultisigResponse, err error) {
	err = c.do("is_multisig", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// PrepareMultisig Prepare a wallet for multisig by generating a multisig string to share with peers.
func (c *client) PrepareMultisig(req *PrepareMultisigRequest) (resp *PrepareMultisigResponse, err error) {
	err = c.do("prepare_multisig", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// MakeMultisig Make a wallet multisig by importing peers multisig string.
func (c *client) MakeMultisig(req *MakeMultisigRequest) (resp *MakeMultisigResponse, err error) {
	err = c.do("make_multisig", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ExportMultisigInfo Export multisig info for other participants.
func (c *client) ExportMultisigInfo(req *ExportMultisigInfoRequest) (resp *ExportMultisigInfoResponse, err error) {
	err = c.do("export_multisig_info", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// ImportMultisigInfo Import multisig info from other participants.
func (c *client) ImportMultisigInfo(req *ImportMultisigInfoRequest) (resp *ImportMultisigInfoResponse, err error) {
	err = c.do("import_multisig_info", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// FinalizeMultisig Turn this wallet into a multisig wallet, extra step for N-1/N wallets.
func (c *client) FinalizeMultisig(req *FinalizeMultisigRequest) (resp *FinalizeMultisigResponse, err error) {
	err = c.do("finalize_multisig", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SignMultisig Sign a transaction in multisig.
func (c *client) SignMultisig(req *SignMultisigRequest) (resp *SignMultisigResponse, err error) {
	err = c.do("sign_multisig", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SubmitMultisig Submit a signed multisig transaction.
func (c *client) SubmitMultisig(req *SubmitMultisigRequest) (resp *SubmitMultisigResponse, err error) {
	err = c.do("submit_multisig", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetVersion Get RPC version Major & Minor integer-format, where Major is the first 16 bits and Minor the last 16 bits.
func (c *client) GetVersion(req *GetVersionRequest) (resp *GetVersionResponse, err error) {
	err = c.do("get_version", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

//#####################
// Other RPC Methods
//#####################
