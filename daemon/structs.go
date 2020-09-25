package daemon

import (
	"github.com/konraddical2/gonero"
)

//#####################
// General RPC Structs
//#####################

// BlockHeader A structure containing block header information.
type BlockHeader struct {
	// The block size in bytes.
	BlockSize uint64 `json:"block_size"`
	// The number of blocks succeeding this block on the blockchain. A larger number means an older block.
	Depth uint64 `json:"depth"`
	// The strength of the Monero network based on mining power.
	Difficulty uint64 `json:"difficulty"`
	// The hash of this block.
	Hash string `json:"hash"`
	// The number of blocks preceding this block on the blockchain.
	Height uint64 `json:"height"`
	// The major version of the monero protocol at this block height.
	MajorVersion uint64 `json:"major_version"`
	// The minor version of the monero protocol at this block height.
	MinorVersion uint64 `json:"minor_version"`
	// a cryptographic random one-time number used in mining a Monero block.
	Nonce uint64 `json:"nonce"`
	// Number of transactions in the block, not counting the coinbase tx.
	NumTxes uint64 `json:"num_txes"`
	// Usually false. If true, this block is not part of the longest chain.
	OrphanStatus bool `json:"orphan_status"`
	// The hash of the block immediately preceding this block in the chain.
	PrevHash string `json:"prev_hash"`
	// The amount of new atomic units generated in this block and rewarded to the miner. Note: 1 XMR = 1e12 atomic units.
	Reward uint64 `json:"reward"`
	// The unix time at which the block was recorded into the blockchain.
	Timestamp uint64 `json:"timestamp"`
}

//#####################
// JSON RPC Structs
//#####################

// GetBlockCountRequest is a struct for GetBlockCount() requests
type GetBlockCountRequest struct {
	// None
}

