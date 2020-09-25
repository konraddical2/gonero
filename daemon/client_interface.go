package daemon

// Client is a monerod daemon RPC client
type Client interface {
	// JSON RPC Methods
	// Look up how many blocks are in the longest chain known to the node.
	GetBlockCount(*GetBlockCountRequest) (*GetBlockCountResponse, error)
	// Look up a block's hash by its height.
	OnGetBlockHash(*OnGetBlockHashRequest) (*OnGetBlockHashResponse, error)
	// Get a block template on which mining a new block.
	GetBlockTemplate(*GetBlockTemplateRequest) (*GetBlockTemplateResponse, error)
	// Submit a mined block to the network.
	SubmitBlock(*SubmitBlockRequest) (*SubmitBlockResponse, error)
	// Block header information for the most recent block is easily retrieved with this method. No inputs are needed.
	GetLastBlockHeader(*GetLastBlockHeaderRequest) (*GetLastBlockHeaderResponse, error)
	// Block header information can be retrieved using either a block's hash or height. This method includes a block's hash as an input parameter to retrieve basic information about the block.
	GetBlockHeaderByHash(*GetBlockHeaderByHashRequest) (*GetBlockHeaderByHashResponse, error)
	// Similar to get_block_header_by_hash above, this method includes a block's height as an input parameter to retrieve basic information about the block.
	GetBlockHeaderByHeight(*GetBlockHeaderByHeightRequest) (*GetBlockHeaderByHeightResponse, error)
	// Similar to get_block_header_by_height above, but for a range of blocks. This method includes a starting block height and an ending block height as parameters to retrieve basic information about the range of blocks.
	GetBlockHeadersRange(*GetBlockHeadersRangeRequest) (*GetBlockHeadersRangeResponse, error)
	// Full block information can be retrieved by either block height or hash, like with the above block header calls. For full block information, both lookups use the same method, but with different input parameters.
	GetBlock(*GetBlockRequest) (*GetBlockResponse, error)
	// Retrieve information about incoming and outgoing connections to your node.
	GetConnections(*GetConnectionsRequest) (*GetConnectionsResponse, error)
	// Retrieve general information about the state of your node and the network.
	GetInfo(*GetInfoRequest) (*GetInfoResponse, error)
	// Look up information regarding hard fork voting and readiness.
	HardForkInfo(*HardForkInfoRequest) (*HardForkInfoResponse, error)
	// Ban another node by IP.
	SetBans(*SetBansRequest) (*SetBansResponse, error)
	// Get list of banned IPs.
	GetBans(*GetBansRequest) (*GetBansResponse, error)
	// Flush tx ids from transaction pool
	FlushTxpool(*FlushTxpoolRequest) (*FlushTxpoolResponse, error)
	// Get a histogram of output amounts. For all amounts (possibly filtered by parameters), gives the number of outputs on the chain for that amount. RingCT outputs counts as 0 amount.
	GetOutputHistogram(*GetOutputHistogramRequest) (*GetOutputHistogramResponse, error)
	// Give the node current version
	GetVersion(*GetVersionRequest) (*GetVersionResponse, error)
	// Get the coinbase amount and the fees amount for n last blocks starting at particular height
	GetCoinbaseTxSum(*GetCoinbaseTxSumRequest) (*GetCoinbaseTxSumResponse, error)
	// Gives an estimation on fees per byte.
	GetFeeEstimate(*GetFeeEstimateRequest) (*GetFeeEstimateResponse, error)
	// Display alternative chains seen by the node.
	GetAlternateChains(*GetAlternateChainsRequest) (*GetAlternateChainsResponse, error)
	// Relay a list of transaction IDs.
	RelayTx(*RelayTxRequest) (*RelayTxResponse, error)
	// Get synchronisation informations
	SyncInfo(*SyncInfoRequest) (*SyncInfoResponse, error)
	// Get all transaction pool backlog
	GetTxpoolBacklog(*GetTxpoolBacklogRequest) (*GetTxpoolBacklogResponse, error)
	// None
	GetOutputDistribution(*GetOutputDistributionRequest) (*GetOutputDistributionResponse, error)

	// Other RPC Methods
	// Get the node's current height.
	GetHeight(*GetHeightRequest) (*GetHeightResponse, error)
	// Look up one or more transactions by hash.
	GetTransactions(*GetTransactionsRequest) (*GetTransactionsResponse, error)
	// Get the known blocks hashes which are not on the main chain.
	GetAltBlocksHashes(*GetAltBlocksHashesRequest) (*GetAltBlocksHashesResponse, error)
	// Check if outputs have been spent using the key image associated with the output.
	IsKeyImageSpent(*IsKeyImageSpentRequest) (*IsKeyImageSpentResponse, error)
	// Broadcast a raw transaction to the network.
	SendRawTransaction(*SendRawTransactionRequest) (*SendRawTransactionResponse, error)
	// Start mining on the daemon.
	StartMining(*StartMiningRequest) (*StartMiningResponse, error)
	// Stop mining on the daemon.
	StopMining(*StopMiningRequest) (*StopMiningResponse, error)
	// Get the mining status of the daemon.
	MiningStatus(*MiningStatusRequest) (*MiningStatusResponse, error)
	// Save the blockchain. The blockchain does not need saving and is always saved when modified, however it does a sync to flush the filesystem cache onto the disk for safety purposes against Operating System or Harware crashes.
	SaveBc(*SaveBcRequest) (*SaveBcResponse, error)
	// Get the known peers list.
	GetPeerList(*GetPeerListRequest) (*GetPeerListResponse, error)
	// Set the log hash rate display mode.
	SetLogHashRate(*SetLogHashRateRequest) (*SetLogHashRateResponse, error)
	// Set the daemon log level. By default, log level is set to 0.
	SetLogLevel(*SetLogLevelRequest) (*SetLogLevelResponse, error)
	// Set the daemon log categories. Categories are represented as a comma separated list of <Category>:<level> (similarly to syslog standard <Facility>:<Severity-level>), where:
	SetLogCategories(*SetLogCategoriesRequest) (*SetLogCategoriesResponse, error)
	// Show information about valid transactions seen by the node but not yet mined into a block, as well as spent key image information for the txpool in the node's memory.
	GetTransactionPool(*GetTransactionPoolRequest) (*GetTransactionPoolResponse, error)
	// Get hashes from transaction pool. Binary request.
	GetTransactionPoolHashesBin(*GetTransactionPoolHashesBinRequest) (*GetTransactionPoolHashesBinResponse, error)
	// Get the transaction pool statistics.
	GetTransactionPoolStats(*GetTransactionPoolStatsRequest) (*GetTransactionPoolStatsResponse, error)
	// Send a command to the daemon to safely disconnect and shut down.
	StopDaemon(*StopDaemonRequest) (*StopDaemonResponse, error)
	// Get daemon bandwidth limits.
	GetLimit(*GetLimitRequest) (*GetLimitResponse, error)
	// Set daemon bandwidth limits.
	SetLimit(*SetLimitRequest) (*SetLimitResponse, error)
	// Limit number of Outgoing peers.
	OutPeers(*OutPeersRequest) (*OutPeersResponse, error)
	// Limit number of Incoming peers.
	InPeers(*InPeersRequest) (*InPeersResponse, error)
	// Get outputs.
	GetOuts(*GetOutsRequest) (*GetOutsResponse, error)
	// Update daemon.
	Update(*UpdateRequest) (*UpdateResponse, error)
	// Generate blocks in Regtest mode
	GenerateBlocks(*GenerateBlocksRequest) (*GenerateBlocksResponse, error)
}
