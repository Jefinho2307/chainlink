package smoke

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ethereum/mercury/exchanger"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups/mercury"
	"github.com/stretchr/testify/require"
)

func TestMercurySmoke(t *testing.T) {
	l := zerolog.New(zerolog.NewTestWriter(t))

	testEnv := mercury.NewMercuryTestEnv(t, "smoke")
	testEnv.SetupFullMercuryEnv(nil, nil)

	var (
		feedId      = testEnv.Config.FeedId
		feedIdBytes = mercury.StringToByte32(feedId)
	)

	t.Run("test mercury server has report for the latest block number", func(t *testing.T) {
		t.Skip("askdlASDSADSADSASDSDASASDA")

		latestBlockNum, err := testEnv.EvmClient.LatestBlockNumber(context.Background())
		_ = latestBlockNum
		require.NoError(t, err, "Err getting latest block number")
		report, _, err := testEnv.MSClient.GetReports(feedId, latestBlockNum-5)
		require.NoError(t, err, "Error getting report from Mercury Server")
		require.NotEmpty(t, report.ChainlinkBlob, "Report response does not contain chainlinkBlob")
		// TODO: decode the report and validate it has correct fields
	})

	t.Run("test report verfification using Exchanger.ResolveTradeWithReport call", func(t *testing.T) {
		order := mercury.Order{
			FeedID:       feedIdBytes,
			CurrencySrc:  mercury.StringToByte32("1"),
			CurrencyDst:  mercury.StringToByte32("2"),
			AmountSrc:    big.NewInt(1),
			MinAmountDst: big.NewInt(2),
			Sender:       common.HexToAddress("c7ca5f083dce8c0034e9a6033032ec576d40b222"),
			Receiver:     common.HexToAddress("c7ca5f083dce8c0034e9a6033032ec576d40bf45"),
		}

		// Commit to a trade
		commitmentHash := mercury.CreateCommitmentHash(order)
		err := testEnv.ExchangerContract.CommitTrade(commitmentHash)
		require.NoError(t, err)

		// Resove the trade and get mercry server url
		encodedCommitment, err := mercury.CreateEncodedCommitment(order)
		require.NoError(t, err)
		mercuryUrlPath, err := testEnv.ExchangerContract.ResolveTrade(encodedCommitment)
		require.NoError(t, err)
		// feedIdHex param is still not fixed in the Exchanger contract. Should be feedIDHex
		fixedMerucyrUrlPath := strings.Replace(mercuryUrlPath, "feedIdHex", "feedIDHex", -1)

		// Get report from mercury server
		msClient := client.NewMercuryServerClient(
			testEnv.Config.MSLocalUrl, testEnv.Config.MSAdminId, testEnv.Config.MSAdminKey)
		report, _, err := msClient.CallGet(fmt.Sprintf("/client%s", fixedMerucyrUrlPath))
		l.Info().Msgf("Got response from Mercury server: %s", report)
		require.NoError(t, err, "Error getting report from Mercury Server")
		require.NotEmpty(t, report["chainlinkBlob"], "Report response does not contain chainlinkBlob")
		reportBlob := report["chainlinkBlob"].(string)

		// Resolve the trade with report
		reportBytes, err := hex.DecodeString(reportBlob[2:])
		require.NoError(t, err)
		receipt, err := testEnv.ExchangerContract.ResolveTradeWithReport(reportBytes, encodedCommitment)
		require.NoError(t, err)

		// Get transaction logs
		exchangerABI, err := abi.JSON(strings.NewReader(exchanger.ExchangerABI))
		require.NoError(t, err)
		tradeExecuted := map[string]interface{}{}
		err = exchangerABI.UnpackIntoMap(tradeExecuted, "TradeExecuted", receipt.Logs[1].Data)
		require.NoError(t, err)
		l.Info().Interface("TradeExecuted", tradeExecuted).Msg("ResolveTradeWithReport logs")
	})
}