// GetBlockCountResponse is a struct for GetBlockCount() responses
type GetBlockCountResponse struct {
	// Number of blocks in longest chain seen by the node.
	Count uint64 `json:"count"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// OnGetBlockHashRequest is a struct for OnGetBlockHash() requests
// block height (int array of length 1)
type OnGetBlockHashRequest [1]uint64

// OnGetBlockHashResponse is a struct for OnGetBlockHash() responses
// Block hash
type OnGetBlockHashResponse string

// GetBlockTemplateRequest is a struct for GetBlockTemplate() requests
type GetBlockTemplateRequest struct {
	// Address of wallet to receive coinbase transactions if block is successfully mined.
	WalletAddress string `json:"wallet_address"`
	// Reserve size.
	ReserveSize uint64 `json:"reserve_size"`
}

// GetBlockTemplateResponse is a struct for GetBlockTemplate() responses
type GetBlockTemplateResponse struct {
	// Blob on which to try to mine a new block.
	BlocktemplateBlob string `json:"blocktemplate_blob"`
	// Blob on which to try to find a valid nonce.
	BlockhashingBlob string `json:"blockhashing_blob"`
	// Difficulty of next block.
	Difficulty uint64 `json:"difficulty"`
	// Coinbase reward expected to be received if block is successfully mined.
	ExpectedReward uint64 `json:"expected_reward"`
	// Height on which to mine.
	Height uint64 `json:"height"`
	// Hash of the most recent block on which to mine the next block.
	PrevHash string `json:"prev_hash"`
	// Reserved offset.
	ReservedOffset uint64 `json:"reserved_offset"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// SubmitBlockRequest is a struct for SubmitBlock() requests
// list of block blobs which have been mined.  See get_block_template to get a blob on which to mine.
type SubmitBlockRequest []string

// SubmitBlockResponse is a struct for SubmitBlock() responses
type SubmitBlockResponse struct {
	// Block submit status.
	Status string `json:"status"`
}

// GetLastBlockHeaderRequest is a struct for GetLastBlockHeader() requests
type GetLastBlockHeaderRequest struct {
	// None
}

// GetLastBlockHeaderResponse is a struct for GetLastBlockHeader() responses
type GetLastBlockHeaderResponse struct {
	// A structure containing block header information.
	BlockHeader BlockHeader `json:"block_header"`
	// The block size in bytes.
	BlockSize uint64 `json:"block_size"`
	// The number of blocks succeeding this block on the blockchain. A larger number means an older block.
	Depth uint64 `json:"depth"`
	// The strength of the Monero network based on mining power.
	Difficulty uint64 `json:"difficulty"`
	// The hash of this block.
	Hash string `json:"hash"`
	// The number of blocks preceding this block on the blockchain.
	Height uint64 `json:"height"`
	// The major version of the monero protocol at this block height.
	MajorVersion uint64 `json:"major_version"`
	// The minor version of the monero protocol at this block height.
	MinorVersion uint64 `json:"minor_version"`
	// a cryptographic random one-time number used in mining a Monero block.
	Nonce uint64 `json:"nonce"`
	// Number of transactions in the block, not counting the coinbase tx.
	NumTxes uint64 `json:"num_txes"`
	// Usually false. If true, this block is not part of the longest chain.
	OrphanStatus bool `json:"orphan_status"`
	// The hash of the block immediately preceding this block in the chain.
	PrevHash string `json:"prev_hash"`
	// The amount of new atomic units generated in this block and rewarded to the miner. Note: 1 XMR = 1e12 atomic units.
	Reward uint64 `json:"reward"`
	// The unix time at which the block was recorded into the blockchain.
	Timestamp uint64 `json:"timestamp"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetBlockHeaderByHashRequest is a struct for GetBlockHeaderByHash() requests
type GetBlockHeaderByHashRequest struct {
	// The block's sha256 hash.
	Hash string `json:"hash"`
}

// GetBlockHeaderByHashResponse is a struct for GetBlockHeaderByHash() responses
type GetBlockHeaderByHashResponse struct {
	// A structure containing block header information.
	BlockHeader BlockHeader `json:"block_header"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetBlockHeaderByHeightRequest is a struct for GetBlockHeaderByHeight() requests
type GetBlockHeaderByHeightRequest struct {
	// The block's height.
	Height uint64 `json:"height"`
}

// GetBlockHeaderByHeightResponse is a struct for GetBlockHeaderByHeight() responses
type GetBlockHeaderByHeightResponse struct {
	// A structure containing block header information.
	BlockHeader BlockHeader `json:"block_header"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetBlockHeadersRangeRequest is a struct for GetBlockHeadersRange() requests
type GetBlockHeadersRangeRequest struct {
	// The starting block's height.
	StartHeight uint64 `json:"start_height"`
	// The ending block's height.
	EndHeight uint64 `json:"end_height"`
}

// GetBlockHeadersRangeResponse is a struct for GetBlockHeadersRange() responses
type GetBlockHeadersRangeResponse struct {
	// array of block_header (a structure containing block header information. See get_last_block_header).
	Headers []BlockHeader `json:"headers"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetBlockRequest is a struct for GetBlock() requests
type GetBlockRequest struct {
	// The block's height.
	Height uint64 `json:"height"`
	// The block's hash.
	Hash string `json:"hash"`
}

// GetBlockResponse is a struct for GetBlock() responses
type GetBlockResponse struct {
	// Hexadecimal blob of block information.
	Blob string `json:"blob"`
	// A structure containing block header information.
	BlockHeader BlockHeader `json:"block_header"`
	// JSON formatted block details
	JSON string `json:"json"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode,
	// and is therefore not trusted (true),
	// or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// JSONBlock XXX?
type JSONBlock struct {
	// Same as in block header.
	MajorVersion uint64 `json:"major_version"`
	// Same as in block header.
	MinorVersion uint64 `json:"minor_version"`
	// Same as in block header.
	Timestamp uint64 `json:"timestamp"`
	// Same as prev_hash in block header.
	PrevID uint64 `json:"prev_id"`
	// Same as in block header.
	Nonce uint64 `json:"nonce"`
	// Miner Transaction information
	MinerTx MinerTransaction `json:"miner_tx"`
	// List of hashes of non coinbase transactions in the block.
	// If there are no other transactions, this will be an empty list.
	TxHashes []string `json:"tx_hashes"`
}

// MinerTransaction contains information about a miner transaction
type MinerTransaction struct {
	// Transaction version number.
	Version uint64 `json:"version"`
	// The block height when the coinbase transaction becomes spendable.
	UnlockTime uint64 `json:"unlock_time"`
	// List of transaction inputs
	Vin []MinerTransactionInput `json:"vin"`
	// List of transaction outputs
	Vout []MinerTransactionOutput `json:"vout"`
	// Usually called the "transaction ID" but can be used to include any random 32 byte/64 character hex string.
	Extra string `json:"extra"`
	// Contain signatures of tx signers. Coinbased txs do not have signatures.
	Signatures []string `json:"signatures"`
}

// MinerTransactionInput is a struct that contains information about inputs to a mining transaction
// TODO XXX test the nesting of height
type MinerTransactionInput struct {
	// Miner txs are coinbase txs, or "gen".
	Gen string `json:"gen"`
	// This block height, a.k.a. when the coinbase is generated.
	Height uint64 `json:"height"`
}

// MinerTransactionOutput is a struct that contains information about ouputs of a mining transaction
// TODO XXX test the nesting of Key
type MinerTransactionOutput struct {
	// The amount of the output, in atomic units.
	Amount gonero.AtomicXMR `json:"amount"`
	// --
	Target string `json:"target"`
	// --
	Key string `json:"key"`
}

// GetConnectionsRequest is a struct for GetConnections() requests
type GetConnectionsRequest struct {
	// None
}

// GetConnectionsResponse is a struct for GetConnections() responses
type GetConnectionsResponse struct {
	// List of all connections and their info
	Connections []Connection `json:"connections"`
}

// Connection is
type Connection struct {
	// The peer's address, actually IPv4 & port
	Address string `json:"address"`
	// Average bytes of data downloaded by node.
	AvgDownload uint64 `json:"avg_download"`
	// Average bytes of data uploaded by node.
	AvgUpload uint64 `json:"avg_upload"`
	// The connection ID
	ConnectionID string `json:"connection_id"`
	// Current bytes downloaded by node.
	CurrentDownload uint64 `json:"current_download"`
	// Current bytes uploaded by node.
	CurrentUpload uint64 `json:"current_upload"`
	// The peer height
	Height uint64 `json:"height"`
	// The peer host
	Host string `json:"host"`
	// Is the node getting information from your node?
	Incoming bool `json:"incoming"`
	// The node's IP address.
	IP string `json:"ip"`
	// unsigned int
	LiveTime uint64 `json:"live_time"`
	// boolean
	LocalIP bool `json:"local_ip"`
	// boolean
	Localhost bool `json:"localhost"`
	// The node's ID on the network.
	PeerID string `json:"peer_id"`
	// The port that the node is using to connect to the network.
	Port string `json:"port"`
	// unsigned int
	RecvCount uint64 `json:"recv_count"`
	// unsigned int
	RecvIdleTime uint64 `json:"recv_idle_time"`
	// unsigned int
	SendCount uint64 `json:"send_count"`
	// unsigned int
	SendIdleTime uint64 `json:"send_idle_time"`
	// string
	State string `json:"state"`
	// unsigned int
	SupportFlags uint64 `json:"support_flags"`
}

// GetInfoRequest is a struct for GetInfo() requests
type GetInfoRequest struct {
	// None
}

// GetInfoResponse is a struct for GetInfo() responses
type GetInfoResponse struct {
	// Number of alternative blocks to main chain.
	AltBlocksCount uint64 `json:"alt_blocks_count"`
	// Maximum allowed block size
	BlockSizeLimit uint64 `json:"block_size_limit"`
	// Median block size of latest 100 blocks
	BlockSizeMedian uint64 `json:"block_size_median"`
	// bootstrap node to give immediate usability to wallets while syncing by proxying RPC to it. (Note: the replies may be untrustworthy).
	BootstrapDaemonAddress string `json:"bootstrap_daemon_address"`
	// Cumulative difficulty of all blocks in the blockchain.
	CumulativeDifficulty uint64 `json:"cumulative_difficulty"`
	// Network difficulty (analogous to the strength of the network)
	Difficulty uint64 `json:"difficulty"`
	// Available disk space on the node.
	FreeSpace uint64 `json:"free_space"`
	// Grey Peerlist Size
	GreyPeerlistSize uint64 `json:"grey_peerlist_size"`
	// Current length of longest chain known to daemon.
	Height uint64 `json:"height"`
	// Current length of the local chain of the daemon.
	HeightWithoutBootstrap uint64 `json:"height_without_bootstrap"`
	// Number of peers connected to and pulling from your node.
	IncomingConnectionsCount uint64 `json:"incoming_connections_count"`
	// States if the node is on the mainnet (true) or not (false).
	Mainnet bool `json:"mainnet"`
	// States if the node is offline (true) or online (false).
	Offline bool `json:"offline"`
	// Number of peers that you are connected to and getting information from.
	OutgoingConnectionsCount uint64 `json:"outgoing_connections_count"`
	// Number of RPC client connected to the daemon (Including this RPC request).
	RPCConnectionsCount uint64 `json:"rpc_connections_count"`
	// States if the node is on the stagenet (true) or not (false).
	Stagenet bool `json:"stagenet"`
	// Start time of the daemon, as UNIX time.
	StartTime uint64 `json:"start_time"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// Current target for next proof of work.
	Target uint64 `json:"target"`
	// The height of the next block in the chain.
	TargetHeight uint64 `json:"target_height"`
	// States if the node is on the testnet (true) or not (false).
	Testnet bool `json:"testnet"`
	// Hash of the highest block in the chain.
	TopBlockHash string `json:"top_block_hash"`
	// Total number of non-coinbase transaction in the chain.
	TxCount uint64 `json:"tx_count"`
	// Number of transactions that have been broadcast but not included in a block.
	TxPoolSize uint64 `json:"tx_pool_size"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
	// States if a bootstrap node has ever been used since the daemon started.
	WasBootstrapEverUsed bool `json:"was_bootstrap_ever_used"`
	// White Peerlist Size
	WhitePeerlistSize uint64 `json:"white_peerlist_size"`
}

// HardForkInfoRequest is a struct for HardForkInfo() requests
type HardForkInfoRequest struct {
	// None
}

// HardForkInfoResponse is a struct for HardForkInfo() responses
type HardForkInfoResponse struct {
	// Block height at which hard fork would be enabled if voted in.
	EarliestHeight uint64 `json:"earliest_height"`
	// Tells if hard fork is enforced.
	Enabled bool `json:"enabled"`
	// Current hard fork state: 0 (There is likely a hard fork), 1 (An update is needed to fork properly), or 2 (Everything looks good).
	State uint64 `json:"state"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// Minimum percent of votes to trigger hard fork. Default is 80.
	Threshold uint64 `json:"threshold"`
	// The major block version for the fork.
	Version uint64 `json:"version"`
	// Number of votes towards hard fork.
	Votes uint64 `json:"votes"`
	// Hard fork voting status.
	Voting uint64 `json:"voting"`
	// Number of blocks over which current votes are cast. Default is 10080 blocks.
	Window uint64 `json:"window"`
}

// SetBansRequest is a struct for SetBans() requests
type SetBansRequest struct {
	// A list of nodes to ban
	Bans []Ban `json:"bans"`
}

// SetBansResponse is a struct for SetBans() responses
type SetBansResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// Ban is a struct that contains the banning status of a node
type Ban struct {
	// Host to ban (IP in A.B.C.D form - will support I2P address in the future).
	Host string `json:"host"`
	// IP address to ban, in Int format.
	IP uint64 `json:"ip"`
	// Set true to ban.
	Ban bool `json:"ban"`
	// Number of seconds to ban node.
	Seconds uint64 `json:"seconds"`
}

// GetBansRequest is a struct for GetBans() requests
type GetBansRequest struct {
	// None
}

// GetBansResponse is a struct for GetBans() responses
type GetBansResponse struct {
	// List of banned nodes
	Bans []Ban `json:"bans"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// FlushTxpoolRequest is a struct for FlushTxpool() requests
type FlushTxpoolRequest struct {
	// Optional, list of transactions IDs to flush from pool (all tx ids flushed if empty).
	Txids []string `json:"txids,omitempty"`
}

// FlushTxpoolResponse is a struct for FlushTxpool() responses
type FlushTxpoolResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// GetOutputHistogramRequest is a struct for GetOutputHistogram() requests
type GetOutputHistogramRequest struct {
	// list of unsigned int
	Amounts []uint64 `json:"amounts"`
	// unsigned int
	MinCount uint64 `json:"min_count"`
	// unsigned int
	MaxCount uint64 `json:"max_count"`
	// boolean
	Unlocked bool `json:"unlocked"`
	// unsigned int
	RecentCutoff uint64 `json:"recent_cutoff"`
}

// GetOutputHistogramResponse is a struct for GetOutputHistogram() responses
type GetOutputHistogramResponse struct {
	// list of histogram entries
	Histogram []Histogram `json:"histogram"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// Histogram is a histogram of output amounts.
// For all amounts (possibly filtered by parameters),
// gives the number of outputs on the chain for that amount.
// RingCT outputs counts as 0 amount.
type Histogram struct {
	// Output amount in atomic units
	Amount gonero.AtomicXMR `json:"amount"`
	// unsigned int
	TotalInstances uint64 `json:"total_instances"`
	// unsigned int
	UnlockedInstances uint64 `json:"unlocked_instances"`
	// unsigned int
	RecentInstances uint64 `json:"recent_instances"`
}

// GetVersionRequest is a struct for GetVersion() requests
type GetVersionRequest struct {
	// None
}

// GetVersionResponse is a struct for GetVersion() responses
type GetVersionResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
	// unsigned int
	Version uint64 `json:"version"`
}

// GetCoinbaseTxSumRequest is a struct for GetCoinbaseTxSum() requests
type GetCoinbaseTxSumRequest struct {
	// Block height from which getting the amounts
	Height uint64 `json:"height"`
	// number of blocks to include in the sum
	Count uint64 `json:"count"`
}

// GetCoinbaseTxSumResponse is a struct for GetCoinbaseTxSum() responses
type GetCoinbaseTxSumResponse struct {
	// amount of coinbase reward in atomic units
	EmissionAmount gonero.AtomicXMR `json:"emission_amount"`
	// amount of fees in atomic units
	FeeAmount gonero.AtomicXMR `json:"fee_amount"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// GetFeeEstimateRequest is a struct for GetFeeEstimate() requests
type GetFeeEstimateRequest struct {
	// Optional
	GraceBlocks uint64 `json:"grace_blocks,omitempty"`
}

// GetFeeEstimateResponse is a struct for GetFeeEstimate() responses
type GetFeeEstimateResponse struct {
	// Amount of fees estimated per byte in atomic units
	Fee uint64 `json:"fee"`
	// Final fee should be rounded up to an even multiple of this value
	QuantizationMask uint64 `json:"quantization_mask"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetAlternateChainsRequest is a struct for GetAlternateChains() requests
type GetAlternateChainsRequest struct {
	// None
}

// GetAlternateChainsResponse is a struct for GetAlternateChains() responses
type GetAlternateChainsResponse struct {
	// array of chains
	Chains []Chain `json:"chains"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// Chain is a struct for GetAlternateChains()
type Chain struct {
	//  the block hash of the first diverging block of this alternative chain.
	BlockHash string `json:"block_hash"`
	//  the cumulative difficulty of all blocks in the alternative chain.
	Difficulty uint64 `json:"difficulty"`
	//  the block height of the first diverging block of this alternative chain.
	Height uint64 `json:"height"`
	//  the length in blocks of this alternative chain, after divergence.
	Length uint64 `json:"length"`
}

// RelayTxRequest is a struct for RelayTx() requests
type RelayTxRequest struct {
	// list of transaction IDs to relay
	Txids []string `json:"txids"`
}

// RelayTxResponse is a struct for RelayTx() responses
type RelayTxResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// SyncInfoRequest is a struct for SyncInfo() requests
type SyncInfoRequest struct {
	// None
}

// SyncInfoResponse is a struct for SyncInfo() responses
type SyncInfoResponse struct {
	// unsigned int
	Height uint64 `json:"height"`
	// info;structure of connection info, as defined in get_connections
	Peers []SyncPeer `json:"peers"`
	// array of span structure (optional, absent if node is fully synced)
	Spans []Span `json:"spans"`
	// Id of connection
	ConnectionID string `json:"connection_id"`
	// number of blocks in that span
	Nblocks uint64 `json:"nblocks"`
	// connection rate
	Rate uint64 `json:"rate"`
	// peer address the node is downloading (or has downloaded) than span from
	RemoteAddress string `json:"remote_address"`
	// total number of bytes in that span's blocks (including txes)
	Size uint64 `json:"size"`
	// connection speed
	Speed uint64 `json:"speed"`
	// block height of the first block in that span
	StartBlockHeight uint64 `json:"start_block_height"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// target height the node is syncing from (optional, absent if node is fully synced)
	TargetHeight uint64 `json:"target_height,omitempty"`
}

// SyncPeer is a struct for SyncInfo()
type SyncPeer struct {
	// connection info, as defined in get_connections
	Info Connection `json:"info"`
}

// Span is a struct for SyncInfo()
type Span struct {
	//  Id of connection
	ConnectionID string `json:"connection_id"`
	//  number of blocks in that span
	Nblocks uint64 `json:"nblocks"`
	//  connection rate
	Rate uint64 `json:"rate"`
	//  peer address the node is downloading (or has downloaded) than span from
	RemoteAddress string `json:"remote_address"`
	//  total number of bytes in that span's blocks (including txes)
	Size uint64 `json:"size"`
	//  connection speed
	Speed uint64 `json:"speed"`
	//  block height of the first block in that span
	StartBlockHeight uint64 `json:"start_block_height"`
}

// GetTxpoolBacklogRequest is a struct for GetTxpoolBacklog() requests
type GetTxpoolBacklogRequest struct {
	// None
}

// GetTxpoolBacklogResponse is a struct for GetTxpoolBacklog() responses
// FIXME currently failing
type GetTxpoolBacklogResponse struct {
	// array of tx_backlog_entry structures (in binary form)
	Backlog []TxBacklogEntry `json:"backlog"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// TxBacklogEntry is a struct for GetTxpoolBacklog()
// FIXME currently failing
type TxBacklogEntry struct {
	// unsigned int (in binary form)
	BlobSize uint `json:"blob_size"`
	// unsigned int (in binary form)
	Fee uint `json:"fee"`
	// unsigned int (in binary form)
	TimeInPool uint `json:"time_in_pool"`
}

// GetOutputDistributionRequest is a struct for GetOutputDistribution() requests
type GetOutputDistributionRequest struct {
	// amounts to look for
	Amounts []gonero.AtomicXMR `json:"amounts"`
	// (optional, default is false) States if the result should be cumulative (true) or not (false)
	Cumulative bool `json:"cumulative,omitempty"`
	// (optional, default is 0) starting height to check from
	FromHeight uint64 `json:"from_height,omitempty"`
	// (optional, default is 0) ending height to check up to
	ToHeight uint64 `json:"to_height,omitempty"`
}

// GetOutputDistributionResponse is a struct for GetOutputDistribution() responses
type GetOutputDistributionResponse struct {
	// array of OuputDistribution structures
	Distributions []OuputDistribution `json:"distributions"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// OuputDistribution is a struct for GetOutputDistribution()
type OuputDistribution struct {
	// unsigned int
	Amount gonero.AtomicXMR `json:"amount"`
	// unsigned int
	Base uint64 `json:"base"`
	// unsigned int
	Distribution []uint64 `json:"distribution"`
	// unsigned int
	StartHeight uint64 `json:"start_height"`
}

//#####################
// Other RPC Structs
//#####################

// GetHeightRequest is a struct for GetHeight() requests
type GetHeightRequest struct {
	// None
}

// GetHeightResponse is a struct for GetHeight() responses
type GetHeightResponse struct {
	// Current length of longest chain known to daemon.
	Height uint64 `json:"height"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetTransactionsRequest is a struct for GetTransactions() requests
type GetTransactionsRequest struct {
	// List of transaction hashes to look up.
	TxsHashes []string `json:"txs_hashes"`
	// Optional (false by default). If set true, the returned transaction information will be decoded rather than binary.
	DecodeAsJSON bool `json:"decode_as_json,omitempty"`
	// Optional (false by default).
	Prune bool `json:"prune,omitempty"`
}

// GetTransactionsResponse is a struct for GetTransactions() responses
type GetTransactionsResponse struct {
	//  (Optional - returned if not empty) Transaction hashes that could not be found.
	MissedTx []string `json:"missed_tx,omitempty"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// array of entry
	Txs []Transaction `json:"txs"`
	// Full transaction information as a hex string (old compatibility parameter)
	TxsAsHex string `json:"txs_as_hex"`
	// (Optional - returned if set in inputs. Old compatibility parameter) List of transaction as in as_json above:
	TxsAsJSON string `json:"txs_as_json,omitempty"`
}

// Transaction is a struct for GetTransactions()
type Transaction struct {
	//  Full transaction information as a hex string.
	AsHex string `json:"as_hex"`
	//  List of transaction info:
	AsJSON string `json:"as_json"`
	//  block height including the transaction
	BlockHeight uint64 `json:"block_height"`
	//  Unix time at chich the block has been added to the blockchain
	BlockTimestamp uint64 `json:"block_timestamp"`
	//  States if the transaction is a double
	DoubleSpendSeen bool `json:"double_spend_seen"`
	//  States if the transaction is in pool (true) or included in a block (false)
	InPool bool `json:"in_pool"`
	//  transaction indexes
	OutputIndices []uint64 `json:"output_indices"`
	//  transaction hash
	TxHash string `json:"tx_hash"`
}

// JSONTransaction is a struct containing transaction information
// XXX
type JSONTransaction struct {
	// Transaction version
	Version uint64 `json:"version"`
	// If not 0, this tells when a transaction output is spendable.
	UnlockTime uint64 `json:"unlock_time"`
	// List of inputs into transaction
	Vin []TransactionInput `json:"vin"`
	// List of outputs from transaction
	Vout []MinerTransactionOutput `json:"vout"`
	// Usually called the "payment ID" but can be used to include any random 32 bytes.
	Extra string `json:"extra"`
	// List of signatures used in ring signature to hide the true origin of the transaction.
	Signatures []string `json:"signatures"`
}

// TransactionInput is a struct that contains the inputs to a transaction
type TransactionInput struct {
	// The public key of the previous output spent in this transaction.
	Key string `json:"key"`
	// The amount of the input, in atomic units.
	Amount gonero.AtomicXMR `json:"amount"`
	// A list of integer offets to the input.
	KeyOffsets int64 `json:"key_offsets"`
	// The key image for the given input
	KImage string `json:"k_image"`
}

// TransactionOuput XXX
type TransactionOuput struct {
	// Amount of transaction output, in atomic units.
	Amount gonero.AtomicXMR `json:"amount"`
	// Output destination information
	Target Target `json:"target"`
}

// Target XXX
type Target struct {
	// The stealth public key of the receiver.
	// Whoever owns the private key associated with this key controls this transaction output.
	Key string `json:"key"`
}

// GetAltBlocksHashesRequest is a struct for GetAltBlocksHashes() requests
type GetAltBlocksHashesRequest struct {
	// None
}

// GetAltBlocksHashesResponse is a struct for GetAltBlocksHashes() responses
type GetAltBlocksHashesResponse struct {
	// list of alternative blocks hashes to main chain
	BlksHashes []string `json:"blks_hashes"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// IsKeyImageSpentRequest is a struct for IsKeyImageSpent() requests
type IsKeyImageSpentRequest struct {
	// List of key image hex strings to check.
	KeyImages []string `json:"key_images"`
}

// IsKeyImageSpentResponse is a struct for IsKeyImageSpent() responses
type IsKeyImageSpentResponse struct {
	// List of statuses for each image checked. Statuses are follows: 0 = unspent, 1 = spent in blockchain, 2 = spent in transaction pool
	SpentStatus []uint64 `json:"spent_status"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// SendRawTransactionRequest is a struct for SendRawTransaction() requests
type SendRawTransactionRequest struct {
	// Full transaction information as hexadecimal string.
	TxAsHex string `json:"tx_as_hex"`
	// Stop relaying transaction to other nodes (default is false).
	DoNotRelay bool `json:"do_not_relay"`
}

// SendRawTransactionResponse is a struct for SendRawTransaction() responses
type SendRawTransactionResponse struct {
	// Transaction is a double spend (true) or not (false).
	DoubleSpend bool `json:"double_spend"`
	// Fee is too low (true) or OK (false).
	FeeTooLow bool `json:"fee_too_low"`
	// Input is invalid (true) or valid (false).
	InvalidInput bool `json:"invalid_input"`
	// Output is invalid (true) or valid (false).
	InvalidOutput bool `json:"invalid_output"`
	// Mixin count is too low (true) or OK (false).
	LowMixin bool `json:"low_mixin"`
	// Transaction is a standard ring transaction (true) or a ring confidential transaction (false).
	NotRct bool `json:"not_rct"`
	// Transaction was not relayed (true) or relayed (false).
	NotRelayed bool `json:"not_relayed"`
	// Transaction uses more money than available (true) or not (false).
	Overspend bool `json:"overspend"`
	// Additional information. Currently empty or "Not relayed" if transaction was accepted but not relayed.
	Reason string `json:"reason"`
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
	// Transaction size is too big (true) or OK (false).
	TooBig bool `json:"too_big"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// StartMiningRequest is a struct for StartMining() requests
type StartMiningRequest struct {
	// States if the mining should run in background (true) or foreground (false).
	DoBackgroundMining bool `json:"do_background_mining"`
	// States if batery state (on laptop) should be ignored (true) or not (false).
	IgnoreBattery bool `json:"ignore_battery"`
	// Account address to mine to.
	MinerAddress string `json:"miner_address"`
	// Number of mining thread to run.
	ThreadsCount uint64 `json:"threads_count"`
}

// StartMiningResponse is a struct for StartMining() responses
type StartMiningResponse struct {
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
}

// StopMiningRequest is a struct for StopMining() requests
type StopMiningRequest struct {
	// None
}

// StopMiningResponse is a struct for StopMining() responses
type StopMiningResponse struct {
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
}

// MiningStatusRequest is a struct for MiningStatus() requests
type MiningStatusRequest struct {
	// None
}

// MiningStatusResponse is a struct for MiningStatus() responses
type MiningStatusResponse struct {
	// States if mining is enabled (true) or disabled (false).
	Active bool `json:"active"`
	// Account address daemon is mining to. Empty if not mining.
	Address string `json:"address"`
	// States if the mining is running in background (true) or foreground (false).
	IsBackgroundMiningEnabled bool `json:"is_background_mining_enabled"`
	// Mining power in hashes per seconds.
	Speed uint64 `json:"speed"`
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
	// Number of running mining threads.
	ThreadsCount uint64 `json:"threads_count"`
}

// SaveBcRequest is a struct for SaveBc() requests
type SaveBcRequest struct {
	// None
}

// SaveBcResponse is a struct for SaveBc() responses
type SaveBcResponse struct {
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
}

// GetPeerListRequest is a struct for GetPeerList() requests
type GetPeerListRequest struct {
	// None
}

// GetPeerListResponse is a struct for GetPeerList() responses
type GetPeerListResponse struct {
	// array of offline peer structure
	GrayList []Peer `json:"gray_list"`
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
	// array of online peer structure, as above.
	WhiteList []Peer `json:"white_list"`
}

// Peer is a struct for GetPeerList()
type Peer struct {
	//  IP address in string (integer) format TODO FIXME
	Host string `json:"host"`
	//  Peer id
	ID uint64 `json:"id"`
	//  IP address in integer format
	IP uint64 `json:"ip"`
	//  unix time at which the peer has been seen for the last time
	LastSeen uint64 `json:"last_seen"`
	//  TCP port the peer is using to connect to monero network.
	Port uint64 `json:"port"`
}

// SetLogHashRateRequest is a struct for SetLogHashRate() requests
type SetLogHashRateRequest struct {
	// States if hash rate logs should be visible (true) or hidden (false)
	Visible bool `json:"visible"`
}

// SetLogHashRateResponse is a struct for SetLogHashRate() responses
type SetLogHashRateResponse struct {
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
}

// SetLogLevelRequest is a struct for SetLogLevel() requests
type SetLogLevelRequest struct {
	// daemon log level to set from 0 (less verbose) to 4 (most verbose)
	Level int64 `json:"level"`
}

// SetLogLevelResponse is a struct for SetLogLevel() responses
type SetLogLevelResponse struct {
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
}

// SetLogCategoriesRequest is a struct for SetLogCategories() requests
type SetLogCategoriesRequest struct {
	// Optional, daemon log categories to enable
	Categories string `json:"categories,omitempty"`
}

// SetLogCategoriesResponse is a struct for SetLogCategories() responses
type SetLogCategoriesResponse struct {
	// daemon log enabled categories
	Categories string `json:"categories"`
	// General RPC error code. "OK" means everything looks good. Any other value means that something went wrong.
	Status RPCStatus `json:"status"`
}

// GetTransactionPoolRequest is a struct for GetTransactionPool() requests
type GetTransactionPoolRequest struct {
	// None
}

// GetTransactionPoolResponse is a struct for GetTransactionPool() responses
type GetTransactionPoolResponse struct {
	// List pf spent output key image
	SpentKeyImages []SpentKeyImage `json:"spent_key_images"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// List of transactions in the mempool are not in a block on the main chain at the moment
	Transactions []MempoolTransaction `json:"transactions"`
}

// SpentKeyImage is the key image for a a spent transaction
type SpentKeyImage struct {
	// Key image.
	IDHash string `json:"id_hash"`
	// tx hashes of the txes (usually one) spending that key image.
	TxsHashes []string `json:"txs_hashes"`
}

// MempoolTransaction is a struct containing information pertaining
// to a transaction currently in the mempool
type MempoolTransaction struct {
	// The size of the full transaction blob.
	BlobSize uint64 `json:"blob_size"`
	// States if this transaction has been seen as double spend.
	DoubleSpendSeen bool `json:"double_spend_seen"`
	// States if this transaction should not be relayed
	DoNotRelay bool `json:"do_not_relay"`
	// The amount of the mining fee included in the transaction, in atomic units.
	Fee uint64 `json:"fee"`
	// The transaction ID hash.
	IDHash string `json:"id_hash"`
	// States if the tx was included in a block at least once (true) or not (false).
	KeptByBlock bool `json:"kept_by_block"`
	// If the transaction validation has previously failed, this tells at what height that occurred.
	LastFailedHeight uint64 `json:"last_failed_height"`
	// Like the previous, this tells the previous transaction ID hash.
	LastFailedIDHash string `json:"last_failed_id_hash"`
	// Last unix time at which the transaction has been relayed.
	LastRelayedTime uint64 `json:"last_relayed_time"`
	// Tells the height of the most recent block with an output used in this transaction.
	MaxUsedBlockHeight uint64 `json:"max_used_block_height"`
	// Tells the hash of the most recent block with an output used in this transaction.
	MaxUsedBlockHash string `json:"max_used_block_hash"`
	// The Unix time that the transaction was first seen on the network by the node.
	ReceiveTime uint64 `json:"receive_time"`
	// States if this transaction has been relayed
	Relayed bool `json:"relayed"`
	// Hexadecimal blob represnting the transaction.
	TxBlob string `json:"tx_blob"`
	// JSON structure of all information in the transaction
	TxJSON string `json:"tx_json"`
}

// JSONMempoolTransaction XXX
type JSONMempoolTransaction struct {
	JSONTransaction
	// Ring signatures:
	RctSignatures  []RctSignature `json:"rct_signatures"`
	RctsigPrunable RctsigPrunable `json:"rctsig_prunable"`
}

// RctSignature XXX
type RctSignature struct {
	Type   string `json:"type"`
	TxnFee uint64 `json:"txnFee"`
	// array of Diffie Helman Elipctic curves structures
	EcdhInfo EcdhInfo `json:"ecdhInfo"`
	OutPk    string   `json:"outPk"`
}

// EcdhInfo XXX
type EcdhInfo struct {
	// String
	Mask string `json:"mask"`
	// String
	Amount string `json:"amount"`
}

// RctsigPrunable XXX
type RctsigPrunable struct {
	RangeSigs []RangeSig `json:"rangeSigs"`
	MGs       []MG       `json:"MGs"`
}

// RangeSig XXX
type RangeSig struct {
	Asig string `json:"asig"`
	Ci   string `json:"Ci"`
}

// MG XXX
type MG struct {
	// array of arrays of two strings.
	Ss [][2]string `json:"ss"`
	// String
	Cc string `json:"cc"`
}

// GetTransactionPoolHashesBinRequest is a struct for GetTransactionPoolHashesBin() requests
type GetTransactionPoolHashesBinRequest struct {
	// None
}

// GetTransactionPoolHashesBinResponse is a struct for GetTransactionPoolHashesBin() responses
type GetTransactionPoolHashesBinResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// binary array of transaction hashes.
	TxHashes [][]byte `json:"tx_hashes"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// GetTransactionPoolStatsRequest is a struct for GetTransactionPoolStats() requests
type GetTransactionPoolStatsRequest struct {
	// None
}

// GetTransactionPoolStatsResponse is a struct for GetTransactionPoolStats() responses
type GetTransactionPoolStatsResponse struct {
	// bytes_max;unsigned int;Max transaction size in pool bytes_med - unsigned int;Median transaction size in pool bytes_min - unsigned int;Min transaction size in pool bytes_total - unsigned int;total size of all transactions in pool histo - structure txpool_histo as follows:          txs - unsigned int;number of transactions bytes - unsigned int;size in bytes.   histo_98pc unsigned int;the time 98% of txes are "younger" than num_10m unsigned int;number of transactions in pool for more than 10 minutes num_double_spends unsigned int;number of double spend transactions num_failing unsigned int;number of failing transactions num_not_relayed unsigned int;number of non-relayed transactions oldest unsigned int;unix time of the oldest transaction in the pool txs_total unsigned int;total number of transactions.
	PoolStats PoolStats `json:"pool_stats"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// PoolStats is a struct for GetTransactionPoolStats()
type PoolStats struct {
	// Max transaction size in pool
	BytesMax uint64 `json:"bytes_max"`
	// Median transaction size in pool
	BytesMed uint64 `json:"bytes_med"`
	// Min transaction size in pool
	BytesMin uint64 `json:"bytes_min"`
	// total size of all transactions in pool
	BytesTotal uint64 `json:"bytes_total"`
	// txs;unsigned int;number of transactions bytes - unsigned int;size in bytes.
	Histo TxpoolHisto `json:"histo"`
	// unsigned int; the time 98% of txes are "younger" than
	Histo98pc uint64 `json:"histo_98pc"`
	// unsigned int; number of transactions in pool for more than 10 minutes
	Num10m uint64 `json:"num_10m"`
	// unsigned int; number of double spend transactions
	NumDoubleSpends uint64 `json:"num_double_spends"`
	// unsigned int; number of failing transactions
	NumFailing uint64 `json:"num_failing"`
	// unsigned int; number of non-relayed transactions
	NumNotRelayed uint64 `json:"num_not_relayed"`
	// unsigned int; unix time of the oldest transaction in the pool
	Oldest uint64 `json:"oldest"`
	// unsigned int; total number of transactions.
	TxsTotal uint64 `json:"txs_total"`
}

// TxpoolHisto is a struct for GetTransactionPoolStats()
type TxpoolHisto struct {
	//  number of transactions
	Txs uint64 `json:"txs"`
	//  size in bytes.
	Bytes uint64 `json:"bytes"`
}

// StopDaemonRequest is a struct for StopDaemon() requests
type StopDaemonRequest struct {
	// None
}

// StopDaemonResponse is a struct for StopDaemon() responses
type StopDaemonResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// GetLimitRequest is a struct for GetLimit() requests
type GetLimitRequest struct {
	// None
}

// GetLimitResponse is a struct for GetLimit() responses
type GetLimitResponse struct {
	// Download limit in kBytes per second
	LimitDown uint64 `json:"limit_down"`
	// Upload limit in kBytes per second
	LimitUp uint64 `json:"limit_up"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// SetLimitRequest is a struct for SetLimit() requests
type SetLimitRequest struct {
	// Download limit in kBytes per second (-1 reset to default, 0 don't change the current limit)
	LimitDown int64 `json:"limit_down"`
	// Upload limit in kBytes per second (-1 reset to default, 0 don't change the current limit)
	LimitUp int64 `json:"limit_up"`
}

// SetLimitResponse is a struct for SetLimit() responses
type SetLimitResponse struct {
	// Download limit in kBytes per second
	LimitDown uint64 `json:"limit_down"`
	// Upload limit in kBytes per second
	LimitUp uint64 `json:"limit_up"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// OutPeersRequest is a struct for OutPeers() requests
type OutPeersRequest struct {
	// Max number of outgoing peers
	OutPeers uint64 `json:"out_peers"`
}

// OutPeersResponse is a struct for OutPeers() responses
type OutPeersResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// InPeersRequest is a struct for InPeers() requests
type InPeersRequest struct {
	// Max number of incoming peers
	InPeers uint64 `json:"in_peers"`
}

// InPeersResponse is a struct for InPeers() responses
type InPeersResponse struct {
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
}

// GetOutsRequest is a struct for GetOuts() requests
type GetOutsRequest struct {
	// index - unsigned int
	Outputs Outputs `json:"outputs"`
	// unsigned int
	Amount gonero.AtomicXMR `json:"amount"`
	// unsigned int
	Index uint64 `json:"index"`
	// If true, a txid will included for each output in the response.
	GetTxid bool `json:"get_txid"`
}

// GetOutsResponse is a struct for GetOuts() responses
type GetOutsResponse struct {
	// height;unsigned int;block height of the output key - String;the public key of the output mask - String txid - String;transaction id unlocked - boolean;States if output is locked (false) or not (true)
	Outs Outs `json:"outs"`
	// block height of the output
	Height uint64 `json:"height"`
	// the public key of the output
	Key string `json:"key"`
	// String
	Mask string `json:"mask"`
	// transaction id
	Txid string `json:"txid"`
	// States if output is locked (false) or not (true)
	Unlocked bool `json:"unlocked"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if the result is obtained using the bootstrap mode, and is therefore not trusted (true), or when the daemon is fully synced (false).
	Untrusted bool `json:"untrusted"`
}

// Outputs is a struct for GetOuts()
type Outputs struct {
	//
	Amount gonero.AtomicXMR `json:"amount"`
	//
	Index uint64 `json:"index"`
}

// Outs is a struct for GetOuts()
type Outs struct {
	//  block height of the output
	Height uint64 `json:"height"`
	//  the public key of the output
	Key string `json:"key"`
	// --
	Mask string `json:"mask"`
	//  transaction id
	Txid string `json:"txid"`
	//  States if output is locked (false) or not (true)
	Unlocked bool `json:"unlocked"`
}

// UpdateRequest is a struct for Update() requests
type UpdateRequest struct {
	// command to use, either check or download
	Command string `json:"command"`
	// Optional, path where to download the update.
	Path string `json:"path,omitempty"`
}

// UpdateResponse is a struct for Update() responses
type UpdateResponse struct {
	// string
	AutoURI string `json:"auto_uri"`
	// string
	Hash string `json:"hash"`
	// path to download the update
	Path string `json:"path"`
	// General RPC error code. "OK" means everything looks good.
	Status RPCStatus `json:"status"`
	// States if an update is available to download (true) or not (false)
	Update bool `json:"update"`
	// string
	UserURI string `json:"user_uri"`
	// Version available for download.
	Version string `json:"version"`
}

////////////////
// Regtest only
////////////////

// GenerateBlocksRequest is a struct for GenerateBlocks requests
// This is only a valid command in Regtest mode!
type GenerateBlocksRequest struct {
	AmountOfBlocks uint   `json:"amount_of_blocks"`
	WalletAddress  string `json:"wallet_address"`
}

// GenerateBlocksResponse is a struct for GenerateBlocks responses
// This is only a valid command in Regtest mode!
type GenerateBlocksResponse struct {
	Height    uint64    `json:"height"`
	Status    RPCStatus `json:"status"`
	Untrusted bool      `json:"untrusted"`
}
