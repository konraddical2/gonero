package wallet

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
	cfg, _ := gonero.NewRPCConfig("https", "localhost", 6061, "", "", "")
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
		{method: "SetDaemon", request: &SetDaemonRequest{}},
		{method: "GetBalance", request: &GetBalanceRequest{}},
		{method: "GetAddress", request: &GetAddressRequest{}},
		{method: "GetAddressIndex", request: &GetAddressIndexRequest{}},
		{method: "CreateAddress", request: &CreateAddressRequest{}},
		{method: "LabelAddress", request: &LabelAddressRequest{}},
		{method: "ValidateAddress", request: &ValidateAddressRequest{}},
		{method: "GetAccounts", request: &GetAccountsRequest{}},
		{method: "CreateAccount", request: &CreateAccountRequest{}},
		{method: "LabelAccount", request: &LabelAccountRequest{}},
		{method: "GetAccountTags", request: &GetAccountTagsRequest{}},
		{method: "TagAccounts", request: &TagAccountsRequest{}},
		{method: "UntagAccounts", request: &UntagAccountsRequest{}},
		{method: "SetAccountTagDescription", request: &SetAccountTagDescriptionRequest{}},
		{method: "GetHeight", request: &GetHeightRequest{}},
		{method: "Transfer", request: &TransferRequest{}},
		{method: "TransferSplit", request: &TransferSplitRequest{}},
		{method: "SignTransfer", request: &SignTransferRequest{}},
		{method: "SubmitTransfer", request: &SubmitTransferRequest{}},
		{method: "SweepDust", request: &SweepDustRequest{}},
		{method: "SweepAll", request: &SweepAllRequest{}},
		{method: "SweepSingle", request: &SweepSingleRequest{}},
		{method: "RelayTx", request: &RelayTxRequest{}},
		{method: "Store", request: &StoreRequest{}},
		{method: "GetPayments", request: &GetPaymentsRequest{}},
		{method: "GetBulkPayments", request: &GetBulkPaymentsRequest{}},
		{method: "IncomingTransfers", request: &IncomingTransfersRequest{}},
		{method: "QueryKey", request: &QueryKeyRequest{}},
		{method: "MakeIntegratedAddress", request: &MakeIntegratedAddressRequest{}},
		{method: "SplitIntegratedAddress", request: &SplitIntegratedAddressRequest{}},
		//{method: "StopWallet", request: &StopWalletRequest{}},
		{method: "RescanBlockchain", request: &RescanBlockchainRequest{}},
		{method: "SetTxNotes", request: &SetTxNotesRequest{}},
		{method: "GetTxNotes", request: &GetTxNotesRequest{}},
		{method: "SetAttribute", request: &SetAttributeRequest{}},
		{method: "GetAttribute", request: &GetAttributeRequest{}},
		{method: "GetTxKey", request: &GetTxKeyRequest{}},
		{method: "CheckTxKey", request: &CheckTxKeyRequest{}},
		{method: "GetTxProof", request: &GetTxProofRequest{}},
		{method: "CheckTxProof", request: &CheckTxProofRequest{}},
		{method: "GetSpendProof", request: &GetSpendProofRequest{}},
		{method: "CheckSpendProof", request: &CheckSpendProofRequest{}},
		{method: "GetReserveProof", request: &GetReserveProofRequest{}},
		{method: "CheckReserveProof", request: &CheckReserveProofRequest{}},
		{method: "GetTransfers", request: &GetTransfersRequest{}},
		{method: "GetTransferByTxid", request: &GetTransferByTxidRequest{}},
		{method: "DescribeTransfer", request: &DescribeTransferRequest{}},
		{method: "Sign", request: &SignRequest{}},
		{method: "Verify", request: &VerifyRequest{}},
		{method: "ExportOutputs", request: &ExportOutputsRequest{}},
		{method: "ImportOutputs", request: &ImportOutputsRequest{}},
		{method: "ExportKeyImages", request: &ExportKeyImagesRequest{}},
		{method: "ImportKeyImages", request: &ImportKeyImagesRequest{}},
		{method: "MakeURI", request: &MakeURIRequest{}},
		{method: "ParseURI", request: &ParseURIRequest{}},
		{method: "GetAddressBook", request: &GetAddressBookRequest{}},
		{method: "AddAddressBook", request: &AddAddressBookRequest{}},
		{method: "EditAddressBook", request: &EditAddressBookRequest{}},
		{method: "DeleteAddressBook", request: &DeleteAddressBookRequest{}},
		{method: "Refresh", request: &RefreshRequest{}},
		{method: "AutoRefresh", request: &AutoRefreshRequest{}},
		{method: "RescanSpent", request: &RescanSpentRequest{}},
		{method: "StartMining", request: &StartMiningRequest{}},
		{method: "StopMining", request: &StopMiningRequest{}},
		{method: "GetLanguages", request: &GetLanguagesRequest{}},
		{method: "CreateWallet", request: &CreateWalletRequest{}},
		{method: "GenerateFromKeys", request: &GenerateFromKeysRequest{}},
		{method: "OpenWallet", request: &OpenWalletRequest{}},
		{method: "RestoreDeterministicWallet", request: &RestoreDeterministicWalletRequest{}},
		//{method: "CloseWallet", request: &CloseWalletRequest{}},
		{method: "ChangeWalletPassword", request: &ChangeWalletPasswordRequest{}},
		{method: "IsMultisig", request: &IsMultisigRequest{}},
		{method: "PrepareMultisig", request: &PrepareMultisigRequest{}},
		{method: "MakeMultisig", request: &MakeMultisigRequest{}},
		{method: "ExportMultisigInfo", request: &ExportMultisigInfoRequest{}},
		{method: "ImportMultisigInfo", request: &ImportMultisigInfoRequest{}},
		{method: "FinalizeMultisig", request: &FinalizeMultisigRequest{}},
		{method: "SignMultisig", request: &SignMultisigRequest{}},
		{method: "SubmitMultisig", request: &SubmitMultisigRequest{}},
		{method: "GetVersion", request: &GetVersionRequest{}},
	}

	clType := reflect.ValueOf(cl)
	for _, tc := range tests {
		t.Run(tc.method, func(t *testing.T) {
			vals := clType.MethodByName(tc.method).Call([]reflect.Value{reflect.ValueOf(tc.request)})
			if vals[1].Interface() != nil {
				err := vals[1].Interface().(error)
				assert.NotEqual(t, "http status 404", err.Error())
				fmt.Println(tc.method, ":\n", err)
			}
		})
	}
}
