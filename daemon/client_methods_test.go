package daemon

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/konraddical2/gonero"
	"github.com/stretchr/testify/assert"
)

var cl Client

func setup() {
	cfg, _ := gonero.NewRPCConfig("", "localhost", StagenetPort, "", "", "")
	cl = New(cfg)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestClientMethods(t *testing.T) {
	type test struct {
		method  string
		request interface{}
	}

	tests := []test{
		{method: "GetBlockCount", request: &GetBlockCountRequest{}},
		{method: "OnGetBlockHash", request: &OnGetBlockHashRequest{}},
		{method: "GetBlockTemplate", request: &GetBlockTemplateRequest{}},
		{method: "SubmitBlock", request: &SubmitBlockRequest{}},
		{method: "GetLastBlockHeader", request: &GetLastBlockHeaderRequest{}},
		{method: "GetBlockHeaderByHash", request: &GetBlockHeaderByHashRequest{}},
		{method: "GetBlockHeaderByHeight", request: &GetBlockHeaderByHeightRequest{}},
		{method: "GetBlockHeadersRange", request: &GetBlockHeadersRangeRequest{}},
		{method: "GetBlock", request: &GetBlockRequest{}},
		{method: "GetConnections", request: &GetConnectionsRequest{}},
		{method: "GetInfo", request: &GetInfoRequest{}},
		{method: "HardForkInfo", request: &HardForkInfoRequest{}},
		{method: "SetBans", request: &SetBansRequest{}},
		{method: "GetBans", request: &GetBansRequest{}},
		{method: "FlushTxpool", request: &FlushTxpoolRequest{}},
		{method: "GetOutputHistogram", request: &GetOutputHistogramRequest{}},
		{method: "GetVersion", request: &GetVersionRequest{}},
		{method: "GetCoinbaseTxSum", request: &GetCoinbaseTxSumRequest{}},
		{method: "GetFeeEstimate", request: &GetFeeEstimateRequest{}},
		{method: "GetAlternateChains", request: &GetAlternateChainsRequest{}},
		{method: "RelayTx", request: &RelayTxRequest{}},
		{method: "SyncInfo", request: &SyncInfoRequest{}},
		{method: "GetTxpoolBacklog", request: &GetTxpoolBacklogRequest{}},
		{method: "GetOutputDistribution", request: &GetOutputDistributionRequest{}},
		{method: "GetHeight", request: &GetHeightRequest{}},
		{method: "GetTransactions", request: &GetTransactionsRequest{}},
		{method: "GetAltBlocksHashes", request: &GetAltBlocksHashesRequest{}},
		{method: "IsKeyImageSpent", request: &IsKeyImageSpentRequest{}},
		{method: "SendRawTransaction", request: &SendRawTransactionRequest{}},
		{method: "StartMining", request: &StartMiningRequest{}},
		{method: "StopMining", request: &StopMiningRequest{}},
		{method: "MiningStatus", request: &MiningStatusRequest{}},
		{method: "SaveBc", request: &SaveBcRequest{}},
		{method: "GetPeerList", request: &GetPeerListRequest{}},
		{method: "SetLogHashRate", request: &SetLogHashRateRequest{}},
		{method: "SetLogLevel", request: &SetLogLevelRequest{}},
		{method: "SetLogCategories", request: &SetLogCategoriesRequest{}},
		{method: "GetTransactionPool", request: &GetTransactionPoolRequest{}},
		{method: "GetTransactionPoolHashesBin", request: &GetTransactionPoolHashesBinRequest{}},
		{method: "GetTransactionPoolStats", request: &GetTransactionPoolStatsRequest{}},
		//{method: "StopDaemon", request: &StopDaemonRequest{}},
		{method: "GetLimit", request: &GetLimitRequest{}},
		{method: "SetLimit", request: &SetLimitRequest{}},
		{method: "OutPeers", request: &OutPeersRequest{}},
		{method: "InPeers", request: &InPeersRequest{}},
		{method: "GetOuts", request: &GetOutsRequest{}},
		{method: "Update", request: &UpdateRequest{}},
		{method: "GenerateBlocks", request: &GenerateBlocksRequest{}},
	}

	clType := reflect.ValueOf(cl)
	for _, tc := range tests {
		t.Run(tc.method, func(t *testing.T) {
			vals := clType.MethodByName(tc.method).Call([]reflect.Value{reflect.ValueOf(tc.request)})
			err := vals[1].Interface()
			if err != nil {
				assert.NotEqual(t, "http status 404", err.(error).Error())
				fmt.Println(tc.method, ":\n", err)
			}
		})

	}

	t.Run("GenerateBlocks should return an error when not connected to regtest", func(t *testing.T) {
		_, err := cl.GenerateBlocks(&GenerateBlocksRequest{})
		assert.Error(t, err)
		assert.EqualError(t, err, "-13: Regtest required when generating blocks")
	})
}
