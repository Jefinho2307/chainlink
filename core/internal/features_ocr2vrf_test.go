package internal_test

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/libocr/commontypes"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	ocr2dkg "github.com/smartcontractkit/ocr2vrf/dkg"
	"github.com/smartcontractkit/ocr2vrf/ocr2vrf"
	ocr2vrftypes "github.com/smartcontractkit/ocr2vrf/types"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/kyber/v3"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_v3_aggregator_contract"
	dkg_wrapper "github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/dkg"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/vrf_beacon_consumer"
	vrf_wrapper "github.com/smartcontractkit/chainlink/core/gethwrappers/ocr2vrf/generated/vrf_beacon_coordinator"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/dkgencryptkey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/dkgsignkey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/validate"
	"github.com/smartcontractkit/chainlink/core/services/ocrbootstrap"
)

type ocr2vrfUniverse struct {
	owner   *bind.TransactOpts
	backend *backends.SimulatedBackend

	dkgAddress common.Address
	dkg        *dkg_wrapper.DKG

	coordinatorAddress common.Address
	coordinator        *vrf_wrapper.VRFBeaconCoordinator

	linkAddress common.Address
	link        *link_token_interface.LinkToken

	consumerAddress common.Address
	consumer        *vrf_beacon_consumer.BeaconVRFConsumer

	feedAddress common.Address
	feed        *mock_v3_aggregator_contract.MockV3AggregatorContract
}

func setupOCR2VRFContracts(
	t *testing.T, beaconPeriod int64, keyID [32]byte, consumerShouldFail bool) ocr2vrfUniverse {
	owner := testutils.MustNewSimTransactor(t)
	genesisData := core.GenesisAlloc{
		owner.From: {
			Balance: assets.Ether(100),
		},
	}
	b := backends.NewSimulatedBackend(genesisData, ethconfig.Defaults.Miner.GasCeil*2)

	// deploy OCR2VRF contracts, which have the following deploy order:
	// * link token
	// * link/eth feed
	// * DKG
	// * VRF
	// * VRF consumer
	linkAddress, _, link, err := link_token_interface.DeployLinkToken(
		owner, b)
	require.NoError(t, err)

	b.Commit()

	feedAddress, _, feed, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(
		owner, b, 18, assets.GWei(1e7)) // 0.01 eth per link
	require.NoError(t, err)

	b.Commit()

	dkgAddress, _, dkg, err := dkg_wrapper.DeployDKG(owner, b)
	require.NoError(t, err)

	b.Commit()

	coordinatorAddress, _, coordinator, err := vrf_wrapper.DeployVRFBeaconCoordinator(
		owner, b, linkAddress, big.NewInt(beaconPeriod), dkgAddress, keyID)
	require.NoError(t, err)

	b.Commit()

	consumerAddress, _, consumer, err := vrf_beacon_consumer.DeployBeaconVRFConsumer(
		owner, b, coordinatorAddress, consumerShouldFail, big.NewInt(beaconPeriod))
	require.NoError(t, err)

	b.Commit()

	_, err = dkg.AddClient(owner, keyID, coordinatorAddress)
	require.NoError(t, err)

	// Achieve finality depth so the CL node can work properly.
	for i := 0; i < 20; i++ {
		b.Commit()
	}

	return ocr2vrfUniverse{
		owner:              owner,
		backend:            b,
		dkgAddress:         dkgAddress,
		dkg:                dkg,
		coordinatorAddress: coordinatorAddress,
		coordinator:        coordinator,
		linkAddress:        linkAddress,
		link:               link,
		consumerAddress:    consumerAddress,
		consumer:           consumer,
		feedAddress:        feedAddress,
		feed:               feed,
	}
}

