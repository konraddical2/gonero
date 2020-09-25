package daemon

//#####################
// General RPC Methods
//#####################

//#####################
// JSON RPC Methods
//#####################

// GetBlockCount Look up how many blocks are in the longest chain known to the node.
func (c *client) GetBlockCount(req *GetBlockCountRequest) (resp *GetBlockCountResponse, err error) {
	err = c.do("get_block_count", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// OnGetBlockHash Look up a block's hash by its height.
func (c *client) OnGetBlockHash(req *OnGetBlockHashRequest) (resp *OnGetBlockHashResponse, err error) {
	err = c.do("on_get_block_hash", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBlockTemplate Get a block template on which mining a new block.
func (c *client) GetBlockTemplate(req *GetBlockTemplateRequest) (resp *GetBlockTemplateResponse, err error) {
	err = c.do("get_block_template", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SubmitBlock Submit a mined block to the network.
func (c *client) SubmitBlock(req *SubmitBlockRequest) (resp *SubmitBlockResponse, err error) {
	err = c.do("submit_block", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetLastBlockHeader Block header information for the most recent block is easily retrieved with this method. No inputs are needed.
func (c *client) GetLastBlockHeader(req *GetLastBlockHeaderRequest) (resp *GetLastBlockHeaderResponse, err error) {
	err = c.do("get_last_block_header", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBlockHeaderByHash Block header information can be retrieved using either a block's hash or height. This method includes a block's hash as an input parameter to retrieve basic information about the block.
func (c *client) GetBlockHeaderByHash(req *GetBlockHeaderByHashRequest) (resp *GetBlockHeaderByHashResponse, err error) {
	err = c.do("get_block_header_by_hash", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBlockHeaderByHeight Similar to get_block_header_by_hash above, this method includes a block's height as an input parameter to retrieve basic information about the block.
func (c *client) GetBlockHeaderByHeight(req *GetBlockHeaderByHeightRequest) (resp *GetBlockHeaderByHeightResponse, err error) {
	err = c.do("get_block_header_by_height", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBlockHeadersRange Similar to get_block_header_by_height above, but for a range of blocks. This method includes a starting block height and an ending block height as parameters to retrieve basic information about the range of blocks.
func (c *client) GetBlockHeadersRange(req *GetBlockHeadersRangeRequest) (resp *GetBlockHeadersRangeResponse, err error) {
	err = c.do("get_block_headers_range", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBlock Full block information can be retrieved by either block height or hash, like with the above block header calls. For full block information, both lookups use the same method, but with different input parameters.
func (c *client) GetBlock(req *GetBlockRequest) (resp *GetBlockResponse, err error) {
	err = c.do("get_block", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetConnections Retrieve information about incoming and outgoing connections to your node.
func (c *client) GetConnections(req *GetConnectionsRequest) (resp *GetConnectionsResponse, err error) {
	err = c.do("get_connections", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetInfo Retrieve general information about the state of your node and the network.
func (c *client) GetInfo(req *GetInfoRequest) (resp *GetInfoResponse, err error) {
	err = c.do("get_info", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// HardForkInfo Look up information regarding hard fork voting and readiness.
func (c *client) HardForkInfo(req *HardForkInfoRequest) (resp *HardForkInfoResponse, err error) {
	err = c.do("hard_fork_info", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetBans Ban another node by IP.
func (c *client) SetBans(req *SetBansRequest) (resp *SetBansResponse, err error) {
	err = c.do("set_bans", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetBans Get list of banned IPs.
func (c *client) GetBans(req *GetBansRequest) (resp *GetBansResponse, err error) {
	err = c.do("get_bans", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// FlushTxpool Flush tx ids from transaction pool
func (c *client) FlushTxpool(req *FlushTxpoolRequest) (resp *FlushTxpoolResponse, err error) {
	err = c.do("flush_txpool", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetOutputHistogram Get a histogram of output amounts. For all amounts (possibly filtered by parameters), gives the number of outputs on the chain for that amount. RingCT outputs counts as 0 amount.
func (c *client) GetOutputHistogram(req *GetOutputHistogramRequest) (resp *GetOutputHistogramResponse, err error) {
	err = c.do("get_output_histogram", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetVersion Give the node current version
func (c *client) GetVersion(req *GetVersionRequest) (resp *GetVersionResponse, err error) {
	err = c.do("get_version", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetCoinbaseTxSum Get the coinbase amount and the fees amount for n last blocks starting at particular height
func (c *client) GetCoinbaseTxSum(req *GetCoinbaseTxSumRequest) (resp *GetCoinbaseTxSumResponse, err error) {
	err = c.do("get_coinbase_tx_sum", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetFeeEstimate Gives an estimation on fees per byte.
func (c *client) GetFeeEstimate(req *GetFeeEstimateRequest) (resp *GetFeeEstimateResponse, err error) {
	err = c.do("get_fee_estimate", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAlternateChains Display alternative chains seen by the node.
func (c *client) GetAlternateChains(req *GetAlternateChainsRequest) (resp *GetAlternateChainsResponse, err error) {
	err = c.do("get_alternate_chains", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// RelayTx Relay a list of transaction IDs.
func (c *client) RelayTx(req *RelayTxRequest) (resp *RelayTxResponse, err error) {
	err = c.do("relay_tx", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SyncInfo Get synchronisation informations
func (c *client) SyncInfo(req *SyncInfoRequest) (resp *SyncInfoResponse, err error) {
	err = c.do("sync_info", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTxpoolBacklog Get all transaction pool backlog
func (c *client) GetTxpoolBacklog(req *GetTxpoolBacklogRequest) (resp *GetTxpoolBacklogResponse, err error) {
	err = c.do("get_txpool_backlog", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetOutputDistribution None
func (c *client) GetOutputDistribution(req *GetOutputDistributionRequest) (resp *GetOutputDistributionResponse, err error) {
	err = c.do("get_output_distribution", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

//#####################
// Other RPC Methods
//#####################

// GetHeight Get the node's current height.
func (c *client) GetHeight(req *GetHeightRequest) (resp *GetHeightResponse, err error) {
	err = c.doSlash("/get_height", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTransactions Look up one or more transactions by hash.
func (c *client) GetTransactions(req *GetTransactionsRequest) (resp *GetTransactionsResponse, err error) {
	err = c.doSlash("/get_transactions", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAltBlocksHashes Get the known blocks hashes which are not on the main chain.
func (c *client) GetAltBlocksHashes(req *GetAltBlocksHashesRequest) (resp *GetAltBlocksHashesResponse, err error) {
	err = c.doSlash("/get_alt_blocks_hashes", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// IsKeyImageSpent Check if outputs have been spent using the key image associated with the output.
func (c *client) IsKeyImageSpent(req *IsKeyImageSpentRequest) (resp *IsKeyImageSpentResponse, err error) {
	err = c.doSlash("/is_key_image_spent", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SendRawTransaction Broadcast a raw transaction to the network.
func (c *client) SendRawTransaction(req *SendRawTransactionRequest) (resp *SendRawTransactionResponse, err error) {
	err = c.doSlash("/send_raw_transaction", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// StartMining Start mining on the daemon.
func (c *client) StartMining(req *StartMiningRequest) (resp *StartMiningResponse, err error) {
	err = c.doSlash("/start_mining", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// StopMining Stop mining on the daemon.
func (c *client) StopMining(req *StopMiningRequest) (resp *StopMiningResponse, err error) {
	err = c.doSlash("/stop_mining", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// MiningStatus Get the mining status of the daemon.
func (c *client) MiningStatus(req *MiningStatusRequest) (resp *MiningStatusResponse, err error) {
	err = c.doSlash("/mining_status", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SaveBc Save the blockchain. The blockchain does not need saving and is always saved when modified, however it does a sync to flush the filesystem cache onto the disk for safety purposes against Operating System or Harware crashes.
func (c *client) SaveBc(req *SaveBcRequest) (resp *SaveBcResponse, err error) {
	err = c.doSlash("/save_bc", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetPeerList Get the known peers list.
func (c *client) GetPeerList(req *GetPeerListRequest) (resp *GetPeerListResponse, err error) {
	err = c.doSlash("/get_peer_list", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetLogHashRate Set the log hash rate display mode.
func (c *client) SetLogHashRate(req *SetLogHashRateRequest) (resp *SetLogHashRateResponse, err error) {
	err = c.doSlash("/set_log_hash_rate", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetLogLevel Set the daemon log level. By default, log level is set to 0.
func (c *client) SetLogLevel(req *SetLogLevelRequest) (resp *SetLogLevelResponse, err error) {
	err = c.doSlash("/set_log_level", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetLogCategories Set the daemon log categories. Categories are represented as a comma separated list of <Category>:<level> (similarly to syslog standard <Facility>:<Severity-level>), where:
func (c *client) SetLogCategories(req *SetLogCategoriesRequest) (resp *SetLogCategoriesResponse, err error) {
	err = c.doSlash("/set_log_categories", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTransactionPool Show information about valid transactions seen by the node but not yet mined into a block, as well as spent key image information for the txpool in the node's memory.
func (c *client) GetTransactionPool(req *GetTransactionPoolRequest) (resp *GetTransactionPoolResponse, err error) {
	err = c.doSlash("/get_transaction_pool", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTransactionPoolHashesBin Get hashes from transaction pool. Binary request.
func (c *client) GetTransactionPoolHashesBin(req *GetTransactionPoolHashesBinRequest) (resp *GetTransactionPoolHashesBinResponse, err error) {
	err = c.doSlash("/get_transaction_pool_hashes.bin", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetTransactionPoolStats Get the transaction pool statistics.
func (c *client) GetTransactionPoolStats(req *GetTransactionPoolStatsRequest) (resp *GetTransactionPoolStatsResponse, err error) {
	err = c.doSlash("/get_transaction_pool_stats", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// StopDaemon Send a command to the daemon to safely disconnect and shut down.
func (c *client) StopDaemon(req *StopDaemonRequest) (resp *StopDaemonResponse, err error) {
	err = c.doSlash("/stop_daemon", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetLimit Get daemon bandwidth limits.
func (c *client) GetLimit(req *GetLimitRequest) (resp *GetLimitResponse, err error) {
	err = c.doSlash("/get_limit", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// SetLimit Set daemon bandwidth limits.
func (c *client) SetLimit(req *SetLimitRequest) (resp *SetLimitResponse, err error) {
	err = c.doSlash("/set_limit", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// OutPeers Limit number of Outgoing peers.
func (c *client) OutPeers(req *OutPeersRequest) (resp *OutPeersResponse, err error) {
	err = c.doSlash("/out_peers", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// InPeers Limit number of Incoming peers.
func (c *client) InPeers(req *InPeersRequest) (resp *InPeersResponse, err error) {
	err = c.doSlash("/in_peers", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetOuts Get outputs.
func (c *client) GetOuts(req *GetOutsRequest) (resp *GetOutsResponse, err error) {
	err = c.doSlash("/get_outs", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// Update Update daemon.
func (c *client) Update(req *UpdateRequest) (resp *UpdateResponse, err error) {
	err = c.doSlash("/update", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}

//#####################
// Regtest RPC Methods
//#####################
// GenerateBlocks mines a given number of blocs for the given address
// NOTE: This is only a valid command in Regtest mode!
func (c *client) GenerateBlocks(req *GenerateBlocksRequest) (resp *GenerateBlocksResponse, err error) {
	err = c.do("generateblocks", &req, &resp)
	if err != nil {
		return nil, err
	}
	return
}
