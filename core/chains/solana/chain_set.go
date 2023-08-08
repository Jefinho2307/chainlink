package solana

import (
	"github.com/pkg/errors"
	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink-relay/pkg/logger"
	"github.com/smartcontractkit/chainlink-relay/pkg/loop"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/db"

	"github.com/smartcontractkit/chainlink/v2/core/chains"
)

// ChainSetOpts holds options for configuring a ChainSet.
type ChainServiceOpts struct {
	Logger            logger.Logger
	KeyStore          loop.Keystore
	ChainNodeStatuser ConfigStater
}

func (o *ChainServiceOpts) Validate() (err error) {
	required := func(s string) error {
		return errors.Errorf("%s is required", s)
	}
	if o.Logger == nil {
		err = multierr.Append(err, required("Logger"))
	}
	if o.KeyStore == nil {
		err = multierr.Append(err, required("KeyStore"))
	}
	if o.ChainNodeStatuser == nil {
		err = multierr.Append(err, required("Configs"))
	}
	return
}

func (o *ChainServiceOpts) ConfigsAndLogger() (chains.Statuser[db.Node], logger.Logger) {
	return o.ChainNodeStatuser, o.Logger
}

func NewTOMLChain(cfg *SolanaConfig, o ChainServiceOpts) (solana.Chain, error) {
	if !cfg.IsEnabled() {
		return nil, errors.Errorf("cannot create new chain with ID %s, the chain is disabled", *cfg.ChainID)
	}
	c, err := newChain(*cfg.ChainID, cfg, o)
	if err != nil {
		return nil, err
	}
	return c, nil
}

/*
func NewChainSet(opts ChainSetOpts, cfgs SolanaConfigs) (solana.ChainSet, error) {
	solChains := map[string]solana.Chain{}
	var err error
	for _, chain := range cfgs {
		if !chain.IsEnabled() {
			continue
		}
		var err2 error
		solChains[*chain.ChainID], err2 = opts.NewTOMLChain(chain)
		if err2 != nil {
			err = multierr.Combine(err, err2)
			continue
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to load some Solana chains")
	}
	return chains.NewChainSet[db.Node, solana.Chain](solChains, &opts)
}
*/
