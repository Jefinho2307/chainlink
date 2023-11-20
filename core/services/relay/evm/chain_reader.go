package evm

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

type ChainReaderService interface {
	services.ServiceCtx
	commontypes.ChainReader
}

type chainReader struct {
	lggr       logger.Logger
	contractID common.Address
	lp         logpoller.LogPoller
}

// NewChainReaderService constructor for ChainReader, returns nil if there are any errors.
func NewChainReaderService(lggr logger.Logger, lp logpoller.LogPoller, ropts *types.RelayOpts) (*chainReader, error) {
	relayConfig, err := ropts.RelayConfig()
	if err != nil {
		return nil, fmt.Errorf("failed parsing RelayConfig: %w", err)
	}

	if !common.IsHexAddress(ropts.ContractID) {
		return nil, fmt.Errorf("invalid contractID, expected hex address")
	}
	contractID := common.HexToAddress(ropts.ContractID)

	if relayConfig.ChainReader == nil {
		return nil, errors.ErrUnsupported
	}

	if err = validateChainReaderConfig(*relayConfig.ChainReader); err != nil {
		return nil, err
	}

	return &chainReader{lggr.Named("ChainReader"), contractID, lp}, nil
}

func (cr *chainReader) Name() string { return cr.lggr.Name() }

func (cr *chainReader) initialize() error {
	// Initialize chain reader, start cache polling loop, etc.
	return nil
}

func (cr *chainReader) Start(ctx context.Context) error {
	if err := cr.initialize(); err != nil {
		return fmt.Errorf("Failed to initialize ChainReader: %w", err)
	}
	return nil
}

func (cr *chainReader) Close() error { return nil }

func (cr *chainReader) Ready() error { return nil }

func (cr *chainReader) HealthReport() map[string]error {
	return map[string]error{cr.Name(): nil}
}

func (cr *chainReader) GetLatestValue(ctx context.Context, bc commontypes.BoundContract, method string, params any, returnVal any) error {
	return fmt.Errorf("Unimplemented method GetLatestValue called %w", errors.ErrUnsupported)
}

func (cr *chainReader) Encode(ctx context.Context, item any, itemType string) (ocrtypes.Report, error) {
	return nil, fmt.Errorf("Unimplemented method Encode called %w", errors.ErrUnsupported)
}

func (cr *chainReader) Decode(_ context.Context, raw []byte, into any, itemType string) error {
	return fmt.Errorf("Unimplemented method Decode called %w", errors.ErrUnsupported)
}

func (cr *chainReader) GetMaxEncodingSize(ctx context.Context, n int, itemType string) (int, error) {
	return 0, fmt.Errorf("Unimplemented method GetMaxDecodingSize called %w", errors.ErrUnsupported)
}

func (cr *chainReader) GetMaxDecodingSize(ctx context.Context, n int, itemType string) (int, error) {
	return 0, fmt.Errorf("Unimplemented method GetMaxDecodingSize called %w", errors.ErrUnsupported)
}

func validateChainReaderConfig(cfg types.ChainReaderConfig) error {
	for contractName, chainContractReader := range cfg.ChainContractReaders {
		abi, err := abi.JSON(strings.NewReader(chainContractReader.ContractABI))
		if err != nil {
			return err
		}

		for chainReadingDefinitionName, chainReaderDefinition := range chainContractReader.ChainReaderDefinitions {
			switch chainReaderDefinition.ReadType {
			case types.Method:
				err = validateMethods(abi, chainReaderDefinition)
			case types.Event:
				err = validateEvents(abi, chainReaderDefinition)
			default:
				return fmt.Errorf("%w: invalid chain reader definition read type: %d", commontypes.ErrInvalidConfig, chainReaderDefinition.ReadType)
			}
			if err != nil {
				return fmt.Errorf("%w: invalid chain reader config for contract: %q chain reading definition: %q, err: %w", commontypes.ErrInvalidConfig, contractName, chainReadingDefinitionName, err)
			}
		}
	}

	return nil
}

func validateEvents(contractABI abi.ABI, chainReaderDefinition types.ChainReaderDefinition) error {
	event, methodExists := contractABI.Events[chainReaderDefinition.ChainSpecificName]
	if !methodExists {
		return fmt.Errorf("method: %s doesn't exist", chainReaderDefinition.ChainSpecificName)
	}

	if !areChainReaderArgumentsValid(event.Inputs, chainReaderDefinition.ReturnValues) {
		var abiEventInputsNames []string
		for _, input := range event.Inputs {
			abiEventInputsNames = append(abiEventInputsNames, input.Name)
		}
		return fmt.Errorf("return values: [%s] don't match abi event inputs: [%s]", strings.Join(chainReaderDefinition.ReturnValues, ","), strings.Join(abiEventInputsNames, ","))
	}

	var abiEventIndexedInputs []abi.Argument
	for _, eventInput := range event.Inputs {
		if eventInput.Indexed {
			abiEventIndexedInputs = append(abiEventIndexedInputs, eventInput)
		}
	}

	var chainReaderEventParams []string
	for chainReaderEventParam := range chainReaderDefinition.Params {
		chainReaderEventParams = append(chainReaderEventParams, chainReaderEventParam)
	}

	if !areChainReaderArgumentsValid(abiEventIndexedInputs, chainReaderEventParams) {
		var abiEventIndexedInputsNames []string
		for _, abiEventIndexedInput := range abiEventIndexedInputs {
			abiEventIndexedInputsNames = append(abiEventIndexedInputsNames, abiEventIndexedInput.Name)
		}
		return fmt.Errorf("params: [%s] don't match abi event indexed inputs: [%s]", strings.Join(chainReaderEventParams, ","), strings.Join(abiEventIndexedInputsNames, ","))
	}
	return nil
}

func validateMethods(abi abi.ABI, chainReaderDefinition types.ChainReaderDefinition) error {
	method, methodExists := abi.Methods[chainReaderDefinition.ChainSpecificName]
	if !methodExists {
		return fmt.Errorf("method: %q doesn't exist", chainReaderDefinition.ChainSpecificName)
	}

	var methodNames []string
	for methodName := range chainReaderDefinition.Params {
		methodNames = append(methodNames, methodName)
	}

	if !areChainReaderArgumentsValid(method.Inputs, methodNames) {
		var abiMethodInputs []string
		for _, input := range method.Inputs {
			abiMethodInputs = append(abiMethodInputs, input.Name)
		}
		return fmt.Errorf("params: [%s] don't match abi method inputs: [%s]", strings.Join(methodNames, ","), strings.Join(abiMethodInputs, ","))
	}

	if !areChainReaderArgumentsValid(method.Outputs, chainReaderDefinition.ReturnValues) {
		var abiMethodOutputs []string
		for _, input := range method.Outputs {
			abiMethodOutputs = append(abiMethodOutputs, input.Name)
		}
		return fmt.Errorf("return values: [%s] don't match abi method outputs: [%s]", strings.Join(chainReaderDefinition.ReturnValues, ","), strings.Join(abiMethodOutputs, ","))
	}

	return nil
}

func areChainReaderArgumentsValid(contractArgs []abi.Argument, chainReaderArgs []string) bool {
	for _, chArgName := range chainReaderArgs {
		found := false
		for _, contractArg := range contractArgs {
			if chArgName == contractArg.Name {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