func TestIntegration_OCR2VRF(t *testing.T) {
	t.Parallel()

	keyID := randomKeyID(t)
	uni := setupOCR2VRFContracts(t, 5, keyID, false)

	t.Log("Creating bootstrap node")

	bootstrapNodePort := randomPort(t)
	bootstrapNode := setupNodeOCR2(t, uni.owner, bootstrapNodePort, "bootstrap", uni.backend)
	numNodes := 5

	t.Log("Creating OCR2 nodes")
	var (
		oracles        []confighelper2.OracleIdentityExtra
		transmitters   []common.Address
		onchainPubKeys []common.Address
		kbs            []ocr2key.KeyBundle
		apps           []*cltest.TestApplication
		dkgEncrypters  []dkgencryptkey.Key
		dkgSigners     []dkgsignkey.Key
	)
	for i := 0; i < numNodes; i++ {
		node := setupNodeOCR2(t, uni.owner, bootstrapNodePort+uint16(i+1), fmt.Sprintf("ocr2vrforacle%d", i), uni.backend)
		// Supply the bootstrap IP and port as a V2 peer address
		node.config.Overrides.P2PV2Bootstrappers = []commontypes.BootstrapperLocator{
			{PeerID: bootstrapNode.peerID, Addrs: []string{
				fmt.Sprintf("127.0.0.1:%d", bootstrapNodePort),
			}},
		}

		dkgSignKey, err := node.app.GetKeyStore().DKGSign().Create()
		require.NoError(t, err)

		dkgEncryptKey, err := node.app.GetKeyStore().DKGEncrypt().Create()
		require.NoError(t, err)

		kbs = append(kbs, node.keybundle)
		apps = append(apps, node.app)
		transmitters = append(transmitters, node.transmitter)
		dkgEncrypters = append(dkgEncrypters, dkgEncryptKey)
		dkgSigners = append(dkgSigners, dkgSignKey)
		onchainPubKeys = append(onchainPubKeys, common.BytesToAddress(node.keybundle.PublicKey()))
		oracles = append(oracles, confighelper2.OracleIdentityExtra{
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  node.keybundle.PublicKey(),
				TransmitAccount:   ocrtypes2.Account(node.transmitter.String()),
				OffchainPublicKey: node.keybundle.OffchainPublicKey(),
				PeerID:            node.peerID,
			},
			ConfigEncryptionPublicKey: node.keybundle.ConfigEncryptionPublicKey(),
		})
	}

	t.Log("starting ticker to commit blocks")
	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()
	go func() {
		for range tick.C {
			uni.backend.Commit()
		}
	}()

	blockBeforeConfig, err := uni.backend.BlockByNumber(context.Background(), nil)
	require.NoError(t, err)

	t.Log("Setting DKG config before block:", blockBeforeConfig.Number().String())

	// set config for dkg
	setDKGConfig(
		t,
		uni,
		onchainPubKeys,
		transmitters,
		1,
		oracles,
		dkgSigners,
		dkgEncrypters,
		keyID,
	)

	t.Log("Adding bootstrap node job")
	err = bootstrapNode.app.Start(testutils.Context(t))
	require.NoError(t, err)
	defer bootstrapNode.app.Stop()

	chainSet := bootstrapNode.app.GetChains().EVM
	require.NotNil(t, chainSet)
	bootstrapJobSpec := fmt.Sprintf(`
type				= "bootstrap"
name				= "bootstrap"
relay				= "evm"
schemaVersion		= 1
contractID			= "%s"
[relayConfig]
chainID 			= 1337
`, uni.dkgAddress.Hex())
	t.Log("Creating bootstrap job:", bootstrapJobSpec)
	ocrJob, err := ocrbootstrap.ValidatedBootstrapSpecToml(bootstrapJobSpec)
	require.NoError(t, err)
	err = bootstrapNode.app.AddJobV2(context.Background(), &ocrJob)
	require.NoError(t, err)

	t.Log("Creating OCR2VRF jobs")
	var jobIDs []int32
	for i := 0; i < numNodes; i++ {
		err = apps[i].Start(testutils.Context(t))
		require.NoError(t, err)
		defer apps[i].Stop()

		jobSpec := fmt.Sprintf(`
type                 = "offchainreporting2"
schemaVersion        = 1
name                 = "ocr2 vrf integration test"
maxTaskDuration      = "30s"
contractID           = "%s"
ocrKeyBundleID       = "%s"
relay                = "evm"
pluginType           = "ocr2vrf"
transmitterID        = "%s"

[relayConfig]
chainID              = 1337

[pluginConfig]
dkgEncryptionPublicKey = "%s"
dkgSigningPublicKey    = "%s"
dkgKeyID               = "%s"
dkgContractAddress     = "%s"

linkEthFeedAddress     = "%s"
confirmationDelays     = %s # This is an array
lookbackBlocks         = %d # This is an integer
`, uni.coordinatorAddress.String(),
			kbs[i].ID(),
			transmitters[i],
			dkgEncrypters[i].PublicKeyString(),
			dkgSigners[i].PublicKeyString(),
			hex.EncodeToString(keyID[:]),
			uni.dkgAddress.String(),
			uni.feedAddress.String(),
			"[1, 2, 3, 4, 5, 6, 7, 8]", // conf delays
			1000,                       // lookback blocks
		)
		t.Log("Creating OCR2VRF job with spec:", jobSpec)
		ocrJob, err := validate.ValidatedOracleSpecToml(apps[i].Config, jobSpec)
		require.NoError(t, err)
		err = apps[i].AddJobV2(context.Background(), &ocrJob)
		require.NoError(t, err)
		jobIDs = append(jobIDs, ocrJob.ID)
	}

	t.Log("jobs added, running log poller replay")

	// Once all the jobs are added, replay to ensure we have the configSet logs.
	for _, app := range apps {
		require.NoError(t, app.Chains.EVM.Chains()[0].LogPoller().Replay(context.Background(), blockBeforeConfig.Number().Int64()))
	}
	require.NoError(t, bootstrapNode.app.Chains.EVM.Chains()[0].LogPoller().Replay(context.Background(), blockBeforeConfig.Number().Int64()))

	t.Log("Waiting for DKG key to get written")
	// poll until a DKG key is written to the contract
	// at that point we can start sending VRF requests
	var emptyKH [32]byte
	emptyHash := crypto.Keccak256Hash(emptyKH[:])
	gomega.NewWithT(t).Eventually(func() bool {
		kh, err := uni.coordinator.SProvingKeyHash(&bind.CallOpts{
			Context: testutils.Context(t),
		})
		require.NoError(t, err)
		t.Log("proving keyhash:", hexutil.Encode(kh[:]))
		return crypto.Keccak256Hash(kh[:]) != emptyHash
	}, testutils.WaitTimeout(t), 5*time.Second).Should(gomega.BeTrue())

	t.Log("DKG key written, setting VRF config")

	// set config for vrf now that dkg is ready
	setVRFConfig(
		t,
		uni,
		onchainPubKeys,
		transmitters,
		1,
		oracles,
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
		keyID)

	t.Log("Sending VRF request")

	// Send a VRF request and mine it
	_, err = uni.consumer.TestRequestRandomness(uni.owner, 2, 1, big.NewInt(1))
	_, err = uni.consumer.TestRequestRandomnessFulfillment(uni.owner, 1, 1, big.NewInt(2), 50_000, []byte{})
	require.NoError(t, err)

	uni.backend.Commit()

	t.Log("waiting for fulfillment")

	// poll until we're able to redeem the randomness without reverting
	// at that point, it's been fulfilled
	gomega.NewWithT(t).Eventually(func() bool {
		_, err1 := uni.consumer.TestRedeemRandomness(uni.owner, big.NewInt(0))
		t.Logf("TestRedeemRandomness err: %+v", err1)
		return err1 == nil
	}, testutils.WaitTimeout(t), 5*time.Second).Should(gomega.BeTrue())

	// Mine block after redeeming randomness
	uni.backend.Commit()

	// poll until we're able to verify that consumer contract has stored randomness as expected
	// First arg is the request ID, which starts at zero, second is the index into
	// the random words.
	gomega.NewWithT(t).Eventually(func() bool {
		rw1, err1 := uni.consumer.SReceivedRandomnessByRequestID(nil, big.NewInt(0), big.NewInt(0))
		t.Logf("TestRedeemRandomness 1st word err: %+v", err1)
		rw2, err2 := uni.consumer.SReceivedRandomnessByRequestID(nil, big.NewInt(0), big.NewInt(1))
		t.Logf("TestRedeemRandomness 2nd word err: %+v", err2)
		rw3, err3 := uni.consumer.SReceivedRandomnessByRequestID(nil, big.NewInt(1), big.NewInt(0))
		t.Logf("FulfillRandomness 1st word err: %+v", err3)
		t.Log("randomness from redeemRandomness:", rw1.String(), rw2.String())
		t.Log("randomness from fulfillRandomness:", rw3.String())
		return err1 == nil && err2 == nil && err3 == nil
	}, testutils.WaitTimeout(t), 5*time.Second).Should(gomega.BeTrue())
}

