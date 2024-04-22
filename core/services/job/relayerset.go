package job

import (
	"github.com/smartcontractkit/chainlink-common/pkg/types/relayerset"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
)

type RelayerSet struct {
	relayers chainlink.RelayerChainInteroperators
}


func (r *RelayerSet) Get(id relayerset.RelayID) (relayerset.Relayer, error) {




}

GetAll() ([]Relayer, error)