func setDKGConfig(
	t *testing.T,
	uni ocr2vrfUniverse,
	onchainPubKeys []common.Address,
	transmitters []common.Address,
	f uint8,
	oracleIdentities []confighelper2.OracleIdentityExtra,
	signKeys []dkgsignkey.Key,
	encryptKeys []dkgencryptkey.Key,
	keyID [32]byte,
) {
	var (
		signingPubKeys []kyber.Point
		encryptPubKeys []kyber.Point
	)
	for i := range signKeys {
		signingPubKeys = append(signingPubKeys, signKeys[i].PublicKey)
		encryptPubKeys = append(encryptPubKeys, encryptKeys[i].PublicKey)
	}

	offchainConfig, err := ocr2dkg.OffchainConfig(
		encryptPubKeys,
		signingPubKeys,
		&altbn_128.G1{},
		&ocr2vrftypes.PairingTranslation{
			Suite: &altbn_128.PairingSuite{},
		})
	require.NoError(t, err)
	onchainConfig, err := ocr2dkg.OnchainConfig(keyID)
	require.NoError(t, err)

	var schedule []int
	for range oracleIdentities {
		schedule = append(schedule, 1)
	}

	_, _, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		30*time.Second,
		10*time.Second,
		10*time.Second,
		20*time.Second,
		20*time.Second,
		3,
		schedule,
		oracleIdentities,
		offchainConfig,
		50*time.Millisecond,
		10*time.Second,
		10*time.Second,
		100*time.Millisecond,
		1*time.Second,
		int(f),
		onchainConfig)
	require.NoError(t, err)

	_, err = uni.dkg.SetConfig(uni.owner, onchainPubKeys, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	require.NoError(t, err)

	uni.backend.Commit()
}

func setVRFConfig(
	t *testing.T,
	uni ocr2vrfUniverse,
	onchainPubKeys []common.Address,
	transmitters []common.Address,
	f uint8,
	oracleIdentities []confighelper2.OracleIdentityExtra,
	confDelaysSl []int,
	keyID [32]byte,
) {
	offchainConfig := ocr2vrf.OffchainConfig()

	confDelays := make(map[uint32]struct{})
	for _, c := range confDelaysSl {
		confDelays[uint32(c)] = struct{}{}
	}

	onchainConfig := ocr2vrf.OnchainConfig(confDelays)

	var schedule []int
	for range oracleIdentities {
		schedule = append(schedule, 1)
	}

	_, _, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		30*time.Second,
		10*time.Second,
		10*time.Second,
		20*time.Second,
		20*time.Second,
		3,
		schedule,
		oracleIdentities,
		offchainConfig,
		50*time.Millisecond,
		10*time.Second,
		10*time.Second,
		100*time.Millisecond,
		1*time.Second,
		int(f),
		onchainConfig)
	require.NoError(t, err)

	_, err = uni.coordinator.SetConfig(
		uni.owner, onchainPubKeys, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	require.NoError(t, err)

	uni.backend.Commit()
}

func randomKeyID(t *testing.T) (r [32]byte) {
	_, err := rand.Read(r[:])
	require.NoError(t, err)
	return
}

func randomPort(t *testing.T) uint16 {
	p, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint16))
	require.NoError(t, err)
	return uint16(p.Uint64())
}
